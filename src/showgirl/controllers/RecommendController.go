/*
 * Auto generated by code_generator
 * Please do not modify it.
 */
package controllers

import (
	"showgirl/models/utils"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/golang/protobuf/proto"
	"showgirl/client"
	"runtime/debug"
)

type RecommendController struct {
	beego.Controller
}

func (this *RecommendController) DoResponse(
	commonReq   *client.CommonReq,
    rspHeader   *client.STRspHeader,
	rspInfoData []byte,
	errno       client.EErrorTypeDef,
	errmsg      string,
	errdetail   string) {
	if err := recover(); err!=nil {
		beego.Critical("Panic catched, err:\n", err, string(debug.Stack()))
		errno = client.EErrorTypeDef_PROGRAM_EXCEPTION_ERROR
		errmsg = utils.ERRMSG_EXCEPTION_CATCHED
		errdetail = "Panic catched"
	}
	if rspHeader == nil {
		rspHeader = &client.STRspHeader{
			ErrNo: errno.Enum(),
			ErrMsg: proto.String(errmsg),
			ErrDetail: proto.String(errdetail),
		}
	}

	this.Ctx.Output.Header("Access-Control-Allow-Origin", "*")

	commonRes := &client.CommonRsp{
		UserTrustInfo: commonReq.UserTrustInfo,
		RspHeader: rspHeader,
	}
	
	if len(rspInfoData) > 0 {
		commonRes.RspInfo = rspInfoData
	}
	
	data, err := proto.Marshal(commonRes)
	if err != nil {
		beego.Critical("Marshaling commonRes error: ", err)
		return
	}
	this.Ctx.Output.Body(data)
}


func (this *RecommendController) QueryRecommendList() {
	commonReq := &client.CommonReq {}
	var rspHeader *client.STRspHeader
	var rspInfo *client.STQueryRecommendListRsp
	rspInfoData := []byte{}

	errno := client.EErrorTypeDef_CHECK_CONTENT_ERROR
	errmsg := utils.ERRMSG_CLIENT_EXCEPTION
	errdetail := "Param check failed."
	
	defer func() {
		this.DoResponse(commonReq, rspHeader, rspInfoData, errno, errmsg, errdetail)
	}()

	for {
		//Step1: 解析请求
		data := this.Ctx.Input.RequestBody
		//bodyLen := len(data)
		//utils.Debug(0, "Controller Request QueryRecommendList bodyLen = %d, body = %x", bodyLen, data)
		err := proto.Unmarshal(data, commonReq)
		if err != nil {
			errno = client.EErrorTypeDef_RPC_INTERFACE_ABNORMAL
			errdetail = fmt.Sprintf("Unmarshaling RequestBody failed. err:%v", err)
			errmsg = utils.ERRMSG_EXCEPTION_CATCHED
			beego.Critical(errdetail)
			break
		}
		reqHeader := commonReq.UserTrustInfo
		reqInfo := &client.STQueryRecommendListReq{}
		err = proto.Unmarshal(commonReq.ReqInfo, reqInfo)
		if err != nil {
			errno = client.EErrorTypeDef_RPC_INTERFACE_ABNORMAL
			errmsg = utils.ERRMSG_EXCEPTION_CATCHED
			errdetail = fmt.Sprintf("Unmarshaling ReqInfo failed. err:%v", err)
			beego.Critical(errdetail)
			break
		}
		
		//Step2: 调用对应的Handler
		utils.JInfo(reqHeader.GetFlowId(), "Request", utils.LMap{"method":"QueryRecommendList", "reqHeader": reqHeader,"reqInfo": reqInfo})
		//process reqHeader and reqInfo, remember set rspHeader and rspInfo
		rspHeader, rspInfo = HandleQueryRecommendList(reqHeader, reqInfo)
		utils.JInfo(reqHeader.GetFlowId(), "Response", utils.LMap{"method":"QueryRecommendList", "rspHeader": rspHeader,"rspInfo": rspInfo})
		
		//Step3: 将RspInfo进行PB序列化
		if rspInfo != nil {
			rspInfoData, err = proto.Marshal(rspInfo)
			if err != nil {
				errno = client.EErrorTypeDef_GENERATE_CONTENT_ERROR
				errdetail = fmt.Sprintf("rspInfoData Marshal failed. err:%v ", err)
				errmsg = utils.ERRMSG_EXCEPTION_CATCHED
				beego.Critical(errdetail)
				break
			}
		}
		
		break
	}
}

func (this *RecommendController) StyleImageList() {
	commonReq := &client.CommonReq {}
	var rspHeader *client.STRspHeader
	var rspInfo *client.STStyleImageListRsp
	rspInfoData := []byte{}

	errno := client.EErrorTypeDef_CHECK_CONTENT_ERROR
	errmsg := utils.ERRMSG_CLIENT_EXCEPTION
	errdetail := "Param check failed."
	
	defer func() {
		this.DoResponse(commonReq, rspHeader, rspInfoData, errno, errmsg, errdetail)
	}()

	for {
		//Step1: 解析请求
		data := this.Ctx.Input.RequestBody
		//bodyLen := len(data)
		//utils.Debug(0, "Controller Request StyleImageList bodyLen = %d, body = %x", bodyLen, data)
		err := proto.Unmarshal(data, commonReq)
		if err != nil {
			errno = client.EErrorTypeDef_RPC_INTERFACE_ABNORMAL
			errdetail = fmt.Sprintf("Unmarshaling RequestBody failed. err:%v", err)
			errmsg = utils.ERRMSG_EXCEPTION_CATCHED
			beego.Critical(errdetail)
			break
		}
		reqHeader := commonReq.UserTrustInfo
		reqInfo := &client.STStyleImageListReq{}
		err = proto.Unmarshal(commonReq.ReqInfo, reqInfo)
		if err != nil {
			errno = client.EErrorTypeDef_RPC_INTERFACE_ABNORMAL
			errmsg = utils.ERRMSG_EXCEPTION_CATCHED
			errdetail = fmt.Sprintf("Unmarshaling ReqInfo failed. err:%v", err)
			beego.Critical(errdetail)
			break
		}
		
		//Step2: 调用对应的Handler
		utils.JInfo(reqHeader.GetFlowId(), "Request", utils.LMap{"method":"StyleImageList", "reqHeader": reqHeader,"reqInfo": reqInfo})
		//process reqHeader and reqInfo, remember set rspHeader and rspInfo
		rspHeader, rspInfo = HandleStyleImageList(reqHeader, reqInfo)
		utils.JInfo(reqHeader.GetFlowId(), "Response", utils.LMap{"method":"StyleImageList", "rspHeader": rspHeader,"rspInfo": rspInfo})
		
		//Step3: 将RspInfo进行PB序列化
		if rspInfo != nil {
			rspInfoData, err = proto.Marshal(rspInfo)
			if err != nil {
				errno = client.EErrorTypeDef_GENERATE_CONTENT_ERROR
				errdetail = fmt.Sprintf("rspInfoData Marshal failed. err:%v ", err)
				errmsg = utils.ERRMSG_EXCEPTION_CATCHED
				beego.Critical(errdetail)
				break
			}
		}
		
		break
	}
}
