package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"udp-multiplayer-go/proto/pb"

	"google.golang.org/protobuf/proto"
)

var (
	host        = flag.String("host", "127.0.0.1:8080", "IP и порт UDP сервера")
	typeconnect = flag.String("type", "default", "Тип клиента: default, leave")
	uuid        = flag.String("uuid", "", "uuid для leave")
	// mapId = flag.Int64("mapId", 1, "ID карты, целочисленный")
	// user  = flag.String("user", "username", "Юзернейм")
	// dx    = flag.Float64("dx", 0.1, "Скорость по иксу")
	// dy    = flag.Float64("dy", 0.1, "Скорость по игрику")
)

func main() {
	flag.Parse()

	conn, err := net.Dial("udp", *host)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	if *typeconnect == "leave" {
		req, _ := proto.Marshal(&pb.Request{
			Type: pb.RequestType_LEAVE,
			Leave: &pb.Request_LEAVE{
				Uuid: *uuid,
			},
		})

		_, err = conn.Write(req)
		if err != nil {
			log.Println(err)
		}

		return
	}

	fmt.Println("Подключение к серверу... Успешно...")
	fmt.Println("=== Авторизация ===")
	fmt.Print("Логин: ")

	var login string
	fmt.Scan(&login)

	fmt.Print("Целочисленный ID карты: ")
	var mapId int
	fmt.Scan(&mapId)

	req, _ := proto.Marshal(&pb.Request{
		Type: pb.RequestType_JOIN,
		Get:  &pb.Request_GET{},
		Join: &pb.Request_JOIN{
			Login: login,
			MapId: int64(mapId),
		},
	})

	_, err = conn.Write(req)
	fmt.Println("Попытка подключения к серверу, отправлено", len(req), "байт")
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

	uuid := data.Join.GetUuid()

	content, err := json.Marshal(&data)
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println(string(content))

	for {
		fmt.Print("Press ENTER для отправки данных")
		bufio.NewReader(os.Stdin).ReadString('\n')

		fmt.Print("\nВведите новый X: ")
		var x, y, dx, dy float64
		fmt.Scan(&x)

		fmt.Print("Введите новый Y: ")
		fmt.Scan(&y)

		fmt.Print("Введите скорость по X: ")
		fmt.Scan(&dx)

		fmt.Print("Введите скорость по Y: ")
		fmt.Scan(&dy)
		fmt.Println()

		req, _ := proto.Marshal(&pb.Request{
			Type: pb.RequestType_GET,
			Get: &pb.Request_GET{
				Uuid:   uuid,
				Health: 200,
				X:      x,
				Y:      y,
				Dx:     dx,
				Dy:     dy,
			},
		})

		_, err = conn.Write(req)
		if err != nil {
			log.Println(err)
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
		json, _ := json.Marshal(&data)
		fmt.Println(string(json), ":: bytes:", n)
	}
}
