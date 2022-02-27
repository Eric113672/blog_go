package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type IndexData struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

func indexHtml(w http.ResponseWriter, r *http.Request) {
	t := template.New("index.html")
	//viewPath, _ := os.Getwd()
	//fmt.Println("nowPath", viewPath)
	t, _ = t.ParseFiles(`D:\blog_go\template\index.html`)
	var indexData IndexData
	indexData.Title = "go博客"
	indexData.Desc = "无框架实现博客"
	err := t.Execute(w, indexData)
	fmt.Println(err)
}

func main() {
	//程序入口 采用本地8080端口
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/", indexHtml)
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
