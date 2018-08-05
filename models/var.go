package models

import (
	"gopkg.in/mgo.v2"
	"GoodsMaster-Go/sevices"
)

var (
	con *mgo.Session
	errMsg map[string]string
)

func init() {
	/*数据库连接*/
	con = mongodb.Connect()

	/*错误信息字典*/
	errMsg := make(map[string]string)
	// 数据库
	errMsg["db_connect_failure"] = "数据库连接失败"
	errMsg["db_select_failure"] = "数据查询失败"
	errMsg["db_select_empty"] = "没有查询结果"
	// 用户
	errMsg["user_not_found"] = "未找到用户"
}