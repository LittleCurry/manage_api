package vm

import "time"

type TxOwnerUsers struct{
	OwnerUsers	[]TxOwnerUser	`json:"ownerusers"`
	Total		int64		`json:"total"`
}

type TxOwnerUser struct{
	Id		int		`json:"id"`
	Loginname	string		`json:"loginname"`
	Nickname	string		`json:"nickname"`
	Create		string		`json:"createdate"`
	LastLogin	string		`json:"lastlogindate"`
	LockersOwned	int		`json:"lockersowned"`
}

func(this *TxOwnerUser) CreateDate(createDate time.Time){
	this.Create = createDate.Format("2006-01-02")
}

func(this *TxOwnerUser) LastLoginDate(lastLoginDate time.Time){
	this.LastLogin = lastLoginDate.Format("2006-01-02")
}