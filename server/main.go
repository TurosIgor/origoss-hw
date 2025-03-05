package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)

	port := ":8080"
	http.ListenAndServe(port, nil)
}

func handler(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Hello, World!"))
}
