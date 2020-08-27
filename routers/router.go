package routers

import (
	"all-pay/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//bitmall
	beego.Router("/bitmall/api/v1/callback", &controllers.BitMallController{}, "post:CallBack")
	beego.Router("/bitmall/api/v1/ack", &controllers.BitMallController{}, "post:Ack")
	//ccpay
	beego.Router("/ccpay/api/v1/callback", &controllers.CCPayController{}, "post:CallBack")
	beego.Router("/ccpay/api/v1/ack", &controllers.CCPayController{}, "post:Ack")
}
