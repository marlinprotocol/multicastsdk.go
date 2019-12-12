package main

import (
 "fmt"
 "marlinsdk"
)

type cbInterface struct {
}

func (cbIface cbInterface) Did_recv_message(mc_w marlinsdk.ExportedTypeMulticastWrapper, message string, message_length uint64,  channel string, channel_length uint64, message_id uint64) {
	fmt.Println("Go did recv message")
}

func (cbIface cbInterface) Did_subscribe(mc_w marlinsdk.ExportedTypeMulticastWrapper, channel string, channel_length uint64) {
	fmt.Println("Go did subscribe channel: " + channel)
	marlinsdk.SendMessageOnChannel(mc_w, channel, "Go message: subscribed")
}

func (cbIface cbInterface) Did_unsubscribe(mc_w marlinsdk.ExportedTypeMulticastWrapper, channel string, channel_length uint64) {
}

func main() {
	beacon_addr := "127.0.0.1:9002"
	discovery_addr := "127.0.0.1:7002"
	pubsub_addr := "127.0.0.1:7000"

	// m2 := CreateMulticastClient(beacon_addr, discovery_addr, pubsub_addr);
	marlinsdk.CreateMulticastClient(cbInterface{}, beacon_addr, discovery_addr, pubsub_addr);

	beacon_addr = "127.0.0.1:9002"
	discovery_addr = "127.0.0.1:8002"
	pubsub_addr = "127.0.0.1:8000"

	// m1 := CreateMulticastClient(beacon_addr, discovery_addr, pubsub_addr);
	marlinsdk.CreateMulticastClient(cbInterface{}, beacon_addr, discovery_addr, pubsub_addr);

	marlinsdk.RunEventLoop();

	fmt.Println(beacon_addr)
}
