/*
* @Author: lishuang
* @Date:   2022/3/14 21:09
 */

package service

import (
	"blog_go/dao"
	"blog_go/models"
)

func SearchPost(condition string) []models.SearchResp {
	posts, _ := dao.GetPostSearch(condition)
	var searchResps []models.SearchResp
	for _, post := range posts {
		searchResps = append(searchResps, models.SearchResp{
			post.Pid,
			post.Title,
		})
	}
	return searchResps
}
