package handle

import (
	"google.golang.org/protobuf/proto"
	"log"
	"net"
	"udp-multiplayer-go/internal/data"
	"udp-multiplayer-go/proto/pb"
)

func HandleGet(req *pb.Request, conn *net.UDPConn, addr *net.UDPAddr) {
	get := req.Get

	data.MapList.UpdateUser(get.Uuid, get.X, get.Y, get.Dx, get.Dy)

	mapId := data.MapList.GetMapIdByUserId(get.Uuid)

	users := data.MapList.ToProto(mapId)

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
