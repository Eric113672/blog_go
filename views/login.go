/*
* @Author: lishuang
* @Date:   2022/3/8 17:25
 */

package views

import (
	"blog_go/common"
	"blog_go/config"
	"net/http"
)

func (*HTMLApi) Login(w http.ResponseWriter, r *http.Request) {
	login := common.Template.Login

	login.WriteData(w, config.Cfg.Viewer)
}
