package main

import (
	. "goForum/routes"
	. "goForum/config"

	"log"
	"net/http"
)

func main() {
	startWebServer()
}

// 启动 Web 服务器
func startWebServer() {
	// 初始化全局配置
	config := LoadConfig()
	router := NewRouter()

	// 处理静态资源文件
	assets := http.FileServer(http.Dir(config.App.Static))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", assets))

	server := http.Server{
		Addr: config.App.Address,
		Handler: router,
	}

	log.Println("Starting HTTP service at " + config.App.Address)

	// 启动协程监听请求
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Error: ", err.Error())
	}
}
