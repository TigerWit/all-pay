package cli

import (
	"all-pay/channel/bitmall"
	"all-pay/models"
)

type chClient interface {
	common
	deal
	account
	confirm
}

//公共接口
type common interface {
	Timestamp()
}

//交易相关接口
type deal interface {
	Price()
	Buy(*models.ChannelRequest) *models.ChannelResponse
	Sell(*models.ChannelRequest) *models.ChannelResponse
	Order()
	UserOrder()
}

//账户相关接口
type account interface {
	Balance()
	Withdraw(*models.ChannelRequest) *models.ChannelResponse
	WithdrawInfo()
}

//订单结果通知回调&确认提现
type confirm interface {
	CallBack([]byte) (*models.NotifyParam, error)
	Ack([]byte) (*models.AckParam, error)
	ProcessSuccess() interface{}
}

var Client chClient

func init() {
	Client = newChClient()
}

func newChClient() chClient {
	return bitmall.NewClient()
}
