package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	port = flag.String("port", "8080", "http port default 80")
)

func main() {
	flag.Parse()
	log.Printf("starting, will listen on port %s", *port)

	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler).Methods("GET", "POST")
	http.Handle("/", r)

	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%s", *port), r))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s: %s", r.Method, r.URL.String())
	_, _ = w.Write([]byte("hello"))
}
