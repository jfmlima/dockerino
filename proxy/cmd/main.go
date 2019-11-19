package main

import (
	"log"
	"net/http"
	"os"
	"proxy/internal"
	"time"
)

func main() {
	port := os.Getenv("PORT")
	println(port)
	srv := &http.Server{
		Handler: internal.Router{}.Routes(),
		Addr: ":" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}