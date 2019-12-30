package main

import (
	"gitlab.com/marlinprotocol/multicastsdk.go"
)


type Delegate struct {}

func (delegate *Delegate) DidRecvMessage(
	client *multicastsdk.Client,
	message string,
	message_length uint64,
	channel string,
	channel_length uint64,
	message_id uint64,
) {

}

func (delegate *Delegate) DidSubscribe(
	client *multicastsdk.Client,
	channel string,
	channel_length uint64,
) {

}

func (delegate *Delegate) DidUnsubscribe(
	client *multicastsdk.Client,
	channel string,
	channel_length uint64,
) {

}


func main() {
	keypair := multicastsdk.NewKeypair()
	delegate := Delegate{}
	client := multicastsdk.NewClient(
		&keypair.StaticSk[0],
		"127.0.0.1:9002",
		"127.0.0.1:8002",
		"127.0.0.1:8000",
		&delegate,
	)
	defer client.Destroy()

	multicastsdk.RunEventLoop()
}
