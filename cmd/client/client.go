package main

import (
	"bufio"
	"github.com/golang/protobuf/proto"
	"log"
	"net"
	"udp-multiplayer-go/proto/pb"
)

func main() {
	udpAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8080")

	if err != nil {
		log.Panicln(err)
	}

	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		log.Panicln(err)
	}

	req, _ := proto.Marshal(&pb.Request{
		Type: pb.RequestType_JOIN,
		Get:  &pb.Request_GET{},
		Join: &pb.Request_JOIN{
			Login:  "zethange",
			Health: 200,
			StartX: 2,
			StartY: 3,
			MapId:  1,
		},
	})

	_, err = conn.Write(req)
	log.Println("send...", len(req))
	if err != nil {
		log.Panicln(err)
	}

	_, err = bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Println(err)
		return
	}
}
