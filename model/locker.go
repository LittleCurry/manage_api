package model

import (
	"time"
)

type Locker struct {
	Id             int       `xorm:"not null pk autoincr INT(11)"`
	ShortId        string    `xorm:"not null VARCHAR(6)"`
	Qrcode         string    `xorm:"not null VARCHAR(45)"`
	Simid          string    `xorm:"not null VARCHAR(45)"`
	PhoneNumber    string    `xorm:"not null VARCHAR(13)"`
	OwnerProfileId int       `xorm:"not null index INT(11)"`
	OwnerUserId    int       `xorm:"not null index INT(11)"`
	Status         int       `xorm:"not null INT(11)"`
	Version        string    `xorm:"VARCHAR(45)"`
	Active         int       `xorm:"not null TINYINT(2)"`
	CreateDate     time.Time `xorm:"DATE"`
	ActiveDate     time.Time `xorm:"DATE"`
	Address        string    `xorm:"VARCHAR(200)"`
	Longitude      float64   `xorm:"DOUBLE"`
	Latitude       float64   `xorm:"DOUBLE"`
}
