package multicastsdk

/*
#cgo CFLAGS: -Wall
#include <stdlib.h>
#include "./internal/csdk/MarlinMulticastClient.h"
#cgo LDFLAGS: ./internal/csdk/libmulticastsdk_c.a
#cgo LDFLAGS: ./internal/csdk/libmarlin-net.a
#cgo LDFLAGS: ./internal/csdk/libsodium.a
#cgo LDFLAGS: ./internal/csdk/libuv_a.a
#cgo LDFLAGS: -lstdc++ -lm
*/
import "C"

func RunEventLoop() {
	C.marlin_multicast_run_event_loop()
}


type Keypair struct {
	StaticSk [32]C.uint8_t
	StaticPk [32]C.uint8_t
}

func NewKeypair() *Keypair {
	keypair := new(Keypair)

	C.marlin_multicast_create_keypair(&keypair.StaticPk[0], &keypair.StaticSk[0])

	return keypair
}
