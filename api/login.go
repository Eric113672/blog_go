/*
* @Author: lishuang
* @Date:   2022/3/9 17:05
 */

package api

import (
	"blog_go/common"
	"blog_go/service"
	"net/http"
)

func (*Api) Login(w http.ResponseWriter, r *http.Request) {
	// 接收用户名与密码
	params := common.GetRequestJsonParam(r)
	userName := params["username"].(string)
	passwd := params["passwd"].(string)
	loginRes, err := service.Login(userName, passwd)
	if err != nil {
		common.Error(w, err)
		return
	}
	//log.Println(loginRes.UserInfo.UserName, loginRes.UserInfo.Uid)
	common.Success(w, loginRes)
}
