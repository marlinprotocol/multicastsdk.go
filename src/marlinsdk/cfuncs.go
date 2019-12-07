// Wrappers for Go callback functions to be passed into C.

package marlinsdk

/*
#include <stdlib.h>
#include "/Users/amolagrawal/ProjectMarlin/marlinSDN/multicastsdk.cpp/include/marlin/multicast/MulticastClientWrapper.h"

extern void go_did_recv_message(
	MarlinMulticastClientWrapper_t* mc_w,
	const char* message,
	uint64_t message_length,
	const char* channel,
	uint64_t channel_length,
	uint64_t message_id);

void did_recv_message_cgo (
	MarlinMulticastClientWrapper_t* mc_w,
	const char* message,
	uint64_t message_length,
	const char* channel,
	uint64_t channel_length,
	uint64_t message_id) {

	go_did_recv_message(mc_w, message, message_length, channel, channel_length, message_id);
}
*/
import "C"
