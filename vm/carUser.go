package vm

type TxCarUsers struct {
	CarUsers []TxCarUser        `json:"carusers"`
	Total    int64                `json:"total"`
}

type TxCarUser struct {
	Id           int                `json:"id"`
	WxOpenid     string                `json:"openid"`
	WxSessionKey string                `json:"sessionkey"`
	ExpireAt     int                `json:"expireat"`
	Money        string                `json:"money"`
}
