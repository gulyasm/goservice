package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/enbritely/heartbeat-golang"
	"github.com/gorilla/mux"
)

type M struct {
	handler http.Handler
}

func (m M) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	start := time.Now()
	m.handler.ServeHTTP(rw, r)
	log.Printf("%s served in %s\n", r.URL, time.Since(start))
}

func NewM(h http.Handler) http.Handler {
	return M{h}
}

func IndexHandler(rw http.ResponseWriter, r *http.Request) {
	msg := struct {
		Message string `json:"message"`
	}{
		"Hello world!",
	}
	json.NewEncoder(rw).Encode(msg)
}

func main() {
	address := os.Getenv("LISTENING_ADDRESS")
	hAddress := os.Getenv("HEARTBEAT_ADDRESS")
	log.SetPrefix("[service] ")
	go heartbeat.RunHeartbeatService(hAddress)
	r := mux.NewRouter()
	r.HandleFunc("/", IndexHandler)
	log.Println("Service request at: " + address)
	log.Println(http.ListenAndServe(address, NewM(r)))
}
