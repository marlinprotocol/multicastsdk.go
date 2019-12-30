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

type ClientDelegate interface {
	DidRecvMessage(client *Client, message string, message_length uint64, channel string, channel_length uint64, message_id uint64)
	DidSubscribe(client *Client, channel string, channel_length uint64)
	DidUnsubscribe(client *Client, channel string, channel_length uint64)
}

type Client struct {
	cClient *C.MarlinMulticastClient_t
	cClientDelegate *C.MarlinMulticastClientDelegate_t
	delegate ClientDelegate
}

func NewClient(
	staticSk *C.uchar,
	beacon_addr string,
	discovery_addr string,
	pubsub_addr string,
	delegate ClientDelegate,
) *Client {
	client := new(Client)
	client.delegate = delegate

	cstr_beacon_addr := C.CString(beacon_addr)
	defer C.free(unsafe.Pointer(cstr_beacon_addr))

	cstr_discovery_addr := C.CString(discovery_addr)
	defer C.free(unsafe.Pointer(cstr_discovery_addr))

	cstr_pubsub_addr := C.CString(pubsub_addr)
	defer C.free(unsafe.Pointer(cstr_pubsub_addr))

	client.cClientDelegate = newClientDelgate()

	client.cClient = C.marlin_multicast_client_create(
		staticSk,
		cstr_beacon_addr,
		cstr_discovery_addr,
		cstr_pubsub_addr,
	)
	C.marlin_multicast_client_set_delegate(client.cClient, client.cClientDelegate)

	mapstore.Save(unsafe.Pointer(client.cClient), client)

	return client
}

func (client *Client) Destroy() {
	C.marlin_multicast_clientdelegate_destroy(client.cClientDelegate)
	C.marlin_multicast_client_destroy(client.cClient)
}

func (client *Client) SendMessageOnChannel(
	channel string,
	message string,
) {
	cstr_channel := C.CString(channel)
	defer C.free(unsafe.Pointer(cstr_channel))

	cstr_message := C.CString(message)
	defer C.free(unsafe.Pointer(cstr_message))

	C.marlin_multicast_client_send_message_on_channel(
		client.cClient,
		cstr_channel,
		C.uint64_t(len(channel)),
		cstr_message,
		C.uint64_t(len(message)),
	)
}

func newClientDelgate() *C.MarlinMulticastClientDelegate_t {
	delegate := C.marlin_multicast_clientdelegate_create()

	C.marlin_multicast_clientdelegate_set_did_recv_message(
		delegate,
		C.did_recv_message_func(C.did_recv_message_cgo),
	)
	C.marlin_multicast_clientdelegate_set_did_subscribe(
		delegate,
		C.did_subscribe_func(C.did_subscribe_cgo),
	)
	C.marlin_multicast_clientdelegate_set_did_unsubscribe(
		delegate,
		C.did_unsubscribe_func(C.did_unsubscribe_cgo),
	)

	return delegate
}
