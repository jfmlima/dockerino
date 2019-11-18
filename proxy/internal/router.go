package internal

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type Router struct {}

func (r Router) Routes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})

	static := Handler{StaticPath: "static", IndexPath: "index.html"}
	router.PathPrefix("/").Handler(static)

	return router
}