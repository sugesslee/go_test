package main

import (
	"context"
	"fmt"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p/p2p/protocol/ping"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// create a background context (i.e. one that never cancels)
	ctx := context.Background()

	// start a libp2p node with default setting
	//host, err := libp2p.New(ctx)

	// start a libp2p node that listens on TCP port 2000 on the IPv4
	// loopback interface
	//host, err := libp2p.New(ctx, libp2p.ListenAddrStrings("/ip4/127.0.0.1/tcp/63072"))

	// start a libp2p node that listens on a random local TCP port,
	// but without running the built-in ping protocol
	host, err := libp2p.New(ctx, libp2p.ListenAddrStrings("/ip4/127.0.0.1/tcp/63072"), libp2p.Ping(false))

	if err != nil {
		panic(err)
	}

	//fmt.Println("Listen address: ", host.Addrs())
	//github.com/libp2p/go-libp2p-core/peer.Info
	peerInfo := &peer.AddrInfo{
		ID:    host.ID(),
		Addrs: host.Addrs(),
	}

	multiAddrs, err := peer.AddrInfoToP2pAddrs(peerInfo)
	if err != nil {
		panic(err)
	}
	fmt.Println("libp2p node address: ", multiAddrs)

	// wait for a SIGINT or SIGTERM signal
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	fmt.Println("Received signal, shutting down...")

	// shutdown the node
	if err = host.Close(); err != nil {
		panic(err)
	}

	// configure our own ping protocol
	pingService := &ping.PingService{Host: host}
	host.SetStreamHandler(ping.ID, pingService.PingHandler)
}
