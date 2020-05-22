package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	server := http.NewServeMux()

	server.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			fmt.Fprintf(w, "Hello, %q", "GET")
		case "POST":
			fmt.Fprintf(w, "Hello, %q", "POST")
		}
	})

	log.Println("start server at 8080")
	log.Fatal(http.ListenAndServe(":8080", server))
}
