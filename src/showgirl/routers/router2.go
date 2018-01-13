package routers

import (
	"showgirl/controllers"

	"github.com/astaxie/beego"
)

func init() {

	beego.Router("/v2.0/*", &controllers.GoFront2Controller{}, "post:Post")
	beego.Router("/v2.0/*", &controllers.GoFront2Controller{}, "options:Options")
}
