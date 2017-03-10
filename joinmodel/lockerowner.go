package joinmodel

import "git.iguiyu.com/park/locker-admin-back/model"

type LockerOwner struct {
	model.Locker `xorm:"extends"`
	Username string `xorm:"not null VARCHAR(45)"`
}

func (LockerOwner) TableName() string {
	return "locker"
}