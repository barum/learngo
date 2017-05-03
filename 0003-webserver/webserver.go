package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", myroot)
	http.ListenAndServe("localhost:8080", nil)

}

func myroot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Bahram Web Server")

}
