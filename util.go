package multicastsdk

/*
#include <stdlib.h>
#include "./internal/csdk/MarlinMulticastClient.h"
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
