package controllers

import (
	//"github.com/astaxie/beego"
	"showgirl/client"
	"showgirl/models/Recommend"
	"showgirl/models/utils"

	"github.com/golang/protobuf/proto"
)

//auto generated.
func HandleQueryRecommendList(hdr *client.STUserTrustInfo, req *client.STQueryRecommendListReq) (*client.STRspHeader, *client.STQueryRecommendListRsp) {

	rspHeader := &client.STRspHeader{
		ErrNo:  client.EErrorTypeDef_RESULT_OK.Enum(),
		ErrMsg: proto.String("success"),
	}

	stQueryRecommendListRsp := &client.STQueryRecommendListRsp{}

	var pQueryRecommendListRsp *client.STQueryRecommendListRsp

	for {

		//查询推荐列表
		arry, err := Recommend.QueryImageListByRecommendCategory([]int32{int32(client.ERecommendTypeDef_BANNER_TYPE), int32(client.ERecommendTypeDef_NEW_TYPE),
			int32(client.ERecommendTypeDef_RECOMMEND_TYPE)}, hdr.GetFlowId())
		if err != nil {
			utils.Warn(hdr.GetFlowId(), "HandleQueryRecommendList QueryImageListByRecommendCategory error, req = %s, err = %s", req.String(), err.Error())
			rspHeader.ErrNo = client.EErrorTypeDef_SYS_INTERNAL_ERROR.Enum()
			rspHeader.ErrMsg = proto.String("系统繁忙，请稍候重试")
			break
		}

		for idx := range arry {
			if arry[idx].Category == int32(client.ERecommendTypeDef_BANNER_TYPE) {
				stQueryRecommendListRsp.BannerList = arry[idx].ArryStyle
			} else if arry[idx].Category == int32(client.ERecommendTypeDef_NEW_TYPE) {
				stQueryRecommendListRsp.NewImageList = arry[idx].ArryStyle
			} else if arry[idx].Category == int32(client.ERecommendTypeDef_RECOMMEND_TYPE) {
				stQueryRecommendListRsp.RecommendImageList = arry[idx].ArryStyle
			}
		}

		pQueryRecommendListRsp = stQueryRecommendListRsp
		break
	}

	return rspHeader, pQueryRecommendListRsp
}

//auto generated.
func HandleStyleImageList(hdr *client.STUserTrustInfo, req *client.STStyleImageListReq) (*client.STRspHeader, *client.STStyleImageListRsp) {

	rspHeader := &client.STRspHeader{
		ErrNo:  client.EErrorTypeDef_RESULT_OK.Enum(),
		ErrMsg: proto.String("success"),
	}

	stStyleImageListRsp := &client.STStyleImageListRsp{}

	var pStyleImageListRsp *client.STStyleImageListRsp

	for {

		//查询图片列表
		var err error
		stStyleImageListRsp.Urls, err = Recommend.QueryImageListByStyleID(req.GetStyleID(), hdr.GetFlowId())
		if err != nil {
			utils.Warn(hdr.GetFlowId(), "HandleStyleImageList QueryImageListByStyleID error, req = %s, err = %s", req.String(), err.Error())
			rspHeader.ErrNo = client.EErrorTypeDef_SYS_INTERNAL_ERROR.Enum()
			rspHeader.ErrMsg = proto.String("系统繁忙，请稍候重试")
			break
		}

		pStyleImageListRsp = stStyleImageListRsp
		break
	}

	return rspHeader, pStyleImageListRsp

}
