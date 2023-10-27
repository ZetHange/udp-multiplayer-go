package main

import (
	"bufio"
	"flag"
	"log"
	"net"
	"udp-multiplayer-go/proto/pb"

	"google.golang.org/protobuf/proto"
)

var (
	user = flag.String("user", "username", "Юзернейм")
	dx   = flag.Float64("dx", 0.1, "Скорость по иксу")
	dy   = flag.Float64("dy", 0.1, "Скорость по игрику")
)

func main() {
	flag.Parse()

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
			Login:  *user,
			Health: 200,
			StartX: 0,
			StartY: 0,
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
