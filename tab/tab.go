package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
)

var (
	deviceMap = map[string]string{
		"40.60": "http://192.168.40.60:8899/debug/pprof/", // 设备 3 的 pprof 地址
		"40.61": "http://192.168.40.61:8899/debug/pprof/", // 设备 1 的 pprof 地址
		"40.62": "http://192.168.40.62:8899/debug/pprof/", // 设备 2 的 pprof 地址
		"40.63": "http://192.168.40.63:8899/debug/pprof/", // 设备 3 的 pprof 地址
	}
	mu sync.RWMutex // 保护设备列表
)

func proxyHandler(w http.ResponseWriter, r *http.Request) {
	device := r.URL.Query().Get("device")
	if device == "" {
		http.Error(w, "请指定设备参数 ?device=<设备名>", http.StatusBadRequest)
		return
	}

	fmt.Println("-----------1>", device)
	mu.RLock()
	target, exists := deviceMap[device]
	mu.RUnlock()

	if !exists {
		http.Error(w, "设备不存在", http.StatusNotFound)
		return
	}

	fmt.Println("-----------2>", target)
	proxyURL, err := url.Parse(target)
	if err != nil {
		http.Error(w, "目标地址解析失败", http.StatusInternalServerError)
		return
	}

	fmt.Println("-----------3>", proxyURL.String())
	proxy := httputil.NewSingleHostReverseProxy(proxyURL)
	proxy.ModifyResponse = func(resp *http.Response) error {
		log.Printf("请求已转发到目标服务器: %s", resp.Request.URL)
		return err
	}

	fmt.Println(r.URL)
	fmt.Println(r.RequestURI)
	fmt.Println(r.Host)

	u, err := url.Parse("")
	if err != nil {
		log.Fatal(err)
	}

	r.URL = u
	proxy.ServeHTTP(w, r)
	fmt.Println(r.Body)
}

func devicesHandler(w http.ResponseWriter, r *http.Request) {
	mu.RLock()
	defer mu.RUnlock()

	ds := make([]*Device, 0)
	for name, addr := range deviceMap {
		ds = append(ds, &Device{
			Url:  addr,
			Name: name,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(ds); err != nil {
		log.Fatal(err)
	}
}

type Device struct {
	Url  string `json:"url"`
	Name string `json:"name"`
}

func main() {
	// 提供静态文件服务
	staticDir := "./static" // 前端代码存放的路径
	http.Handle("/", http.FileServer(http.Dir(staticDir)))

	// API 路由
	http.HandleFunc("/proxy", proxyHandler)
	http.HandleFunc("/devices", devicesHandler)

	port := ":8080"
	log.Printf("服务已启动，监听端口 %s，前端路径: %v", port, staticDir)
	log.Fatal(http.ListenAndServe(port, nil))
}
