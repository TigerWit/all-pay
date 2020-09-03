package bitmall

import (
	"all-pay/models"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
	"net/http"
	"net/url"
)

type client struct {
	RedirectUrl string
	CallbackUrl string
}

func NewClient() *client {
	return &client{
		RedirectUrl: "https://xxx.com/redirect.html?channel=4usdt",
		CallbackUrl: "https://xxx.com/callback?channel=4usdt",
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

func (c *client) Buy(crq *models.ChannelRequest) *models.ChannelResponse {
	client := &http.Client{}
	param := models.ConvertReq(crq).(*models.BuyParam)
	param.RedirectUrl = c.RedirectUrl
	param.CallbackUrl = c.CallbackUrl
	bytesData, _ := json.Marshal(param)
	req, _ := http.NewRequest("POST", fmt.Sprintf("https://%s/api/v1/order/buy", beego.AppConfig.String("4usdt")), bytes.NewReader(bytesData))
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	return parseResponse(body)
}

func (c *client) Sell(crq *models.ChannelRequest) *models.ChannelResponse {
	client := &http.Client{}
	param := models.ConvertReq(crq).(*models.SellParam)
	param.RedirectUrl = c.RedirectUrl
	param.CallbackUrl = c.CallbackUrl
	bytesData, _ := json.Marshal(param)
	req, _ := http.NewRequest("POST", fmt.Sprintf("https://%s/api/v1/order/sell", beego.AppConfig.String("4usdt")), bytes.NewReader(bytesData))
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	return parseResponse(body)
}

func (c *client) Order() {
	params := url.Values{}
	Url, err := url.Parse(fmt.Sprintf("https://%s/api/v1/order", beego.AppConfig.String("4usdt")))
	if err != nil {
		return
	}
	params.Set("merchant_order_no", "")
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

func (c *client) UserOrder() {
	params := url.Values{}
	Url, err := url.Parse(fmt.Sprintf("https://%s/api/v1/user/order", beego.AppConfig.String("4usdt")))
	if err != nil {
		return
	}
	params.Set("area_code", "")
	params.Set("mobile", "")
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

func (c *client) Withdraw(crq *models.ChannelRequest) *models.ChannelResponse {
	client := &http.Client{}
	param := models.ConvertReq(crq).(*models.WithdrawParam)
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

func (c *client) CallBack(input []byte) (*models.NotifyParam, error) {
	var param *models.NotifyParam
	var err error
	if err = json.Unmarshal(input, param); err == nil {
		return param, nil
	}
	return nil, err
}

func (c *client) Ack(input []byte) (*models.AckParam, error) {
	var param *models.AckParam
	var err error
	if err = json.Unmarshal(input, param); err == nil {
		return param, nil
	}
	return nil, err
}

func (c *client) ProcessSuccess() *models.BitMallReply {
	return &models.BitMallReply{
		0,
		"SUCCESS",
		nil,
		nil,
	}
}

func parseResponse(resp []byte) *models.ChannelResponse {
	var crsp *models.ChannelResponse
	if err := json.Unmarshal([]byte(string(resp)), crsp); err == nil {
		return crsp
	}
	return nil
}
