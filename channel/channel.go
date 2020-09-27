package channel

type MetaData struct {
	ClientId    string
	RedirectUrl string
	CallbackUrl string
	Sign        string
	SignVersion string
}

type OrderType int

const (
	Buy OrderType = iota
	Sell
	Settlement
)

type ChannelRequest struct {
	ChannelID          int       `json:channel_id`
	Type               OrderType `json:"type"`
	Currency           string    `json:"currency"`
	Count              float64   `json:"count"`
	LegalCurrency      string    `json:"legal_currency"`
	ClientId           string    `json:"client_id"`
	MerchantOrderNo    string    `json:"merchant_order_no"`
	KycLevel           string    `json:"kyc_level"`
	KycName            string    `json:"kyc_name"`
	KycIdType          string    `json:"kyc_id_type"`
	KycIdNo            string    `json:"kyc_id_no"`
	AreaCode           string    `json:"area_code"`
	Mobile             string    `json:"mobile"`
	Email              string    `json:"email"`
	RedirectUrl        string    `json:"redirect_url"`
	CallbackUrl        string    `json:"callback_url"`
	Timestamp          string    `json:"timestamp"`
	Sign               string    `json:"sign"`
	SignVersion        string    `json:"sign_version"`
	PayType            string    `json:"payType"`
	BankName           string    `json:"bank_name"`
	BankBranchName     string    `json:"bank_branch_name"`
	BankAccountNo      string    `json:"bank_account_no"`
	AlipayAccountNo    string    `json:"alipay_account_no"`
	AlipayQrlink       string    `json:"alipay_qrlink"`
	Address            string    `json:"address"`
	Chain              string    `json:"chain"`
	MerchantWithdrawNo string    `json:"merchant_withdraw_no"`
}

type ChannelResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
	Meta interface{} `json:"meta"`
}

type ChannelNotifyInfo struct {
	MerchantOrderNo string `json:"merchant_order_no"`
	OrderNo         string `json:"order_no"`
	Count           int    `json:"count"`
}

type ChannelAckInfo struct {
}
