package Pay

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/xml"
	"errors"
	"fmt"
	"showgirl/client"
	"showgirl/models/Account"
	"showgirl/models/mysql"
	"showgirl/models/utils"
	"strconv"
	"strings"
	"time"

	"github.com/gogo/protobuf/proto"
)

const WXAppID = "wx67cb0c8656d9400c"
const WXMchID = "1462801602"
const NotifyURL = "https://grassua.site/callpack/pay"
const PaySignKey = "xxx"
const WxPayURL = "https://api.mch.weixin.qq.com/pay/unifiedorder"

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

//WXPayCreateOrderRsp 支付回包
type WXPayCreateOrderRsp struct {
	ReturnCode string `xml:"return_code"`  //返回码，通讯标识
	ReturnMsg  string `xml:"return_msg"`   //通讯错误码
	AppID      string `xml:"appid"`        //小程序id
	MchID      string `xml:"mch_id"`       //商户号
	DeviceInfo string `xml:"device_info"`  //设备号
	NonceStr   string `xml:"nonce_str"`    //随机字符串
	Sign       string `xml:"sign"`         //签名
	ResultCode string `xml:"result_code"`  //业务结果
	ErrCode    string `xml:"err_code"`     //错误代码
	ErrCodeDes string `xml:"err_code_des"` //错误返回的信息描述
	TradeType  string `xml:"trade_type"`   //交易类型
	PrepayID   string `xml:"prepay_id"`    //预支付交易会话标识
}

//WXPayNotifyInfo 微信支付通知
type WXPayNotifyInfo struct {
	ReturnCode         string `xml:"return_code"`          //返回码，通讯标识
	ReturnMsg          string `xml:"return_msg"`           //通讯错误码
	AppID              string `xml:"appid"`                //小程序id
	MchID              string `xml:"mch_id"`               //商户号
	DeviceInfo         string `xml:"device_info"`          //设备号
	NonceStr           string `xml:"nonce_str"`            //随机字符串
	Sign               string `xml:"sign"`                 //签名
	SignType           string `xml:"sign_type"`            //签名类型
	ResultCode         string `xml:"result_code"`          //业务结果
	ErrCode            string `xml:"err_code"`             //错误代码
	ErrCodeDes         string `xml:"err_code_des"`         //错误返回的信息描述
	OpenID             string `xml:"openid"`               //用户标识
	IsSubscribe        string `xml:"is_subscribe"`         //是否关注公众号
	TradeType          string `xml:"trade_type"`           //交易类型
	BankType           string `xml:"bank_type"`            //银行类型
	TotalFee           int32  `xml:"total_fee"`            //总金额，单位分
	SettleMentTotalFee int32  `xml:"settlement_total_fee"` //应结订单金额，单位分
	FeeType            string `xml:"fee_type"`             //货币类型
	CashFee            int32  `xml:"cash_fee"`             //现金支付金额
	CashFeeType        string `xml:"cash_fee_type"`        //现金支付货币类型
	CouponFee          int32  `xml:"coupon_fee"`           //代金券金额
	CouponCount        int32  `xml:"coupon_count"`         //代金券数量
	TransactionID      string `xml:"transaction_id"`       //微信支付订单号
	OutTradeNo         string `xml:"out_trade_no"`         //商户订单号
	Attach             string `xml:"attach"`               //商家数据包
	TimeEnd            string `xml:"time_end"`             //支付完成时间
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

	//序列化
	reqBody, err := xml.Marshal(stCreateOrder)
	if err != nil {
		utils.Warn(flowid, "CreateTransaction xml.Marshal error, FeeAmount = %d, error = %s",
			FeeAmount, err.Error())
		return stCreateTransaction, err
	}

	//构造请求URL
	rspData, err := utils.SendAndRecv("POST", WxPayURL, bytes.NewBuffer(reqBody), nil, flowid)
	if err != nil {
		utils.Warn(flowid, "CreateTransaction SendAndRecv error, WxPayURL = %s, reqBody = %s, err = %s",
			WxPayURL, reqBody, err.Error())
		return stCreateTransaction, err
	}

	utils.Debug(flowid, "CreateTransaction debug, WxPayURL = %s, reqBody = %s, rspData = %s",
		WxPayURL, reqBody, rspData)

	//解包
	stCreateOrderRsp := &WXPayCreateOrderRsp{}
	err = xml.Unmarshal(rspData, stCreateOrderRsp)
	if err != nil {
		utils.Warn(flowid, "CreateTransaction xml.Unmarshal error, WxPayURL = %s, reqBody = %s, rspData = %s, err = %s",
			WxPayURL, reqBody, rspData, err.Error())
		return stCreateTransaction, err
	}

	//通信错误
	if stCreateOrderRsp.ReturnCode != "SUCCESS" {
		utils.Warn(flowid, "CreateTransaction check ReturnCode error, WxPayURL = %s, reqBody = %s, rspData = %s",
			WxPayURL, reqBody, rspData)
		return stCreateTransaction, errors.New("check rsp error")
	}

	stCreateTransaction.AppId = proto.String(stCreateOrderRsp.AppID)
	stCreateTransaction.MchID = proto.String(stCreateOrderRsp.MchID)
	stCreateTransaction.DeviceInfo = proto.String(stCreateOrderRsp.DeviceInfo)
	stCreateTransaction.NonceStr = proto.String(stCreateOrderRsp.NonceStr)
	stCreateTransaction.Sign = proto.String(stCreateOrderRsp.Sign)
	stCreateTransaction.ResultCode = proto.String(stCreateOrderRsp.ResultCode)
	stCreateTransaction.ErrCode = proto.String(stCreateOrderRsp.ErrCode)
	stCreateTransaction.ErrCodeDes = proto.String(stCreateOrderRsp.ErrCodeDes)
	stCreateTransaction.TradeType = proto.String(stCreateOrderRsp.TradeType)
	stCreateTransaction.PrepayID = proto.String(stCreateOrderRsp.PrepayID)

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

//UpdateUserChargeNum 更新用户充值金额
func UpdateUserChargeNum(WxOpenID string, TotalFee int32) error {

	//查询用户unionid
	WxUnionID := Account.QueryWXUnionIDByWXOpenID(WxOpenID, int64(0))
	if len(WxUnionID) <= 0 {
		return errors.New("query account error")
	}

	//新建mysql实例
	o := mysql.NewShowgirlOrm()

	sSQL := fmt.Sprintf("update ShowGirlAccountInfo set ChargeNum = ChargeNum + %d where WxUnionID = %q",
		TotalFee, WxUnionID)
	_, err := o.Raw(sSQL).Exec()
	if err != nil {
		utils.Warn(0, "UpdateUserChargeNum update account error, sql = %s, err = %s", sSQL, err.Error())
		return err
	}

	utils.Debug(0, "UpdateUserChargeNum debug, sql = %s, WxOpenID = %s, WxUnionID = %s, TotalFee = %d",
		sSQL, WxOpenID, WxUnionID, TotalFee)

	return nil

}

//InsertPayFlow 写入充值流水
func InsertPayFlow(stPayNotifyInfo WXPayNotifyInfo) error {

	//新建mysql实例
	o := mysql.NewShowgirlOrm()

	sSQL := ConstructInsertPayFlowSQL(stPayNotifyInfo)
	if len(sSQL) <= 0 {
		return errors.New("construct sql error")
	}
	_, err := o.Raw(sSQL).Exec()
	if err != nil {
		utils.Warn(0, "InsertPayFlow insert flow error, sql = %s, err = %s", sSQL, err.Error())
		return err
	}

	utils.Debug(0, "InsertPayFlow debug, sql = %s", sSQL)
	return nil
}

//ConstructInsertPayFlowSQL 构建写入充值流水SQL
func ConstructInsertPayFlowSQL(stPayNotifyInfo WXPayNotifyInfo) string {

	return fmt.Sprintf("insert into PayNotifyFlow(Id,ReturnCode,ReturnMsg,AppID,"+
		"MchID,DeviceInfo,NonceStr,Sign,SignType,ResultCode,ErrCode,ErrCodeDes,OpenID,"+
		"IsSubscribe,TradeType,BankType,TotalFee,SettleMentTotalFee,FeeType,CashFee,"+
		"CashFeeType,CouponFee,CouponCount,TransactionID,OutTradeNo,Attach,TimeEnd,"+
		"CreateTime,UpdateTime) values(null,%q,%q,%q,%q,%q, %q,%q,%q,%q,%q,"+
		"%q,%q,%q,%q,%q, %d,%d,%q,%d,%q, %d,%d,%q,%q,%q, %q,Now(),null)",
		stPayNotifyInfo.ReturnCode,
		stPayNotifyInfo.ReturnMsg,
		stPayNotifyInfo.AppID,
		stPayNotifyInfo.MchID,
		stPayNotifyInfo.DeviceInfo,
		stPayNotifyInfo.NonceStr,
		stPayNotifyInfo.Sign,
		stPayNotifyInfo.SignType,
		stPayNotifyInfo.ResultCode,
		stPayNotifyInfo.ErrCode,
		stPayNotifyInfo.ErrCodeDes,
		stPayNotifyInfo.OpenID,
		stPayNotifyInfo.IsSubscribe,
		stPayNotifyInfo.TradeType,
		stPayNotifyInfo.BankType,
		stPayNotifyInfo.TotalFee,
		stPayNotifyInfo.SettleMentTotalFee,
		stPayNotifyInfo.FeeType,
		stPayNotifyInfo.CashFee,
		stPayNotifyInfo.CashFeeType,
		stPayNotifyInfo.CouponFee,
		stPayNotifyInfo.CouponCount,
		stPayNotifyInfo.TransactionID,
		stPayNotifyInfo.OutTradeNo,
		stPayNotifyInfo.Attach,
		stPayNotifyInfo.TimeEnd)

}
