package handle

import (
	"log"
	"net"
	"udp-multiplayer-go/internal/data"
	"udp-multiplayer-go/proto/pb"

	"google.golang.org/protobuf/proto"
)

func Leave(req *pb.Request, conn *net.UDPConn, addr *net.UDPAddr) {
	uuid := req.Leave.GetUuid()

	ok := data.Leave(uuid)

	data, err := proto.Marshal(&pb.Response{
		Leave: &pb.Response_LEAVE{
			Ok: ok,
		},
	})
	if err != nil {
		log.Println(err)
	}

	conn.Write(data)
}
