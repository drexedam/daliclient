package main

import (
	"github.com/drexedam/daliclient"
	"fmt"
)

func main() {
	usbdali := &daliclient.Usbdali{}

	// Open a connection
	if err := usbdali.Connect("localhost"); err != nil {
		panic(err)
	}

	// Don't forget to close the connection
	defer usbdali.Close()

	// Send a broadcast to dim down
	if err := usbdali.Send(daliclient.MakeBroadcastCmd(0x00, daliclient.CmdDimDown)); err != nil {
		panic(err)
	}

	// Wait for response
	var resp []byte
	var err error
	if resp, err = usbdali.Receive(); err != nil {
		panic(err)
	}

	fmt.Println(resp)
}
