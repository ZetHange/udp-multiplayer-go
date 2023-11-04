package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"udp-multiplayer-go/internal/data"
	"udp-multiplayer-go/internal/utils"
)

func ApiStart(port int) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data, err := json.Marshal(map[string]interface{}{
			"stats": &Metrics,
			"maps":  data.MapList.GetMaps(),
			"users": &data.UserList.Users,
			"oko":   &utils.Oko.Users,
		})
		if err != nil {
			log.Println(err)
			return
		}
		w.Header().Add("Content-Type", "application/json")
		w.Write(data)
	})

	http.HandleFunc("/map", func(w http.ResponseWriter, r *http.Request) {
		data, err := json.Marshal(&data.MapList.MapList)

		if err != nil {
			log.Println(err)
			return
		}
		w.Header().Add("Content-Type", "application/json")
		w.Write(data)
	})
	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		data, err := json.Marshal(&data.UserList.Users)
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
