package model

type CarUser struct {
	Id           int    `xorm:"not null pk autoincr INT(11)"`
	WxOpenid     string `xorm:"not null VARCHAR(128)"`
	WxSessionKey string `xorm:"not null VARCHAR(128)"`
	ExpireAt     int    `xorm:"not null INT(11)"`
	Money        string `xorm:"DECIMAL(10,2)"`
}
