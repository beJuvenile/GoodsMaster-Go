package models

import (
	"errors"
	"gopkg.in/mgo.v2/bson"
)

type Settings struct {
	Id 		bson.ObjectId 	`bson:"_id"`
	Name 	string 			`bson:"name"`
	Value 	interface{} 	`bson:"value"`
}

func GetSettingByName(name string) (s *Settings, err error) {
	if con == nil {
		return nil, errors.New(errMsg["db_connect_failure"])
	}

	c := con.DB("goods_master").C("settings")
	err = c.Find(bson.M{"name": name}).One(&s)
	if err != nil {
		return nil, errors.New(errMsg["db_select_failure"])
	} else {
		return s, nil
	}
}
