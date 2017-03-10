package model

type Recharge struct {
	Id         int    `xorm:"not null pk autoincr INT(11)"`
	CarUserId  int    `xorm:"not null INT(11)"`
	OrderId    string `xorm:"not null VARCHAR(45)"`
	CreateTime int    `xorm:"not null INT(11)"`
	PayTime    int    `xorm:"not null INT(11)"`
	Money      string `xorm:"DECIMAL(10,2)"`
	Status     int    `xorm:"not null TINYINT(1)"`
}
