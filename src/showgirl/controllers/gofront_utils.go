package controllers

import (
	//	"bytes"
	//	"compress/gzip"

	//	"encoding/base64"

	"fmt"

	//	"io/ioutil"
	"math/rand"
	"runtime/debug"
	"showgirl/client"
	"showgirl/models/Session"
	"showgirl/models/utils"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	//	dproxy "github.com/koron/go-dproxy"
)

var G_FlowRand *rand.Rand
var gMaxHttpBodyLength = int(8192)

type CookieUserInfo struct {
	UserKey string `json:"UserKey"`
}

//@breif 获取cookie信息
//return:0表示成功，1表示未找到cookie, 2表示cookie时间过期, 3:数据异常
func getCookieData(path string, cookie string, flowid int64) (*client.STUserTrustInfo, int) {

	//解析cookie
	stUserCookieInfo, err := Session.UnmarshalCookie(cookie, flowid)
	if err != nil {
		if err.Error() == "cookie is empty" {
			return nil, 1
		}

		utils.Warn(flowid, "getCookieData cookie = %s, err = %s", cookie, err.Error())

		return nil, 3
	}

	//校验时间
	CurTime := time.Now().Unix()
	if CurTime > (stUserCookieInfo.CurTime + int64(utils.GetConfigByInt("CommonConfig::CookieActiveTime"))) {
		//cookie已过期
		return nil, 2
	}

	stUserStrustInfo := &client.STUserTrustInfo{
		UserID: proto.String(stUserCookieInfo.WXUnionID),
		Url:    proto.String(stUserCookieInfo.Url),
		Name:   proto.String(stUserCookieInfo.NickName),
		FlowId: proto.Int64(flowid),
	}

	utils.Debug(flowid, "getCookieData debug, path = %s, cookie = %s, UserStrustInfo = %s",
		path, cookie, stUserStrustInfo.String())

	return stUserStrustInfo, 0
}

type GoFrontController struct {
	beego.Controller
}

func (this *GoFrontController) DoResponse(jsonMarshaler *jsonpb.Marshaler, clientRspPB *client.CommonClientRsp, FlowIdHeader int64) {
	if err := recover(); err != nil {
		beego.Critical("Panic catched, err:\n", err, string(debug.Stack()))
		clientRspPB.RspHeader = &client.STRspHeader{
			ErrNo:     client.EErrorTypeDef_PROGRAM_EXCEPTION_ERROR.Enum(),
			ErrDetail: proto.String("Panic catched"),
			ErrMsg:    proto.String(utils.ERRMSG_EXCEPTION_CATCHED),
		}
	}

	flowid := int64(FlowIdHeader)
	clientRspPB.RspHeader.FlowId = proto.Int64(FlowIdHeader)
	clientRspJson, err := jsonMarshaler.MarshalToString(clientRspPB)
	if err != nil {
		errStr := fmt.Sprintf("pb2json failed, err:%s", err)
		utils.Warn(flowid, "%s", errStr)
		errNo := *clientRspPB.RspHeader.ErrNo
		errMsg := *clientRspPB.RspHeader.ErrMsg
		if *clientRspPB.RspHeader.ErrNo == client.EErrorTypeDef_RESULT_OK {
			errNo = client.EErrorTypeDef_GENERATE_CONTENT_ERROR
			errMsg = errStr
		}
		jsonString := fmt.Sprintf("{\"RspHeader\":{\"ErrNo\":%d, \"ErrMsg\":\"%s\"}}", errNo, errMsg)
		this.Ctx.Output.Header("FlowId", strconv.FormatInt(FlowIdHeader, 10))
		this.Ctx.Output.Header("Errno", fmt.Sprintf("%d", errNo))
		this.Ctx.WriteString(jsonString)
		return
	}

	//返回json字符串
	this.Ctx.Output.Header("FlowId", strconv.FormatInt(FlowIdHeader, 10))
	this.Ctx.Output.Header("Errno", fmt.Sprintf("%d", clientRspPB.RspHeader.GetErrNo()))
	this.Ctx.WriteString(clientRspJson)
	return
}

func init() {
	//生成随机数种子
	G_FlowRand = rand.New(rand.NewSource(time.Now().UnixNano()))
}
