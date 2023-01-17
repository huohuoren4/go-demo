package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

var logger *log.Logger

func init() {
	file := "/var/golang-demo/log/test-app.log"
	err := os.MkdirAll(filepath.Dir(file), 0755)
	if err != nil {
		log.Fatalln(err)
	}
	logFile, err01 := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err01 != nil {
		log.Fatalln(err)
	}
	logger = log.New(logFile, "", log.Ldate|log.Ltime|log.Lmicroseconds)
}

// 增加memory 多少M, 持续一定时间
func TestMem(mem int, duration int64) {
	s := make([][M]byte, mem)
	PrintLog("TestMem 开始增加内存...")
	for i := 0; i < mem; i++ {
		for j := 0; j < M; j++ {
			s[i][j] = 'c'
		}
	}
	PrintLog("TestMem 增加内存结束...")
	time.Sleep(time.Second * time.Duration(duration))
	s[0][0] = 'a'
}

// CPU使用率xx%( 30% )持续一段时间
func TestCpu(cupUsage int64, duration int64) {
	var t, t1 int64
	t = time.Now().Unix() + duration
	for {
		if time.Now().Unix() > t {
			return
		}
		t1 = time.Now().UnixMilli() + cupUsage
		for {
			if time.Now().UnixMilli() > t1 {
				break
			}
		}
		time.Sleep(time.Millisecond * time.Duration(100-cupUsage))
	}
}

// 定义一个map来实现路由转发
type MyHandler struct {
	Mux map[string]func(http.ResponseWriter, *http.Request)
}

func (handler *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//实现路由的转发
	if h, ok := handler.Mux[r.URL.String()]; ok {
		//用这个handler实现路由转发，相应的路由调用相应func
		h(w, r)
		return
	}
	msg := fmt.Sprintf(`{
    "error": "请求地址不存在!!! 本镜像可访问路径: [ '/', '/foo' ], 可访问端口: [ '80', '8080' ], 容器日志目录: [ /var/golang-demo/log/ ]"
}
`)
	PrintLog(fmt.Sprintf("%v  %v  %v%v    %v", r.Method, http.StatusNotFound, r.Host, r.URL, r.Header))
	_, err := io.WriteString(w, msg)
	if err != nil {
		log.Fatal(err)
	}
}

func HomeHandler(response http.ResponseWriter, request *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}
	s := `{
    "msg":"welcome to home page, 我是主机[ %v ]!!! 本镜像可访问路径: [ '/', '/foo' ], 可访问端口: [ '80', '8080' ], 容器日志可挂载目录: [ '/var/golang-demo/log/' ]", 
    "request_api":"%v %v%v",
    "request_header":"%v"
}
`
	PrintLog(fmt.Sprintf("%v  %v  %v%v    %v", request.Method, http.StatusOK, request.Host, request.URL, request.Header))
	json := fmt.Sprintf(s, hostname, request.Method, request.Host, request.URL, request.Header)
	_, err = response.Write([]byte(json))
	if err != nil {
		log.Fatal(err)
	}
}

func FooHandler(response http.ResponseWriter, request *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}
	s := `{
    "msg":"welcome to foo page, 我是主机[ %v ]!!! 本镜像可访问路径: [ '/', '/foo' ], 可访问端口: [ '80', '8080' ], 容器日志可挂载目录: [ '/var/golang-demo/log/' ]", 
    "request_api":"%v %v%v",
    "request_header":"%v"
}
`
	PrintLog(fmt.Sprintf("%v  %v  %v%v    %v", request.Method, http.StatusOK, request.Host, request.URL, request.Header))
	json := fmt.Sprintf(s, hostname, request.Method, request.Host, request.URL, request.Header)
	_, err = response.Write([]byte(json))
	if err != nil {
		log.Fatal(err)
	}
}

// 自定义服务器
func MyServer(port int) {
	handler := MyHandler{
		Mux: make(map[string]func(http.ResponseWriter, *http.Request)),
	}
	handler.Mux["/"] = HomeHandler
	handler.Mux["/foo"] = FooHandler
	addr := fmt.Sprintf("0.0.0.0:%v", port)
	server := http.Server{
		Addr:    addr,
		Handler: &handler,
	}
	PrintLog("服务器启动成功, 请访问 http://" + addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func PrintLog(msg string) {
	logger.Println(msg)
}

func FreeLog() {
	for {
		PrintLog("服务器运行正常, 日志正常写入!!!")
		time.Sleep(time.Second * 3)
	}
}

const M = 1 << 20

func main() {
	PrintLog("程序运行开始...")
	rand.Seed(time.Now().UnixNano())
	go MyServer(8080)
	go MyServer(80)
	// 空闲日志, 每3秒写入一次
	go FreeLog()
	// CPU增加持续时长
	duration := int64(600)
	go func() {
		for {
			cpuUsage := (rand.Int63n(7) + 2) * 10
			PrintLog(fmt.Sprintf("容器cpu使用率: %v %%, 持续时间: %vs", cpuUsage, duration))
			TestCpu(cpuUsage, duration)
		}
	}()

	var memUsage = 0
	for {
		memUsage = (rand.Intn(3) + 1) * 100
		PrintLog(fmt.Sprintf("容器内存增加: %v M, 持续时间: %vs", memUsage, duration))
		TestMem(memUsage, duration)
	}
}
