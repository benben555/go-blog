package views

import (
	"errors"
	"go-blog/common"
	"go-blog/service"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func (*HTMLApi) Category(w http.ResponseWriter, r *http.Request) {
	categoryTemplate := common.Template.Category
	//http://localhost:8080/c/1 1参数 分类的id
	path := r.URL.Path
	cidStr := strings.TrimPrefix(path, "/c/")
	cid, err := strconv.Atoi(cidStr)
	if err != nil {
		categoryTemplate.WriteError(w, errors.New("路径不匹配"))
		return
	}
	if err := r.ParseForm(); err != nil {
		log.Println("表单获取失败", err)
		categoryTemplate.WriteError(w, errors.New("系统错误，请联系管理员"))
		return
	}
	pageStr := r.Form.Get("page")
	if pageStr == "" {
		pageStr = "1"
	}
	page, _ := strconv.Atoi(pageStr)
	//每页显示的数量
	pageSize := 10

	categoryResponse, err := service.GetPostByCategoryId(cid, page, pageSize)
	if err != nil {
		categoryTemplate.WriteError(w, err)
		return
	}
	categoryTemplate.WriteData(w, categoryResponse)
}
