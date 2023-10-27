package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"udp-multiplayer-go/proto/pb"

	"google.golang.org/protobuf/proto"
)

var (
	host  = flag.String("host", "127.0.0.1:8080", "IP и порт UDP сервера")
	mapId = flag.Int64("mapId", 1, "ID карты, целочисленный")
	user  = flag.String("user", "username", "Юзернейм")
	dx    = flag.Float64("dx", 0.1, "Скорость по иксу")
	dy    = flag.Float64("dy", 0.1, "Скорость по игрику")
)

func main() {
	flag.Parse()

	conn, err := net.Dial("udp", *host)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	req, _ := proto.Marshal(&pb.Request{
		Type: pb.RequestType_JOIN,
		Get:  &pb.Request_GET{},
		Join: &pb.Request_JOIN{
			Login: *user,
			MapId: *mapId,
		},
	})

	_, err = conn.Write(req)
	log.Println("Data sending..., bytes count:", len(req))
	if err != nil {
		log.Panicln(err)
	}

	var buf = make([]byte, 1024)
	n, err := bufio.NewReader(conn).Read(buf)
	if err != nil {
		log.Panicln(err)
	}

	var data pb.Response
	if err = proto.Unmarshal(buf[0:n], &data); err != nil {
		log.Panicln(err)
	}
	content, err := json.Marshal(&data)
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println(string(content), "bytes:", n, "json bytes:", len(content))
}
