/*
 * Auto generated by code_generator
 * Please do not modify it.
 */
package controllers

import (
	//"bytes"
	//"compress/gzip"
	//"crypto/md5"
	//"encoding/base64"
	//"encoding/hex"
	"encoding/json"
	//"crypto/hmac"
	//"crypto/sha1"
	"fmt"
	//"github.com/astaxie/beego"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	//dproxy "github.com/koron/go-dproxy"
	//"io/ioutil"
	//"math/rand"
	"showgirl/models/utils"
	//"time"
	"showgirl/client"
	//"showgirl/models/redis"
	//"runtime/debug"
	"strconv"
	"strings"
	"net/url"
)




func (this *GoFrontController) Account_ThirdPartyWXLogin() {
	jsonMarshaler := &jsonpb.Marshaler{
		EnumsAsInts: true,  //整数是否整形显示		
		EmitDefaults: true, //是否显示值为0的字段		
		OrigName: false,    //是否显示proto名字
	}

	clientRspPB := &client.CommonClientRsp {
		RspHeader: &client.STRspHeader {
			ErrNo:  client.EErrorTypeDef_RESULT_OK.Enum(),
			ErrMsg: proto.String("success"),
		},
		RspJson: nil,
	}
	FlowIdHeader, _ := strconv.ParseInt(this.Ctx.Input.Header("FlowId"), 10, 64)
	if FlowIdHeader == 0 {
		FlowIdHeader = int64(G_FlowRand.Int31())
	}
	flowid := int64(FlowIdHeader)
	
	defer func() {
		this.DoResponse(jsonMarshaler, clientRspPB, FlowIdHeader)
	}()

	for {
		body := this.Ctx.Input.RequestBody
		bodyLen := len(body)
		utils.Debug(flowid, "GoFront Request ThirdPartyWXLogin bodyLen: %d", bodyLen)
		if bodyLen == 0 {
			utils.Warn(flowid, "GoFrontController post check failed, body is empty")
			clientRspPB.RspHeader.ErrNo = client.EErrorTypeDef_CHECK_CONTENT_ERROR.Enum()
			clientRspPB.RspHeader.ErrMsg = proto.String(utils.ERRMSG_CLIENT_EXCEPTION)
			break
		}
		
		//构造空cookie信息
		TrustInfo := &client.STUserTrustInfo{
			UserID: proto.String(""),
			Url: proto.String(""),
			Name: proto.String(""),
			FlowId: proto.Int64(flowid),
		}
		
		//打印请求json
		utils.Debug(flowid, "Account_ThirdPartyWXLogin req json data = %s", body)
		
		//JSON转PB
		reqPB := &client.STThirdPartyWXLoginReq{}
		err := jsonpb.UnmarshalString(string(body[:len(body)]), reqPB)
		if err != nil {
			errStr := fmt.Sprintf("parse body to STThirdPartyWXLoginReq failed, err:%s", err)
			utils.Warn(flowid, "%s", errStr)
			//example:unknown field "AdditionalReq" in STQueryUserAttrReq
			//if !(strings.HasPrefix(err.Error(), "unknown field") && 
			//	strings.HasSuffix(err.Error(), "in client.STThirdPartyWXLoginReq")) {
			if !(strings.HasPrefix(err.Error(), "unknown field")) {
				clientRspPB.RspHeader.ErrNo = client.EErrorTypeDef_CHECK_CONTENT_ERROR.Enum()
				clientRspPB.RspHeader.ErrMsg = proto.String(utils.ERRMSG_CLIENT_EXCEPTION)
				clientRspPB.RspHeader.ErrDetail = proto.String(errStr)
				break
			}
		}

		//设置httpheader
		header := make(map[string]string)
		header["FlowId"] = this.Ctx.Input.Header("FlowId")
		
		utils.Debug(int64(FlowIdHeader), "GoFront Transmit ThirdPartyWXLogin TrustInfo: %s | reqPB: %s | reqJson: %s", 
			TrustInfo.String(), reqPB.String(), body)

		c :=  client.AccountClient{}
		rspHdr, rspBody, err := c.ThirdPartyWXLogin(TrustInfo, reqPB, header)
		if err != nil {
			if strings.HasPrefix(err.Error(), client.RPC_UNMARSHAL_ABNORMAL_PREFIX) ||
			 	strings.HasPrefix(err.Error(), client.RPC_MARSHAL_ABNORMAL_PREFIX) {
				clientRspPB.RspHeader.ErrNo = client.EErrorTypeDef_RPC_INTERFACE_ABNORMAL.Enum()
				clientRspPB.RspHeader.ErrMsg = proto.String(utils.ERRMSG_EXCEPTION_CATCHED)
			} else {
				clientRspPB.RspHeader.ErrNo = client.EErrorTypeDef_RPC_FAILED_ERROR.Enum()
				clientRspPB.RspHeader.ErrMsg = proto.String(utils.ERRMSG_SYSTEM_BUSY)
			}
			errStr := fmt.Sprintf("rpc ThirdPartyWXLogin failed, err:%s", err)
			utils.Warn(flowid, "%s", errStr)
			clientRspPB.RspHeader.ErrDetail = proto.String(errStr)
			break
		}

		//rspBody转json
		rspBodyJson, err := jsonMarshaler.MarshalToString(rspBody)
		if err != nil {
			errStr := fmt.Sprintf("pb2json failed, err:%s", err)
			utils.Warn(flowid, "%s", errStr)
			clientRspPB.RspHeader.ErrNo = client.EErrorTypeDef_GENERATE_CONTENT_ERROR.Enum()
			clientRspPB.RspHeader.ErrMsg = proto.String(utils.ERRMSG_EXCEPTION_CATCHED)
			clientRspPB.RspHeader.ErrDetail = proto.String(errStr)
			break
		}

		//转json
		clientRspPB.RspJson = &rspBodyJson
		clientRspPB.RspHeader = rspHdr
		
		utils.Debug(int64(FlowIdHeader), "GoFront Transmit rsp ThirdPartyWXLogin header:%s | rspPB:%s", rspHdr.String(), rspBodyJson)
		
		break
	}
}

func (this *GoFrontController) Account_GetMyInfo() {
	jsonMarshaler := &jsonpb.Marshaler{
		EnumsAsInts: true,  //整数是否整形显示		
		EmitDefaults: true, //是否显示值为0的字段		
		OrigName: false,    //是否显示proto名字
	}

	clientRspPB := &client.CommonClientRsp {
		RspHeader: &client.STRspHeader {
			ErrNo:  client.EErrorTypeDef_RESULT_OK.Enum(),
			ErrMsg: proto.String("success"),
		},
		RspJson: nil,
	}
	FlowIdHeader, _ := strconv.ParseInt(this.Ctx.Input.Header("FlowId"), 10, 64)
	if FlowIdHeader == 0 {
		FlowIdHeader = int64(G_FlowRand.Int31())
	}
	flowid := int64(FlowIdHeader)
	
	defer func() {
		this.DoResponse(jsonMarshaler, clientRspPB, FlowIdHeader)
	}()

	for {
		body := this.Ctx.Input.RequestBody
		bodyLen := len(body)
		utils.Debug(flowid, "GoFront Request GetMyInfo bodyLen: %d", bodyLen)
		if bodyLen == 0 {
			utils.Warn(flowid, "GoFrontController post check failed, body is empty")
			clientRspPB.RspHeader.ErrNo = client.EErrorTypeDef_CHECK_CONTENT_ERROR.Enum()
			clientRspPB.RspHeader.ErrMsg = proto.String(utils.ERRMSG_CLIENT_EXCEPTION)
			break
		}
		
		//普通Public接口，取出并验证cookie
		cookie := this.Ctx.Input.Header("UserKey")
		//如果从userkey没找到cookie,就从cookie里拿
		if len(cookie) == 0 {
			TmpUserJson := this.Ctx.GetCookie("UserInfo")
			if len(TmpUserJson) > 0 {
				TmpUserInfo, _ := url.QueryUnescape(TmpUserJson)
				stCookieUserInfo := &CookieUserInfo{}
				err := json.Unmarshal([]byte(TmpUserInfo), stCookieUserInfo)
				if err != nil {
					utils.Warn(flowid, "Cookie.UserInfo_Unmarshal_error, method=GetMyInfo  UserInfo=%v, err=%s", TmpUserInfo, err.Error())
				}
				utils.Debug(flowid, "GoFront Request GetMyInfo debug, CookieUserInfo = %v, TmpUserInfo = %s", stCookieUserInfo, TmpUserInfo)
				cookie = stCookieUserInfo.UserKey
			}
		}
		
		utils.Debug(flowid, "go front get cookie, cookie = %s", cookie)
		//redis中取cookie对应的信息
		TrustInfo, ret := getCookieData(this.Ctx.Request.URL.RequestURI(), cookie, flowid)
		
		if ret == 1 || ret == 2 {
			utils.Warn(flowid, "Check_cookie_failed, cookie=%s, ret=%d", cookie, ret)
			clientRspPB.RspHeader.ErrNo = client.EErrorTypeDef_CHECK_COOKIE_ERROR.Enum()
			clientRspPB.RspHeader.ErrMsg = proto.String(utils.ERRMSG_CHECK_COOKIE_ERROR)
			break
		} else if ret == 3 {
			utils.Warn(flowid, "getCookieData_from_redis_failed, cookie=%s, ret=%d", cookie, ret)
			clientRspPB.RspHeader.ErrNo = client.EErrorTypeDef_SYS_INTERNAL_ERROR.Enum()
			clientRspPB.RspHeader.ErrMsg = proto.String(utils.ERRMSG_SYSTEM_BUSY)
			clientRspPB.RspHeader.ErrDetail = proto.String("getCookieData from redis failed")
			break
		}
		
		channel := this.Ctx.Input.Header("Channel")
		utils.Debug(flowid, "stat_channel UserID = %s, Url = %s, NickName = %s", 
			channel, TrustInfo.GetUserID(), TrustInfo.GetUrl(), TrustInfo.GetName())
		
		//打印请求json
		utils.Debug(flowid, "Account_GetMyInfo req json data = %s", body)
		
		//JSON转PB
		reqPB := &client.STGetMyInfoReq{}
		err := jsonpb.UnmarshalString(string(body[:len(body)]), reqPB)
		if err != nil {
			errStr := fmt.Sprintf("parse body to STGetMyInfoReq failed, err:%s", err)
			utils.Warn(flowid, "%s", errStr)
			//example:unknown field "AdditionalReq" in STQueryUserAttrReq
			//if !(strings.HasPrefix(err.Error(), "unknown field") && 
			//	strings.HasSuffix(err.Error(), "in client.STGetMyInfoReq")) {
			if !(strings.HasPrefix(err.Error(), "unknown field")) {
				clientRspPB.RspHeader.ErrNo = client.EErrorTypeDef_CHECK_CONTENT_ERROR.Enum()
				clientRspPB.RspHeader.ErrMsg = proto.String(utils.ERRMSG_CLIENT_EXCEPTION)
				clientRspPB.RspHeader.ErrDetail = proto.String(errStr)
				break
			}
		}

		//设置httpheader
		header := make(map[string]string)
		header["FlowId"] = this.Ctx.Input.Header("FlowId")
		
		utils.Debug(int64(FlowIdHeader), "GoFront Transmit GetMyInfo TrustInfo: %s | reqPB: %s | reqJson: %s", 
			TrustInfo.String(), reqPB.String(), body)

		c :=  client.AccountClient{}
		rspHdr, rspBody, err := c.GetMyInfo(TrustInfo, reqPB, header)
		if err != nil {
			if strings.HasPrefix(err.Error(), client.RPC_UNMARSHAL_ABNORMAL_PREFIX) ||
			 	strings.HasPrefix(err.Error(), client.RPC_MARSHAL_ABNORMAL_PREFIX) {
				clientRspPB.RspHeader.ErrNo = client.EErrorTypeDef_RPC_INTERFACE_ABNORMAL.Enum()
				clientRspPB.RspHeader.ErrMsg = proto.String(utils.ERRMSG_EXCEPTION_CATCHED)
			} else {
				clientRspPB.RspHeader.ErrNo = client.EErrorTypeDef_RPC_FAILED_ERROR.Enum()
				clientRspPB.RspHeader.ErrMsg = proto.String(utils.ERRMSG_SYSTEM_BUSY)
			}
			errStr := fmt.Sprintf("rpc GetMyInfo failed, err:%s", err)
			utils.Warn(flowid, "%s", errStr)
			clientRspPB.RspHeader.ErrDetail = proto.String(errStr)
			break
		}

		//rspBody转json
		rspBodyJson, err := jsonMarshaler.MarshalToString(rspBody)
		if err != nil {
			errStr := fmt.Sprintf("pb2json failed, err:%s", err)
			utils.Warn(flowid, "%s", errStr)
			clientRspPB.RspHeader.ErrNo = client.EErrorTypeDef_GENERATE_CONTENT_ERROR.Enum()
			clientRspPB.RspHeader.ErrMsg = proto.String(utils.ERRMSG_EXCEPTION_CATCHED)
			clientRspPB.RspHeader.ErrDetail = proto.String(errStr)
			break
		}

		//转json
		clientRspPB.RspJson = &rspBodyJson
		clientRspPB.RspHeader = rspHdr
		
		utils.Debug(int64(FlowIdHeader), "GoFront Transmit rsp GetMyInfo header:%s | rspPB:%s", rspHdr.String(), rspBodyJson)
		
		break
	}
}



func (this *GoFrontController) Pay_CreateTransaction() {
	jsonMarshaler := &jsonpb.Marshaler{
		EnumsAsInts: true,  //整数是否整形显示		
		EmitDefaults: true, //是否显示值为0的字段		
		OrigName: false,    //是否显示proto名字
	}

	clientRspPB := &client.CommonClientRsp {
		RspHeader: &client.STRspHeader {
			ErrNo:  client.EErrorTypeDef_RESULT_OK.Enum(),
			ErrMsg: proto.String("success"),
		},
		RspJson: nil,
	}
	FlowIdHeader, _ := strconv.ParseInt(this.Ctx.Input.Header("FlowId"), 10, 64)
	if FlowIdHeader == 0 {
		FlowIdHeader = int64(G_FlowRand.Int31())
	}
	flowid := int64(FlowIdHeader)
	
	defer func() {
		this.DoResponse(jsonMarshaler, clientRspPB, FlowIdHeader)
	}()

	for {
		body := this.Ctx.Input.RequestBody
		bodyLen := len(body)
		utils.Debug(flowid, "GoFront Request CreateTransaction bodyLen: %d", bodyLen)
		if bodyLen == 0 {
			utils.Warn(flowid, "GoFrontController post check failed, body is empty")
			clientRspPB.RspHeader.ErrNo = client.EErrorTypeDef_CHECK_CONTENT_ERROR.Enum()
			clientRspPB.RspHeader.ErrMsg = proto.String(utils.ERRMSG_CLIENT_EXCEPTION)
			break
		}
		
		//普通Public接口，取出并验证cookie
		cookie := this.Ctx.Input.Header("UserKey")
		//如果从userkey没找到cookie,就从cookie里拿
		if len(cookie) == 0 {
			TmpUserJson := this.Ctx.GetCookie("UserInfo")
			if len(TmpUserJson) > 0 {
				TmpUserInfo, _ := url.QueryUnescape(TmpUserJson)
				stCookieUserInfo := &CookieUserInfo{}
				err := json.Unmarshal([]byte(TmpUserInfo), stCookieUserInfo)
				if err != nil {
					utils.Warn(flowid, "Cookie.UserInfo_Unmarshal_error, method=CreateTransaction  UserInfo=%v, err=%s", TmpUserInfo, err.Error())
				}
				utils.Debug(flowid, "GoFront Request CreateTransaction debug, CookieUserInfo = %v, TmpUserInfo = %s", stCookieUserInfo, TmpUserInfo)
				cookie = stCookieUserInfo.UserKey
			}
		}
		
		utils.Debug(flowid, "go front get cookie, cookie = %s", cookie)
		//redis中取cookie对应的信息
		TrustInfo, ret := getCookieData(this.Ctx.Request.URL.RequestURI(), cookie, flowid)
		
		if ret == 1 || ret == 2 {
			utils.Warn(flowid, "Check_cookie_failed, cookie=%s, ret=%d", cookie, ret)
			clientRspPB.RspHeader.ErrNo = client.EErrorTypeDef_CHECK_COOKIE_ERROR.Enum()
			clientRspPB.RspHeader.ErrMsg = proto.String(utils.ERRMSG_CHECK_COOKIE_ERROR)
			break
		} else if ret == 3 {
			utils.Warn(flowid, "getCookieData_from_redis_failed, cookie=%s, ret=%d", cookie, ret)
			clientRspPB.RspHeader.ErrNo = client.EErrorTypeDef_SYS_INTERNAL_ERROR.Enum()
			clientRspPB.RspHeader.ErrMsg = proto.String(utils.ERRMSG_SYSTEM_BUSY)
			clientRspPB.RspHeader.ErrDetail = proto.String("getCookieData from redis failed")
			break
		}
		
		channel := this.Ctx.Input.Header("Channel")
		utils.Debug(flowid, "stat_channel UserID = %s, Url = %s, NickName = %s", 
			channel, TrustInfo.GetUserID(), TrustInfo.GetUrl(), TrustInfo.GetName())
		
		//打印请求json
		utils.Debug(flowid, "Pay_CreateTransaction req json data = %s", body)
		
		//JSON转PB
		reqPB := &client.STCreateTransactionReq{}
		err := jsonpb.UnmarshalString(string(body[:len(body)]), reqPB)
		if err != nil {
			errStr := fmt.Sprintf("parse body to STCreateTransactionReq failed, err:%s", err)
			utils.Warn(flowid, "%s", errStr)
			//example:unknown field "AdditionalReq" in STQueryUserAttrReq
			//if !(strings.HasPrefix(err.Error(), "unknown field") && 
			//	strings.HasSuffix(err.Error(), "in client.STCreateTransactionReq")) {
			if !(strings.HasPrefix(err.Error(), "unknown field")) {
				clientRspPB.RspHeader.ErrNo = client.EErrorTypeDef_CHECK_CONTENT_ERROR.Enum()
				clientRspPB.RspHeader.ErrMsg = proto.String(utils.ERRMSG_CLIENT_EXCEPTION)
				clientRspPB.RspHeader.ErrDetail = proto.String(errStr)
				break
			}
		}

		//设置httpheader
		header := make(map[string]string)
		header["FlowId"] = this.Ctx.Input.Header("FlowId")
		
		utils.Debug(int64(FlowIdHeader), "GoFront Transmit CreateTransaction TrustInfo: %s | reqPB: %s | reqJson: %s", 
			TrustInfo.String(), reqPB.String(), body)

		c :=  client.PayClient{}
		rspHdr, rspBody, err := c.CreateTransaction(TrustInfo, reqPB, header)
		if err != nil {
			if strings.HasPrefix(err.Error(), client.RPC_UNMARSHAL_ABNORMAL_PREFIX) ||
			 	strings.HasPrefix(err.Error(), client.RPC_MARSHAL_ABNORMAL_PREFIX) {
				clientRspPB.RspHeader.ErrNo = client.EErrorTypeDef_RPC_INTERFACE_ABNORMAL.Enum()
				clientRspPB.RspHeader.ErrMsg = proto.String(utils.ERRMSG_EXCEPTION_CATCHED)
			} else {
				clientRspPB.RspHeader.ErrNo = client.EErrorTypeDef_RPC_FAILED_ERROR.Enum()
				clientRspPB.RspHeader.ErrMsg = proto.String(utils.ERRMSG_SYSTEM_BUSY)
			}
			errStr := fmt.Sprintf("rpc CreateTransaction failed, err:%s", err)
			utils.Warn(flowid, "%s", errStr)
			clientRspPB.RspHeader.ErrDetail = proto.String(errStr)
			break
		}

		//rspBody转json
		rspBodyJson, err := jsonMarshaler.MarshalToString(rspBody)
		if err != nil {
			errStr := fmt.Sprintf("pb2json failed, err:%s", err)
			utils.Warn(flowid, "%s", errStr)
			clientRspPB.RspHeader.ErrNo = client.EErrorTypeDef_GENERATE_CONTENT_ERROR.Enum()
			clientRspPB.RspHeader.ErrMsg = proto.String(utils.ERRMSG_EXCEPTION_CATCHED)
			clientRspPB.RspHeader.ErrDetail = proto.String(errStr)
			break
		}

		//转json
		clientRspPB.RspJson = &rspBodyJson
		clientRspPB.RspHeader = rspHdr
		
		utils.Debug(int64(FlowIdHeader), "GoFront Transmit rsp CreateTransaction header:%s | rspPB:%s", rspHdr.String(), rspBodyJson)
		
		break
	}
}



func (this *GoFrontController) Recommend_QueryRecommendList() {
	jsonMarshaler := &jsonpb.Marshaler{
		EnumsAsInts: true,  //整数是否整形显示		
		EmitDefaults: true, //是否显示值为0的字段		
		OrigName: false,    //是否显示proto名字
	}

	clientRspPB := &client.CommonClientRsp {
		RspHeader: &client.STRspHeader {
			ErrNo:  client.EErrorTypeDef_RESULT_OK.Enum(),
			ErrMsg: proto.String("success"),
		},
		RspJson: nil,
	}
	FlowIdHeader, _ := strconv.ParseInt(this.Ctx.Input.Header("FlowId"), 10, 64)
	if FlowIdHeader == 0 {
		FlowIdHeader = int64(G_FlowRand.Int31())
	}
	flowid := int64(FlowIdHeader)
	
	defer func() {
		this.DoResponse(jsonMarshaler, clientRspPB, FlowIdHeader)
	}()

	for {
		body := this.Ctx.Input.RequestBody
		bodyLen := len(body)
		utils.Debug(flowid, "GoFront Request QueryRecommendList bodyLen: %d", bodyLen)
		if bodyLen == 0 {
			utils.Warn(flowid, "GoFrontController post check failed, body is empty")
			clientRspPB.RspHeader.ErrNo = client.EErrorTypeDef_CHECK_CONTENT_ERROR.Enum()
			clientRspPB.RspHeader.ErrMsg = proto.String(utils.ERRMSG_CLIENT_EXCEPTION)
			break
		}
		
		//构造空cookie信息
		TrustInfo := &client.STUserTrustInfo{
			UserID: proto.String(""),
			Url: proto.String(""),
			Name: proto.String(""),
			FlowId: proto.Int64(flowid),
		}
		
		//打印请求json
		utils.Debug(flowid, "Recommend_QueryRecommendList req json data = %s", body)
		
		//JSON转PB
		reqPB := &client.STQueryRecommendListReq{}
		err := jsonpb.UnmarshalString(string(body[:len(body)]), reqPB)
		if err != nil {
			errStr := fmt.Sprintf("parse body to STQueryRecommendListReq failed, err:%s", err)
			utils.Warn(flowid, "%s", errStr)
			//example:unknown field "AdditionalReq" in STQueryUserAttrReq
			//if !(strings.HasPrefix(err.Error(), "unknown field") && 
			//	strings.HasSuffix(err.Error(), "in client.STQueryRecommendListReq")) {
			if !(strings.HasPrefix(err.Error(), "unknown field")) {
				clientRspPB.RspHeader.ErrNo = client.EErrorTypeDef_CHECK_CONTENT_ERROR.Enum()
				clientRspPB.RspHeader.ErrMsg = proto.String(utils.ERRMSG_CLIENT_EXCEPTION)
				clientRspPB.RspHeader.ErrDetail = proto.String(errStr)
				break
			}
		}

		//设置httpheader
		header := make(map[string]string)
		header["FlowId"] = this.Ctx.Input.Header("FlowId")
		
		utils.Debug(int64(FlowIdHeader), "GoFront Transmit QueryRecommendList TrustInfo: %s | reqPB: %s | reqJson: %s", 
			TrustInfo.String(), reqPB.String(), body)

		c :=  client.RecommendClient{}
		rspHdr, rspBody, err := c.QueryRecommendList(TrustInfo, reqPB, header)
		if err != nil {
			if strings.HasPrefix(err.Error(), client.RPC_UNMARSHAL_ABNORMAL_PREFIX) ||
			 	strings.HasPrefix(err.Error(), client.RPC_MARSHAL_ABNORMAL_PREFIX) {
				clientRspPB.RspHeader.ErrNo = client.EErrorTypeDef_RPC_INTERFACE_ABNORMAL.Enum()
				clientRspPB.RspHeader.ErrMsg = proto.String(utils.ERRMSG_EXCEPTION_CATCHED)
			} else {
				clientRspPB.RspHeader.ErrNo = client.EErrorTypeDef_RPC_FAILED_ERROR.Enum()
				clientRspPB.RspHeader.ErrMsg = proto.String(utils.ERRMSG_SYSTEM_BUSY)
			}
			errStr := fmt.Sprintf("rpc QueryRecommendList failed, err:%s", err)
			utils.Warn(flowid, "%s", errStr)
			clientRspPB.RspHeader.ErrDetail = proto.String(errStr)
			break
		}

		//rspBody转json
		rspBodyJson, err := jsonMarshaler.MarshalToString(rspBody)
		if err != nil {
			errStr := fmt.Sprintf("pb2json failed, err:%s", err)
			utils.Warn(flowid, "%s", errStr)
			clientRspPB.RspHeader.ErrNo = client.EErrorTypeDef_GENERATE_CONTENT_ERROR.Enum()
			clientRspPB.RspHeader.ErrMsg = proto.String(utils.ERRMSG_EXCEPTION_CATCHED)
			clientRspPB.RspHeader.ErrDetail = proto.String(errStr)
			break
		}

		//转json
		clientRspPB.RspJson = &rspBodyJson
		clientRspPB.RspHeader = rspHdr
		
		utils.Debug(int64(FlowIdHeader), "GoFront Transmit rsp QueryRecommendList header:%s | rspPB:%s", rspHdr.String(), rspBodyJson)
		
		break
	}
}

func (this *GoFrontController) Recommend_StyleImageList() {
	jsonMarshaler := &jsonpb.Marshaler{
		EnumsAsInts: true,  //整数是否整形显示		
		EmitDefaults: true, //是否显示值为0的字段		
		OrigName: false,    //是否显示proto名字
	}

	clientRspPB := &client.CommonClientRsp {
		RspHeader: &client.STRspHeader {
			ErrNo:  client.EErrorTypeDef_RESULT_OK.Enum(),
			ErrMsg: proto.String("success"),
		},
		RspJson: nil,
	}
	FlowIdHeader, _ := strconv.ParseInt(this.Ctx.Input.Header("FlowId"), 10, 64)
	if FlowIdHeader == 0 {
		FlowIdHeader = int64(G_FlowRand.Int31())
	}
	flowid := int64(FlowIdHeader)
	
	defer func() {
		this.DoResponse(jsonMarshaler, clientRspPB, FlowIdHeader)
	}()

	for {
		body := this.Ctx.Input.RequestBody
		bodyLen := len(body)
		utils.Debug(flowid, "GoFront Request StyleImageList bodyLen: %d", bodyLen)
		if bodyLen == 0 {
			utils.Warn(flowid, "GoFrontController post check failed, body is empty")
			clientRspPB.RspHeader.ErrNo = client.EErrorTypeDef_CHECK_CONTENT_ERROR.Enum()
			clientRspPB.RspHeader.ErrMsg = proto.String(utils.ERRMSG_CLIENT_EXCEPTION)
			break
		}
		
		//构造空cookie信息
		TrustInfo := &client.STUserTrustInfo{
			UserID: proto.String(""),
			Url: proto.String(""),
			Name: proto.String(""),
			FlowId: proto.Int64(flowid),
		}
		
		//打印请求json
		utils.Debug(flowid, "Recommend_StyleImageList req json data = %s", body)
		
		//JSON转PB
		reqPB := &client.STStyleImageListReq{}
		err := jsonpb.UnmarshalString(string(body[:len(body)]), reqPB)
		if err != nil {
			errStr := fmt.Sprintf("parse body to STStyleImageListReq failed, err:%s", err)
			utils.Warn(flowid, "%s", errStr)
			//example:unknown field "AdditionalReq" in STQueryUserAttrReq
			//if !(strings.HasPrefix(err.Error(), "unknown field") && 
			//	strings.HasSuffix(err.Error(), "in client.STStyleImageListReq")) {
			if !(strings.HasPrefix(err.Error(), "unknown field")) {
				clientRspPB.RspHeader.ErrNo = client.EErrorTypeDef_CHECK_CONTENT_ERROR.Enum()
				clientRspPB.RspHeader.ErrMsg = proto.String(utils.ERRMSG_CLIENT_EXCEPTION)
				clientRspPB.RspHeader.ErrDetail = proto.String(errStr)
				break
			}
		}

		//设置httpheader
		header := make(map[string]string)
		header["FlowId"] = this.Ctx.Input.Header("FlowId")
		
		utils.Debug(int64(FlowIdHeader), "GoFront Transmit StyleImageList TrustInfo: %s | reqPB: %s | reqJson: %s", 
			TrustInfo.String(), reqPB.String(), body)

		c :=  client.RecommendClient{}
		rspHdr, rspBody, err := c.StyleImageList(TrustInfo, reqPB, header)
		if err != nil {
			if strings.HasPrefix(err.Error(), client.RPC_UNMARSHAL_ABNORMAL_PREFIX) ||
			 	strings.HasPrefix(err.Error(), client.RPC_MARSHAL_ABNORMAL_PREFIX) {
				clientRspPB.RspHeader.ErrNo = client.EErrorTypeDef_RPC_INTERFACE_ABNORMAL.Enum()
				clientRspPB.RspHeader.ErrMsg = proto.String(utils.ERRMSG_EXCEPTION_CATCHED)
			} else {
				clientRspPB.RspHeader.ErrNo = client.EErrorTypeDef_RPC_FAILED_ERROR.Enum()
				clientRspPB.RspHeader.ErrMsg = proto.String(utils.ERRMSG_SYSTEM_BUSY)
			}
			errStr := fmt.Sprintf("rpc StyleImageList failed, err:%s", err)
			utils.Warn(flowid, "%s", errStr)
			clientRspPB.RspHeader.ErrDetail = proto.String(errStr)
			break
		}

		//rspBody转json
		rspBodyJson, err := jsonMarshaler.MarshalToString(rspBody)
		if err != nil {
			errStr := fmt.Sprintf("pb2json failed, err:%s", err)
			utils.Warn(flowid, "%s", errStr)
			clientRspPB.RspHeader.ErrNo = client.EErrorTypeDef_GENERATE_CONTENT_ERROR.Enum()
			clientRspPB.RspHeader.ErrMsg = proto.String(utils.ERRMSG_EXCEPTION_CATCHED)
			clientRspPB.RspHeader.ErrDetail = proto.String(errStr)
			break
		}

		//转json
		clientRspPB.RspJson = &rspBodyJson
		clientRspPB.RspHeader = rspHdr
		
		utils.Debug(int64(FlowIdHeader), "GoFront Transmit rsp StyleImageList header:%s | rspPB:%s", rspHdr.String(), rspBodyJson)
		
		break
	}
}


