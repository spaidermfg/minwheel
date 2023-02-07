package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Println("start running...")
	http.HandleFunc("/trigger", HandleRequest)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	log.Println(r.URL.Path)
	log.Printf("processing of requests from %v...\n", r.URL.Path)
	fmt.Fprintf(w, "Hello World!")
}
