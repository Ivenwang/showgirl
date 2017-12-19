package controllers

import (
	"encoding/xml"
	"runtime/debug"
	"showgirl/models/Pay"
	"showgirl/models/utils"

	"github.com/astaxie/beego"
)

//NotifyPayController 支付通知控制器
type NotifyPayController struct {
	beego.Controller
}

//NotifyPayRsp 支付通知回包结构
type NotifyPayRsp struct {
	ReturnCode string `xml:"return_code"` //返回码，通讯标识
	ReturnMsg  string `xml:"return_msg"`  //通讯错误码
}

//DoResponse 支付回包
func (mc *NotifyPayController) DoResponse(returncode string, returnmsg string) {

	if err := recover(); err != nil {
		beego.Critical("DoResponse Panic catched, err:\n", err, string(debug.Stack()))
		return
	}

	rsp := NotifyPayRsp{
		ReturnCode: returncode,
		ReturnMsg:  returnmsg,
	}
	//序列化
	rspXML, err := xml.Marshal(rsp)
	if err != nil {
		beego.Warn("DoResponse xml.Marshal error, return_code = %s, return_msg = %s, error = %s",
			returncode, returnmsg, err.Error())
		return
	}

	utils.Debug(0, "DoResponse debug, rsp = %v, rspXML = %s", rsp, string(rspXML))

	mc.Ctx.Output.Body(rspXML)

	return
}

//NotifyPayInfo 通知支付信息
func (mc *NotifyPayController) NotifyPayInfo() {
	returnCode := "SUCCESS"
	returnMsg := "OK"

	defer func() {
		mc.DoResponse(returnCode, returnMsg)
	}()

	data := mc.Ctx.Input.RequestBody

	//解析通知数据结构
	stWXPayNotifyInfo := &Pay.WXPayNotifyInfo{}
	err := xml.Unmarshal(data, stWXPayNotifyInfo)
	if err != nil {
		beego.Warn("NotifyPayInfo xml.Unmarshal error, notifyBody = %s, err = %s",
			data, err.Error())
		returnCode = "FAIL"
		returnMsg = "unmarshal error"
		return
	}

	//通信错误
	if stWXPayNotifyInfo.ReturnCode == "SUCCESS" && stWXPayNotifyInfo.ResultCode == "SUCCESS" {
		//更新用户充值金额
		err = Pay.UpdateUserChargeNum(stWXPayNotifyInfo.OpenID, stWXPayNotifyInfo.TotalFee)
		if err != nil {
			beego.Warn("NotifyPayInfo UpdateUserChargeNum error, notifyBody = %s, err = %s",
				data, err.Error())
			return
		}
	}

	//记录流水
	err = Pay.InsertPayFlow(*stWXPayNotifyInfo)
	if err != nil {
		beego.Warn("NotifyPayInfo InsertPayFlow error, notifyBody = %s, err = %s",
			data, err.Error())
		return
	}

	return
}
