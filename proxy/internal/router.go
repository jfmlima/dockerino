package internal

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type Router struct {}

func (r Router) Routes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})

	router.HandleFunc("/api/whoami", func(w http.ResponseWriter, r *http.Request) {
		apiUrl, err := url.Parse("http://be:8080")
		if err != nil {
			println(err)
		}

		r.Header.Set("X-Forwarded-Host", r.Header.Get("Host"))

		r.URL.Host = apiUrl.Host
		r.URL.Scheme = apiUrl.Scheme
		r.Host = apiUrl.Host

		httputil.NewSingleHostReverseProxy(apiUrl).ServeHTTP(w, r)
	})

	static := Handler{StaticPath: "static", IndexPath: "index.html"}
	router.PathPrefix("/").Handler(static)

	return router
}