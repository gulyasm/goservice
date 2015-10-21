package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/enbritely/heartbeat-golang"
	"github.com/gorilla/mux"
)

func IndexHandler(rw http.ResponseWriter, r *http.Request) {
	msg := struct {
		Message string
	}{
		"Hello world",
	}
	json.NewEncoder(rw).Encode(msg)
}
func main() {
	address := os.Getenv("LISTENING_ADDRESS")
	hAddress := os.Getenv("HEARTBEAT_ADDRESS")
	go heartbeat.RunHeartbeatService(hAddress)
	r := mux.NewRouter()
	r.HandleFunc("/", IndexHandler)
	log.Println(http.ListenAndServe(address, r))
}
