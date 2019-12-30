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

extern void did_recv_message_cgo (
	MarlinMulticastClient_t* client,
	const char* message,
	uint64_t message_length,
	const char* channel,
	uint64_t channel_length,
	uint64_t message_id);

extern void did_subscribe_cgo (
	MarlinMulticastClient_t* client,
	const char* channel,
	uint64_t channel_length);

extern void did_unsubscribe_cgo (
	MarlinMulticastClient_t* client,
	const char* channel,
	uint64_t channel_length);
*/
import "C"

import (
	"gitlab.com/marlinprotocol/multicastsdk.go/internal/mapstore"
	"unsafe"
)

//export go_did_recv_message
func go_did_recv_message(
	cClient *C.MarlinMulticastClient_t,
	message *C.char,
	message_length C.int,
	channel *C.char,
	channel_length C.int,
	message_id C.int,
) {
	client := mapstore.Restore(unsafe.Pointer(cClient)).(*Client)
	cbIface := client.delegate
	cbIface.DidRecvMessage(
		client,
		C.GoString(message),
		uint64(message_length),
		C.GoString(channel),
		uint64(channel_length),
		uint64(message_id))
}

//export go_did_subscribe
func go_did_subscribe(
	cClient *C.MarlinMulticastClient_t,
	channel *C.char,
	channel_length C.int,
) {
	client := mapstore.Restore(unsafe.Pointer(cClient)).(*Client)
	cbIface := client.delegate
	cbIface.DidSubscribe(
		client,
		C.GoString(channel),
		uint64(channel_length))
}

//export go_did_unsubscribe
func go_did_unsubscribe(
	cClient *C.MarlinMulticastClient_t,
	channel *C.char,
	channel_length C.int,
) {
	client := mapstore.Restore(unsafe.Pointer(cClient)).(*Client)
	cbIface := client.delegate
	cbIface.DidUnsubscribe(
		client,
		C.GoString(channel),
		uint64(channel_length))
}
