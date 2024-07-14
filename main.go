package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/aeilang/full/frontend"
)

func main() {
	assets, _ := frontend.Assets()

	fs := http.FileServer(http.FS(assets))

	r := http.NewServeMux()
	r.Handle("GET /", SPA(fs))

	r.HandleFunc("GET /api", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("anoter"))
	})
	r.HandleFunc("GET /dashboard", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Dashborad"))
	})

	server := http.Server{
		Addr:    ":8888",
		Handler: r,
	}

	log.Println("listen to 8888")
	server.ListenAndServe()
}

func SPA(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.Path

		if !strings.Contains(url, ".") {
			r.URL.Path = "/"
		}

		next.ServeHTTP(w, r)
	})
}
