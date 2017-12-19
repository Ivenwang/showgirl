package routers

import (
	"showgirl/controllers"

	"github.com/astaxie/beego"
)

func init() {

	beego.Router("/callpack/pay", &controllers.NotifyPayController{}, "post:NotifyPayInfo")
}
