package main

import (
	"fmt"
	"net"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		// r.RemoteAddr includes ip and port number
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			http.Error(w, "Server Error", http.StatusInternalServerError)
			return
		}
		fmt.Fprintln(w, "API Rate Limiter working good!")
		fmt.Fprintf(w, "Your IP Address with port: %s\n", r.RemoteAddr) // with port
		fmt.Fprintf(w, "Your IP Address: %s", ip)                       // without port
	})

	fmt.Println("API Rate Limiter working on 8080 port")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("API Rate Limiter not started", err)
	}
}
