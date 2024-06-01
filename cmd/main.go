package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	port := "8000"

	r.HandleFunc("/ping", PingHandler).Methods("GET")

	fmt.Println("Server starting on port:", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

func PingHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong"))
}
