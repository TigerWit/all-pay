package cli

import "all-pay/channel/bitmall"

type chClient interface {
	common
	deal
	account
}

//公共接口
type common interface {
	Timestamp()
}

//交易相关接口
type deal interface {
	Price()
	Buy()
	Sell()
	Order()
	UserOrder()
}

//账户相关接口
type account interface {
	Balance()
	Withdraw()
	WithdrawInfo()
}

var Client chClient

func init() {
	Client = newChClient()
}

func newChClient() chClient {
	return bitmall.NewClient()
}
