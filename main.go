package main

import (
	"fmt"
	"net/http"
	"os"
)

// HTTP 处理函数
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, HTTP World!")
}

// 启动 HTTP 服务器
func startHTTPServer() {
	http.HandleFunc("/", helloHandler)
	fmt.Println("Starting HTTP server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting HTTP server:", err)
		os.Exit(1)
	}
}

func main() {
	go startHTTPServer()

	// 阻塞主协程，使程序持续运行
	select {}
}
