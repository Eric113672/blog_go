/*
* @Author: lishuang
* @Date:   2022/3/3 11:57
 */

package models

//  从数据库获取的
type Category struct {
	Cid      int
	Name     string
	CreateAt string
	UpdateAt string
}

type CategoryResponse struct {
	*HomeResponse
	CategoryName string
}
