// Code generated by protoc-gen-go. DO NOT EDIT.
// source: RecommendSvr.proto

package client

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ERecommendTypeDef int32

const (
	ERecommendTypeDef_BANNER_TYPE    ERecommendTypeDef = 1
	ERecommendTypeDef_NEW_TYPE       ERecommendTypeDef = 2
	ERecommendTypeDef_RECOMMEND_TYPE ERecommendTypeDef = 3
)

var ERecommendTypeDef_name = map[int32]string{
	1: "BANNER_TYPE",
	2: "NEW_TYPE",
	3: "RECOMMEND_TYPE",
}
var ERecommendTypeDef_value = map[string]int32{
	"BANNER_TYPE":    1,
	"NEW_TYPE":       2,
	"RECOMMEND_TYPE": 3,
}

func (x ERecommendTypeDef) Enum() *ERecommendTypeDef {
	p := new(ERecommendTypeDef)
	*p = x
	return p
}
func (x ERecommendTypeDef) String() string {
	return proto.EnumName(ERecommendTypeDef_name, int32(x))
}
func (x *ERecommendTypeDef) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ERecommendTypeDef_value, data, "ERecommendTypeDef")
	if err != nil {
		return err
	}
	*x = ERecommendTypeDef(value)
	return nil
}
func (ERecommendTypeDef) EnumDescriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

// 推荐系统协议
type STImageListInfo struct {
	AlbumName        *string `protobuf:"bytes,1,opt,name=AlbumName" json:"AlbumName,omitempty"`
	StyleID          *int32  `protobuf:"zigzag32,2,opt,name=StyleID" json:"StyleID,omitempty"`
	Url              *string `protobuf:"bytes,3,opt,name=Url" json:"Url,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *STImageListInfo) Reset()                    { *m = STImageListInfo{} }
func (m *STImageListInfo) String() string            { return proto.CompactTextString(m) }
func (*STImageListInfo) ProtoMessage()               {}
func (*STImageListInfo) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *STImageListInfo) GetAlbumName() string {
	if m != nil && m.AlbumName != nil {
		return *m.AlbumName
	}
	return ""
}

func (m *STImageListInfo) GetStyleID() int32 {
	if m != nil && m.StyleID != nil {
		return *m.StyleID
	}
	return 0
}

func (m *STImageListInfo) GetUrl() string {
	if m != nil && m.Url != nil {
		return *m.Url
	}
	return ""
}

// 查询推荐列表请求
type STQueryRecommendListReq struct {
	QueryStartPos    *int32 `protobuf:"zigzag32,1,opt,name=QueryStartPos" json:"QueryStartPos,omitempty"`
	QueryNumber      *int32 `protobuf:"zigzag32,2,opt,name=QueryNumber" json:"QueryNumber,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *STQueryRecommendListReq) Reset()                    { *m = STQueryRecommendListReq{} }
func (m *STQueryRecommendListReq) String() string            { return proto.CompactTextString(m) }
func (*STQueryRecommendListReq) ProtoMessage()               {}
func (*STQueryRecommendListReq) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *STQueryRecommendListReq) GetQueryStartPos() int32 {
	if m != nil && m.QueryStartPos != nil {
		return *m.QueryStartPos
	}
	return 0
}

func (m *STQueryRecommendListReq) GetQueryNumber() int32 {
	if m != nil && m.QueryNumber != nil {
		return *m.QueryNumber
	}
	return 0
}

// 查询推荐列表响应
type STQueryRecommendListRsp struct {
	BannerList         []*STImageListInfo `protobuf:"bytes,1,rep,name=BannerList" json:"BannerList,omitempty"`
	NewImageList       []*STImageListInfo `protobuf:"bytes,2,rep,name=NewImageList" json:"NewImageList,omitempty"`
	RecommendImageList []*STImageListInfo `protobuf:"bytes,3,rep,name=RecommendImageList" json:"RecommendImageList,omitempty"`
	XXX_unrecognized   []byte             `json:"-"`
}

func (m *STQueryRecommendListRsp) Reset()                    { *m = STQueryRecommendListRsp{} }
func (m *STQueryRecommendListRsp) String() string            { return proto.CompactTextString(m) }
func (*STQueryRecommendListRsp) ProtoMessage()               {}
func (*STQueryRecommendListRsp) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

func (m *STQueryRecommendListRsp) GetBannerList() []*STImageListInfo {
	if m != nil {
		return m.BannerList
	}
	return nil
}

func (m *STQueryRecommendListRsp) GetNewImageList() []*STImageListInfo {
	if m != nil {
		return m.NewImageList
	}
	return nil
}

func (m *STQueryRecommendListRsp) GetRecommendImageList() []*STImageListInfo {
	if m != nil {
		return m.RecommendImageList
	}
	return nil
}

// 查询相册id对应的图片信息请求
type STStyleImageListReq struct {
	StyleID          *int32 `protobuf:"zigzag32,1,req,name=StyleID" json:"StyleID,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *STStyleImageListReq) Reset()                    { *m = STStyleImageListReq{} }
func (m *STStyleImageListReq) String() string            { return proto.CompactTextString(m) }
func (*STStyleImageListReq) ProtoMessage()               {}
func (*STStyleImageListReq) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{3} }

func (m *STStyleImageListReq) GetStyleID() int32 {
	if m != nil && m.StyleID != nil {
		return *m.StyleID
	}
	return 0
}

// 查询相册id对应的图片信息响应
type STStyleImageListRsp struct {
	Urls             []string `protobuf:"bytes,1,rep,name=Urls" json:"Urls,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *STStyleImageListRsp) Reset()                    { *m = STStyleImageListRsp{} }
func (m *STStyleImageListRsp) String() string            { return proto.CompactTextString(m) }
func (*STStyleImageListRsp) ProtoMessage()               {}
func (*STStyleImageListRsp) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{4} }

func (m *STStyleImageListRsp) GetUrls() []string {
	if m != nil {
		return m.Urls
	}
	return nil
}

func init() {
	proto.RegisterType((*STImageListInfo)(nil), "client.STImageListInfo")
	proto.RegisterType((*STQueryRecommendListReq)(nil), "client.STQueryRecommendListReq")
	proto.RegisterType((*STQueryRecommendListRsp)(nil), "client.STQueryRecommendListRsp")
	proto.RegisterType((*STStyleImageListReq)(nil), "client.STStyleImageListReq")
	proto.RegisterType((*STStyleImageListRsp)(nil), "client.STStyleImageListRsp")
	proto.RegisterEnum("client.ERecommendTypeDef", ERecommendTypeDef_name, ERecommendTypeDef_value)
}

func init() { proto.RegisterFile("RecommendSvr.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 305 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xd1, 0x4a, 0xf3, 0x40,
	0x10, 0x85, 0xd9, 0xe4, 0xe7, 0xd7, 0x4e, 0xaa, 0x35, 0x53, 0xa4, 0xb9, 0x0c, 0x11, 0x24, 0x28,
	0xf6, 0x42, 0x9f, 0xa0, 0xb5, 0x2b, 0x14, 0xec, 0x5a, 0x93, 0x14, 0xf1, 0x4a, 0xd2, 0x3a, 0x95,
	0x42, 0x36, 0x89, 0x9b, 0xad, 0xd2, 0x57, 0xf2, 0x29, 0x25, 0x09, 0xc6, 0x22, 0xda, 0xcb, 0xfd,
	0x66, 0xce, 0xd9, 0x39, 0x07, 0x30, 0xa0, 0x45, 0x26, 0x25, 0xa5, 0xcf, 0xe1, 0x9b, 0xea, 0xe7,
	0x2a, 0xd3, 0x19, 0xfe, 0x5f, 0x24, 0x2b, 0x4a, 0xb5, 0x37, 0x84, 0x4e, 0x18, 0x8d, 0x65, 0xfc,
	0x42, 0xb7, 0xab, 0x42, 0x8f, 0xd3, 0x65, 0x86, 0x36, 0xb4, 0x06, 0xc9, 0x7c, 0x2d, 0x45, 0x2c,
	0xc9, 0x61, 0x2e, 0xf3, 0x5b, 0xd8, 0x81, 0xbd, 0x50, 0x6f, 0x12, 0x1a, 0x8f, 0x1c, 0xc3, 0x65,
	0xbe, 0x8d, 0x16, 0x98, 0x33, 0x95, 0x38, 0x66, 0x39, 0xf5, 0x38, 0xf4, 0xc2, 0xe8, 0x7e, 0x4d,
	0x6a, 0xd3, 0x7c, 0x54, 0x9a, 0x05, 0xf4, 0x8a, 0xc7, 0x70, 0x50, 0x0d, 0x42, 0x1d, 0x2b, 0x3d,
	0xcd, 0x8a, 0xca, 0xcf, 0xc6, 0x2e, 0x58, 0x15, 0x16, 0x6b, 0x39, 0x27, 0x55, 0x7b, 0x7a, 0x1f,
	0xec, 0x0f, 0x9f, 0x22, 0xc7, 0x73, 0x80, 0x61, 0x9c, 0xa6, 0xa4, 0x4a, 0xe0, 0x30, 0xd7, 0xf4,
	0xad, 0xcb, 0x5e, 0xbf, 0xce, 0xd0, 0xff, 0x19, 0xe0, 0x02, 0xda, 0x82, 0xde, 0x1b, 0xe6, 0x18,
	0xbb, 0xd7, 0xaf, 0xb6, 0x0a, 0xfa, 0x16, 0x99, 0x3b, 0x45, 0xde, 0x29, 0x74, 0xc3, 0xa8, 0xee,
	0xe4, 0x8b, 0x97, 0x79, 0xb7, 0x8a, 0x62, 0xae, 0xe1, 0xdb, 0xde, 0xc9, 0x2f, 0x7b, 0x45, 0x8e,
	0x6d, 0xf8, 0x37, 0x53, 0x49, 0x51, 0x25, 0x69, 0x9d, 0xdd, 0x80, 0xcd, 0x9b, 0x13, 0xa2, 0x4d,
	0x4e, 0x23, 0x5a, 0x62, 0x07, 0xac, 0xe1, 0x40, 0x08, 0x1e, 0x3c, 0x45, 0x8f, 0x53, 0x7e, 0xc4,
	0xb0, 0x0d, 0xfb, 0x82, 0x3f, 0xd4, 0x2f, 0x03, 0x11, 0x0e, 0x03, 0x7e, 0x7d, 0x37, 0x99, 0x70,
	0x31, 0xaa, 0x99, 0xf9, 0x19, 0x00, 0x00, 0xff, 0xff, 0x98, 0x59, 0x27, 0x81, 0xe9, 0x01, 0x00,
	0x00,
}
