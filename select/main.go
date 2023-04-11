// package main

// //基于select的多路复用

// func main() {

// }

package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

func main() {
	log.Println("start running...")
	http.HandleFunc("/deploy/console", HandleConsoleRequest)
	http.HandleFunc("/deploy/dass3", HandleDass3Request)
	http.HandleFunc("/deploy", HandleFunc)
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func HandleFunc(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	log.Println("path: ", r.Host, r.URL.Path)
	go HandleConsoleRequest(w, r)
	go HandleDass3Request(w, r)
}

func HandleConsoleRequest(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	log.Println("path: ", r.Host, r.URL.Path)
	output, err := exec.Command("/bin/sh", "-c", "./dass-console-build.sh").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "%v", string(output))
	fmt.Fprintln(w, "deployment dassConsole succeeded!!!")
}

func HandleDass3Request(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	log.Println("path: ", r.Host, r.URL.Path)
	output, err := exec.Command("/bin/sh", "-c", "./deploy_dass3.sh").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "%v", string(output))
	fmt.Fprintln(w, "deployment dass3 succeded!!!")
}
