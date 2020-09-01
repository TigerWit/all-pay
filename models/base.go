package models

import "encoding/json"

type OrderType int

const (
	Buy OrderType = iota
	Sell
	Settlement
)

type ChannelRequest struct {
	Type               int     `json:"type"`
	Currency           string  `json:"currency"`
	Count              float64 `json:"count"`
	LegalCurrency      string  `json:"legal_currency"`
	ClientId           string  `json:"client_id"`
	MerchantOrderNo    string  `json:"merchant_order_no"`
	KycLevel           string  `json:"kyc_level"`
	KycName            string  `json:"kyc_name"`
	KycIdType          string  `json:"kyc_id_type"`
	KycIdNo            string  `json:"kyc_id_no"`
	AreaCode           string  `json:"area_code"`
	Mobile             string  `json:"mobile"`
	Email              string  `json:"email"`
	RedirectUrl        string  `json:"redirect_url"`
	CallbackUrl        string  `json:"callback_url"`
	Timestamp          string  `json:"timestamp"`
	Sign               string  `json:"sign"`
	SignVersion        string  `json:"sign_version"`
	PayType            string  `json:"payType"`
	BankName           string  `json:"bank_name"`
	BankBranchName     string  `json:"bank_branch_name"`
	BankAccountNo      string  `json:"bank_account_no"`
	AlipayAccountNo    string  `json:"alipay_account_no"`
	AlipayQrlink       string  `json:"alipay_qrlink"`
	Address            string  `json:"address"`
	Chain              string  `json:"chain"`
	MerchantWithdrawNo string  `json:"merchant_withdraw_no"`
}

func ConvertReq(req *ChannelRequest) interface{} {
	switch req.Type {
	case int(Buy):
		return cvtBuy(req)
	case int(Sell):
		return cvtSell(req)
	case int(Settlement):
		return cvtWithdraw(req)
	}
	return nil
}

func cvtBuy(req *ChannelRequest) *BuyParam {
	reqj, _ := json.Marshal(req)
	buyParam := new(BuyParam)
	_ = json.Unmarshal(reqj, buyParam)
	return buyParam
}

func cvtSell(req *ChannelRequest) *SellParam {
	reqj, _ := json.Marshal(req)
	sellParam := new(SellParam)
	_ = json.Unmarshal(reqj, sellParam)
	return sellParam
}

func cvtWithdraw(req *ChannelRequest) *WithdrawParam {
	reqj, _ := json.Marshal(req)
	withdrawParam := new(WithdrawParam)
	_ = json.Unmarshal(reqj, withdrawParam)
	return withdrawParam
}

type ChannelResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
	Meta interface{} `json:"meta"`
}