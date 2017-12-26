package controllers

import (
	//"github.com/astaxie/beego"
	"showgirl/client"
	"showgirl/models/Image"
	"showgirl/models/utils"

	"github.com/golang/protobuf/proto"
)

//auto generated.
func HandleQueryStyleList(hdr *client.STUserTrustInfo, req *client.STQueryStyleListReq) (*client.STRspHeader, *client.STQueryStyleListRsp) {

	rspHeader := &client.STRspHeader{
		ErrNo:  client.EErrorTypeDef_RESULT_OK.Enum(),
		ErrMsg: proto.String("success"),
	}

	stQueryStyleListRsp := &client.STQueryStyleListRsp{}

	var pQueryStyleListRsp *client.STQueryStyleListRsp

	for {

		//查询相册列表
		var err error
		var total int32
		stQueryStyleListRsp.StyleList, total, err = Image.QueryStyleInfoList(req.GetQueryBegin(),
			req.GetQueryNum(), hdr.GetFlowId())
		if err != nil {
			utils.Warn(hdr.GetFlowId(), "HandleQueryStyleList QueryStyleInfoList error, req = %s, err = %s", req.String(), err.Error())
			rspHeader.ErrNo = client.EErrorTypeDef_SYS_INTERNAL_ERROR.Enum()
			rspHeader.ErrMsg = proto.String("系统繁忙，请稍候重试")
			break
		}

		stQueryStyleListRsp.Total = proto.Int32(total)

		pQueryStyleListRsp = stQueryStyleListRsp
		break
	}

	return rspHeader, pQueryStyleListRsp

}

//auto generated.
func HandleQueryResourceList(hdr *client.STUserTrustInfo, req *client.STQueryResourceListReq) (*client.STRspHeader, *client.STQueryResourceListRsp) {

	rspHeader := &client.STRspHeader{
		ErrNo:  client.EErrorTypeDef_RESULT_OK.Enum(),
		ErrMsg: proto.String("success"),
	}

	stQueryResourceListRsp := &client.STQueryResourceListRsp{}

	var pQueryResourceListRsp *client.STQueryResourceListRsp

	for {

		//查询资源列表
		var err error
		stQueryResourceListRsp.UrlList, err = Image.QueryResourceInfoList(req.GetStyleID(), hdr.GetFlowId())
		if err != nil {
			utils.Warn(hdr.GetFlowId(), "HandleQueryResourceList QueryResourceInfoList error, req = %s, err = %s", req.String(), err.Error())
			rspHeader.ErrNo = client.EErrorTypeDef_SYS_INTERNAL_ERROR.Enum()
			rspHeader.ErrMsg = proto.String("系统繁忙，请稍候重试")
			break
		}

		pQueryResourceListRsp = stQueryResourceListRsp
		break
	}

	return rspHeader, pQueryResourceListRsp

}

//auto generated.
func HandleCreateStyle(hdr *client.STUserTrustInfo, req *client.STCreateStyleReq) (*client.STRspHeader, *client.STCreateStyleRsp) {

	rspHeader := &client.STRspHeader{
		ErrNo:  client.EErrorTypeDef_RESULT_OK.Enum(),
		ErrMsg: proto.String("success"),
	}

	stCreateStyleRsp := &client.STCreateStyleRsp{}

	var pCreateStyleRsp *client.STCreateStyleRsp

	for {

		//创建相册
		err := Image.CreateStyle(req.GetStyleName(), req.GetStyleType(), hdr.GetFlowId())
		if err != nil {
			utils.Warn(hdr.GetFlowId(), "HandleCreateStyle CreateStyle error, req = %s, err = %s", req.String(), err.Error())
			rspHeader.ErrNo = client.EErrorTypeDef_SYS_INTERNAL_ERROR.Enum()
			rspHeader.ErrMsg = proto.String("系统繁忙，请稍候重试")
			break
		}

		pCreateStyleRsp = stCreateStyleRsp
		break
	}

	return rspHeader, pCreateStyleRsp

}

//auto generated.
func HandleUploadImage(hdr *client.STUserTrustInfo, req *client.STUploadImageReq) (*client.STRspHeader, *client.STUploadImageRsp) {

	rspHeader := &client.STRspHeader{
		ErrNo:  client.EErrorTypeDef_RESULT_OK.Enum(),
		ErrMsg: proto.String("success"),
	}

	for {

		//上传并设置相册图片
		err := Image.UploadAndSetImage(req.GetImage(), req.GetStyleID(), hdr.GetFlowId())
		if err != nil {
			utils.Warn(hdr.GetFlowId(), "HandleUploadImage UploadAndSetImage error, req = %s, err = %s", req.String(), err.Error())
			rspHeader.ErrNo = client.EErrorTypeDef_SYS_INTERNAL_ERROR.Enum()
			rspHeader.ErrMsg = proto.String("系统繁忙，请稍候重试")
			break
		}

		break
	}

	return rspHeader, nil

}

//auto generated.
func HandleDeleteStyle(hdr *client.STUserTrustInfo, req *client.STDeleteStyleReq) (*client.STRspHeader, *client.STDeleteStyleRsp) {

	rspHeader := &client.STRspHeader{
		ErrNo:  client.EErrorTypeDef_RESULT_OK.Enum(),
		ErrMsg: proto.String("success"),
	}

	stDeleteStyleRsp := &client.STDeleteStyleRsp{}

	var pDeleteStyleRsp *client.STDeleteStyleRsp

	for {

		//删除相册
		err := Image.DeleteStyle(req.GetStyleID(), hdr.GetFlowId())
		if err != nil {
			utils.Warn(hdr.GetFlowId(), "HandleDeleteStyle DeleteStyle error, req = %s, err = %s", req.String(), err.Error())
			rspHeader.ErrNo = client.EErrorTypeDef_SYS_INTERNAL_ERROR.Enum()
			rspHeader.ErrMsg = proto.String("系统繁忙，请稍候重试")
			break
		}

		pDeleteStyleRsp = stDeleteStyleRsp
		break
	}

	return rspHeader, pDeleteStyleRsp

}

//auto generated.
func HandleDeleteResource(hdr *client.STUserTrustInfo, req *client.STDeleteResourceReq) (*client.STRspHeader, *client.STDeleteResourceRsp) {

	rspHeader := &client.STRspHeader{
		ErrNo:  client.EErrorTypeDef_RESULT_OK.Enum(),
		ErrMsg: proto.String("success"),
	}

	stDeleteResourceRsp := &client.STDeleteResourceRsp{}

	var pDeleteResourceRsp *client.STDeleteResourceRsp

	for {

		//删除相册图片资源
		err := Image.DeleteResource(req.GetImageID(), hdr.GetFlowId())
		if err != nil {
			utils.Warn(hdr.GetFlowId(), "HandleDeleteResource DeleteStyle error, req = %s, err = %s", req.String(), err.Error())
			rspHeader.ErrNo = client.EErrorTypeDef_SYS_INTERNAL_ERROR.Enum()
			rspHeader.ErrMsg = proto.String("系统繁忙，请稍候重试")
			break
		}

		pDeleteResourceRsp = stDeleteResourceRsp
		break
	}

	return rspHeader, pDeleteResourceRsp

}
