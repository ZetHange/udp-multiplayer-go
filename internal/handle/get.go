package handle

import (
	"net"
	"udp-multiplayer-go/proto/pb"
)

func HandleGet(req *pb.Request, conn *net.UDPConn, addr *net.UDPAddr) {
	// get := req.Get
}
