package main

import (
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
	"strings"
)

// Clipboard 结构体用于存储粘贴板数据
type Clipboard struct {
	Data string `json:"data"`
}

// clipboard 实际存储粘贴板数据的变量
var clipboard Clipboard

// 处理首页的HTTP GET请求
func homeHandler(w http.ResponseWriter, r *http.Request) {
	// 如果请求方法是GET，则显示首页
	if r.Method == "GET" {
		homeTemplate.Execute(w, nil)
	} else if r.Method == "POST" {
		log.Println("-----", r.Method, r.Body)
		all, err2 := io.ReadAll(r.Body)
		if err2 != nil {
			log.Fatal(err2)
		}
		defer r.Body.Close()

		log.Println("======", string(all))
		//body := r.GetBody

		after := strings.Split(string(all), "=")[1]
		log.Println("======", after)
		if after != string(all) {
			log.Println("HHHHHHHHHHHH", after)
			w.WriteHeader(http.StatusOK)
			return
		}

		// 如果请求方法是POST，则更新粘贴板数据
		err := json.NewDecoder(r.Body).Decode(&clipboard)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

// 处理获取粘贴板数据的HTTP GET请求
func getClipboardHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(clipboard)
	}
}

// 首页模板
var homeTemplate = template.Must(template.New("home").Parse(`
<!DOCTYPE html>
<html>
<head>
    <title>Web Clipboard</title>
</head>
<body>
    <form action="/" method="post">
        <textarea name="data" rows="10" cols="50">{{.data}}</textarea>
        <input type="submit" value="Submit">
    </form>
</body>
</html>`))

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/clipboard", getClipboardHandler)
	http.ListenAndServe(":8080", nil)
}
