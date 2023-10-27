package app

import (
	"net"
	"udp-multiplayer-go/internal/handle"
	"udp-multiplayer-go/proto/pb"
)

func Handle(req *pb.Request, conn *net.UDPConn, addr *net.UDPAddr) {
	switch req.Type {
	case pb.RequestType_GET:
		handle.HandleGet(req, conn, addr)
		return
	case pb.RequestType_JOIN:
		handle.Join(req, conn, addr)
		return
	}
}
