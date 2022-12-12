package main

import (
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HelloDocker(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Halo Docker! \n")
}

func serveHTTP(addr string) {
	router := mux.NewRouter()
	router.HandleFunc("/hello", HelloDocker)
	err := http.ListenAndServe(addr, router)
	if err != nil {
		log.Fatalf("error listen to port %v \n", addr)
	}
	log.Printf("Listening to %v", addr)
}

func main() {
	serveHTTP(":8080")
}
