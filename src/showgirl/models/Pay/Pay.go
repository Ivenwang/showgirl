package Pay

import (
	"crypto/md5"
	"encoding/hex"
	"showgirl/client"
	"showgirl/models/utils"
	"strconv"
	"strings"
	"time"
)

const WXAppID = "wx67cb0c8656d9400c"
const WXMchID = "xxx"
const NotifyURL = "https://grassua.site/callpack/pay"
const PaySignKey = "xxx"

//WXPayCreateOrderReq 创建支付订单
type WXPayCreateOrderReq struct {
	AppID          string `xml:"appid"`            //小程序id
	MchID          string `xml:"mch_id"`           //商户号
	DeviceInfo     string `xml:"device_info"`      //设备号
	NonceStr       string `xml:"nonce_str"`        //随机字符串
	Sign           string `xml:"sign"`             //签名
	SignType       string `xml:"sign_type"`        //签名类型
	Body           string `xml:"body"`             //商品描述
	Detail         string `xml:"detail"`           //商品详情
	Attach         string `xml:"attach"`           //附加数据
	OutTradeNo     string `xml:"out_trade_no"`     //商户订单号
	FeeType        string `xml:"fee_type"`         //货币类型
	TotalFee       int32  `xml:"total_fee"`        //总金额
	SpbillCreateIP string `xml:"spbill_create_ip"` //终端IP
	TimeStart      string `xml:"time_start"`       //交易起始时间
	TimeExpire     string `xml:"time_expire"`      //交易结束时间
	GoodsTag       string `xml:"goods_tag"`        //商品标记
	NotifyURL      string `xml:"notify_url"`       //通知地址
	TradeType      string `xml:"trade_type"`       //交易类型
	LimitPay       string `xml:"limit_pay"`        //指定支付方式
	OpenID         string `xml:"openid"`           //用户标识
}

//CreateTransaction 创建订单
func CreateTransaction(IPAddr string, FeeAmount int32, OpenID string, flowid int64) (*client.STCreateTransactionRsp, error) {

	if FeeAmount <= 0 {
		return nil, nil
	}

	//构造请求支付结构
	stCreateTransaction := &client.STCreateTransactionRsp{}

	stCreateOrder := WXPayCreateOrderReq{
		AppID:          WXAppID,
		MchID:          WXMchID,
		DeviceInfo:     "WEB",
		NonceStr:       string(utils.Krand(32, utils.KC_RAND_KIND_ALL)),
		SignType:       "MD5",
		Body:           "美女-特权",
		OutTradeNo:     string(utils.Krand(32, utils.KC_RAND_KIND_ALL)),
		FeeType:        "CNY",
		TotalFee:       FeeAmount,
		SpbillCreateIP: IPAddr,
		TimeStart:      time.Now().Format("20060102150405"),
		TimeExpire:     time.Unix(time.Now().Unix()+86400, 0).Format("20060102150405"),
		NotifyURL:      NotifyURL,
		TradeType:      "JSAPI",
		OpenID:         OpenID,
	}

	stCreateOrder.Sign = GenSignByPay(stCreateOrder, flowid)

	//发送支付请求

	return stCreateTransaction, nil
}

//GenSignByPay 构建支付签名
func GenSignByPay(stCreateOrder WXPayCreateOrderReq, flowid int64) string {

	//1、构造stringA
	stringA := ""
	if len(stCreateOrder.AppID) > 0 {
		stringA += "appid=" + stCreateOrder.AppID + "&"
	}
	if len(stCreateOrder.Body) > 0 {
		stringA += "body=" + stCreateOrder.Body + "&"
	}
	if len(stCreateOrder.DeviceInfo) > 0 {
		stringA += "device_info=" + stCreateOrder.DeviceInfo + "&"
	}
	if len(stCreateOrder.FeeType) > 0 {
		stringA += "fee_type=" + stCreateOrder.FeeType + "&"
	}
	if len(stCreateOrder.MchID) > 0 {
		stringA += "mch_id=" + stCreateOrder.MchID + "&"
	}
	if len(stCreateOrder.NonceStr) > 0 {
		stringA += "nonce_str=" + stCreateOrder.NonceStr + "&"
	}
	if len(stCreateOrder.NotifyURL) > 0 {
		stringA += "notify_url=" + stCreateOrder.NotifyURL + "&"
	}
	if len(stCreateOrder.OpenID) > 0 {
		stringA += "openid=" + stCreateOrder.OpenID + "&"
	}
	if len(stCreateOrder.OutTradeNo) > 0 {
		stringA += "out_trade_no=" + stCreateOrder.OutTradeNo + "&"
	}
	if len(stCreateOrder.Sign) > 0 {
		stringA += "sign=" + stCreateOrder.Sign + "&"
	}
	if len(stCreateOrder.SignType) > 0 {
		stringA += "sign_type=" + stCreateOrder.SignType + "&"
	}
	if len(stCreateOrder.SpbillCreateIP) > 0 {
		stringA += "spbill_create_ip=" + stCreateOrder.SpbillCreateIP + "&"
	}
	if len(stCreateOrder.TimeExpire) > 0 {
		stringA += "time_expire=" + stCreateOrder.TimeExpire + "&"
	}
	if len(stCreateOrder.TimeStart) > 0 {
		stringA += "time_start=" + stCreateOrder.TimeStart + "&"
	}
	if stCreateOrder.TotalFee > 0 {
		stringA += "total_fee=" + strconv.FormatInt(int64(stCreateOrder.TotalFee), 10) + "&"
	}
	if len(stCreateOrder.TradeType) > 0 {
		stringA += "trade_type=" + stCreateOrder.TradeType + "&"
	}

	//2、构造stringSignTemp
	stringSignTemp := stringA + "key=" + PaySignKey

	md5Ctx := md5.New()
	md5Ctx.Write([]byte(stringSignTemp))
	cipherStr := md5Ctx.Sum(nil)

	return strings.ToUpper(hex.EncodeToString(cipherStr))

}
