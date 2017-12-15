package Recommend

import (
	"showgirl/client"
	"showgirl/models/mysql"
	"showgirl/models/utils"
	"strconv"

	"github.com/gogo/protobuf/proto"
)

type CategaryArryStyleInfo struct {
	Category  int32
	ArryStyle []*client.STImageListInfo
}

type StyleBaseInfo struct {
	StyleName string //名称
	Url       string //封面图
}

//QueryImageListByRecommendCategory 根据分类查询个人秀列表
func QueryImageListByRecommendCategory(CategoryList []int32, flowid int64) ([]CategaryArryStyleInfo, error) {

	arry := []CategaryArryStyleInfo{}

	if len(CategoryList) <= 0 {
		return arry, nil
	}

	//新建mysql实例
	o := mysql.NewShowgirlOrm()

	var DBStyleURL []string
	var DBStyleID []int32
	var DBShowName []string
	var DBCategory []int32

	sQuerySQL := "select StyleUrl,StyleShowInfo.Id,StyleShowInfo.ShowName,RecommendCategory" +
		" from StyleContentInfo,StyleShowInfo where" +
		" StyleContentInfo.ShowID=StyleShowInfo.Id and" +
		" StyleContentInfo.RecommendCategory in ("

	for idx := range CategoryList {
		sQuerySQL += strconv.FormatInt(int64(CategoryList[idx]), 10)
		sQuerySQL += ","
	}

	sQuerySQL = sQuerySQL[:len(sQuerySQL)-1]
	sQuerySQL += ") group by StyleShowInfo.Id"

	utils.Debug(flowid, "QueryImageListByRecommendCategory debug, sql = %s", sQuerySQL)

	num, err := o.Raw(sQuerySQL).QueryRows(&DBStyleURL, &DBStyleID, &DBShowName, &DBCategory)
	if err != nil {
		utils.Warn(flowid, "QueryImageListByRecommendCategory QueryRows error, sql = %s, err = %s",
			sQuerySQL, err.Error())
		return arry, err
	}

	if num <= 0 {
		utils.Debug(flowid, "QueryImageListByRecommendCategory no found account, sql = %s",
			sQuerySQL)
		return arry, nil
	}

	RecommendMap := make(map[int32]map[int32]StyleBaseInfo)

	for idx := range DBStyleURL {

		mapName, ok := RecommendMap[DBCategory[idx]]
		if !ok {
			mapName = make(map[int32]StyleBaseInfo)
			RecommendMap[DBCategory[idx]] = mapName
		}

		mapName[DBStyleID[idx]] = StyleBaseInfo{StyleName: DBShowName[idx], Url: DBStyleURL[idx]}
	}

	for Category, value := range RecommendMap {
		stCategoryInfo := CategaryArryStyleInfo{}
		stCategoryInfo.Category = Category
		for StyleID, BaseInfo := range value {
			ImageInfo := &client.STImageListInfo{}
			ImageInfo.StyleID = proto.Int32(StyleID)
			ImageInfo.AlbumName = proto.String(BaseInfo.StyleName)
			ImageInfo.Url = proto.String(BaseInfo.Url)
			stCategoryInfo.ArryStyle = append(stCategoryInfo.ArryStyle, ImageInfo)
		}

		arry = append(arry, stCategoryInfo)
	}

	return arry, nil
}

//QueryImageListByStyleID 根据StyleID查询图片列表
func QueryImageListByStyleID(StyleID int32, flowid int64) ([]string, error) {

	//新建mysql实例
	o := mysql.NewShowgirlOrm()

	var DBStyleURL []string

	sQuerySQL := "select StyleUrl from StyleContentInfo where ShowID = " + strconv.FormatInt(int64(StyleID), 10)

	utils.Debug(flowid, "QueryImageListByStyleID debug, sql = %s", sQuerySQL)

	num, err := o.Raw(sQuerySQL).QueryRows(&DBStyleURL)
	if err != nil {
		utils.Warn(flowid, "QueryImageListByStyleID QueryRows error, sql = %s, err = %s",
			sQuerySQL, err.Error())
		return []string{}, err
	}

	if num <= 0 {
		utils.Debug(flowid, "QueryImageListByStyleID no found data, sql = %s",
			sQuerySQL)
		return []string{}, nil
	}

	return DBStyleURL, nil
}
