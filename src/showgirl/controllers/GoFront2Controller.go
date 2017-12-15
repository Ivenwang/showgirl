package controllers

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/astaxie/beego"
	//"github.com/golang/protobuf/jsonpb"
	//"tiantian/client"
	//"github.com/golang/protobuf/proto"
	//"fmt"
	"showgirl/client"
	"showgirl/models/utils"
	"strings"

	"fmt"

	"net/http"
	"time"

	"github.com/astaxie/beego/httplib"
	"github.com/bitly/go-simplejson"
)

var (
	httpSettingSpecial httplib.BeegoHTTPSettings
	gTransportSepcial  *http.Transport

	gTimeOutSecSpecial = 60
	gMaxIdleConns      = 10
)

type GoFront2Controller struct {
	beego.Controller
}

func (this *GoFront2Controller) Post() {
	port := beego.AppConfig.String("httpport")

	targetUrl := "http://127.0.0.1:" + port + "/v1.0" + strings.Replace(this.Ctx.Request.URL.Path, "/v2.0", "", -1)
	if len(this.Ctx.Request.URL.RawQuery) > 0 {
		targetUrl += "?" + this.Ctx.Request.URL.RawQuery
	}

	request := httplib.Post(targetUrl)
	request.Setting(httpSettingSpecial)
	request.SetTransport(gTransportSepcial)

	// 获取请求的ETag
	reqETag := this.Ctx.Input.Header("If-None-Match")
	//透传所有HTTP HEADER
	for key, value := range this.Ctx.Request.Header {
		request.Header(key, value[0])
	}
	request.Body(this.Ctx.Input.RequestBody)

	utils.Debug(0, "GoFront2Controller body debug, path = %s, body = %s",
		this.Ctx.Request.URL.Path, this.Ctx.Input.RequestBody)

	ret, err := request.Bytes()
	if err != nil {
		utils.Warn(0, targetUrl+"\t "+err.Error())
		return
	}

	json, err := simplejson.NewJson(ret)
	if err != nil {
		utils.Warn(0, targetUrl+"\t "+err.Error())
		return
	}

	//utils.Debug(0, "GoFront2Controller etag req = %s, targetUrl = %s", reqETag, targetUrl)

	var etag string
	bytes, err := json.Get("RspJson").Bytes()
	if err == nil && bytes != nil {
		newBody, _ := simplejson.NewJson(bytes)
		json.Set("RspJson", newBody)

		//1、由于flowid每次都不一样，所以不可以使用整体body做etag
		//2、由于不同错误码的时候，RspJson有时是一样的，所以不能只使用RspJson做etag
		//使用RspHeader.ErrNo + RspJson做etag
		headerErrNo := json.Get("RspHeader").Get("ErrNo")
		BodyBytes := fmt.Sprintf("%v%v", headerErrNo, bytes)

		//设置Errno自定义HttpHeader
		flowId := json.Get("RspHeader").Get("FlowId")
		if headerErrNo != nil {
			this.Ctx.Output.Header("Errno", fmt.Sprintf("%v", headerErrNo.Interface()))
		}
		if flowId != nil {
			this.Ctx.Output.Header("FlowId", fmt.Sprintf("%v", flowId.Interface()))
		}

		//utils.Debug(0, "GoFront2Controller etag debug, BodyBytes = %s, targetUrl = %s", BodyBytes, targetUrl)

		// 计算ETag
		sum := md5.Sum([]byte(BodyBytes))
		etag = hex.EncodeToString(sum[:])

		// change to weak etag
		// etag = "W/" + etag

		// 判断返回的code，只有等于200时才判断ETag是否相同
		retCode := json.Get("RspHeader").Get("ErrNo").MustInt()
		if int32(retCode) == int32(client.EErrorTypeDef_RESULT_OK) && etag == reqETag {
			// 判断ETag是否相同，相同直接返回304
			this.Ctx.ResponseWriter.WriteHeader(304)
			return
		}
	}

	//utils.Debug(0, "GoFront2Controller etag rsp = %s, targetUrl = %s", etag, targetUrl)

	// 设置ETag
	this.Ctx.Output.Header("ETag", etag)

	this.Data["json"] = json
	this.ServeJSON()
}

func init() {
	conns, _ := beego.AppConfig.Int("CommonConfig::MaxIdleConnsPerHost")
	timeoutex, _ := beego.AppConfig.Int("CommonConfig::HttpTimeOutSpecial")

	if conns > 0 {
		gMaxIdleConns = conns
	}

	if timeoutex > 0 {
		gTimeOutSecSpecial = timeoutex
	}

	httpSettingSpecial = httplib.BeegoHTTPSettings{
		ConnectTimeout:   time.Duration(gTimeOutSecSpecial) * time.Second,
		ReadWriteTimeout: time.Duration(gTimeOutSecSpecial) * time.Second,
	}

	gTransportSepcial = &http.Transport{
		//连接池最大的Idle连接数
		MaxIdleConnsPerHost: gMaxIdleConns,
		//请求超时设置
		ResponseHeaderTimeout: time.Duration(gTimeOutSecSpecial) * time.Second,
		//连接超时设置，不设置连接的deadline
		Dial: client.TimeoutDialer(time.Duration(gTimeOutSecSpecial) * time.Second),
	}

}
