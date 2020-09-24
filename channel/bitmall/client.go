package bitmall

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/TigerWit/all-pay/channel"
	"github.com/astaxie/beego"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type client struct {
	*channel.MetaData
}

func NewClient(md *channel.MetaData) *client {
	return &client{
		md,
	}
}

func (c *client) Timestamp() {
	resp, err := http.Get(fmt.Sprintf("https://%s/api/v1/timestamp", beego.AppConfig.String("4usdt")))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	fmt.Println(resp.StatusCode)
	if resp.StatusCode == 200 {
		fmt.Println("ok")
	}
}

func (c *client) Price() {
	params := url.Values{}
	Url, err := url.Parse(fmt.Sprintf("https://%s/api/v1/price", beego.AppConfig.String("4usdt")))
	if err != nil {
		return
	}
	params.Set("currency", "")
	params.Set("legal_currency", "")
	params.Set("client_id", "")
	params.Set("timestamp", "")
	params.Set("sign", "")
	params.Set("sign_version", "")
	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	fmt.Println(urlPath)
	resp, err := http.Get(urlPath)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func (c *client) Buy(crq *channel.ChannelRequest) *channel.ChannelResponse {
	client := &http.Client{}
	param := convertReq(crq).(*BuyParam)
	param.RedirectUrl = c.RedirectUrl
	param.CallbackUrl = c.CallbackUrl
	bytesData, _ := json.Marshal(param)
	req, _ := http.NewRequest("POST", fmt.Sprintf("https://%s/api/v1/order/buy", beego.AppConfig.String("4usdt")), bytes.NewReader(bytesData))
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	return parseResponse(body)
}

func (c *client) Sell(crq *channel.ChannelRequest) *channel.ChannelResponse {
	client := &http.Client{}
	param := convertReq(crq).(*SellParam)
	param.RedirectUrl = c.RedirectUrl
	param.CallbackUrl = c.CallbackUrl
	bytesData, _ := json.Marshal(param)
	req, _ := http.NewRequest("POST", fmt.Sprintf("https://%s/api/v1/order/sell", beego.AppConfig.String("4usdt")), bytes.NewReader(bytesData))
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	return parseResponse(body)
}

func (c *client) Order(merchantOrderNo string) *channel.ChannelResponse {
	params := url.Values{}
	Url, err := url.Parse(fmt.Sprintf("https://%s/api/v1/order", beego.AppConfig.String("4usdt")))
	if err != nil {
		return nil
	}
	params.Set("merchant_order_no", merchantOrderNo)
	params.Set("client_id", c.ClientId)
	params.Set("timestamp", time.Now().String())
	params.Set("sign", c.Sign)
	params.Set("sign_version", c.SignVersion)
	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	fmt.Println(urlPath)
	resp, err := http.Get(urlPath)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return parseResponse(body)
}

func (c *client) UserOrder(mobile string) *channel.ChannelResponse {
	params := url.Values{}
	Url, err := url.Parse(fmt.Sprintf("https://%s/api/v1/user/order", beego.AppConfig.String("4usdt")))
	if err != nil {
		return nil
	}
	params.Set("area_code", "")
	params.Set("mobile", mobile)
	params.Set("client_id", c.ClientId)
	params.Set("timestamp", time.Now().String())
	params.Set("sign", c.Sign)
	params.Set("sign_version", c.SignVersion)
	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	fmt.Println(urlPath)
	resp, err := http.Get(urlPath)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return parseResponse(body)
}

func (c *client) Balance() {
	params := url.Values{}
	Url, err := url.Parse(fmt.Sprintf("https://%s/api/v1/account/balance", beego.AppConfig.String("4usdr")))
	if err != nil {
		return
	}
	params.Set("client_id", "")
	params.Set("timestamp", "")
	params.Set("sign", "")
	params.Set("sign_version", "")
	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	fmt.Println(urlPath)
	resp, err := http.Get(urlPath)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func (c *client) Withdraw(crq *channel.ChannelRequest) *channel.ChannelResponse {
	client := &http.Client{}
	param := convertReq(crq).(*WithdrawParam)
	bytesData, _ := json.Marshal(param)
	req, _ := http.NewRequest("POST", fmt.Sprintf("https://%s/api/v1/account/withdraw", beego.AppConfig.String("4usdt")), bytes.NewReader(bytesData))
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	return parseResponse(body)
}

func (c *client) WithdrawInfo() {
	params := url.Values{}
	Url, err := url.Parse(fmt.Sprintf("https://%s/api/v1/account/withdraw/info", beego.AppConfig.String("4usdt")))
	if err != nil {
		return
	}
	params.Set("merchant_withdraw_no", "")
	params.Set("client_id", "")
	params.Set("timestamp", "")
	params.Set("sign", "")
	params.Set("sign_version", "")
	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	fmt.Println(urlPath)
	resp, err := http.Get(urlPath)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func (c *client) CallBack(input []byte) (*channel.ChannelNotifyInfo, error) {
	var param *NotifyParam
	var err error
	if err = json.Unmarshal(input, param); err == nil {
		return wrapNotifyInfo(param), nil
	}
	return nil, err
}

func wrapNotifyInfo(param *NotifyParam) *channel.ChannelNotifyInfo {
	return &channel.ChannelNotifyInfo{}
}

func (c *client) Ack(input []byte) (*channel.ChannelAckInfo, error) {
	var param *AckParam
	var err error
	if err = json.Unmarshal(input, param); err == nil {
		return wrapAckInfo(param), nil
	}
	return nil, err
}

func wrapAckInfo(param *AckParam) *channel.ChannelAckInfo {
	return &channel.ChannelAckInfo{}
}

func (c *client) ProcessSuccess() interface{} {
	return &BitMallReply{
		0,
		"SUCCESS",
		nil,
		nil,
	}
}

func parseResponse(resp []byte) *channel.ChannelResponse {
	var crsp *channel.ChannelResponse
	if err := json.Unmarshal([]byte(string(resp)), crsp); err == nil {
		return crsp
	}
	return nil
}
