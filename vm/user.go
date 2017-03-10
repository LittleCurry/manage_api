package vm

type RxUser struct{
	LoginName	string `json:"loginname"`
	Password	string `json:"password"`
}

type TxUser struct{
	LoginName	string `json:"loginname"`
	SessionKey	string `json:"sessionkey"`
}