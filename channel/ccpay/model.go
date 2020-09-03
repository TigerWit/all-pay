package ccpay

type CCNotifyParam struct {
	TradeType      string `json:"tradeType"`
	BusinessType   int    `json:"businessType"`
	CoinName       string `json:"coinName"`
	UniqTradeNo    string `json:"uniqTradeNo"`
	TradeNo        string `json:"tradeNo"`
	OutIndex       int64  `json:"outIndex"`
	TradeAmount    string `json:"tradeAmount"`
	TradeFee       string `json:"tradeFee"`
	FromAddress    string `json:"fromAddress"`
	ToAddress      string `json:"toAddress"`
	TxStatus       string `json:"txStatus"`
	BlockTimestamp int64  `json:"blockTimestamp"`
	BlockHeight    int64  `json:"blockHeight"`
	TxHash         string `json:"txHash"`
}
type CCAckParam struct {
	Address            string `json:"address"`
	Chain              string `json:"chain"`
	Count              string `json:"count"`
	CreatedTime        string `json:"created_time"`
	Currency           string `json:"currency"`
	Decimal            string `json:"decimal"`
	Fee                string `json:"fee"`
	Memo               string `json:"memo"`
	MerchantWithdrawNo string `json:"merchant_withdraw_no"`
	OrderNo            string `json:"order_no"`
	Side               string `json:"side"`
}

type CCReply struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
