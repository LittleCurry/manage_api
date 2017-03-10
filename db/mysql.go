package db

import "github.com/go-xorm/xorm"

var engine *xorm.Engine

func InitMySQL() {
	//Xorm
	engine, _ = xorm.NewEngine("mysql", "lock_test:lock_test@tcp(wx.iguiyu.com:56706)/lock_test?charset=utf8")
	engine.Ping()
	engine.DB().SetMaxIdleConns(100)
	engine.ShowSQL(true)
}

func GetMySQL() *xorm.Engine {
	return engine
}
