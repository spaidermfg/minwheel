package main

import (
	"encoding/base64"
	"html/template"
	"log"
	"net/http"
)

const tpl = `
<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8">
		<title>{{.Title}}</title>
	</head>
	<body>
		{{range .Items}}<div>{{ . }}</div>{{else}}<div><strong>no rows</strong></div>{{end}}
	</body>
</html>`

type Base64Web struct {
}

func main() {
	b := Base64Web{}
	b.web()
}

func (b *Base64Web) web() {
	http.HandleFunc("/", handlerFunc)
	http.HandleFunc("/hello", indexHandler)
	http.HandleFunc("/demo", handlerDemo)
	http.ListenAndServe(":8080", nil)
}

func handlerFunc(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("is-control", "yes")
	files, _ := template.New("").ParseFiles("text.html")
	files.Execute(w, nil)
	w.WriteHeader(http.StatusOK)
}

func handlerDemo(w http.ResponseWriter, r *http.Request) {
	check := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}
	t, err := template.New("webpage").Parse(tpl)
	check(err)

	data := struct {
		Title string
		Items []string
	}{
		Title: "My page",
		Items: []string{
			"My photos",
			"My blog",
		},
	}

	err = t.Execute(w, data)
	check(err)

	noItems := struct {
		Title string
		Items []string
	}{
		Title: "My another page",
		Items: []string{},
	}

	err = t.Execute(w, noItems)
	check(err)
}

// 数据结构，用于模板渲染
type PageData struct {
	Input    string
	Output   string
	IsEncode bool
}

// 处理主页请求
func indexHandler(w http.ResponseWriter, r *http.Request) {
	data := PageData{}

	// 如果有表单提交
	if r.Method == "POST" {
		// 解析表单参数
		r.ParseForm()

		// 获取表单参数
		input := r.Form.Get("input")
		isEncode := r.Form.Get("action") == "encode"

		// 根据用户选择进行加密或解密
		if isEncode {
			data.Output = base64.StdEncoding.EncodeToString([]byte(input))
		} else {
			decoded, err := base64.StdEncoding.DecodeString(input)
			if err != nil {
				data.Output = "Invalid Base64 input"
			} else {
				data.Output = string(decoded)
			}
		}

		// 设置页面数据
		data.Input = input
		data.IsEncode = isEncode
	}

	// 渲染 HTML 模板并返回给客户端
	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.Execute(w, data)
}
