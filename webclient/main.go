package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
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

type GetAppInfoReq struct {
	Platform string `json:"platform"`
	Name     string `json:"name"`
	Version  string `json:"version"`
}

const (
	LOGIN_URL    = "http://192.168.0.1/login"
	APP_URL      = "http://192.168.0.1/driver/page"
	APP_INFO_URL = "http://192.168.0.1/appconfig"
	ACCOUNT      = "sysadmin"
	PASSWORD     = "sysadmin"
	LOGINTYPE    = "PASSWORD"
)

// 请求类型，请求地址，鉴权key，请求体
func CreateHttpRequest[T any](method, reqUrl, autherKey string, body T) ([]byte, error) {
	var reqBody io.Reader
	if method == http.MethodGet {
		t := reflect.TypeOf(body)
		if t.Kind() != reflect.Ptr || t.Elem().Kind() != reflect.Struct {
			panic("params must be a pointer to struct")
		}

		params := url.Values{}
		v := reflect.ValueOf(body).Elem()
		for i := 0; i < v.NumField(); i++ {
			params.Set(strings.ToLower(v.Type().Field(i).Name), strings.ToLower(v.Field(i).String()))
		}
		reqUrl = fmt.Sprintf("%s?%s", reqUrl, params.Encode())
	} else {
		b, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		reqBody = strings.NewReader(string(b))
	}

	req, err := http.NewRequestWithContext(context.Background(), method, reqUrl, reqBody)
	if err != nil {
		return nil, err
	}

	req.Header.Add("content-type", "application/json")
	if autherKey != "" {
		req.Header.Add("Authorization", "Bearer "+autherKey)
	}

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

// 校验鉴权key是否过期
func JudgeAutherExpired() {

}

func main() {
	//登陆
	cookie, err := LoginCloudPlat()
	if err != nil {
		panic(err)
	}
	log.Println("cookie: ", cookie.Data, cookie.Success)

	//获取可安装应用
	appData := &GetAppList{Platform: "linux", Limit: "10", Page: "2"}
	b, err := CreateHttpRequest(http.MethodGet, APP_URL, cookie.Data, appData)
	if err != nil {
		panic(err)
	}
	log.Println(string(b))
	fmt.Println("")

	//获取应用详细信息
	infoData := &GetAppInfoReq{Platform: "linux", Name: "GB_Modbus_Poll_TCP", Version: "v2.0.0-beta5"}
	b2, err := CreateHttpRequest(http.MethodGet, APP_INFO_URL, cookie.Data, infoData)
	if err != nil {
		panic(err)
	}
	log.Println(string(b2))
}
