package main

import (
	"context"
	"fmt"
)
import "github.com/libp2p/go-libp2p"

func main() {
	// create a background context (i.e. one that never cancels)
	ctx := context.Background()

	// start a libp2p node with default settings
	node, err := libp2p.New(ctx,
		libp2p.ListenAddrStrings("/ip4/127.0.0.1/tcp/2000"),
	)
	if err != nil {
		panic(err)
	}

	// print the node's listening addresses
	fmt.Printf("Node Id= %v, Node Addrs= %v", node.ID(), node.Addrs())

	// shut the node down
	err = node.Close()
	if err != nil {
		panic(err)
	}

}
