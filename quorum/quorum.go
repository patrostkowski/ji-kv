package quorum

import "net"

type Entity struct {
	name string
	addr string
}

type QuorumList struct {
	name string
	addr net.Addr
}
