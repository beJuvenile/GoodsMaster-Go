package mongodb

import (
	"gopkg.in/mgo.v2"
	"sync"
)

var (
	ins *mgo.Session
	conErr error
	once sync.Once
	dbUrl string
)

// 实例化Mongo
func Connect() *mgo.Session {
	once.Do(func() {
		dbUrl = "mongodb://myuser:mypass@localhost:40001,otherhost:40001/mydb"
		ins, conErr = mgo.Dial("localhost:27017")

		if conErr != nil {
			ins = nil
		}
	})

	return ins
}


// 销毁实例
func Close()  {
	ins.Close()
}