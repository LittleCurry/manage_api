package model

type LockerParking struct {
	Id              int    `xorm:"not null pk autoincr INT(11)"`
	ParkingUserId   int    `xorm:"not null INT(11)"`
	LockerId        int    `xorm:"not null index INT(11)"`
	BeginTime       int    `xorm:"INT(10)"`
	EndTime         int    `xorm:"INT(10)"`
	WechatMessageId string `xorm:"VARCHAR(50)"`
	Money           string `xorm:"DECIMAL(10,2)"`
	Status          int    `xorm:"TINYINT(2)"`
}
