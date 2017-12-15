package Account

import (
	"encoding/json"
	"showgirl/models/utils"
	"time"
)

const WXLoginKey = "wx67cb0c8656d9400c"
const WXSecretKey = "b0d94d697e372ba770f81a2b78d3028e"

type WxAccessTokenInfo struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	OpenId       string `json:"openid"`
	Scope        string `json:"scope"`
	UnionId      string `json:"unionid"`
}

type WxUserDataInfo struct {
	OpenId     string   `json:"openid"`
	NickName   string   `json:"nickname"`
	Sex        int64    `json:"sex"`
	Province   string   `json:"province"`
	City       string   `json:"city"`
	Country    string   `json:"country"`
	HeadImgUrl string   `json:"headimgurl"`
	Privilege  []string `json:"privilege"`
	UnionId    string   `json:"unionid"`
}

//GetAccessTokenByWxKey 获取微信信息
func GetAccessTokenByWxKey(code string, flowid int64) (string, string, string, error) {

	//构造请求URL
	reqURL := "https://api.weixin.qq.com/sns/oauth2/access_token?appid=" + WXLoginKey + "&secret=" + WXSecretKey + "&code=" + code + "&grant_type=authorization_code"

	map_param := make(map[string]string)

	utils.Debug(flowid, "GetAccessTokenByWxKey reqUrl = %s", reqURL)

	rspData, err := utils.SendAndRecv("GET", reqURL, nil, map_param, flowid)
	if err != nil {
		utils.Warn(flowid, "GetAccessTokenByWxKey SendAndRecv error, reqUrl = %s, err = %s", reqURL, err.Error())
		return "", "", "", err
	}

	utils.Debug(flowid, "GetAccessTokenByWxKey rspData = %s", rspData)

	stWxAccessTokenInfo := &WxAccessTokenInfo{}

	err = json.Unmarshal(rspData, stWxAccessTokenInfo)
	if err != nil {
		utils.Warn(flowid, "GetAccessTokenByWxKey json Unmarshal error, reqUrl = %s, rspData = %s, err = %s", reqURL, rspData, err.Error())
		return "", "", "", err
	}

	utils.Debug(flowid, "GetAccessTokenByWxKey debug, WxAccessTokenInfo = %v", stWxAccessTokenInfo)

	return stWxAccessTokenInfo.AccessToken, stWxAccessTokenInfo.OpenId, stWxAccessTokenInfo.UnionId, nil
}

//QueryAndSetUserWxInfo 设置用户微信信息
func QueryAndSetUserWxInfo(AccessToken string, WxOpenID string, WxUnionID string, flowid int64) (UserBaseInfo, error) {

	baseInfo := UserBaseInfo{}

	//构造请求URL
	reqURL := "https://api.weixin.qq.com/sns/userinfo?access_token=" + AccessToken + "&openid=" + WxOpenID

	map_param := make(map[string]string)

	utils.Debug(flowid, "QueryAndSetUserWxInfo reqUrl = %s", reqURL)

	rspData, err := utils.SendAndRecv("GET", reqURL, nil, map_param, flowid)
	if err != nil {
		utils.Warn(flowid, "QueryAndSetUserWxInfo SendAndRecv error, reqUrl = %s, err = %s", reqURL, err.Error())
		return baseInfo, err
	}

	stWxUserDataInfo := &WxUserDataInfo{}

	err = json.Unmarshal(rspData, stWxUserDataInfo)
	if err != nil {
		utils.Warn(flowid, "QueryAndSetUserWxInfo json Unmarshal error, reqUrl = %s, rspData = %s, err = %s", reqURL, rspData, err.Error())
		return baseInfo, err
	}

	utils.Debug(flowid, "QueryAndSetUserWxInfo debug, reqURL = %s, rspData = %s", reqURL, rspData)

	//创建账号信息
	err = RegisterAccount(WxOpenID, WxUnionID, stWxUserDataInfo.NickName, stWxUserDataInfo.HeadImgUrl, flowid)
	if err != nil {
		utils.Warn(flowid, "QueryAndSetUserWxInfo RegisterAccount WXOpenID = %s, WXUnionID = %s, NickName = %s, Url = %s, error = %s",
			WxOpenID, WxUnionID, stWxUserDataInfo.NickName, stWxUserDataInfo.HeadImgUrl, err.Error())
		return baseInfo, err
	}

	//构造返回结构
	baseInfo.WxOpenID = WxOpenID
	baseInfo.WxUnionID = WxUnionID
	baseInfo.NickName = stWxUserDataInfo.NickName
	baseInfo.URL = stWxUserDataInfo.HeadImgUrl
	baseInfo.Charge = int32(0)
	baseInfo.LastTime = int64(time.Now().Unix())

	return baseInfo, nil
}
