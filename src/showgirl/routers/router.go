/*
 * Auto generated by client.python
 * Please do not modify it.
 */
package routers

import (
	"github.com/astaxie/beego"
	"os"
	"strings"
	"showgirl/controllers"
)

func init() {
	moduleStr := beego.AppConfig.String("module")
	modules := strings.Split(moduleStr, ",")
	for _, m := range modules {
		switch m {
		
		case "Account":
			beego.Router("/account/thirdpartywxlogin", &controllers.AccountController{}, "post:ThirdPartyWXLogin")
		
			beego.Router("/account/getmyinfo", &controllers.AccountController{}, "post:GetMyInfo")
		
		case "Image":
			beego.Router("/image/querystylelist", &controllers.ImageController{}, "post:QueryStyleList")
		
			beego.Router("/image/queryresourcelist", &controllers.ImageController{}, "post:QueryResourceList")
		
			beego.Router("/image/createstyle", &controllers.ImageController{}, "post:CreateStyle")
		
			beego.Router("/image/uploadimage", &controllers.ImageController{}, "post:UploadImage")
		
			beego.Router("/image/deletestyle", &controllers.ImageController{}, "post:DeleteStyle")
		
			beego.Router("/image/deleteresource", &controllers.ImageController{}, "post:DeleteResource")
		
			beego.Router("/image/updatestyle", &controllers.ImageController{}, "post:UpdateStyle")
		
		case "Pay":
			beego.Router("/pay/createtransaction", &controllers.PayController{}, "post:CreateTransaction")
		
		case "Recommend":
			beego.Router("/recommend/queryrecommendlist", &controllers.RecommendController{}, "post:QueryRecommendList")
		
			beego.Router("/recommend/styleimagelist", &controllers.RecommendController{}, "post:StyleImageList")
		
		case "GoFront":
		
			beego.Router("/v1.0/account/thirdpartywxlogin", &controllers.GoFrontController{}, "post:Account_ThirdPartyWXLogin")
		
			beego.Router("/v1.0/account/getmyinfo", &controllers.GoFrontController{}, "post:Account_GetMyInfo")
		
			beego.Router("/v1.0/image/querystylelist", &controllers.GoFrontController{}, "post:Image_QueryStyleList")
		
			beego.Router("/v1.0/image/queryresourcelist", &controllers.GoFrontController{}, "post:Image_QueryResourceList")
		
			beego.Router("/v1.0/image/createstyle", &controllers.GoFrontController{}, "post:Image_CreateStyle")
		
			beego.Router("/v1.0/image/uploadimage", &controllers.GoFrontController{}, "post:Image_UploadImage")
		
			beego.Router("/v1.0/image/deletestyle", &controllers.GoFrontController{}, "post:Image_DeleteStyle")
		
			beego.Router("/v1.0/image/deleteresource", &controllers.GoFrontController{}, "post:Image_DeleteResource")
		
			beego.Router("/v1.0/image/updatestyle", &controllers.GoFrontController{}, "post:Image_UpdateStyle")
		
			beego.Router("/v1.0/pay/createtransaction", &controllers.GoFrontController{}, "post:Pay_CreateTransaction")
		
			beego.Router("/v1.0/recommend/queryrecommendlist", &controllers.GoFrontController{}, "post:Recommend_QueryRecommendList")
		
			beego.Router("/v1.0/recommend/styleimagelist", &controllers.GoFrontController{}, "post:Recommend_StyleImageList")
		
		case "OpGateway":
		
			beego.Router("/op/account/thirdpartywxlogin", &controllers.OpGatewayController{}, "post:Account_ThirdPartyWXLogin")
		
			beego.Router("/op/account/getmyinfo", &controllers.OpGatewayController{}, "post:Account_GetMyInfo")
		
			beego.Router("/op/image/querystylelist", &controllers.OpGatewayController{}, "post:Image_QueryStyleList")
		
			beego.Router("/op/image/queryresourcelist", &controllers.OpGatewayController{}, "post:Image_QueryResourceList")
		
			beego.Router("/op/image/createstyle", &controllers.OpGatewayController{}, "post:Image_CreateStyle")
		
			beego.Router("/op/image/uploadimage", &controllers.OpGatewayController{}, "post:Image_UploadImage")
		
			beego.Router("/op/image/deletestyle", &controllers.OpGatewayController{}, "post:Image_DeleteStyle")
		
			beego.Router("/op/image/deleteresource", &controllers.OpGatewayController{}, "post:Image_DeleteResource")
		
			beego.Router("/op/image/updatestyle", &controllers.OpGatewayController{}, "post:Image_UpdateStyle")
		
			beego.Router("/op/pay/createtransaction", &controllers.OpGatewayController{}, "post:Pay_CreateTransaction")
		
			beego.Router("/op/recommend/queryrecommendlist", &controllers.OpGatewayController{}, "post:Recommend_QueryRecommendList")
		
			beego.Router("/op/recommend/styleimagelist", &controllers.OpGatewayController{}, "post:Recommend_StyleImageList")
		
		default:
			beego.BeeLogger.Critical("Unknown module. module:%s", m)
			os.Exit(1)
		}
	}

}
