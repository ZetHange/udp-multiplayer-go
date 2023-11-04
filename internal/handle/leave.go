package handle

import (
	"log"
	"net"
	"udp-multiplayer-go/internal/data"
	"udp-multiplayer-go/internal/utils"
	"udp-multiplayer-go/proto/pb"

	"google.golang.org/protobuf/proto"
)

func Leave(req *pb.Request, conn *net.UDPConn, addr *net.UDPAddr) {
	uuid := req.Leave.GetUuid()

	user, ok := data.Leave(uuid)

	utils.Oko.Lock()
	delete(utils.Oko.Users, uuid)
	utils.Oko.Unlock()

	data, err := proto.Marshal(&pb.Response{
		Leave: &pb.Response_LEAVE{
			Ok: ok,
		},
	})
	if err != nil {
		log.Println(err)
	}

	conn.Write(data)

	log.Printf("[DISCONNECT](id: %s) User with login: %s disconnected from map", uuid, user.Login)
}
