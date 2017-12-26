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
func (*STImageListInfo) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

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
func (*STQueryRecommendListReq) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

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
func (*STQueryRecommendListRsp) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2} }

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
func (*STStyleImageListReq) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{3} }

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
func (*STStyleImageListRsp) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{4} }

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
}

func init() { proto.RegisterFile("RecommendSvr.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 255 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xdd, 0x4a, 0xc3, 0x40,
	0x10, 0x85, 0xd9, 0x44, 0x94, 0x4c, 0x2a, 0x25, 0x5b, 0xa4, 0x7b, 0x19, 0x56, 0x90, 0x80, 0x98,
	0x0b, 0x7d, 0x02, 0x8b, 0x5e, 0x04, 0x24, 0x68, 0x37, 0x7d, 0x80, 0xb4, 0x8e, 0x52, 0xd8, 0x9f,
	0xb8, 0xbb, 0x51, 0xfa, 0x4a, 0x3e, 0xa5, 0x24, 0xc1, 0x5a, 0x44, 0x73, 0x7b, 0x66, 0xe6, 0x1b,
	0xbe, 0x03, 0x74, 0x89, 0x1b, 0xa3, 0x14, 0xea, 0x67, 0xf1, 0x6e, 0xf3, 0xc6, 0x1a, 0x6f, 0xe8,
	0xf1, 0x46, 0x6e, 0x51, 0x7b, 0xbe, 0x80, 0xa9, 0xa8, 0x0a, 0x55, 0xbf, 0xe2, 0xc3, 0xd6, 0xf9,
	0x42, 0xbf, 0x18, 0x9a, 0x40, 0x74, 0x2b, 0xd7, 0xad, 0x2a, 0x6b, 0x85, 0x8c, 0xa4, 0x24, 0x8b,
	0xe8, 0x14, 0x4e, 0x84, 0xdf, 0x49, 0x2c, 0xee, 0x58, 0x90, 0x92, 0x2c, 0xa1, 0x31, 0x84, 0x2b,
	0x2b, 0x59, 0xd8, 0x4d, 0xf9, 0x3d, 0xcc, 0x45, 0xf5, 0xd4, 0xa2, 0xdd, 0xed, 0x1f, 0x75, 0xb0,
	0x25, 0xbe, 0xd1, 0x33, 0x38, 0xed, 0x07, 0xc2, 0xd7, 0xd6, 0x3f, 0x1a, 0xd7, 0xf3, 0x12, 0x3a,
	0x83, 0xb8, 0x8f, 0xcb, 0x56, 0xad, 0xd1, 0x0e, 0x4c, 0xfe, 0x49, 0xfe, 0xe1, 0xb8, 0x86, 0x5e,
	0x02, 0x2c, 0x6a, 0xad, 0xd1, 0x76, 0x01, 0x23, 0x69, 0x98, 0xc5, 0xd7, 0xf3, 0x7c, 0x70, 0xc8,
	0x7f, 0x0b, 0x5c, 0xc1, 0xa4, 0xc4, 0x8f, 0x7d, 0xc6, 0x82, 0xf1, 0xf5, 0x9b, 0x83, 0x82, 0x7e,
	0x8e, 0xc2, 0xd1, 0x23, 0x7e, 0x01, 0x33, 0x51, 0x0d, 0x9d, 0x7c, 0xe7, 0x9d, 0xef, 0x41, 0x51,
	0x24, 0x0d, 0xb2, 0x84, 0x9f, 0xff, 0xb1, 0xe7, 0x1a, 0x3a, 0x81, 0xa3, 0x95, 0x95, 0xae, 0x37,
	0x89, 0xbe, 0x02, 0x00, 0x00, 0xff, 0xff, 0xaf, 0x21, 0xae, 0xc2, 0xa1, 0x01, 0x00, 0x00,
}
