/*
* @Author: lishuang
* @Date:   2022/3/3 11:57
 */

package main

import (
	"blog_go/common"
	"blog_go/router"
	"log"
	"net/http"
)

func init() {
	//模板加载
	common.LoadTemplate()
}
func main() {
	//程序入口，一个项目 只能有一个入口
	//web程序，http协议 ip port
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	//路由
	router.Router()
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
