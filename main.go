package main

import (
	"blog/internal/config"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	config.LoadConfig()

	r := mux.NewRouter()
	port := config.AppConfig.Port

	r.HandleFunc("/ping", PingHandler).Methods("GET")

	fmt.Println("Server starting on port:", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

func PingHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong"))
}
