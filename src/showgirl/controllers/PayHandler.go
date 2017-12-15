package controllers

import (
	//"github.com/astaxie/beego"
	"showgirl/client"

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

		pCreateTransactionRsp = stCreateTransactionRsp

		break
	}

	return rspHeader, pCreateTransactionRsp

}
