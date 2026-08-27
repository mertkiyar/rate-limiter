package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "API Rate Limiter working good!")
	})

	fmt.Println("API Rate Limiter working on 8080 port")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("API Rate Limiter not started", err)
	}
}
