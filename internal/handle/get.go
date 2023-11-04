package handle

import (
	"log"
	"net"
	"time"
	"udp-multiplayer-go/internal/data"
	"udp-multiplayer-go/internal/utils"
	"udp-multiplayer-go/proto/pb"

	"google.golang.org/protobuf/proto"
)

func HandleGet(req *pb.Request, conn *net.UDPConn, addr *net.UDPAddr) {
	get := req.Get

	utils.Oko.Lock()
	utils.Oko.Users[get.Uuid] = time.Now()
	utils.Oko.Unlock()

	ok := data.UpdateUser(get.Uuid, get.Dx, get.Dy)
	if !ok {
		return
	}

	mapId := data.MapList.GetMapIdByUserId(get.Uuid)

	users := data.MapList.ToPb(mapId)

	data, err := proto.Marshal(&pb.Response{
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
}
