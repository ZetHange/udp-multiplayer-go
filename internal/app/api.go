package app

import (
	"encoding/json"
	"log"
	"net/http"
	data2 "udp-multiplayer-go/internal/data"
)

func ApiStart() {
	http.HandleFunc("/maps", func(w http.ResponseWriter, r *http.Request) {
		data, err := json.Marshal(data2.Maps)
		if err != nil {
			log.Println(err)
		}
		w.Header().Add("Content-Type", "application/json")
		w.Write(data)
	})

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Println(err)
		return
	}
}
