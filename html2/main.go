package main

import (
	"encoding/base64"
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	// 设置路由
	http.HandleFunc("/encrypt", encryptHandler)
	http.HandleFunc("/decrypt", decryptHandler)
	http.HandleFunc("/index", indexHandler)

	// 启动服务
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("html2/index.html"))
	tmpl.Execute(w, nil)
}

func encryptHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// 加载加密页面模板
		tmpl := template.Must(template.ParseFiles("html2/encrypt.html"))
		tmpl.Execute(w, nil)
	} else if r.Method == "POST" {
		// 解析表单数据
		r.ParseForm()
		data := r.Form.Get("data")

		fmt.Println("----------------------", data)
		// 执行加密操作
		encryptedData := base64.StdEncoding.EncodeToString([]byte(data))

		// 返回加密后的数据给前端
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"result": "` + encryptedData + `"}`))
	}
}

func decryptHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// 加载解密页面模板
		tmpl := template.Must(template.ParseFiles("html2/encrypt.html"))
		tmpl.Execute(w, nil)
	} else if r.Method == "POST" {
		// 解析表单数据
		r.ParseForm()
		encodedData := r.Form.Get("data")

		// 执行解密操作
		decodedData, err := base64.StdEncoding.DecodeString(encodedData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// 返回解密后的数据给前端
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"result": "` + string(decodedData) + `"}`))
	}
}
