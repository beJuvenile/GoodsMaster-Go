package mongodb

import (
	"gopkg.in/mgo.v2"
	"sync"
)

var (
	ins *mgo.Session
	conErr error
	once sync.Once
)

// 实例化Mongo
func Connect() *mgo.Session {
	once.Do(func() {
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