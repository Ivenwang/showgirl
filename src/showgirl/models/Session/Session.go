package Session

import (
	"encoding/json"
	"errors"
	"showgirl/models/utils"
	"time"
)

type UserCookieInfo struct {
	WXOpenID  string `json:"openid"`
	WXUnionID string `json:"unionid"`
	NickName  string `json:"name"`
	Url       string `json:"url"`
	CurTime   int64  `json:"time"`
}

func UnmarshalCookie(Cookie string, flowid int64) (UserCookieInfo, error) {

	stCookieInfo := UserCookieInfo{}

	//长度校验
	if len(Cookie) <= 0 {
		return stCookieInfo, errors.New("cookie is empty")
	}

	//读取加密key
	key := utils.GetConfigByString("Session::cookiekey")

	//解密Cookie信息
	CookieJson, err := utils.Decrypt(Cookie, key)
	if err != nil {
		utils.Warn(flowid, "GetCookieData Decrypt data error, Cookie = %s, err = %s", Cookie, err.Error())
		return stCookieInfo, err
	}

	//结构化json
	err = json.Unmarshal([]byte(CookieJson), &stCookieInfo)
	if err != nil {
		utils.Warn(flowid, "GetCookieData Unmarshal data error, Cookie = %s, err = %s", Cookie, err.Error())
		return stCookieInfo, err
	}

	return stCookieInfo, nil
}

//获取cookie数据
func GetCookieData(Cookie string, flowid int64) (UserCookieInfo, error) {

	//没必要去redis解析，直接静态解析就可以了
	return UnmarshalCookie(Cookie, flowid)

	// stCookieInfo, err := UnmarshalCookie(Cookie, flowid)
	// if err != nil {
	// 	utils.Warn(flowid, "GetCookieData UnmarshalCookie error, Cookie = %s, err = %s", Cookie, err.Error())
	// 	return stCookieInfo, err
	// }

	// //查询Cookie对应的数据是否存在
	// RedisCookieData, err := redis.HashGet("CookieEx", strconv.FormatInt(stCookieInfo.OpenId, 10), strconv.FormatInt(int64(stCookieInfo.Idx), 10))
	// if err != nil {
	// 	utils.Warn(flowid, "GetCookieData HashGet CookieEx error, stCookieInfo = %v, err = %s", stCookieInfo, err.Error())
	// 	return stCookieInfo, err
	// }

	// if Cookie != RedisCookieData {
	// 	utils.Debug(flowid, "GetCookieData check cookie error, Cookie = %s, RedisCookieData = %s, CookieInfo = %v",
	// 		Cookie, RedisCookieData, stCookieInfo)
	// 	return stCookieInfo, errors.New("cookie expired")
	// }

	// return stCookieInfo, nil
}

//CreateUserCookie 创建用户cookie
func CreateUserCookie(WXOpenID string, WXUnionID string, NickName string, Url string, flowid int64) (string, error) {

	//获取当前时间
	CurTime := time.Now().Unix()

	//创建cookie信息
	stCookie := UserCookieInfo{
		WXOpenID:  WXOpenID,
		WXUnionID: WXUnionID,
		NickName:  NickName,
		Url:       Url,
		CurTime:   CurTime,
	}
	CookieData, err := json.Marshal(stCookie)
	if err != nil {
		utils.Warn(flowid, "CreateUserCookie Marshal error, WXOpenID = %s, WXUnionID = %s, NickName = %s, Url = %s, err = %s",
			WXOpenID, WXUnionID, NickName, Url, err.Error())
		return "", err
	}

	//读取加密key
	key := utils.GetConfigByString("Session::cookiekey")

	//生成cookie
	CookieSecret, err := utils.Encrypt(string(CookieData), key)
	if err != nil {
		utils.Warn(flowid, "CreateUserCookie Encrypt data error, CookieData = %s, err = %s", string(CookieData), err.Error())
		return "", err
	}

	utils.Debug(flowid, "CreateUserCookie debug, CookieData = %s, CookieSecret = %s", string(CookieData), CookieSecret)

	return CookieSecret, nil

}
