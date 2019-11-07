package main

import (
	"context"
	"fmt"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p/p2p/protocol/ping"
	"github.com/multiformats/go-multiaddr"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// create a background context
	ctx := context.Background()
	node, err := libp2p.New(ctx, libp2p.ListenAddrStrings("/ip4/127.0.0.1/tcp/0"), libp2p.Ping(false))
	if err != nil {
		panic(err)
	}
	// config our own ping protocol
	pingService := &ping.PingService{Host: node}
	node.SetStreamHandler(ping.ID, pingService.PingHandler)

	// print the node info
	peerInfo := &peer.AddrInfo{ID: node.ID(), Addrs: node.Addrs()}
	multiaddrs, err := peer.AddrInfoToP2pAddrs(peerInfo)
	if err != nil {
		panic(err)
	}
	fmt.Println("Listen address: ", multiaddrs[0])

	// if remote peer have been passed on the command line, connect to it and send it 5 ping messages, otherwise wait for a signal to stop
	if len(os.Args) > 1 {
		addr, err := multiaddr.NewMultiaddr(os.Args[1])
		if err != nil {
			panic(err)
		}
		addrInfo, err := peer.AddrInfoFromP2pAddr(addr)
		if err != nil {
			panic(err)
		}

		if err = node.Connect(ctx, *addrInfo); err != nil {
			panic(err)
		}

		fmt.Println("sending 5 ping messages to ", addr)
		ch := pingService.Ping(ctx, ping.ID)

		for i := 0; i < 5; i++ {
			res := <-ch
			fmt.Println("pinged ", addr, "in ", res.RTT)
		}
	} else {
		// wait for a SIGINT or SIGTERM signal
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
		<-ch
		fmt.Println("Received signal, shutting down...")
	}

	// shutdown the node
	if err = node.Close(); err != nil {
		panic(err)
	}
}
