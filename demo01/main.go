package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// 查询用户列表
func Get()  {
	resp, err := http.Get("http://localhost:8080/hello?name=123&password=12121")
	if err!=nil {
		log.Fatalln(err)
	}
	defer func() {
		if err :=resp.Body.Close(); err!=nil{
			log.Fatalln(err)
		}
	}()
	data, err1 := ioutil.ReadAll(resp.Body)
	if err1!=nil{
		log.Fatalln(err)
	}
	fmt.Println(string(data))
	for key, val := range resp.Header{
		fmt.Println(key, val)
	}
}

// 查询用户信息
func GetInfo()  {
	resp, err := http.Get("http://localhost:8080/hello/1")
	if err!=nil {
		log.Fatalln(err)
	}
	defer func() {
		if err :=resp.Body.Close(); err!=nil{
			log.Fatalln(err)
		}
	}()
	data, err1 := ioutil.ReadAll(resp.Body)
	if err1!=nil{
		log.Fatalln(err)
	}
	fmt.Println(string(data))
	for key, val := range resp.Header{
		fmt.Println(key, val)
	}
}


// 创建用户
func Post() {
	reqBody :=`{ "name": "123", "password": "123456" }`
	resp, err := http.Post("http://localhost:8080/hello", "application/json", strings.NewReader(reqBody))
	if err!=nil {
		log.Fatalln(err)
	}
	defer func() {
		if err :=resp.Body.Close(); err!=nil{
			log.Fatalln(err)
		}
	}()
	data, err1 := ioutil.ReadAll(resp.Body)
	if err1!=nil{
		log.Fatalln(err)
	}
	fmt.Println(string(data))
	for key, val := range resp.Header{
		fmt.Println(key, val)
	}
}

// 上传用户头像
func PostFile(){
	client := &http.Client{}
	reqBody:=`{ "file": "C:/Users/Administrator/Desktop/" }`
	req, err := http.NewRequest("POST", "http://localhost:8080/hello/1/upload", strings.NewReader(reqBody))
	if err!=nil{
		log.Fatalln(err)
	}
	req.Header.Add("Content-Type", "multipart/form-data")
	req.Header.Set()
	resp, err01 := client.Do(req)
	if err01!=nil{
		log.Fatalln(err)
	}
	data, err02 := ioutil.ReadAll(resp.Body)
	if err02!=nil{
		log.Fatalln(err02)
	}
	fmt.Println(string(data))
	for key, val := range resp.Header {
		fmt.Println(key, val)
	}
}



func main() {
	PostFile()


}