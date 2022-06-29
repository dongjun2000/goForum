package main

import (
	. "goForum/config"
	. "goForum/routes"

	"log"
	"net/http"
)

func main() {
	startWebServer()
}

// 启动 Web 服务器
func startWebServer() {
	// 通过 router.go 中定义的路由器来分发请求
	router := NewRouter()

	// 处理静态资源文件
	assets := http.FileServer(http.Dir(ViperConfig.App.Static))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", assets))

	server := http.Server{
		Addr: ViperConfig.App.Address,
		Handler: router,
	}

	log.Println("Starting HTTP service at " + ViperConfig.App.Address)

	// 启动协程监听请求
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Error: ", err.Error())
	}
}
