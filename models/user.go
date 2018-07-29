package models

import (
	//"strconv"
	//"time"

	"gopkg.in/mgo.v2/bson"
	"errors"
)

type User struct {
	Id       bson.ObjectId `bson:"_id"`
	Nickname string `bson:"nickname"`
	Account string `bson:"account"`
	Password string `bson:"password"`
	Phone string `bson:"phone"`
	CreatedAt bson.MongoTimestamp `bson:"create_at"`
	UpdatedAt bson.MongoTimestamp `bson:"updated_at"`
}

// 查找用户
func GetUserByAccount(account string) (u *User, err error)  {
	if con == nil {
		return nil, errors.New(errMsg["db_connect_failure"])
	}

	c := con.DB("goods_master").C("users")
	err = c.Find(bson.M{"account": account}).One(&u)

	return u, err
}
