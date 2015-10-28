package main

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
	"os"
	"text/template"
	"time"

	"github.com/enbritely/heartbeat-golang"
	"github.com/gorilla/mux"
)

var queries = []Query{}

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

type ErrorMessage struct {
	Error string
}

func ipHandler(rw http.ResponseWriter, r *http.Request) {
	domain := r.URL.Query().Get("domain")
	if domain == "" {
		msg := ErrorMessage{"Empty domain parameter"}
		json.NewEncoder(rw).Encode(msg)
		return
	}
	ips, err := net.LookupIP(domain)
	if err != nil {
		msg := ErrorMessage{"Invalid domain address."}
		json.NewEncoder(rw).Encode(msg)
		return
	}
	msg := IPMessage{
		ips,
	}
	json.NewEncoder(rw).Encode(msg)
	queries = append(queries, Query{domain, ips})
}

type Query struct {
	Domain string
	IPs    []net.IP
}

type PageData struct {
	Build   string
	Queries []Query
}

func uiHandler(rw http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Panic("Error occured parsing the template", err)
	}
	page := PageData{
		Build:   heartbeat.CommitHash,
		Queries: queries,
	}
	if err = tmpl.Execute(rw, page); err != nil {
		log.Panic("Failed to write template", err)
	}

}
func createBaseHandler() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/service/ip", ipHandler)
	r.HandleFunc("/", uiHandler)
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
