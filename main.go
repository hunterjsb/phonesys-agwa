// main.go
package main

import (
	"agwa/handlers"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", handlers.HomeHandler)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
	http.ListenAndServe(":"+getPort(), nil)
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return port
}
