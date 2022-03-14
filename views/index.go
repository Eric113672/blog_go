/*
* @Author: lishuang
* @Date:   2022/3/3 11:57
 */

package views

import (
	"blog_go/common"
	"blog_go/service"
	"errors"
	"log"
	"net/http"
	"strconv"
)

func (*HTMLApi) Index(w http.ResponseWriter, r *http.Request) {
	index := common.Template.Index
	//页面上涉及到的所有的数据，必须有定义
	//数据库查询
	if err := r.ParseForm(); err != nil {
		log.Println("表单获取失败：", err)
		index.WriteError(w, errors.New("系统错误，请联系管理员!!"))
		return
	}
	pageStr := r.Form.Get("page")
	if pageStr == "" {
		pageStr = "1"
	}
	page, _ := strconv.Atoi(pageStr)
	//每页显示的数量
	pageSize := 10

	hr, err := service.GetAllIndexInfo(page, pageSize)
	if err != nil {
		log.Println("Index获取数据出错：", err)
		index.WriteError(w, errors.New("系统错误，请联系管理员!!"))
	}
	index.WriteData(w, hr)
}