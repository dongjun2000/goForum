package main

import (
	. "goForum/routes"

	"log"
	"net/http"
)

func main() {
	startWebServer("8000")
}

// 通过指定端口启动 Web 服务器
func startWebServer(port string) {
	router := NewRouter()

	// 处理静态资源文件
	assets := http.FileServer(http.Dir("public"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", assets))

	server := http.Server{
		Addr: ":8000",
		Handler: router,
	}

	log.Println("Starting HTTP service at " + port)

	// 启动协程监听请求
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Error: ", err)
	}
}
