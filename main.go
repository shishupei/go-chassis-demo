package main

import (
	"fmt"
	"github.com/go-chassis/go-chassis/v2"
	"log"
	"net/http"
)

// HelloWorldHandler 定义一个处理请求的函数
func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	// 设置响应头
	w.Header().Set("Content-Type", "application/json")
	// 返回 HelloWorld 消息
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"message": "Hello, ServiceComb with GoChassis!"}`)
}

func main() {
	// 初始化 GoChassis
	err := chassis.Init()
	if err != nil {
		log.Fatalf("Failed to initialize GoChassis: %v", err)
	}

	// 创建一个 REST 服务，并注册路由
	http.HandleFunc("/hello", HelloWorldHandler)

	// 启动服务
	port := ":8080"
	log.Printf("Starting HelloWorld API on port %s", port)
	err = http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
