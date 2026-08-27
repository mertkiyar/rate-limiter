package main

import (
	"fmt"
	"net"
	"net/http"
)

// map[ip]count format
var visitors = make(map[string]int)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		// r.RemoteAddr includes ip and port number
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			http.Error(w, "Server Error", http.StatusInternalServerError)
			return
		}

		// increse visit count with each login
		visitors[ip]++

		if visitors[ip] > 3 {
			http.Error(w, "Limit exceeded", http.StatusTooManyRequests)
			return
		}

		fmt.Fprintln(w, "API Rate Limiter working good!")
		fmt.Fprintf(w, "Your IP Address with port: %s\n", r.RemoteAddr) // with port
		fmt.Fprintf(w, "Your IP Address: %s\n", ip)                     // without port
		fmt.Fprintf(w, "Visit Count: %d", visitors[ip])                 // visit count
	})

	fmt.Println("API Rate Limiter working on 8080 port")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("API Rate Limiter not started", err)
	}
}
