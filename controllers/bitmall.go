package controllers

import (
	"all-pay/models"
	"encoding/json"
	"github.com/astaxie/beego"
)

type BitMallController struct {
	beego.Controller
}

func (c *BitMallController) CallBack() {
	var param models.NotifyParam
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &param); err == nil {
		c.Data["json"] = models.BitMallReply{
			0,
			"SUCCESS",
			nil,
			nil,
		}
	} else {
		c.Data["json"] = models.BitMallReply{}
	}
	c.ServeJSON()
}

func (c *BitMallController) Ack() {
	var param models.AckParam
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &param); err == nil {
		c.Data["json"] = models.BitMallReply{
			0,
			"SUCCESS",
			nil,
			nil,
		}
	} else {
		c.Data["json"] = models.BitMallReply{}
	}
	c.ServeJSON()
}
