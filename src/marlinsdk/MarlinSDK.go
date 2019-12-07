package marlinsdk

/*
#cgo CFLAGS: -g -Wall
#include <stdlib.h>
#include "/Users/amolagrawal/ProjectMarlin/marlinSDN/multicastsdk.cpp/include/marlin/multicast/MulticastClientWrapper.h"
#cgo LDFLAGS: /Users/amolagrawal/ProjectMarlin/marlinSDN/multicastsdk.cpp/build/libcwrappermulticast.a
#cgo LDFLAGS: /usr/local/lib/libmarlin-net.a
#cgo LDFLAGS: /usr/local/lib/libgtestd.a
#cgo LDFLAGS: /usr/local/lib/libuv_a.a
#cgo LDFLAGS: -lstdc++

extern void did_recv_message_cgo (
	MarlinMulticastClientWrapper_t* mc_w,
	const char* message,
	uint64_t message_length,
	const char* channel,
	uint64_t channel_length,
	uint64_t message_id);

static void inline set_delegate(MarlinMulticastClientDelegate_t* mc_d, type_did_recv_message_func f) {
	mc_d->did_recv_message = f;
}

*/
import "C"

import (
 //	"fmt"
	"unsafe"
)

//export go_did_recv_message
func go_did_recv_message(
	mc_w *C.MarlinMulticastClientWrapper_t,
	message  *C.char,
	message_length C.int,
	channel  *C.char,
	channel_length C.int,
	message_id C.int) {
}


func CreateMulticastClientDelgate() (*C.MarlinMulticastClientDelegate_t) {
	mc_d := C.marlin_multicast_create_multicastclientdelegate()
	C.marlin_multicast_set_did_recv_message(mc_d, C.type_did_recv_message_func(C.did_recv_message_cgo))

	return mc_d
}

/*
takes multicast delegate callback functions
takes beacon address, public address, discovery address
creates_multicast_client_delagate with arguments
take null/empty if those are not provided
*/
func CreateMulticastClient(
	beacon_addr string,
	discovery_addr string,
	pubsub_addr string) (*C.MarlinMulticastClientWrapper_t) {

	cstr_beacon_addr := C.CString(beacon_addr)
	defer C.free(unsafe.Pointer(cstr_beacon_addr))

	cstr_discovery_addr := C.CString(discovery_addr)
	defer C.free(unsafe.Pointer(cstr_discovery_addr))

	cstr_pubsub_addr := C.CString(pubsub_addr)
	defer C.free(unsafe.Pointer(cstr_pubsub_addr))

	mc_d := CreateMulticastClientDelgate()

	mc_w := C.marlin_multicast_create_multicastclientwrapper(cstr_beacon_addr, cstr_discovery_addr, cstr_pubsub_addr)
	C.marlin_multicast_setDelegate(mc_w, mc_d)

	return mc_w
}

func RunEventLoop() {
	C.marlin_multicast_run_event_loop()
}
