package handle

import (
	"log"
	"net"
	"time"
	"udp-multiplayer-go/internal/data"
	"udp-multiplayer-go/internal/utils"
	"udp-multiplayer-go/proto/pb"

	"github.com/google/uuid"
	"google.golang.org/protobuf/proto"
)

func Join(req *pb.Request, conn *net.UDPConn, addr *net.UDPAddr) {
	u, _ := uuid.NewRandom()

	utils.Oko.Lock()
	utils.Oko.Users[u.String()] = time.Now()
	utils.Oko.Unlock()

	user := &data.User{
		Id:     u.String(),
		Login:  req.Join.Login,
		Health: 200,
		X:      req.Join.StartX,
		Y:      req.Join.StartY,
		Dx:     0,
		Dy:     0,
	}

	data.JoinUser(int(req.Join.MapId), user)

	users := data.MapList.ToPb(int(req.Join.MapId))

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

	log.Printf("[CONNECT](id: %s) User with login: %s joined to map %v", u.String(), req.Join.Login, req.Join.MapId)
}
