package Image

import (
	"fmt"
	"showgirl/client"
	"showgirl/models/mysql"
	"showgirl/models/utils"
	"strings"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/golang/protobuf/proto"
)

//QueryStyleInfoList 查询相册列表
func QueryStyleInfoList(QueryBegin int32, QueryNum int32, flowid int64) ([]*client.STStyleInfo, int32, error) {

	//新建mysql实例
	o := mysql.NewShowgirlOrm()

	var DBID []int32
	var DBShowName []string
	var DBRecommendCategory []int32
	var DBCreateTime []int64

	sSQL := "select Id,ShowName,RecommendCategory,CreateTime from StyleShowInfo order by RecommendCategory asc,CreateTime desc"

	num, err := o.Raw(sSQL).QueryRows(&DBID, &DBShowName, &DBRecommendCategory, &DBCreateTime)
	if err != nil {
		utils.Debug(flowid, "QueryStyleInfoList QueryRows error, sql = %s, err = %s",
			sSQL, err.Error())
		return nil, 0, err
	}

	if num <= 0 {
		utils.Debug(flowid, "QueryStyleInfoList no data, sql = %s",
			sSQL)
		return nil, 0, nil
	}

	arry := []*client.STStyleInfo{}

	for idx := range DBID {
		stStyleInfo := &client.STStyleInfo{
			StyleID:    proto.Int32(DBID[idx]),
			StyleName:  proto.String(DBShowName[idx]),
			StyleType:  client.ERecommendTypeDef(DBRecommendCategory[idx]).Enum(),
			CreateTime: proto.Int64(DBCreateTime[idx]),
		}
		arry = append(arry, stStyleInfo)
	}

	//查询总数
	var DBTotal []int32

	sTotalSQL := "select count(1) from StyleShowInfo"
	num, err = o.Raw(sTotalSQL).QueryRows(&DBTotal)
	if err != nil {
		utils.Debug(flowid, "QueryStyleInfoList QueryRows total error, sql = %s, err = %s",
			sSQL, err.Error())
		return nil, 0, err
	}

	if num <= 0 {
		utils.Debug(flowid, "QueryStyleInfoList no data, sql = %s",
			sSQL)
		return nil, 0, nil
	}

	return arry, DBTotal[0], nil

}

//QueryResourceInfoList 查询相册资源列表
func QueryResourceInfoList(StyleID int32, flowid int64) ([]*client.STResourceImageInfo, error) {

	//新建mysql实例
	o := mysql.NewShowgirlOrm()

	var DBID []int32
	var DBStyleURL []string

	sSQL := fmt.Sprintf("select Id,StyleUrl from StyleContentInfo where ShowID = %d order by Id asc", StyleID)

	num, err := o.Raw(sSQL).QueryRows(&DBID, &DBStyleURL)
	if err != nil {
		utils.Debug(flowid, "QueryResourceInfoList QueryRows error, sql = %s, err = %s",
			sSQL, err.Error())
		return nil, err
	}

	if num <= 0 {
		utils.Debug(flowid, "QueryResourceInfoList no data, sql = %s",
			sSQL)
		return nil, nil
	}

	arry := []*client.STResourceImageInfo{}

	for idx := range DBID {
		stResourceInfo := &client.STResourceImageInfo{
			ImageID: proto.Int32(DBID[idx]),
			Url:     proto.String(DBStyleURL[idx]),
		}
		arry = append(arry, stResourceInfo)
	}

	return arry, nil

}

//CreateStyle 创建相册
func CreateStyle(StyleName string, StyleType client.ERecommendTypeDef, flowid int64) error {

	//新建mysql实例
	o := mysql.NewShowgirlOrm()

	sSQL := fmt.Sprintf("insert into StyleShowInfo(Id,ShowName,CreateTime,UpdateTime,RecommendCategory) values(null,%q,%d,Now(),%d)",
		StyleName, time.Now().Unix(), int32(StyleType))
	_, err := o.Raw(sSQL).Exec()
	if err != nil {
		utils.Warn(flowid, "CreateStyle error, sql = %s, err = %s", sSQL, err.Error())
		return err
	}

	return nil
}

//DeleteStyle 删除相册
func DeleteStyle(StyleID int32, flowid int64) error {

	//新建mysql实例
	o := mysql.NewShowgirlOrm()

	sSQL := fmt.Sprintf("delete from StyleShowInfo where Id = %d", StyleID)
	_, err := o.Raw(sSQL).Exec()
	if err != nil {
		utils.Warn(flowid, "DeleteStyle error, sql = %s, err = %s", sSQL, err.Error())
		return err
	}

	return nil
}

//DeleteResource 删除相册图片资源
func DeleteResource(ResourceID int32, flowid int64) error {

	//新建mysql实例
	o := mysql.NewShowgirlOrm()

	sSQL := fmt.Sprintf("delete from StyleContentInfo where Id = %d", ResourceID)
	_, err := o.Raw(sSQL).Exec()
	if err != nil {
		utils.Warn(flowid, "DeleteResource error, sql = %s, err = %s", sSQL, err.Error())
		return err
	}

	return nil
}

//UploadAndSetImage 上传并设置图片资源
func UploadAndSetImage(strImage string, StyleID int32, flowid int64) error {

	imageURL, err := UploadImage(strImage, flowid)
	if err != nil {
		utils.Warn(flowid, "UploadAndSetImage UploadImage error, strImage = %s, err = %s", strImage, err.Error())
		return err
	}

	//新建mysql实例
	o := mysql.NewShowgirlOrm()

	sSQL := fmt.Sprintf("insert into StyleContentInfo(Id,StyleUrl,ShowID,CreateTime,UpdateTime) values(null,%q,%d,%d,Now())",
		imageURL, StyleID, time.Now().Unix())
	_, err = o.Raw(sSQL).Exec()
	if err != nil {
		utils.Warn(flowid, "UploadAndSetImage error, sql = %s, err = %s", sSQL, err.Error())
		return err
	}

	return nil
}

//UploadImage 上传图片
func UploadImage(strImage string, flowid int64) (string, error) {

	client, err := oss.New("https://oss-cn-shanghai.aliyuncs.com", "LTAIxR05Gfh0uBtU", "AccessKeySecret")
	if err != nil {
		utils.Warn(flowid, "UploadImage new oss error, err = %s", err.Error())
		return "", err
	}

	bucket, err := client.Bucket("girlstyle")
	if err != nil {
		utils.Warn(flowid, "UploadImage new bucket error, err = %s", err.Error())
		return "", err
	}

	objectName := string(utils.Krand(16, utils.KC_RAND_KIND_ALL))

	request := &oss.PutObjectRequest{
		ObjectKey: objectName,
		Reader:    strings.NewReader(strImage),
	}
	resp, err := bucket.DoPutObject(request, nil)
	if err != nil {
		utils.Warn(flowid, "UploadImage DoPutObject error, err = %s", err.Error())
		return "", err
	}
	defer resp.Body.Close()

	return "https://girlstyle.oss-cn-shanghai.aliyuncs.com/" + objectName, nil
}
