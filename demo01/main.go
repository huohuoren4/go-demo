package main

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
	"log"
)

// 打印请求信息
func PrintRequest(req *gorequest.SuperAgent) {
	fmt.Printf("请求信息:\n请求的url: %v\n请求方法: %v\n请求头: %v\nContent-Type: %v\nraw数据: %v\n",
		req.Url, req.Method, req.Header, req.ForceType, req.RawString)
}

// 打印响应信息
func PrintResponse(resp gorequest.Response, body string) {
	fmt.Printf("响应信息:\n响应的状态码: %v\n响应头: %v\n响应体: %v\n", resp.StatusCode, resp.Header, body)
}

// 获取用户列表
func GetUserList() {
	req := gorequest.New().Get("http://localhost:8080/hello?name=123&password=12121")
	PrintRequest(req)
	resp, body, err := req.End()
	if err != nil {
		log.Fatal(err)
	}
	PrintResponse(resp, body)
}

// 获取用户详细信息
func GetUserInfo() {
	req := gorequest.New().Get("http://localhost:8080/hello/1")
	PrintRequest(req)
	resp, body, err := req.End()
	if err != nil {
		log.Fatal(err)
	}
	PrintResponse(resp, body)
}

// 创建用户
func PostUser() {
	reqBody := `{ "name": "123", "password": "123456" }`
	req := gorequest.New().Post("http://localhost:8080/hello").Type("json").Send(reqBody)
	PrintRequest(req)
	resp, body, err := req.End()
	if err != nil {
		log.Fatal(err)
	}
	PrintResponse(resp, body)
}

// 上传用户头像
func PostUserIcon() {
	filepath := "C:/Users/Administrator/Desktop/Snipaste_2023-01-11_23-21-01.png"
	// 文件名不能是file
	req := gorequest.New().Post("http://localhost:8080/hello/1/upload").
		Type("multipart").
		Set("User-Agent", "golang").
		SendFile(filepath, "", "image_file")
	PrintRequest(req)
	resp, body, err01 := req.End()
	if err01 != nil {
		log.Fatal(err01)
	}
	PrintResponse(resp, body)
}

// 修改用户信息
func PutUserInfo() {
	reqBody := `{ "name": "123" }`
	req := gorequest.New().Put("http://localhost:8080/hello/1").Type("json").Send(reqBody)
	resp, body, err := req.End()
	PrintRequest(req)
	if err != nil {
		log.Fatal(err)
	}
	PrintResponse(resp, body)
}

func DeleteUser() {
	req := gorequest.New().Delete("http://localhost:8080/hello/1")
	PrintRequest(req)
	resp, body, err := req.End()
	if err != nil {
		log.Fatal(err)
	}
	PrintResponse(resp, body)
}

func main() {
	GetUserList()
	fmt.Println("#######################")
	GetUserInfo()
	fmt.Println("#######################")
	PostUser()
	fmt.Println("#######################")
	PostUserIcon()
	fmt.Println("#######################")
	PutUserInfo()
	fmt.Println("#######################")
	DeleteUser()
}
