package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

var pasteData string

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 如果是POST请求，更新粘贴内容
		if r.Method == "POST" {
			data, err := io.ReadAll(r.Body)
			if err != nil {
				http.Error(w, "无法读取请求体", http.StatusInternalServerError)
				return
			}
			pasteData = strings.TrimSpace(string(data))
			log.Println("[接收到新的消息]:\n", pasteData)
			fmt.Fprintf(w, "粘贴成功")
			return
		}

		// 如果是GET请求，返回粘贴内容页面
		if r.Method == "GET" {
			html := `
				<!DOCTYPE html>
				<html>
				<head>
					<meta charset="utf-8">
					<title>ClipBoard</title>
					<style>
						body {
           					 display: flex;
           					 flex-direction: column;
           					 justify-content: center;
           					 align-items: center;
           					 height: 100vh;
            					margin: 0;
     					   }
						h1, textarea, button {
       					     margin: 10px; /* 调整元素之间的间距 */
						}
					</style>
				</head>
				<body>
 					<h1>Clip Board</h1>
 					<textarea id="paste" rows="20" cols="80"></textarea>
					<button style="height: 50px; width: 400px" onclick="paste()">粘贴</button>
					<script>
						function paste() {
							var data = document.getElementById("paste").value;
							fetch("/", {
								method: 'POST',
								body: data
							}).then(function(response) {
								alert("粘贴成功");x
							});
						}
					</script>
				</body>
				</html>
			`
			w.Header().Set("Content-Type", "text/html")
			fmt.Fprintf(w, html)
			return
		}
	})

	// 获取粘贴板内容
	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, pasteData)
	})

	// 启动HTTP服务器
	fmt.Println("服务器运行在 http://localhost:80")
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Println("启动HTTP服务器失败:", err)
	}
}
