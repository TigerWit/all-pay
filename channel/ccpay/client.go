package ccpay

import (
	"github.com/TigerWit/all-pay/channel"
)

type client struct {
	*channel.MetaData
}

func NewClient(md *channel.MetaData) *client {
	return &client{
		md,
	}
}

func (c *client) Timestamp() {}

func (c *client) Price() {}

func (c *client) Buy(*channel.ChannelRequest) *channel.ChannelResponse {
	return nil
}

func (c *client) Sell(*channel.ChannelRequest) *channel.ChannelResponse {
	return nil
}

func (c *client) Balance() {}

func (c *client) Withdraw(*channel.ChannelRequest) *channel.ChannelResponse {
	return nil
}

func (c *client) WithdrawInfo() {}

func (c *client) Order(merchantOrderNo string) *channel.ChannelResponse{
	return nil
}

func (c *client) UserOrder(mobile string) *channel.ChannelResponse{
	return nil
}

func (c *client) CallBack([]byte) (*channel.ChannelNotifyInfo, error) {
	return nil, nil
}

func (c *client) Ack([]byte) (*channel.ChannelAckInfo, error) {
	return nil, nil
}

func (c *client) ProcessSuccess() interface{} {
	return nil
}
