package handle

import (
	"github.com/golang/protobuf/proto"
	"github.com/google/uuid"
	"log"
	"net"
	"udp-multiplayer-go/internal/data"
	"udp-multiplayer-go/proto/pb"
)

func Join(req *pb.Request, conn *net.UDPConn, addr *net.UDPAddr) {
	u, _ := uuid.NewRandom()

	user := &data.User{
		Id:     u.String(),
		Login:  req.Join.Login,
		Health: int(req.Join.Health),
		X:      req.Join.StartX,
		Y:      req.Join.StartY,
		Dx:     0.5,
		Dy:     0.6,
	}

	data.JoinUser(int(req.Join.MapId), user)

	data, err := proto.Marshal(&pb.Response{
		Join: &pb.Response_JOIN{
			Ok:   true,
			Uuid: u.String(),
		},
	})
	if err != nil {
		log.Println(err)
		return
	}

	_, err = conn.WriteToUDP(data, addr)
	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("[id: %s] user <%s> connect to map %v", u.String(), req.Join.Login, req.Join.MapId)
}
