package app

import (
	"net"
	"udp-multiplayer-go/internal/handle"
	"udp-multiplayer-go/proto/pb"
)

func HandleUdp(req *pb.Request, conn *net.UDPConn, addr *net.UDPAddr) {
	HandleRequest()

	switch req.Type {
	case pb.RequestType_GET:
		handle.HandleGet(req, conn, addr)
		return
	case pb.RequestType_JOIN:
		handle.Join(req, conn, addr)
		return
	case pb.RequestType_LEAVE:
		handle.Leave(req, conn, addr)
		return
	}
}
