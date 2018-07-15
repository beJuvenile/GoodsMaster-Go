package models

import (
	"gopkg.in/mgo.v2"
	"errors"
	"gopkg.in/mgo.v2/bson"
)

var (
	Db *mgo.Session
	DbErr error
)

type Settings struct {
	//Id_ bson.ObjectId `bson:"_id"`
	Name string `bson:"name"`
	Value interface{} `bson:"value"`
} 

func init()  {
	Db, DbErr = mgo.Dial("localhost:27017")

}

func GetItemByName(name string) (s *Settings, err error) {
	if DbErr != nil {
		return nil, errors.New("DB connect error")
	}
	//defer Db.Close()

	c := Db.DB("goods_master").C("settings")
	err = c.Find(bson.M{"name": name}).One(&s)
	if err != nil {
		return nil, errors.New("not found")
	}

	return s, nil
}
