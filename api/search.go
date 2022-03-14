/*
* @Author: lishuang
* @Date:   2022/3/14 21:08
 */

package api

import (
	"blog_go/common"
	"blog_go/service"
	"net/http"
)

func (*Api) SearchPost(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	condition := r.Form.Get("val")
	searchResp := service.SearchPost(condition)
	common.Success(w, searchResp)
}
