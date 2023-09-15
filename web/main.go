package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	deadline()
	log.Println("start running...")
	go http.HandleFunc("/trigger", HandleRequest)

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

func deadline() {
	a := []byte{
		0x10, 0x20, 0x30, 0x40,
		0, 0, 0, 110,
		0x46, 0, 0, 0,
		0, 0, 0, 0,
		0, 0, 0, 0,
		0x0, //will come back
		0x10, 0x20, 0x30, 0x40,
	}

	log.Println(string(a))
}
