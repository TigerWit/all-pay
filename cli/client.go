package cli

import (
	"fmt"
	"github.com/TigerWit/all-pay/channel"
	"github.com/TigerWit/all-pay/channel/bitmall"
	"github.com/TigerWit/all-pay/channel/ccpay"
)

type ChannelClient interface {
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
	Buy(*channel.ChannelRequest) *channel.ChannelResponse
	Sell(*channel.ChannelRequest) *channel.ChannelResponse
	Order(merchantOrderNo string) *channel.ChannelResponse
	UserOrder(mobile string) *channel.ChannelResponse
}

//账户相关接口
type account interface {
	Balance()
	Withdraw(*channel.ChannelRequest) *channel.ChannelResponse
	WithdrawInfo()
}

//订单结果通知回调&确认提现
type confirm interface {
	CallBack([]byte) (*channel.ChannelNotifyInfo, error)
	Ack([]byte) (*channel.ChannelAckInfo, error)
	ProcessSuccess() interface{}
}

func NewChClient(channelID string, meta *channel.MetaData) ChannelClient {
	switch channelID {
	case "4usdt":
		return bitmall.NewClient(meta)
	case "ccpay":
		return ccpay.NewClient(meta)
	default:
		fmt.Println("un support channel!")
		return nil
	}
}
