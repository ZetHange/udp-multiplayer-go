package main

import (
	"bufio"
	"flag"
	"log"
	"math/rand"
	"net"
	"time"
	"udp-multiplayer-go/proto/pb"

	"github.com/google/uuid"
	"google.golang.org/protobuf/proto"
)

var (
	host = flag.String("host", "127.0.0.1:8080", "IP и порт UDP сервера")
)

func main() {
	flag.Parse()

	for {
		conn, err := net.Dial("udp", *host)
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()

		id := uuid.New()
		req, _ := proto.Marshal(&pb.Request{
			Type: pb.RequestType_JOIN,
			Get:  &pb.Request_GET{},
			Join: &pb.Request_JOIN{
				Login: id.String(),
				MapId: 1,
			},
		})
		conn.Write(req)

		var buf = make([]byte, 1024)
		n, err := bufio.NewReader(conn).Read(buf)
		if err != nil {
			log.Panicln(err)
		}

		var data pb.Response
		if err = proto.Unmarshal(buf[0:n], &data); err != nil {
			log.Panicln(err)
		}

		uuid := data.Join.GetUuid()

		go func() {
			for {
				req, _ := proto.Marshal(&pb.Request{
					Type: pb.RequestType_GET,
					Get: &pb.Request_GET{
						Uuid:   uuid,
						Health: 200,
						X:      rand.Float64(),
						Y:      rand.Float64(),
						Dx:     rand.Float64(),
						Dy:     rand.Float64(),
					},
				})

				_, err = conn.Write(req)
				if err != nil {
					log.Println(err)
				}

				// time.Sleep(20 * time.Millisecond)
			}
		}()

		time.Sleep(500 * time.Millisecond)
	}
}
