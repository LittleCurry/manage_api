package model

type AdminUser struct {
	Id        int    `xorm:"not null pk autoincr INT(11)"`
	Loginname string `xorm:"not null VARCHAR(45)"`
	Password  string `xorm:"not null VARCHAR(45)"`
	Username  string `xorm:"not null VARCHAR(45)"`
}
