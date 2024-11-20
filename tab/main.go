package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlePprof)
	fmt.Println("========================> Running")
	if err := http.ListenAndServe(":8181", nil); err != nil {
		log.Fatal("------------>", err)
	}
}

func handlePprof(w http.ResponseWriter, r *http.Request) {
	must := template.Must(template.ParseFiles("index.html"))
	if err := must.Execute(w, nil); err != nil {
		log.Fatal("------------>", err)
	}
}
