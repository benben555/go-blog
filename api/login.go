package api

import (
	"go-blog/common"
	"go-blog/service"
	"net/http"
)

func (*Api) Login(w http.ResponseWriter, r *http.Request) {
	//接收一个用户名和密码，返回对应的json
	params := common.GetRequestJsonParam(r)
	userName := params["username"].(string)
	passwd := params["passwd"].(string)
	loginRes, err := service.Login(userName, passwd)
	if err != nil {
		common.Error(w, err)
		return
	}
	common.Success(w, loginRes)
	print("api.login", loginRes)
}
