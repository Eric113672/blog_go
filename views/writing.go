/*
* @Author: lishuang
* @Date:   2022/3/9 23:07
 */

package views

import (
	"blog_go/common"
	"blog_go/service"
	"net/http"
)

func (*HTMLApi) Writing(w http.ResponseWriter, r *http.Request) {
	writing := common.Template.Writing
	wr := service.Writing()
	writing.WriteData(w, wr)
}
