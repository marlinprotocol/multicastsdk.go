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

extern void did_subscribe_cgo (
	MarlinMulticastClientWrapper_t* mc_w,
	const char* channel,
	uint64_t channel_length);

extern void did_unsubscribe_cgo (
	MarlinMulticastClientWrapper_t* mc_w,
	const char* channel,
	uint64_t channel_length);
*/
import "C"

import (
 //	"fmt"
	"unsafe"
	"mapstore"
)

type ExportedTypeMulticastWrapper *C.MarlinMulticastClientWrapper_t

type marlinSdkCallbackIface interface {
  Did_recv_message(mc_w ExportedTypeMulticastWrapper, message string, message_length uint64,  channel string, channel_length uint64, message_id uint64)
  Did_subscribe(mc_w ExportedTypeMulticastWrapper, channel string, channel_length uint64)
  Did_unsubscribe(mc_w ExportedTypeMulticastWrapper, channel string, channel_length uint64)
}

//export go_did_recv_message
func go_did_recv_message(
	mc_w *C.MarlinMulticastClientWrapper_t,
	message  *C.char,
	message_length C.int,
	channel  *C.char,
	channel_length C.int,
	message_id C.int) {

	cbIface := mapstore.Restore(unsafe.Pointer(mc_w)).(marlinSdkCallbackIface);
	cbIface.Did_recv_message(
		mc_w,
		C.GoString(message),
		uint64(message_length),
		C.GoString(channel),
		uint64(channel_length),
		uint64(message_id))
}

//export go_did_subscribe
func go_did_subscribe(
	mc_w *C.MarlinMulticastClientWrapper_t,
	channel  *C.char,
	channel_length C.int) {

	cbIface := mapstore.Restore(unsafe.Pointer(mc_w)).(marlinSdkCallbackIface);
	cbIface.Did_subscribe(
		mc_w,
		C.GoString(channel),
		uint64(channel_length))
}

//export go_did_unsubscribe
func go_did_unsubscribe(
	mc_w *C.MarlinMulticastClientWrapper_t,
	channel  *C.char,
	channel_length C.int) {

	cbIface := mapstore.Restore(unsafe.Pointer(mc_w)).(marlinSdkCallbackIface);
	cbIface.Did_unsubscribe(
		mc_w,
		C.GoString(channel),
		uint64(channel_length))
}

func CreateMulticastClientDelgate() (*C.MarlinMulticastClientDelegate_t) {
	mc_d := C.marlin_multicast_create_multicastclientdelegate()
	C.marlin_multicast_set_did_recv_message(mc_d, C.type_did_recv_message_func(C.did_recv_message_cgo))
	C.marlin_multicast_set_did_subscribe(mc_d, C.type_did_subscribe_func(C.did_subscribe_cgo))
	C.marlin_multicast_set_did_unsubscribe(mc_d, C.type_did_unsubscribe_func(C.did_unsubscribe_cgo))

	return mc_d
}

/*
takes multicast delegate callback functions
takes beacon address, public address, discovery address
creates_multicast_client_delagate with arguments
take null/empty if those are not provided
*/
func CreateMulticastClient(
	cbIface marlinSdkCallbackIface,
	beacon_addr string,
	discovery_addr string,
	pubsub_addr string) (ExportedTypeMulticastWrapper) {

	cstr_beacon_addr := C.CString(beacon_addr)
	defer C.free(unsafe.Pointer(cstr_beacon_addr))

	cstr_discovery_addr := C.CString(discovery_addr)
	defer C.free(unsafe.Pointer(cstr_discovery_addr))

	cstr_pubsub_addr := C.CString(pubsub_addr)
	defer C.free(unsafe.Pointer(cstr_pubsub_addr))

	mc_d := CreateMulticastClientDelgate()

	mc_w := C.marlin_multicast_create_multicastclientwrapper(cstr_beacon_addr, cstr_discovery_addr, cstr_pubsub_addr)
	C.marlin_multicast_setDelegate(mc_w, mc_d)

	mapstore.Save(unsafe.Pointer(mc_w), cbIface)

	return mc_w
}

func SendMessageOnChannel(
	mc_w ExportedTypeMulticastWrapper,
	channel string,
	message string) {

	cstr_channel := C.CString(channel)
	defer C.free(unsafe.Pointer(cstr_channel))

	cstr_message := C.CString(message)
	defer C.free(unsafe.Pointer(cstr_message))

	C.marlin_multicast_send_message_on_channel(mc_w, cstr_channel, cstr_message, C.ulonglong(len(message)));
}

func RunEventLoop() {
	C.marlin_multicast_run_event_loop()
}
