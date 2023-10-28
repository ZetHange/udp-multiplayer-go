package handle

import (
	"log"
	"net"
	"udp-multiplayer-go/internal/data"
	"udp-multiplayer-go/proto/pb"

	"github.com/google/uuid"
	"google.golang.org/protobuf/proto"
)

func Join(req *pb.Request, conn *net.UDPConn, addr *net.UDPAddr) {
	u, _ := uuid.NewRandom()

	user := &data.User{
		Id:     u.String(),
		Login:  req.Join.Login,
		Health: 200,
		X:      0,
		Y:      0,
		Dx:     0,
		Dy:     0,
	}

	data.JoinUser(int(req.Join.MapId), user)

	users := data.MapList.ToProto(int(req.Join.MapId))

	data, err := proto.Marshal(&pb.Response{
		Join: &pb.Response_JOIN{
			Ok:   true,
			Uuid: u.String(),
		},
		Get: &pb.Response_GET{
			Users: users,
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
