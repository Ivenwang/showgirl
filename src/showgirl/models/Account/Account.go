package Account

import (
	"fmt"
	"showgirl/client"
	"showgirl/models/mysql"
	"showgirl/models/utils"
	"time"
)

//UserBaseInfo 用户基础数据
type UserBaseInfo struct {
	WxOpenID  string `json:"openid"`
	WxUnionID string `json:"unionid"`
	NickName  string `json:"nickname"`
	URL       string `json:"url"`
	Charge    int32  `json:"charge"`
	LastTime  int64  `json:"lasttime"`
}

//QueryAccountCorrectByUnionID 1表示查询不到账号，2表示DB错误
func QueryAccountCorrectByUnionID(WxUnionID string, flowid int64) (UserBaseInfo, client.EFindAccountDef) {

	baseInfo := UserBaseInfo{}

	if len(WxUnionID) <= 0 {
		return baseInfo, client.EFindAccountDef_ACCOUNT_PARAM_ERROR
	}

	//新建mysql实例
	o := mysql.NewShowgirlOrm()

	var DBWxOpenID []string
	var DBWxUnionID []string
	var DBNickName []string
	var DBUserURL []string
	var DBChargeNum []int32
	var DBLastTime []int64

	sSQL := fmt.Sprintf("select WxOpenID,WxUnionID,NickName,UserUrl,ChargeNum,unix_timestamp(UpdateTime) from ShowGirlAccountInfo where WxUnionID = %q",
		WxUnionID)

	num, err := o.Raw(sSQL).QueryRows(&DBWxOpenID, &DBWxUnionID, &DBNickName, &DBUserURL,
		&DBChargeNum, &DBLastTime)
	if err != nil {
		utils.Debug(flowid, "QueryAccountCorrectByUnionID QueryRows error, sql = %s, err = %s",
			sSQL, err.Error())
		return baseInfo, client.EFindAccountDef_ACCOUNT_SYSTEM_ERROR
	} else if num <= 0 {
		utils.Debug(flowid, "QueryAccountCorrectByUnionID no found account, sql = %s",
			sSQL)
		return baseInfo, client.EFindAccountDef_NOT_FOUND_ACCOUNT
	}

	utils.Debug(flowid, "QueryAccountCorrectByUnionID debug, sql = %s", sSQL)

	baseInfo.WxOpenID = DBWxOpenID[0]
	baseInfo.WxUnionID = DBWxUnionID[0]
	baseInfo.NickName = DBNickName[0]
	baseInfo.URL = DBUserURL[0]
	baseInfo.Charge = DBChargeNum[0]
	baseInfo.LastTime = DBLastTime[0]

	return baseInfo, client.EFindAccountDef_FOUND_ACCOUNT_SUCCESS
}

//RegisterAccount 插入账号
func RegisterAccount(WxOpenID string, WxUnionID string, NickName string, URL string, flowid int64) error {

	//新建mysql实例
	o := mysql.NewShowgirlOrm()

	CurTime := time.Now().Unix()

	sSQL := fmt.Sprintf("insert into ShowGirlAccountInfo(Id,WxOpenID,WxUnionID,NickName,UserUrl,ChargeNum,CreateTime,UpdateTime)"+
		" values(null,%q,%q,%q,%q,%d,%d,Now())", WxOpenID, WxUnionID, NickName, URL, 0, CurTime)
	_, err := o.Raw(sSQL).Exec()
	if err != nil {
		utils.Warn(flowid, "RegisterAccount insert account error, sql = %s, err = %s", sSQL, err.Error())
		return err
	}

	return nil
}

//UpdateLastLoginTime 更新登录时间
func UpdateLastLoginTime(WXUnionID string, flowid int64) error {
	//新建mysql实例
	o := mysql.NewShowgirlOrm()

	sSQL := fmt.Sprintf("update ShowGirlAccountInfo set UpdateTime=CURRENT_TIMESTAMP where WxUnionID = %q",
		WXUnionID)

	_, err := o.Raw(sSQL).Exec()
	if err != nil {
		utils.Warn(flowid, "UpdateLastLoginTime update error, err = %s, sql = %s", err.Error(), sSQL)
		return err
	}

	utils.Debug(flowid, "UpdateLastLoginTime update succ, sql = %s", sSQL)
	return nil

}
