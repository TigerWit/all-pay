package models

type BuyParam struct {
	Currency        string  `json:"currency"`
	Count           float64 `json:"count"`
	LegalCurrency   string  `json:"legal_currency"`
	ClientId        string  `json:"client_id"`
	MerchantOrderNo string  `json:"merchant_order_no"`
	KycLevel        string  `json:"kyc_level"`
	KycName         string  `json:"kyc_name"`
	KycIdType       string  `json:"kyc_id_type"`
	KycIdNo         string  `json:"kyc_id_no"`
	AreaCode        string  `json:"area_code"`
	Mobile          string  `json:"mobile"`
	Email           string  `json:"email"`
	RedirectUrl     string  `json:"redirect_url"`
	CallbackUrl     string  `json:"callback_url"`
	Timestamp       string  `json:"timestamp"`
	Sign            string  `json:"sign"`
	SignVersion     string  `json:"sign_version"`
	PayType         string  `json:"payType"`
}

type SellParam struct {
	Currency        string  `json:"currency"`
	Count           float64 `json:"count"`
	LegalCurrency   string  `json:"legal_currency"`
	ClientId        string  `json:"client_id"`
	MerchantOrderNo string  `json:"merchant_order_no"`
	KycLevel        string  `json:"kyc_level"`
	KycName         string  `json:"kyc_name"`
	KycIdType       string  `json:"kyc_id_type"`
	KycIdNo         string  `json:"kyc_id_no"`
	AreaCode        string  `json:"area_code"`
	Mobile          string  `json:"mobile"`
	Email           string  `json:"email"`
	RedirectUrl     string  `json:"redirect_url"`
	CallbackUrl     string  `json:"callback_url"`
	Timestamp       string  `json:"timestamp"`
	Sign            string  `json:"sign"`
	BankName        string  `json:"bank_name"`
	BankBranchName  string  `json:"bank_branch_name"`
	BankAccountNo   string  `json:"bank_account_no"`
	AlipayAccountNo string  `json:"alipay_account_no"`
	AlipayQrlink    string  `json:"alipay_qrlink"`
	SignVersion     string  `json:"sign_version"`
}

type NotifyParam struct {
	MerchantOrderNo string  `json:"merchant_order_no"`
	OrderNo         string  `json:"order_no"`
	Side            string  `json:"side"`
	Currency        string  `json:"currency"`
	Count           float64 `json:"count"`
	Status          string  `json:"status"`
	TradeTime       string  `json:"trade_time"`
	ClientId        string  `json:"client_id"`
	Timestamp       string  `json:"timestamp"`
	Sign            string  `json:"sign"`
	SignVersion     string  `json:"sign_version"`
}

type WithdrawParam struct {
	Address            string  `json:"address"`
	Chain              string  `json:"chain"`
	Count              float64 `json:"count"`
	Currency           string  `json:"currency"`
	MerchantWithdrawNo string  `json:"merchant_withdraw_no"`
	ClientId           string  `json:"client_id"`
	Sign               string  `json:"sign"`
	SignVersion        string  `json:"sign_version"`
	Timestamp          string  `json:"timestamp"`
}

type AckParam struct {
	Address            string  `json:"address"`
	Chain              string  `json:"chain"`
	Count              float64 `json:"count"`
	Currency           string  `json:"currency"`
	Fee                string  `json:"fee"`
	MerchantWithdrawNo string  `json:"merchant_withdraw_no"`
	Decimal            string  `json:"decimal"`
	Side               string  `json:"side"`
	Memo               string  `json:"memo"`
	CreatedTime        string  `json:"created_time"`
	ClientId           string  `json:"client_id"`
	Timestamp          string  `json:"timestamp"`
	SignVersion        string  `json:"sign_version"`
	Sign               string  `json:"sign"`
}
