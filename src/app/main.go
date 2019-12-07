package main

import (
 "fmt"
 "marlinsdk"
)

func main() {
	beacon_addr := "127.0.0.1:9002"
	discovery_addr := "127.0.0.1:7002"
	pubsub_addr := "127.0.0.1:7000"

	// m2 := CreateMulticastClient(beacon_addr, discovery_addr, pubsub_addr);
	marlinsdk.CreateMulticastClient(beacon_addr, discovery_addr, pubsub_addr);

	beacon_addr = "127.0.0.1:9002"
	discovery_addr = "127.0.0.1:8002"
	pubsub_addr = "127.0.0.1:8000"

	// m1 := CreateMulticastClient(beacon_addr, discovery_addr, pubsub_addr);
	marlinsdk.CreateMulticastClient(beacon_addr, discovery_addr, pubsub_addr);

	marlinsdk.RunEventLoop();

	fmt.Println(beacon_addr)
}
