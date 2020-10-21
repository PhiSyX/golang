package chat

import (
	"context"

	cli "github.com/PhiSyX/golang/chatp2p/cli"
)

// --------- //
// Structure //
// --------- //

type ChatState struct {
	Sender *MessageSender

	ctx *context.Context
	api *NodeAPI
}

// -------------- //
// Implémentation //
// -------------- //

func NewState(ctx context.Context, api *NodeAPI, cli_args *cli.CLI) *ChatState {
	state := &ChatState{
		ctx: &ctx,
		api: api,
		Sender: &MessageSender{
			ID:   (*api.Host).ID(),
			Nick: *cli_args.Options.Nick,
		},
	}
	return state
}
