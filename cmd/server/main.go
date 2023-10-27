package main

import (
	"log"
	"net"
	"strconv"
	"udp-multiplayer-go/internal/app"
	"udp-multiplayer-go/internal/data"
	"udp-multiplayer-go/proto/pb"

	"google.golang.org/protobuf/proto"
)

func main() {
	port := 8080

	udpAddr, err := net.ResolveUDPAddr("udp", ":"+strconv.Itoa(port))
	if err != nil {
		log.Fatalln(err)
	}

	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		log.Fatalln(err)
	}

	go app.ApiStart(3000)
	go data.B2Init()
	log.Printf("UDP server started on :%v", port)

	for {
		var buf [1024]byte
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
		app.HandleUdp(&message, conn, addr)
	}
}
