package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"strings"
)

// 模拟登陆
type CloudLoginReq struct {
	Account   string `json:"account"`
	Password  string `json:"password"`
	LoginType string `json:"loginType"`
}

type CloudLoginRes struct {
	Data    string `json:"data"`
	Success bool   `json:"success"`
}

type GetAppList struct {
	Platform string `json:"platform"`
	Limit    string `json:"limit"`
	Page     string `json:"page"`
}

const (
	LOGIN_URL = "http://192.168.20.42/api/authCenter/login"
	APP_URL   = "http://192.168.20.42/api/terminalCenter/app/cloud/driver/page"
	ACCOUNT   = "sysadmin"
	PASSWORD  = "sysadmin"
	LOGINTYPE = "PASSWORD"
)

// 请求类型，请求地址，鉴权key，请求体
func CreateHttpRequest[T any](method, url, key string, body T) ([]byte, error) {
	b, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	reqBody := strings.NewReader(string(b))

	if method == http.MethodGet {
		if reflect.TypeOf(body).Kind() != reflect.Ptr || reflect.TypeOf(body).Elem().Kind() != reflect.Struct {
			panic("params must be a pointer to struct")
		}

		v := reflect.ValueOf(body).Elem()
		for i := 0; i < v.NumField(); i++ {
			field := v.Field(i)
			fmt.Printf("%v - %v\n", v.Field(i).Interface(), field.Type().Name())
		}
	}

	req, err := http.NewRequestWithContext(context.Background(), method, url, reqBody)
	if err != nil {
		return nil, err
	}

	if key != "" {
		req.Header.Add("Authorization", "Bearer "+key)
	}
	req.Header.Add("content-type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return respBody, err
}

// 响应结果，解析结构体
func ParseResponse[T any](respBody []byte, result T) (T, error) {
	err := json.Unmarshal(respBody, result)
	if err != nil {
		return *new(T), err
	}
	return result, err
}

// 请求类型，请求地址，鉴权key，请求体, 解析结构体
func RequestParse[T, V any](method, url, autherKey string, reqBody T, resBody V) (V, error) {
	respBody, err := CreateHttpRequest(method, url, autherKey, &reqBody)
	if err != nil {
		panic(err)
	}

	result, err := ParseResponse(respBody, &resBody)
	if err != nil {
		panic(err)
	}
	return *result, err
}

// 登陆云平台，返回cookie
func LoginCloudPlat() (*CloudLoginRes, error) {
	data := &CloudLoginReq{ACCOUNT, PASSWORD, LOGINTYPE}
	resp, err := RequestParse(http.MethodPost, LOGIN_URL, "", data, &CloudLoginRes{})
	if err != nil {
		return &CloudLoginRes{}, err
	}

	return resp, err
}

func JudgeAutherExpired() {

}

func main() {
	cookie, err := LoginCloudPlat()
	if err != nil {
		panic(err)
	}
	log.Println(cookie.Data, cookie.Success)

	data := &GetAppList{"linux", "1", "10"}
	b, err := CreateHttpRequest(http.MethodGet, APP_URL, cookie.Data, data)
	if err != nil {
		panic(err)
	}
	log.Println(string(b))
	testencode()
}

func testencode() {
	baseURL := "https://httpbin.org/get"
	params := url.Values{}
	params.Set("name", "value")
	params.Set("age", "14")
	reqURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())
	fmt.Println(reqURL)
}
