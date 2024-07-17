package main

import (
	"fmt"
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
		Addr:    ":80",
		Handler: Limit(r),
	}

	log.Println("listen to 80")
	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err)
	}
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

func Limit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}
