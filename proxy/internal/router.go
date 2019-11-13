package internal

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type Router struct {}

func (r Router) Routes() *mux.Router {
	m := mux.NewRouter()

	router := mux.NewRouter()

	router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})

	static := Handler{StaticPath: "build", IndexPath: "index.html"}
	router.PathPrefix("/").Handler(static)

	return m
}