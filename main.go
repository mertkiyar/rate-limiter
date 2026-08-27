package main

import (
	"fmt"
	"net"
	"net/http"
	"time"
)

type visitor struct {
	count    int
	lastSeen time.Time
}

// map[ip]visitor format
var visitors = make(map[string]*visitor)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		// r.RemoteAddr includes ip and port number
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			http.Error(w, "Server Error", http.StatusInternalServerError)
			return
		}

		v, exists := visitors[ip]
		if !exists {
			// if it is not exists, create new
			v = &visitor{count: 0, lastSeen: time.Now()}
			visitors[ip] = v
		}

		// every 10 seconds, reset visit count
		if time.Since(v.lastSeen) > 10*time.Second {
			v.count = 0
		}

		// update last seed time and increase visit count with each login
		v.lastSeen = time.Now()
		v.count++

		if v.count > 3 {
			http.Error(w, "Limit exceeded, please wait 10 seconds to reset", http.StatusTooManyRequests)
			return
		}

		fmt.Fprintln(w, "API Rate Limiter working good!")
		fmt.Fprintf(w, "Your IP Address with port: %s\n", r.RemoteAddr) // with port
		fmt.Fprintf(w, "Your IP Address: %s\n", ip)                     // without port
		fmt.Fprintf(w, "Visit Count: %d", v.count)                      // visit count
	})

	fmt.Println("API Rate Limiter working on 8080 port")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("API Rate Limiter not started", err)
	}
}
