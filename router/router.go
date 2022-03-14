/*
* @Author: lishuang
* @Date:   2022/3/3 11:57
 */

package router

import (
	"blog_go/api"
	"blog_go/views"
	"net/http"
)

func Router() {
	//1. 页面  views 2. api 数据（json） 3. 静态资源
	http.HandleFunc("/", views.HTML.Index)
	//http://localhost:8080/c/1  1参数 分类的id
	http.HandleFunc("/c/", views.HTML.Category)
	// 登录页面
	http.HandleFunc("/login", views.HTML.Login)
	//文章详细内容
	http.HandleFunc("/p/", views.HTML.Detail)
	//写作页面
	http.HandleFunc("/writing", views.HTML.Writing)
	// 归档页面
	http.HandleFunc("/pigeonhole", views.HTML.Pigeonhole)

	// api
	http.HandleFunc("/api/v1/post", api.API.SaveAndUpdatePost)
	http.HandleFunc("/api/v1/login", api.API.Login)
	http.HandleFunc("/api/v1/post/", api.API.GetPost)
	http.HandleFunc("/api/v1/qiniu/token", api.API.QiniuToken)
	http.HandleFunc("/api/v1/post/search", api.API.SearchPost)

	// 静态资源文件服务器  这种静态还是用nginx来搭 效率比直接用后台好
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))
}
