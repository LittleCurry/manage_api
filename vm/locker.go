package vm

import (
	"time"
)

type TxLockers struct{
	Lockers		[]TxLocker      `json:"lockers"`
	Total		int64		`json:"total"`
}

type TxLocker struct {
	Id		int    `json:"id"`
	ShortId		string `json:"shortid"`
	Qrcode		string `json:"qrcode"`
	Simid		string `json:"simid"`
	Version		string `json:"version"`
	Create		string `json:"createdate"`
}

type TxActiveLockers struct{
	ActiveLockers	[]TxActiveLocker	`json:"activelockers"`
	Total		int64			`json:"total"`
}

type TxActiveLocker struct {
	Id		int    `json:"id"`
	ShortId		string `json:"shortid"`
	Qrcode		string `json:"qrcode"`
	Owner		string `json:"owner"`
	Active		string `json:"activedate"`
	Address		string `json:"address"`
}

func(this *TxLocker) CreateDate(createDate time.Time){
	this.Create = createDate.Format("2006-01-02")
}

func(this *TxActiveLocker) Username(userName string){
	this.Owner = userName
}

func(this *TxActiveLocker) ActiveDate(activeDate time.Time){
	this.Active = activeDate.Format("2006-01-02")
}

type RxLocker struct {
	Id             int    `json:"id"`
	ShortId        string `json:"shortid"`
	Qrcode         string `json:"qrcode"`
	Abc            string `json:"abc"`
}

type RxSearch struct {
	Search		string `json:"search"`
}

