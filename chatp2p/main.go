package main

import (
	"context"
	"fmt"
	"log"

	chat "github.com/PhiSyX/golang/chatp2p/chat"
	cli "github.com/PhiSyX/golang/chatp2p/cli"
	dns "github.com/PhiSyX/golang/chatp2p/dns"
	ui "github.com/PhiSyX/golang/chatp2p/ui"

	"github.com/libp2p/go-libp2p"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
)

func create_api(ctx *context.Context, cli_args *cli.CLI, ui *ui.UI) (*chat.NodeAPI, error) {
	var options []libp2p.Option
	options = append(options, libp2p.ListenAddrStrings("/ip4/0.0.0.0/tcp/0", "/ip6/::/tcp/0"))

	host, err := libp2p.New(*ctx, options...)
	if err != nil {
		return nil, err
	}

	if *cli_args.Options.Debug {
		ui.LogTrace("Adresses hôte: ")
		for _, addr := range host.Addrs() {
			ui.LogTrace("  ", addr)
		}
	}

	ps, err := pubsub.NewGossipSub(*ctx, host)
	if err != nil {
		log.Panicln("Impossible de créer le GossipSub: ", err)
	}

	// Multicast DNS
	ui.LogInfo("Initialisation du Multicast DNS")
	err = dns.SetupMulticastDNS(ctx, &host, cli_args, ui)
	if err != nil {
		return nil, err
	}

	return &chat.NodeAPI{Host: &host, PubSub: ps}, nil
}

func main() {
	cli_args := cli.GetCliArgs()
	ui := ui.CreateUIFromCLI(cli_args)

	go ui.Run()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	api, err := create_api(&ctx, cli_args, ui)
	if err != nil {
		log.Panicln("Impossible de créer l'hôte: ", err)
	}

	channel, err := chat.NewChannel(ctx, api.PubSub, cli_args.Options.Channel, (*api.Host).ID())
	if err != nil {
		log.Panicln("Impossible de créer le salon: ", err)
	}

	state := chat.NewState(ctx, api, cli_args)

	ui.Log(fmt.Sprintf("Hello %s, utilise la commande '/help' pour afficher les commandes disponibles", *cli_args.Options.Nick))
	ui.Start(state, channel)

	<-ui.OkChan
}
