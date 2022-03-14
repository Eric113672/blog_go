/*
* @Author: lishuang
* @Date:   2022/3/3 23:24
 */

package dao

import (
	"blog_go/models"
	"log"
)

func GetUserNameById(userId int) string {
	var userName string
	row := DB.QueryRow("select user_name from blog_user where uid=?", userId)
	if row.Err() != nil {
		log.Println(row.Err())
	}
	_ = row.Scan(&userName)

	return userName
}

func GetUser(userName, passwd string) *models.User {
	row := DB.QueryRow("select * from blog_user where user_name=? and passwd=? limit 1", userName, passwd)
	if row.Err() != nil {
		log.Println(row.Err())
		return nil
	}
	var user = &models.User{}
	err := row.Scan(&user.Uid, &user.UserName, &user.Passwd, &user.Avatar, &user.CreateAt, &user.UpdateAt)

	if err != nil {
		log.Println(err)
		return nil
	}
	return user

}
