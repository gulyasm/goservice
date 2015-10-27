package main

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/enbritely/heartbeat-golang"
	"github.com/gorilla/mux"
)

const (
	EnvHeartbeatAddress = "HEARTBEAT_ADDRESS"
	EnvListeningAddress = "LISTENING_ADDRESS"
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

type IPMessage struct {
	IPs []net.IP
}

func getIPs(domain string) []net.IP {
	ips, err := net.LookupIP(domain)
	if err != nil {
		log.Println("Failed to resolve: " + err.Error())
	}
	return ips
}

func IPHandler(rw http.ResponseWriter, r *http.Request) {
	msg := IPMessage{
		getIPs("cnn.com"),
	}
	json.NewEncoder(rw).Encode(msg)
}

func createBaseHandler() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/service/ip", IPHandler)
	return NewM(r)
}

func main() {
	log.SetPrefix("[service] ")

	hAddress := os.Getenv(EnvHeartbeatAddress)
	go heartbeat.RunHeartbeatService(hAddress)

	address := os.Getenv(EnvListeningAddress)
	log.Println("Service request at: " + address)
	log.Println(http.ListenAndServe(address, createBaseHandler()))
}
