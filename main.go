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
	node, err := libp2p.New(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Node Id= %v, Node Addrs= %v", node.ID(), node.Addrs())

	err = node.Close()
	if err != nil {
		panic(err)
	}

}
