package controllers

import (
	//"github.com/astaxie/beego"
	"showgirl/client"
	"showgirl/models/Pay"
	"showgirl/models/utils"

	"github.com/golang/protobuf/proto"
)

//auto generated.
func HandleCreateTransaction(hdr *client.STUserTrustInfo, req *client.STCreateTransactionReq) (*client.STRspHeader, *client.STCreateTransactionRsp) {

	rspHeader := &client.STRspHeader{
		ErrNo:  client.EErrorTypeDef_RESULT_OK.Enum(),
		ErrMsg: proto.String("success"),
	}

	stCreateTransactionRsp := &client.STCreateTransactionRsp{}

	var pCreateTransactionRsp *client.STCreateTransactionRsp

	for {

		if req.GetFeeAmount() <= 0 {
			utils.Warn(hdr.GetFlowId(), "HandleCreateTransaction CreateTransaction check FeeAmount error, req = %s", req.String())
			rspHeader.ErrNo = client.EErrorTypeDef_CHECK_CONTENT_ERROR.Enum()
			rspHeader.ErrMsg = proto.String("支付金额不合法")
			break
		}
		var err error
		stCreateTransactionRsp, err = Pay.CreateTransaction("", req.GetFeeAmount(), req.GetOpenID(), hdr.GetFlowId())
		if err != nil {
			utils.Warn(hdr.GetFlowId(), "HandleCreateTransaction CreateTransaction error, req = %s, err = %s", req.String(), err.Error())
			rspHeader.ErrNo = client.EErrorTypeDef_SYS_INTERNAL_ERROR.Enum()
			rspHeader.ErrMsg = proto.String("请求支付失败")
			break
		}

		pCreateTransactionRsp = stCreateTransactionRsp

		break
	}

	return rspHeader, pCreateTransactionRsp

}
