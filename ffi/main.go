package main

// #cgo CFLAGS: -g -Wall -I${SRCDIR}/clang
// #cgo LDFLAGS: -L${SRCDIR}/clang -lffi
// #include "add.c"
// #include "sub.c"

import "C"

func main() {
	println("Expression arithmétique: 10 + 32 = ", C.add(10, 32))
	println("Expression arithmétique: 80 - 77 = ", C.sub(80, 77))
}
