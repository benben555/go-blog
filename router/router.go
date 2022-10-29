package router

import (
	"go-blog/api"
	"go-blog/views"
	"net/http"
)

func Router() {
	//1.返回页面 views2.api 返回数据3.静态资源
	http.HandleFunc("/", views.HTML.Index)
	//http://localhost:8080/c/1 1参数 分类的id
	http.HandleFunc("/c/", views.HTML.Category)
	http.HandleFunc("/p/", views.HTML.Detail)
	http.HandleFunc("/login", views.HTML.Login)
	http.HandleFunc("/writing", views.HTML.Writing)
	http.HandleFunc("/api/v1/post", api.API.SaveAndUpdatePost)
	http.HandleFunc("/api/v1/post/", api.API.GetPost)
	http.HandleFunc("/api/v1/login", api.API.Login)
	//http.HandleFunc("/", views.HTML.Index)
	//http.HandleFunc("/index.html", indexHtml)
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))
}
