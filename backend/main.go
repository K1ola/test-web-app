package main

import (
	"fmt"
	"log"
	"net/http"
)

func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")

		responseHeader := w.Header()
		responseHeader.Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH")
		responseHeader.Set("Access-Control-Allow-Credentials", "true")
		responseHeader.Set("Access-Control-Allow-Headers", "Content-Type")

		//TODO add white list origins
		responseHeader.Set("Access-Control-Allow-Origin", origin)

		if r.Method == "OPTIONS" {
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"resp": "Hello world from GO!"}`))
	})

	wrappedMux := CORSMiddleware(mux)

	fmt.Println("Started server on port 8888")
	log.Fatal(http.ListenAndServe(":8888", wrappedMux))
}
