package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	data2 "udp-multiplayer-go/internal/data"
)

func ApiStart(port int) {
	http.HandleFunc("/map", func(w http.ResponseWriter, r *http.Request) {
		data, err := json.Marshal(&data2.MapList.MapList)
		if err != nil {
			log.Println(err)
			return
		}
		w.Header().Add("Content-Type", "application/json")
		w.Write(data)
	})
	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		data, err := json.Marshal(&data2.UserList.Users)
		if err != nil {
			log.Println(err)
			return
		}
		w.Header().Add("Content-Type", "application/json")
		w.Write(data)
	})

	log.Printf("API started on :%v\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
	if err != nil {
		log.Fatalln(err)
		return
	}
}
