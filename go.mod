module gostudy

go 1.12

require (
	github.com/davecgh/go-spew v1.1.1
	github.com/garyburd/redigo v1.6.0
	github.com/go-sql-driver/mysql v1.4.1
	github.com/gogo/protobuf v1.2.1
	github.com/golang/protobuf v1.3.2
	github.com/google/go-querystring v1.0.0 // indirect
	github.com/google/uuid v1.1.1
	github.com/grpc-ecosystem/go-grpc-middleware v1.0.0
	github.com/ipfs/go-datastore v0.0.5
	github.com/ipfs/go-log v0.0.1
	github.com/joelnb/sofa v0.1.1
	github.com/libp2p/go-libp2p v0.1.0
	github.com/libp2p/go-libp2p-circuit v0.1.0
	github.com/libp2p/go-libp2p-core v0.0.1
	github.com/libp2p/go-libp2p-crypto v0.1.0
	github.com/libp2p/go-libp2p-discovery v0.1.0
	github.com/libp2p/go-libp2p-host v0.1.0
	github.com/libp2p/go-libp2p-kad-dht v0.1.0
	github.com/libp2p/go-libp2p-net v0.1.0
	github.com/libp2p/go-libp2p-peer v0.2.0
	github.com/libp2p/go-libp2p-peerstore v0.1.0
	github.com/libp2p/go-libp2p-pubsub v0.1.0 // indirect
	github.com/libp2p/go-libp2p-swarm v0.1.0
	github.com/multiformats/go-multiaddr v0.0.4
	github.com/multiformats/go-multiaddr-net v0.0.1
	github.com/op/go-logging v0.0.0-20160315200505-970db520ece7
	github.com/pkg/errors v0.8.1
	github.com/stretchr/objx v0.1.1
	github.com/stretchr/testify v1.3.0
	github.com/sykesm/zap-logfmt v0.0.1
	github.com/syndtr/goleveldb v1.0.0
	github.com/whyrusleeping/go-logging v0.0.0-20170515211332-0457bb6b88fc
	github.com/whyrusleeping/go-smux-multiplex v3.0.16+incompatible // indirect
	github.com/whyrusleeping/go-smux-multistream v2.0.2+incompatible // indirect
	github.com/whyrusleeping/go-smux-yamux v2.0.9+incompatible // indirect
	github.com/whyrusleeping/yamux v1.2.0 // indirect
	github.com/willf/bitset v1.1.10
	github.com/yanyiwu/gojieba v1.1.0
	go.uber.org/multierr v1.3.0 // indirect
	go.uber.org/zap v1.9.1
	golang.org/x/crypto v0.0.0-20191011191535-87dc89f01550
	golang.org/x/net v0.0.0-20191021144547-ec77196f6094
	golang.org/x/text v0.3.2
	google.golang.org/grpc v1.24.0
	labix.org/v2/mgo v0.0.0-20140701140051-000000000287
)

replace (
	golang.org/x/net => github.com/golang/net v0.0.0-20191021144547-ec77196f6094
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190215142949-d0b11bdaac8a
	golang.org/x/text => github.com/golang/text v0.3.2
	golang.org/x/tools => github.com/golang/tools v0.0.0-20190524140312-2c0ae7006135
	golang.org/x/xerrors => github.com/golang/xerrors v0.0.0-20190717185122-a985d3407aa7
	google.golang.org/grpc => github.com/grpc/grpc-go v1.24.0
)
