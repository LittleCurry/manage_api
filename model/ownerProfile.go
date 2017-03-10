package model

type OwnerProfile struct {
	Id          int `xorm:"not null pk autoincr INT(11)"`
	OwnerUserId int `xorm:"not null index INT(11)"`
}
