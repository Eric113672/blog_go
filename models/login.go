/*
* @Author: lishuang
* @Date:   2022/3/9 17:26
 */

package models

type LoginRes struct {
	Token    string   `json:"token"`
	UserInfo UserInfo `json:"user_info"`
}
