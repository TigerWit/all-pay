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

type client struct{}

func NewClient() *client {
	return &client{}
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

func (c *client) Buy() {
	client := &http.Client{}
	param := &models.BuyParam{}
	bytesData, _ := json.Marshal(param)
	req, _ := http.NewRequest("POST", fmt.Sprintf("https://%s/api/v1/order/buy", beego.AppConfig.String("4usdt")), bytes.NewReader(bytesData))
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func (c *client) Sell() {
	client := &http.Client{}
	param := &models.SellParam{}
	bytesData, _ := json.Marshal(param)
	req, _ := http.NewRequest("POST", fmt.Sprintf("https://%s/api/v1/order/sell", beego.AppConfig.String("4usdt")), bytes.NewReader(bytesData))
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
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

func (c *client) Withdraw() {
	client := &http.Client{}
	param := &models.WithdrawParam{}
	bytesData, _ := json.Marshal(param)
	req, _ := http.NewRequest("POST", fmt.Sprintf("https://%s/api/v1/account/withdraw", beego.AppConfig.String("4usdt")), bytes.NewReader(bytesData))
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
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
