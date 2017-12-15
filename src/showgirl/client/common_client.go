package client

import (
	"errors"
	"fmt"
	"net"
	"strings"
	"time"

	"net/http"
	"showgirl/models/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
)

const (
	RPC_MARSHAL_ABNORMAL_PREFIX   = "marshaling_error"
	RPC_UNMARSHAL_ABNORMAL_PREFIX = "unmarshaling_error"
	RPC_CALL_FAILED_PREFIX        = "call_error"
)

var (
	httpSetting httplib.BeegoHTTPSettings
	gTransport  *http.Transport

	httpSettingSpecial httplib.BeegoHTTPSettings
	gTransportSepcial  *http.Transport

	gTimeOutSec        = 5
	gTimeOutSecSpecial = 60
	gMaxIdleConns      = 10
	gTimeoutRetry      = 1
	gIdleTimeOutSec    = 300
)

var G_SpecialSvrMap = make(map[string]int)

//Need to optimise
type CommonClient struct {
}

func (this *CommonClient) Call(svr, path string, req []byte, header map[string]string, flowid int64) (body []byte, err error) {
	defer func() {
		if err := recover(); err != nil {
			errmsg := fmt.Sprintf("CommonClient Call panic error, err:%v", err)
			beego.Critical(errmsg)
			body, err = nil, errors.New(errmsg)
		}
	}()

	for i := 0; i < 1; i++ {
		url := beego.AppConfig.String(svr + "Svr::url")
		if url == "" {
			url = beego.AppConfig.String("DefaultSvr::url")
			if url == "" {
				utils.Warn(flowid, "Get config DefaultSvr::url failed.")
				body, err = nil, errors.New("Get config DefaultSvr::url failed.")
				break
			}
		}
		url += path
		url = strings.ToLower(url)

		//超时重试
		for i := 0; i <= gTimeoutRetry; i++ {
			start := time.Now().UnixNano()

			//do http request
			request := httplib.Post(url)

			//设置http头
			for HeaderKey, HeaderValue := range header {
				request.Header(HeaderKey, HeaderValue)
			}

			if _, found := G_SpecialSvrMap[svr]; found {
				request.Setting(httpSettingSpecial)
				request.SetTransport(gTransportSepcial)
			} else {
				request.Setting(httpSetting)
				request.SetTransport(gTransport)
			}

			//request.Setting(httpSetting)
			//request.SetTransport(gTransport)
			request.Body(req)

			body, err = request.Bytes()

			//milliseconds
			needTime := (time.Now().UnixNano() - start) / 1000 / 1000

			if needTime > 100 {
				utils.Info(flowid, "CommonClient call url = %s, take time = %d ms", url, needTime)
			}

			if err != nil {
				utils.Debug(flowid, "CommonClient call failed. times = %d, url = %s, len(body) = %d, err = %s, take time = %d ms", i+1, url, len(body), err.Error(), needTime)
				continue
			}
			if body != nil {
				//beego.Debug("CommonClient Call done. url:", url, " bodylen:", len(body), " body:", body)
				break
			}
		}

		break
	}
	return body, err
}

// TimeoutDialer returns functions of connection dialer with timeout settings for http.Transport Dial field.
func TimeoutDialer(connTimeout time.Duration) func(net, addr string) (c net.Conn, err error) {
	return func(netw, addr string) (net.Conn, error) {
		conn, err := net.DialTimeout(netw, addr, connTimeout)
		if err != nil {
			return nil, err
		}
		return conn, err
	}
}

func init() {

	conns, _ := beego.AppConfig.Int("CommonConfig::MaxIdleConnsPerHost")
	timeout, _ := beego.AppConfig.Int("CommonConfig::HttpTimeOut")
	timeoutRetry, _ := beego.AppConfig.Int("CommonConfig::HttpTimeOutRetry")
	timeoutex, _ := beego.AppConfig.Int("CommonConfig::HttpTimeOutSpecial")

	if conns > 0 {
		gMaxIdleConns = conns
	}

	if timeout > 0 {
		gTimeOutSec = timeout
	}

	if timeoutex > 0 {
		gTimeOutSecSpecial = timeoutex
	}

	if timeoutRetry > 0 {
		gTimeoutRetry = timeoutRetry
	}

	//设置超时特殊服务
	strSpecialSvrConfig := utils.GetConfigByString("CommonConfig::SpecialSvr")
	SpecialSvrList := strings.Split(strSpecialSvrConfig, ",")

	for _, value := range SpecialSvrList {
		G_SpecialSvrMap[value] = 1
	}

	httpSetting = httplib.BeegoHTTPSettings{
		//UserAgent:        "beegoServer",
		ConnectTimeout:   time.Duration(gTimeOutSec) * time.Second,
		ReadWriteTimeout: time.Duration(gTimeOutSec) * time.Second,
		// Gzip:             true,
		// DumpBody:         true,
		//		Transport: &http.Transport{
		//			//设置http连接复用
		//			MaxIdleConnsPerHost: gMaxIdleConns,
		//		},
	}

	gTransport = &http.Transport{
		//连接池最大的Idle连接数
		MaxIdleConnsPerHost: gMaxIdleConns,
		//请求超时设置
		ResponseHeaderTimeout: time.Duration(gTimeOutSec) * time.Second,
		//连接超时设置，不设置连接的deadline
		Dial: TimeoutDialer(time.Duration(gTimeOutSec) * time.Second),
		//暂时关闭连接复用
		//DisableKeepAlives: true,
		//设置idle超时时间
		IdleConnTimeout: time.Duration(gIdleTimeOutSec) * time.Second,
	}

	httpSettingSpecial = httplib.BeegoHTTPSettings{
		//UserAgent:        "beegoServer",
		ConnectTimeout:   time.Duration(gTimeOutSecSpecial) * time.Second,
		ReadWriteTimeout: time.Duration(gTimeOutSecSpecial) * time.Second,
		// Gzip:             true,
		// DumpBody:         true,
		//		Transport: &http.Transport{
		//			//设置http连接复用
		//			MaxIdleConnsPerHost: gMaxIdleConns,
		//		},
	}

	gTransportSepcial = &http.Transport{
		//连接池最大的Idle连接数
		MaxIdleConnsPerHost: gMaxIdleConns,
		//请求超时设置
		ResponseHeaderTimeout: time.Duration(gTimeOutSecSpecial) * time.Second,
		//连接超时设置，不设置连接的deadline
		Dial: TimeoutDialer(time.Duration(gTimeOutSecSpecial) * time.Second),
		//暂时关闭连接复用
		//DisableKeepAlives: true,
		//设置idle超时时间
		IdleConnTimeout: time.Duration(gIdleTimeOutSec) * time.Second,
	}

}
