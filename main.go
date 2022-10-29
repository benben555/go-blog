package main

import (
	"go-blog/common"
	"go-blog/router"
	"log"
	"net/http"
)

func init() {
	//模版加载
	common.LoadTemplate()
}

//type IndexData struct {
//	Title string `json:"title"`
//	Desc  string `json:"desc"`
//}

//func index(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content", "application/json")
//	var indexData IndexData
//	indexData.Title = "go-blog"
//	indexData.Desc = "10-25"
//	jsonStr, _ := json.Marshal(indexData)
//	w.Write(jsonStr)
//}

func main() {
	//web 程序 http协议 ip port
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	router.Router()
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}

}
