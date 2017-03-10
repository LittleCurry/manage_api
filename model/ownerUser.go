package model

import (
	"time"
)

type OwnerUser struct {
	Id            int       `xorm:"not null pk autoincr INT(11)"`
	Loginname     string    `xorm:"not null VARCHAR(45)"`
	Passwd        string    `xorm:"not null VARCHAR(45)"`
	Nickname      string    `xorm:"VARCHAR(45)"`
	CreateDate    time.Time `xorm:"DATE"`
	LastLoginDate time.Time `xorm:"DATE"`
	LockersOwned  int       `xorm:"INT(11)"`
}
