package controllers

import (
	//"github.com/astaxie/beego"
	"showgirl/client"
	"showgirl/models/Account"
	"showgirl/models/Session"
	"showgirl/models/utils"

	"github.com/golang/protobuf/proto"
)

//auto generated.
func HandleThirdPartyWXLogin(hdr *client.STUserTrustInfo, req *client.STThirdPartyWXLoginReq) (*client.STRspHeader, *client.STThirdPartyWXLoginRsp) {

	rspHeader := &client.STRspHeader{
		ErrNo:  client.EErrorTypeDef_RESULT_OK.Enum(),
		ErrMsg: proto.String("success"),
	}

	stThirdPartyWXLoginRsp := &client.STThirdPartyWXLoginRsp{}

	var pThirdPartyWXLoginRsp *client.STThirdPartyWXLoginRsp

	for {

		//根据code获取access_token
		AccessToken, WXOpenID, WxUnionID, err := Account.GetAccessTokenByWxKey(req.GetCode(), hdr.GetFlowId())
		if err != nil {
			utils.Warn(hdr.GetFlowId(), "HandleThirdPartyWXLogin GetAccessTokenByWxKey error, code = %s, error = %s",
				req.GetCode(), err.Error())
			rspHeader.ErrNo = client.EErrorTypeDef_CHECK_PARAM_ERROR.Enum()
			rspHeader.ErrMsg = proto.String("微信登录失败")
			break
		}

		if len(AccessToken) <= 0 || len(WXOpenID) <= 0 || len(WxUnionID) <= 0 {
			utils.Warn(hdr.GetFlowId(), "HandleThirdPartyWXLogin GetAccessTokenByWxKey error, code = %s, AccessToken = %s, WxUnionId = %s",
				req.GetCode(), AccessToken, WxUnionID)
			rspHeader.ErrNo = client.EErrorTypeDef_CHECK_PARAM_ERROR.Enum()
			rspHeader.ErrMsg = proto.String("校验id错误，请在微信开放平台绑定应用")
			break
		}

		stUserBaseInfo, FindResult := Account.QueryAccountCorrectByUnionID(WxUnionID, hdr.GetFlowId())

		if FindResult == client.EFindAccountDef_ACCOUNT_SYSTEM_ERROR ||
			FindResult == client.EFindAccountDef_ACCOUNT_PARAM_ERROR {
			//DB错误
			utils.Warn(hdr.GetFlowId(), "HandleThirdPartyWXLogin QueryAccountCorrectByUnionID error, FindResult = %d", int32(FindResult))
			rspHeader.ErrNo = client.EErrorTypeDef_SYS_INTERNAL_ERROR.Enum()
			rspHeader.ErrMsg = proto.String("系统繁忙，请稍候重试")
			break
		}

		//如果没找到账户信息
		if FindResult == client.EFindAccountDef_NOT_FOUND_ACCOUNT || len(stUserBaseInfo.WxOpenID) <= 0 ||
			len(stUserBaseInfo.WxUnionID) <= 0 {

			//查询用户微信userinfo，插入db
			stUserBaseInfo, err = Account.QueryAndSetUserWxInfo(WXOpenID, WxUnionID, hdr.GetFlowId())
			if err != nil {
				utils.Warn(hdr.GetFlowId(), "HandleThirdPartyWXLogin QueryAndSetUserWxInfo error, access_token = %s, openid = %s, error = %s", AccessToken, WXOpenID, err.Error())
				rspHeader.ErrNo = client.EErrorTypeDef_SYS_INTERNAL_ERROR.Enum()
				rspHeader.ErrMsg = proto.String("获取用户信息失败，请退出重试")
				break
			}
		}

		//更新用户最后登录时间
		err = Account.UpdateLastLoginTime(WxUnionID, hdr.GetFlowId())
		if err != nil {
			utils.Warn(hdr.GetFlowId(), "HandleThirdPartyWXLogin UpdateLastLoginTime error, WxUnionID = %s, err = %s", WxUnionID, err.Error())
			rspHeader.ErrNo = client.EErrorTypeDef_SYS_INTERNAL_ERROR.Enum()
			rspHeader.ErrMsg = proto.String("系统繁忙，请稍后重试")
			break
		}

		//生成token
		//创建用户cookie
		CookieData, err := Session.CreateUserCookie(WXOpenID, WxUnionID, stUserBaseInfo.NickName, stUserBaseInfo.URL, hdr.GetFlowId())
		if err != nil {
			utils.Warn(hdr.GetFlowId(), "HandleThirdPartyWXLogin CreateUserCookie error, WXOpenID = %s, WXUnionID = %s, NickName = %s, URL = %s, err = %s",
				WXOpenID, WxUnionID, stUserBaseInfo.NickName, stUserBaseInfo.URL, err.Error())
			rspHeader.ErrNo = client.EErrorTypeDef_SYS_INTERNAL_ERROR.Enum()
			rspHeader.ErrMsg = proto.String("系统繁忙，请稍后重试")
			break
		}

		//构造回包结构
		stThirdPartyWXLoginRsp.UserKey = proto.String(CookieData)
		stThirdPartyWXLoginRsp.UserId = proto.String(WxUnionID)
		stThirdPartyWXLoginRsp.NickName = proto.String(stUserBaseInfo.NickName)
		stThirdPartyWXLoginRsp.Url = proto.String(stUserBaseInfo.URL)
		stThirdPartyWXLoginRsp.LastTime = proto.Int64(stUserBaseInfo.LastTime)
		stThirdPartyWXLoginRsp.ChargeNum = proto.Int32(stUserBaseInfo.Charge)
		stThirdPartyWXLoginRsp.WxOpenID = proto.String(WXOpenID)

		pThirdPartyWXLoginRsp = stThirdPartyWXLoginRsp

		break
	}

	return rspHeader, pThirdPartyWXLoginRsp

}

//auto generated.
func HandleGetMyInfo(hdr *client.STUserTrustInfo, req *client.STGetMyInfoReq) (*client.STRspHeader, *client.STGetMyInfoRsp) {

	rspHeader := &client.STRspHeader{
		ErrNo:  client.EErrorTypeDef_RESULT_OK.Enum(),
		ErrMsg: proto.String("success"),
	}

	stGetMyInfoRsp := &client.STGetMyInfoRsp{}

	var pGetMyInfoRsp *client.STGetMyInfoRsp

	for {
		stUserBaseInfo, FindResult := Account.QueryAccountCorrectByUnionID(hdr.GetUserID(), hdr.GetFlowId())

		if FindResult == client.EFindAccountDef_ACCOUNT_SYSTEM_ERROR ||
			FindResult == client.EFindAccountDef_ACCOUNT_PARAM_ERROR {
			//DB错误
			utils.Warn(hdr.GetFlowId(), "HandleThirdPartyWXLogin QueryAccountCorrectByUnionID error, FindResult = %d", int32(FindResult))
			rspHeader.ErrNo = client.EErrorTypeDef_SYS_INTERNAL_ERROR.Enum()
			rspHeader.ErrMsg = proto.String("系统繁忙，请稍候重试")
			break
		}

		//如果没找到账户信息
		if FindResult == client.EFindAccountDef_NOT_FOUND_ACCOUNT || len(stUserBaseInfo.WxOpenID) <= 0 ||
			len(stUserBaseInfo.WxUnionID) <= 0 {
			//没有找到数据
			utils.Warn(hdr.GetFlowId(), "HandleThirdPartyWXLogin QueryAccountCorrectByUnionID error, FindResult = %d", int32(FindResult))
			rspHeader.ErrNo = client.EErrorTypeDef_SYS_INTERNAL_ERROR.Enum()
			rspHeader.ErrMsg = proto.String("没有数据")
			break
		}

		//构造回包结构
		stGetMyInfoRsp.UserId = proto.String(hdr.GetUserID())
		stGetMyInfoRsp.NickName = proto.String(stUserBaseInfo.NickName)
		stGetMyInfoRsp.Url = proto.String(stUserBaseInfo.URL)
		stGetMyInfoRsp.LastTime = proto.Int64(stUserBaseInfo.LastTime)
		stGetMyInfoRsp.ChargeNum = proto.Int32(stUserBaseInfo.Charge)

		pGetMyInfoRsp = stGetMyInfoRsp

		break
	}

	return rspHeader, pGetMyInfoRsp

}
