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
	"sync"
)

const (
	DASS_THREE   = "deploy_dass3"
	DASS_CONSOLE = "deploy_dassc"
)

func main() {
	log.Println("start running...")
	http.HandleFunc("/deploy", HandleFunc)
	http.HandleFunc("/deploy/dass3", HandleDass3Request)
	http.HandleFunc("/deploy/console", HandleConsoleRequest)
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func HandleFunc(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	log.Println("path:", r.Host+r.URL.Path)

	var wg sync.WaitGroup
	apps := []string{DASS_CONSOLE, DASS_THREE}
	results := make(chan string, len(apps))
	for _, app := range apps {
		wg.Add(1)
		go func(appName string) {
			defer wg.Done()
			info, err := deployApp(appName)
			if err != nil {
				results <- err.Error()
				return
			}
			results <- info
		}(app)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for info := range results {
		fmt.Fprintf(w, "%s\n", info)
	}
}

func HandleConsoleRequest(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	log.Println("path:", r.Host+r.URL.Path)
	info, err := deployApp(DASS_CONSOLE)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte(info + "\n" + "deployment dassConsole succeeded!!!"))
}

func HandleDass3Request(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	log.Println("path:", r.Host+r.URL.Path)
	info, err := deployApp(DASS_THREE)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte(info + "\n" + "deployment dass3 succeded!!!"))
}

func deployApp(app string) (string, error) {
	output, err := exec.Command("/bin/sh", "-c", fmt.Sprintf("./%s.sh", app)).Output()
	if err != nil {
		log.Fatal(err)
		return "", fmt.Errorf("Error running deploy script[%v.sh]: %v", app, err)
	}
	return string(output), nil
}
