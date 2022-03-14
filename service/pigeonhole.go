/*
* @Author: lishuang
* @Date:   2022/3/14 14:46
 */

package service

import (
	"blog_go/config"
	"blog_go/dao"
	"blog_go/models"
)

func FindPostPigeonhole() models.PigeonholeRes {
	//查询所有的文章 进行月份的整理
	//查询所有的分类
	posts, _ := dao.GetPostAll()
	pigeonholeMap := make(map[string][]models.Post)
	for _, post := range posts {
		at := post.CreateAt
		month := at.Format("2006-01")
		pigeonholeMap[month] = append(pigeonholeMap[month], post)
	}
	categorys, _ := dao.GetAllCategory()
	return models.PigeonholeRes{
		config.Cfg.Viewer,
		config.Cfg.System,
		categorys,
		pigeonholeMap,
	}
}
