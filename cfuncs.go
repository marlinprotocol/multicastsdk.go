// Wrappers for Go callback functions to be passed into C.

package multicastsdk

/*
#include <stdlib.h>
#include "./internal/csdk/MarlinMulticastClient.h"

extern void go_did_recv_message(
	MarlinMulticastClient_t* client,
	const char* message,
	uint64_t message_length,
	const char* channel,
	uint64_t channel_length,
	uint64_t message_id);

extern void go_did_subscribe(
	MarlinMulticastClient_t* client,
	const char* channel,
	uint64_t channel_length
);

extern void go_did_unsubscribe(
	MarlinMulticastClient_t* client,
	const char* channel,
	uint64_t channel_length
);

void did_recv_message_cgo (
	MarlinMulticastClient_t* client,
	const char* message,
	uint64_t message_length,
	const char* channel,
	uint64_t channel_length,
	uint64_t message_id) {

	go_did_recv_message(client, message, message_length, channel, channel_length, message_id);
}

void did_subscribe_cgo (
	MarlinMulticastClient_t* client,
	const char* channel,
	uint64_t channel_length) {

	go_did_subscribe(client, channel, channel_length);
}

extern void did_unsubscribe_cgo (
	MarlinMulticastClient_t* client,
	const char* channel,
	uint64_t channel_length) {

	go_did_unsubscribe(client, channel, channel_length);
}
*/
import "C"
