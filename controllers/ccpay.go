package controllers

import (
	"all-pay/models"
	"encoding/json"
	"github.com/astaxie/beego"
)

type CCPayController struct {
	beego.Controller
}

func (c *CCPayController) CallBack() {
	var param models.CCNotifyParam
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &param); err == nil {
		c.Data["json"] = models.CCReply{
			0,
			"SUCCESS",
		}
	} else {
		c.Data["json"] = models.CCReply{}
	}
	c.ServeJSON()
}

func (c *CCPayController) Ack() {
	var param models.CCAckParam
	var err error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &param); err == nil {
		c.Data["json"] = models.CCReply{
			0,
			"SUCCESS",
		}
	} else {
		c.Data["json"] = models.CCReply{}
	}
	c.ServeJSON()
}
