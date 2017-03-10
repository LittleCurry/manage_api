package model

type Bill struct {
	Id         int    `xorm:"not null pk autoincr INT(11)"`
	CarUserId  int    `xorm:"INT(11)"`
	CreateTime int    `xorm:"not null INT(11)"`
	Money      string `xorm:"not null DECIMAL(10,2)"`
	Type       int    `xorm:"not null TINYINT(1)"`
	Action     int    `xorm:"TINYINT(2)"`
	Remark     string `xorm:"VARCHAR(255)"`
}
