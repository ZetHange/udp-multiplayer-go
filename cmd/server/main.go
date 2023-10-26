package main

import (
	"github.com/golang/protobuf/proto"
	"log"
	"net"
	"udp-multiplayer-go/internal/app"
	"udp-multiplayer-go/internal/data"
	"udp-multiplayer-go/proto/pb"
)

func main() {
	udpAddr, err := net.ResolveUDPAddr("udp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}

	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		log.Fatalln(err)
	}

	go app.ApiStart()
	go data.B2Init()
	log.Println("API started on :3000")
	log.Println("UDP server started on :8080")

	for {
		var buf [512]byte
		n, addr, err := conn.ReadFromUDP(buf[0:])
		if err != nil {
			log.Println(err)
			continue
		}

		var message pb.Request
		err = proto.Unmarshal(buf[0:n], &message)
		if err != nil {
			log.Println(err)
			continue
		}
		app.Handle(&message, conn, addr)
	}
}
