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

type ImageController struct {
	beego.Controller
}

func (this *ImageController) DoResponse(
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


func (this *ImageController) QueryStyleList() {
	commonReq := &client.CommonReq {}
	var rspHeader *client.STRspHeader
	var rspInfo *client.STQueryStyleListRsp
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
		//utils.Debug(0, "Controller Request QueryStyleList bodyLen = %d, body = %x", bodyLen, data)
		err := proto.Unmarshal(data, commonReq)
		if err != nil {
			errno = client.EErrorTypeDef_RPC_INTERFACE_ABNORMAL
			errdetail = fmt.Sprintf("Unmarshaling RequestBody failed. err:%v", err)
			errmsg = utils.ERRMSG_EXCEPTION_CATCHED
			beego.Critical(errdetail)
			break
		}
		reqHeader := commonReq.UserTrustInfo
		reqInfo := &client.STQueryStyleListReq{}
		err = proto.Unmarshal(commonReq.ReqInfo, reqInfo)
		if err != nil {
			errno = client.EErrorTypeDef_RPC_INTERFACE_ABNORMAL
			errmsg = utils.ERRMSG_EXCEPTION_CATCHED
			errdetail = fmt.Sprintf("Unmarshaling ReqInfo failed. err:%v", err)
			beego.Critical(errdetail)
			break
		}
		
		//Step2: 调用对应的Handler
		utils.JInfo(reqHeader.GetFlowId(), "Request", utils.LMap{"method":"QueryStyleList", "reqHeader": reqHeader,"reqInfo": reqInfo})
		//process reqHeader and reqInfo, remember set rspHeader and rspInfo
		rspHeader, rspInfo = HandleQueryStyleList(reqHeader, reqInfo)
		utils.JInfo(reqHeader.GetFlowId(), "Response", utils.LMap{"method":"QueryStyleList", "rspHeader": rspHeader,"rspInfo": rspInfo})
		
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

func (this *ImageController) QueryResourceList() {
	commonReq := &client.CommonReq {}
	var rspHeader *client.STRspHeader
	var rspInfo *client.STQueryResourceListRsp
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
		//utils.Debug(0, "Controller Request QueryResourceList bodyLen = %d, body = %x", bodyLen, data)
		err := proto.Unmarshal(data, commonReq)
		if err != nil {
			errno = client.EErrorTypeDef_RPC_INTERFACE_ABNORMAL
			errdetail = fmt.Sprintf("Unmarshaling RequestBody failed. err:%v", err)
			errmsg = utils.ERRMSG_EXCEPTION_CATCHED
			beego.Critical(errdetail)
			break
		}
		reqHeader := commonReq.UserTrustInfo
		reqInfo := &client.STQueryResourceListReq{}
		err = proto.Unmarshal(commonReq.ReqInfo, reqInfo)
		if err != nil {
			errno = client.EErrorTypeDef_RPC_INTERFACE_ABNORMAL
			errmsg = utils.ERRMSG_EXCEPTION_CATCHED
			errdetail = fmt.Sprintf("Unmarshaling ReqInfo failed. err:%v", err)
			beego.Critical(errdetail)
			break
		}
		
		//Step2: 调用对应的Handler
		utils.JInfo(reqHeader.GetFlowId(), "Request", utils.LMap{"method":"QueryResourceList", "reqHeader": reqHeader,"reqInfo": reqInfo})
		//process reqHeader and reqInfo, remember set rspHeader and rspInfo
		rspHeader, rspInfo = HandleQueryResourceList(reqHeader, reqInfo)
		utils.JInfo(reqHeader.GetFlowId(), "Response", utils.LMap{"method":"QueryResourceList", "rspHeader": rspHeader,"rspInfo": rspInfo})
		
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

func (this *ImageController) CreateStyle() {
	commonReq := &client.CommonReq {}
	var rspHeader *client.STRspHeader
	var rspInfo *client.STCreateStyleRsp
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
		//utils.Debug(0, "Controller Request CreateStyle bodyLen = %d, body = %x", bodyLen, data)
		err := proto.Unmarshal(data, commonReq)
		if err != nil {
			errno = client.EErrorTypeDef_RPC_INTERFACE_ABNORMAL
			errdetail = fmt.Sprintf("Unmarshaling RequestBody failed. err:%v", err)
			errmsg = utils.ERRMSG_EXCEPTION_CATCHED
			beego.Critical(errdetail)
			break
		}
		reqHeader := commonReq.UserTrustInfo
		reqInfo := &client.STCreateStyleReq{}
		err = proto.Unmarshal(commonReq.ReqInfo, reqInfo)
		if err != nil {
			errno = client.EErrorTypeDef_RPC_INTERFACE_ABNORMAL
			errmsg = utils.ERRMSG_EXCEPTION_CATCHED
			errdetail = fmt.Sprintf("Unmarshaling ReqInfo failed. err:%v", err)
			beego.Critical(errdetail)
			break
		}
		
		//Step2: 调用对应的Handler
		utils.JInfo(reqHeader.GetFlowId(), "Request", utils.LMap{"method":"CreateStyle", "reqHeader": reqHeader,"reqInfo": reqInfo})
		//process reqHeader and reqInfo, remember set rspHeader and rspInfo
		rspHeader, rspInfo = HandleCreateStyle(reqHeader, reqInfo)
		utils.JInfo(reqHeader.GetFlowId(), "Response", utils.LMap{"method":"CreateStyle", "rspHeader": rspHeader,"rspInfo": rspInfo})
		
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

func (this *ImageController) UploadImage() {
	commonReq := &client.CommonReq {}
	var rspHeader *client.STRspHeader
	var rspInfo *client.STUploadImageRsp
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
		//utils.Debug(0, "Controller Request UploadImage bodyLen = %d, body = %x", bodyLen, data)
		err := proto.Unmarshal(data, commonReq)
		if err != nil {
			errno = client.EErrorTypeDef_RPC_INTERFACE_ABNORMAL
			errdetail = fmt.Sprintf("Unmarshaling RequestBody failed. err:%v", err)
			errmsg = utils.ERRMSG_EXCEPTION_CATCHED
			beego.Critical(errdetail)
			break
		}
		reqHeader := commonReq.UserTrustInfo
		reqInfo := &client.STUploadImageReq{}
		err = proto.Unmarshal(commonReq.ReqInfo, reqInfo)
		if err != nil {
			errno = client.EErrorTypeDef_RPC_INTERFACE_ABNORMAL
			errmsg = utils.ERRMSG_EXCEPTION_CATCHED
			errdetail = fmt.Sprintf("Unmarshaling ReqInfo failed. err:%v", err)
			beego.Critical(errdetail)
			break
		}
		
		//Step2: 调用对应的Handler
		utils.JInfo(reqHeader.GetFlowId(), "Request", utils.LMap{"method":"UploadImage", "reqHeader": reqHeader,"reqInfo": reqInfo})
		//process reqHeader and reqInfo, remember set rspHeader and rspInfo
		rspHeader, rspInfo = HandleUploadImage(reqHeader, reqInfo)
		utils.JInfo(reqHeader.GetFlowId(), "Response", utils.LMap{"method":"UploadImage", "rspHeader": rspHeader,"rspInfo": rspInfo})
		
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

func (this *ImageController) DeleteStyle() {
	commonReq := &client.CommonReq {}
	var rspHeader *client.STRspHeader
	var rspInfo *client.STDeleteStyleRsp
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
		//utils.Debug(0, "Controller Request DeleteStyle bodyLen = %d, body = %x", bodyLen, data)
		err := proto.Unmarshal(data, commonReq)
		if err != nil {
			errno = client.EErrorTypeDef_RPC_INTERFACE_ABNORMAL
			errdetail = fmt.Sprintf("Unmarshaling RequestBody failed. err:%v", err)
			errmsg = utils.ERRMSG_EXCEPTION_CATCHED
			beego.Critical(errdetail)
			break
		}
		reqHeader := commonReq.UserTrustInfo
		reqInfo := &client.STDeleteStyleReq{}
		err = proto.Unmarshal(commonReq.ReqInfo, reqInfo)
		if err != nil {
			errno = client.EErrorTypeDef_RPC_INTERFACE_ABNORMAL
			errmsg = utils.ERRMSG_EXCEPTION_CATCHED
			errdetail = fmt.Sprintf("Unmarshaling ReqInfo failed. err:%v", err)
			beego.Critical(errdetail)
			break
		}
		
		//Step2: 调用对应的Handler
		utils.JInfo(reqHeader.GetFlowId(), "Request", utils.LMap{"method":"DeleteStyle", "reqHeader": reqHeader,"reqInfo": reqInfo})
		//process reqHeader and reqInfo, remember set rspHeader and rspInfo
		rspHeader, rspInfo = HandleDeleteStyle(reqHeader, reqInfo)
		utils.JInfo(reqHeader.GetFlowId(), "Response", utils.LMap{"method":"DeleteStyle", "rspHeader": rspHeader,"rspInfo": rspInfo})
		
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

func (this *ImageController) DeleteResource() {
	commonReq := &client.CommonReq {}
	var rspHeader *client.STRspHeader
	var rspInfo *client.STDeleteResourceRsp
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
		//utils.Debug(0, "Controller Request DeleteResource bodyLen = %d, body = %x", bodyLen, data)
		err := proto.Unmarshal(data, commonReq)
		if err != nil {
			errno = client.EErrorTypeDef_RPC_INTERFACE_ABNORMAL
			errdetail = fmt.Sprintf("Unmarshaling RequestBody failed. err:%v", err)
			errmsg = utils.ERRMSG_EXCEPTION_CATCHED
			beego.Critical(errdetail)
			break
		}
		reqHeader := commonReq.UserTrustInfo
		reqInfo := &client.STDeleteResourceReq{}
		err = proto.Unmarshal(commonReq.ReqInfo, reqInfo)
		if err != nil {
			errno = client.EErrorTypeDef_RPC_INTERFACE_ABNORMAL
			errmsg = utils.ERRMSG_EXCEPTION_CATCHED
			errdetail = fmt.Sprintf("Unmarshaling ReqInfo failed. err:%v", err)
			beego.Critical(errdetail)
			break
		}
		
		//Step2: 调用对应的Handler
		utils.JInfo(reqHeader.GetFlowId(), "Request", utils.LMap{"method":"DeleteResource", "reqHeader": reqHeader,"reqInfo": reqInfo})
		//process reqHeader and reqInfo, remember set rspHeader and rspInfo
		rspHeader, rspInfo = HandleDeleteResource(reqHeader, reqInfo)
		utils.JInfo(reqHeader.GetFlowId(), "Response", utils.LMap{"method":"DeleteResource", "rspHeader": rspHeader,"rspInfo": rspInfo})
		
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

func (this *ImageController) UpdateStyle() {
	commonReq := &client.CommonReq {}
	var rspHeader *client.STRspHeader
	var rspInfo *client.STUpdateStyleRsp
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
		//utils.Debug(0, "Controller Request UpdateStyle bodyLen = %d, body = %x", bodyLen, data)
		err := proto.Unmarshal(data, commonReq)
		if err != nil {
			errno = client.EErrorTypeDef_RPC_INTERFACE_ABNORMAL
			errdetail = fmt.Sprintf("Unmarshaling RequestBody failed. err:%v", err)
			errmsg = utils.ERRMSG_EXCEPTION_CATCHED
			beego.Critical(errdetail)
			break
		}
		reqHeader := commonReq.UserTrustInfo
		reqInfo := &client.STUpdateStyleReq{}
		err = proto.Unmarshal(commonReq.ReqInfo, reqInfo)
		if err != nil {
			errno = client.EErrorTypeDef_RPC_INTERFACE_ABNORMAL
			errmsg = utils.ERRMSG_EXCEPTION_CATCHED
			errdetail = fmt.Sprintf("Unmarshaling ReqInfo failed. err:%v", err)
			beego.Critical(errdetail)
			break
		}
		
		//Step2: 调用对应的Handler
		utils.JInfo(reqHeader.GetFlowId(), "Request", utils.LMap{"method":"UpdateStyle", "reqHeader": reqHeader,"reqInfo": reqInfo})
		//process reqHeader and reqInfo, remember set rspHeader and rspInfo
		rspHeader, rspInfo = HandleUpdateStyle(reqHeader, reqInfo)
		utils.JInfo(reqHeader.GetFlowId(), "Response", utils.LMap{"method":"UpdateStyle", "rspHeader": rspHeader,"rspInfo": rspInfo})
		
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
