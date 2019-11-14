package main

import (
	"log"
	"net/http"
	"proxy/internal"
	"time"
)

func main() {
	srv := &http.Server{
		Handler: internal.Router{}.Routes(),
		Addr: ":8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}