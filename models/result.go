/*
* @Author: lishuang
* @Date:   2022/3/9 17:08
 */

package models

type Result struct {
	Error string      `json:"error"`
	Data  interface{} `json:"data"`
	Code  int         `json:"code"`
}
