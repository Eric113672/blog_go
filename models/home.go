/*
* @Author: lishuang
* @Date:   2022/3/3 11:57
 */

package models

import "blog_go/config"

type HomeResponse struct {
	config.Viewer
	Categorys []Category
	Posts     []PostMore
	Total     int
	Page      int
	Pages     []int
	PageEnd   bool
}
