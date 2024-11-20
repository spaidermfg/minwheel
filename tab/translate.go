package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	// pprof 服务器的目标地址
	target := "http://192.168.40.63:8899/debug/pprof/" // 替换为你的 pprof 地址
	proxyURL, err := url.Parse(target)
	if err != nil {
		log.Fatalf("解析目标地址失败: %v", err)
	}

	// 创建一个反向代理
	proxy := httputil.NewSingleHostReverseProxy(proxyURL)

	// 自定义的请求处理器
	proxy.ModifyResponse = func(resp *http.Response) error {
		// 你可以在这里修改响应，例如添加自定义头部或日志
		log.Printf("请求已转发到目标服务器: %s", resp.Request.URL)
		return nil
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 修改请求的 Host，使其与目标地址匹配
		r.Host = proxyURL.Host
		// 转发请求
		proxy.ServeHTTP(w, r)
	})

	// 启动本地服务
	port := ":8181"
	log.Printf("反向代理服务器已启动，监听端口 %s，转发到 %s", port, target)
	log.Fatal(http.ListenAndServe(port, nil))
}
