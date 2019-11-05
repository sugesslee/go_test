package main

import (
	"context"
	"crypto/rand"
	"fmt"

	libp2p "github.com/libp2p/go-libp2p"
	crypto "github.com/libp2p/go-libp2p-crypto"
)

/*
 *
 * <pre>
 *  Version         Date            Author          Description
 * ---------------------------------------------------------------------------------------
 *  1.0.0           2019/11/04     redli        -
 * </pre>
 * @author redli
 * @version 1.0.0 2019/11/4 9:07 AM
 * @date 2019/11/4 9:07 AM
 * @since 1.0.0
 */
func main() {
	// The context governs the lifetime of the libp2p node
	// context上下文控制libp2p节点的生存周期
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// To construct a simple host with all the default settings, just use `New`
	// 构造一个具有所有默认设置的简单host，只需使用“New”方法
	h, err := libp2p.New(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello World, my hosts ID is %s\n", h.ID())
	// Set your own keypair
	// 配置自身的密钥对
	priv, _, err := crypto.GenerateEd25519Key(rand.Reader)
	if err != nil {
		panic(err)
	}

	h2, err := libp2p.New(ctx,
		// Use your own created keypair
		// 使用自身创建的密钥对
		libp2p.Identity(priv),

		// Set your own listen address
		// The config takes an array of addresses, specify as many as you want.
		// 配置自身的监听地址
		// 该配置采用地址数组的形式，想指定多少就可指定多少
		libp2p.ListenAddrStrings("/ip4/0.0.0.0/tcp/9000"),
	)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello World, my second hosts ID is %s\n", h2.ID())
}