package model

type OwnerProfileDetail struct {
	Id             int    `xorm:"not null pk autoincr INT(11)"`
	OwnerProfileId int    `xorm:"not null index INT(11)"`
	OwnerUserId    int    `xorm:"not null index INT(11)"`
	Fulltime       int    `xorm:"not null TINYINT(1)"`
	BeginTime      int    `xorm:"INT(11)"`
	EndTime        int    `xorm:"INT(11)"`
	PricePerUnit   string `xorm:"not null DECIMAL(5,2)"`
	Unit           int    `xorm:"not null TINYINT(4)"`
	CappedPrice    string `xorm:"DECIMAL(5,2)"`
}
