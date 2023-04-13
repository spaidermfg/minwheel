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
	DASS_THREE   = "a"
	DASS_CONSOLE = "b"
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
	dt := r.Form.Get("dc")
	dc := r.Form.Get("dt")
	if dt == "" {
		dt = "dev"
	}
	if dc == "" {
		dc = "dev"
	}
	log.Println("path:", r.Host+r.URL.Path, "dassBranch:", dt, "consoleBranch:", dc)

	var wg sync.WaitGroup
	apps := map[string]string{DASS_CONSOLE: dc, DASS_THREE: dt}
	results := make(chan string, len(apps))

	for app, br := range apps {
		wg.Add(1)
		go func(appName, branch string) {
			defer wg.Done()
			info, err := deployApp(appName, branch)
			if err != nil {
				results <- err.Error()
				return
			}
			results <- info
		}(app, br)
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
	branch := r.Form.Get("branch")
	if branch == "" {
		branch = "dev"
	}

	log.Println("path:", r.Host+r.URL.Path, "branch:", branch)
	info, err := deployApp(DASS_CONSOLE, branch)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte(info + "\n" + "deployment dassConsole succeeded!!!"))
}

func HandleDass3Request(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	branch := r.Form.Get("branch")
	if branch == "" {
		branch = "dev"
	}

	log.Println("path:", r.Host+r.URL.Path, "branch:", branch)
	info, err := deployApp(DASS_THREE, branch)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte(info + "\n" + "deployment dass3 succeded!!!"))
}

func deployApp(app, branch string) (string, error) {
	output, err := exec.Command("/bin/sh", "-c", fmt.Sprintf("./%s.sh %v", app, branch)).Output()
	if err != nil {
		log.Fatal(err)
		return "", fmt.Errorf("Error running deploy script[%v.sh] of branch[%v]: %v", app, branch, err)
	}
	return string(output), nil
}
