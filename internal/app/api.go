package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	data2 "udp-multiplayer-go/internal/data"
)

func ApiStart(port int) {
	http.HandleFunc("/maps", func(w http.ResponseWriter, r *http.Request) {
		data, err := json.Marshal(data2.Maps.Maps)
		if err != nil {
			log.Println(err)
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
