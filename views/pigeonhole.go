/*
* @Author: lishuang
* @Date:   2022/3/14 14:45
 */

package views

import (
	"blog_go/common"
	"blog_go/service"
	"net/http"
)

func (*HTMLApi) Pigeonhole(w http.ResponseWriter, r *http.Request) {
	pigeonhole := common.Template.Pigeonhole

	pigeonholeRes := service.FindPostPigeonhole()
	pigeonhole.WriteData(w, pigeonholeRes)
}
