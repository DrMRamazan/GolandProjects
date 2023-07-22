package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", Handler)
	log.Println("Start :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Ramazan\n"))
}
