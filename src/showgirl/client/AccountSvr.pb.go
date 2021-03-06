// Code generated by protoc-gen-go. DO NOT EDIT.
// source: AccountSvr.proto

/*
Package client is a generated protocol buffer package.

It is generated from these files:
	AccountSvr.proto
	CommonProtocol.proto
	ImageSvr.proto
	PaySvr.proto
	RecommendSvr.proto

It has these top-level messages:
	STThirdPartyWXLoginReq
	STThirdPartyWXLoginRsp
	STGetMyInfoReq
	STGetMyInfoRsp
	STUserTrustInfo
	STCookieInfo
	STRspHeader
	CommonReq
	CommonRsp
	CommonClientRsp
	STQueryStyleListReq
	STStyleInfo
	STQueryStyleListRsp
	STQueryResourceListReq
	STResourceImageInfo
	STQueryResourceListRsp
	STCreateStyleReq
	STCreateStyleRsp
	STUploadImageReq
	STUploadImageRsp
	STDeleteStyleReq
	STDeleteStyleRsp
	STDeleteResourceReq
	STDeleteResourceRsp
	STUpdateStyleReq
	STUpdateStyleRsp
	STCreateTransactionReq
	STCreateTransactionRsp
	STImageListInfo
	STQueryRecommendListReq
	STQueryRecommendListRsp
	STStyleImageListReq
	STStyleImageListRsp
*/
package client

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// 微信登录
type STThirdPartyWXLoginReq struct {
	Code             *string `protobuf:"bytes,1,req,name=code" json:"code,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *STThirdPartyWXLoginReq) Reset()                    { *m = STThirdPartyWXLoginReq{} }
func (m *STThirdPartyWXLoginReq) String() string            { return proto.CompactTextString(m) }
func (*STThirdPartyWXLoginReq) ProtoMessage()               {}
func (*STThirdPartyWXLoginReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *STThirdPartyWXLoginReq) GetCode() string {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return ""
}

type STThirdPartyWXLoginRsp struct {
	UserKey          *string `protobuf:"bytes,1,req,name=UserKey" json:"UserKey,omitempty"`
	UserId           *string `protobuf:"bytes,2,opt,name=UserId" json:"UserId,omitempty"`
	NickName         *string `protobuf:"bytes,3,opt,name=NickName" json:"NickName,omitempty"`
	Url              *string `protobuf:"bytes,4,opt,name=Url" json:"Url,omitempty"`
	LastTime         *int64  `protobuf:"zigzag64,5,opt,name=LastTime" json:"LastTime,omitempty"`
	ChargeNum        *int32  `protobuf:"zigzag32,6,opt,name=ChargeNum" json:"ChargeNum,omitempty"`
	WxOpenID         *string `protobuf:"bytes,7,opt,name=WxOpenID" json:"WxOpenID,omitempty"`
	VipDeadLine      *int64  `protobuf:"zigzag64,8,opt,name=VipDeadLine" json:"VipDeadLine,omitempty"`
	BVip             *bool   `protobuf:"varint,9,opt,name=bVip" json:"bVip,omitempty"`
	PayControl       *bool   `protobuf:"varint,10,opt,name=PayControl" json:"PayControl,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *STThirdPartyWXLoginRsp) Reset()                    { *m = STThirdPartyWXLoginRsp{} }
func (m *STThirdPartyWXLoginRsp) String() string            { return proto.CompactTextString(m) }
func (*STThirdPartyWXLoginRsp) ProtoMessage()               {}
func (*STThirdPartyWXLoginRsp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *STThirdPartyWXLoginRsp) GetUserKey() string {
	if m != nil && m.UserKey != nil {
		return *m.UserKey
	}
	return ""
}

func (m *STThirdPartyWXLoginRsp) GetUserId() string {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return ""
}

func (m *STThirdPartyWXLoginRsp) GetNickName() string {
	if m != nil && m.NickName != nil {
		return *m.NickName
	}
	return ""
}

func (m *STThirdPartyWXLoginRsp) GetUrl() string {
	if m != nil && m.Url != nil {
		return *m.Url
	}
	return ""
}

func (m *STThirdPartyWXLoginRsp) GetLastTime() int64 {
	if m != nil && m.LastTime != nil {
		return *m.LastTime
	}
	return 0
}

func (m *STThirdPartyWXLoginRsp) GetChargeNum() int32 {
	if m != nil && m.ChargeNum != nil {
		return *m.ChargeNum
	}
	return 0
}

func (m *STThirdPartyWXLoginRsp) GetWxOpenID() string {
	if m != nil && m.WxOpenID != nil {
		return *m.WxOpenID
	}
	return ""
}

func (m *STThirdPartyWXLoginRsp) GetVipDeadLine() int64 {
	if m != nil && m.VipDeadLine != nil {
		return *m.VipDeadLine
	}
	return 0
}

func (m *STThirdPartyWXLoginRsp) GetBVip() bool {
	if m != nil && m.BVip != nil {
		return *m.BVip
	}
	return false
}

func (m *STThirdPartyWXLoginRsp) GetPayControl() bool {
	if m != nil && m.PayControl != nil {
		return *m.PayControl
	}
	return false
}

// 获取自己的信息
type STGetMyInfoReq struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *STGetMyInfoReq) Reset()                    { *m = STGetMyInfoReq{} }
func (m *STGetMyInfoReq) String() string            { return proto.CompactTextString(m) }
func (*STGetMyInfoReq) ProtoMessage()               {}
func (*STGetMyInfoReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type STGetMyInfoRsp struct {
	UserId           *string `protobuf:"bytes,1,opt,name=UserId" json:"UserId,omitempty"`
	NickName         *string `protobuf:"bytes,2,opt,name=NickName" json:"NickName,omitempty"`
	Url              *string `protobuf:"bytes,3,opt,name=Url" json:"Url,omitempty"`
	LastTime         *int64  `protobuf:"zigzag64,4,opt,name=LastTime" json:"LastTime,omitempty"`
	ChargeNum        *int32  `protobuf:"zigzag32,5,opt,name=ChargeNum" json:"ChargeNum,omitempty"`
	WxOpenID         *string `protobuf:"bytes,6,opt,name=WxOpenID" json:"WxOpenID,omitempty"`
	VipDeadLine      *int64  `protobuf:"zigzag64,7,opt,name=VipDeadLine" json:"VipDeadLine,omitempty"`
	BVip             *bool   `protobuf:"varint,8,opt,name=bVip" json:"bVip,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *STGetMyInfoRsp) Reset()                    { *m = STGetMyInfoRsp{} }
func (m *STGetMyInfoRsp) String() string            { return proto.CompactTextString(m) }
func (*STGetMyInfoRsp) ProtoMessage()               {}
func (*STGetMyInfoRsp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *STGetMyInfoRsp) GetUserId() string {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return ""
}

func (m *STGetMyInfoRsp) GetNickName() string {
	if m != nil && m.NickName != nil {
		return *m.NickName
	}
	return ""
}

func (m *STGetMyInfoRsp) GetUrl() string {
	if m != nil && m.Url != nil {
		return *m.Url
	}
	return ""
}

func (m *STGetMyInfoRsp) GetLastTime() int64 {
	if m != nil && m.LastTime != nil {
		return *m.LastTime
	}
	return 0
}

func (m *STGetMyInfoRsp) GetChargeNum() int32 {
	if m != nil && m.ChargeNum != nil {
		return *m.ChargeNum
	}
	return 0
}

func (m *STGetMyInfoRsp) GetWxOpenID() string {
	if m != nil && m.WxOpenID != nil {
		return *m.WxOpenID
	}
	return ""
}

func (m *STGetMyInfoRsp) GetVipDeadLine() int64 {
	if m != nil && m.VipDeadLine != nil {
		return *m.VipDeadLine
	}
	return 0
}

func (m *STGetMyInfoRsp) GetBVip() bool {
	if m != nil && m.BVip != nil {
		return *m.BVip
	}
	return false
}

func init() {
	proto.RegisterType((*STThirdPartyWXLoginReq)(nil), "client.STThirdPartyWXLoginReq")
	proto.RegisterType((*STThirdPartyWXLoginRsp)(nil), "client.STThirdPartyWXLoginRsp")
	proto.RegisterType((*STGetMyInfoReq)(nil), "client.STGetMyInfoReq")
	proto.RegisterType((*STGetMyInfoRsp)(nil), "client.STGetMyInfoRsp")
}

func init() { proto.RegisterFile("AccountSvr.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 285 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0xcf, 0xc1, 0x6a, 0xea, 0x40,
	0x14, 0xc6, 0x71, 0x46, 0x63, 0x4c, 0x8e, 0xe2, 0xd5, 0xb9, 0x50, 0x66, 0x19, 0xb2, 0x28, 0x59,
	0xf5, 0x1d, 0x8a, 0x42, 0x91, 0x5a, 0x2b, 0x26, 0x6a, 0xb7, 0xd3, 0xe4, 0x54, 0x87, 0x26, 0x33,
	0xe3, 0x64, 0x2c, 0xcd, 0xa3, 0xf4, 0x85, 0xfa, 0x5c, 0x25, 0x29, 0x05, 0x25, 0x6e, 0xff, 0xe7,
	0x3b, 0x8b, 0x1f, 0x8c, 0xef, 0xd3, 0x54, 0x9d, 0xa4, 0x8d, 0x3f, 0xcc, 0x9d, 0x36, 0xca, 0x2a,
	0xea, 0xa6, 0xb9, 0x40, 0x69, 0xc3, 0x5b, 0xb8, 0x89, 0x93, 0xe4, 0x20, 0x4c, 0xb6, 0xe2, 0xc6,
	0x56, 0xbb, 0x97, 0x85, 0xda, 0x0b, 0xb9, 0xc6, 0x23, 0x1d, 0x82, 0x93, 0xaa, 0x0c, 0x19, 0x09,
	0x3a, 0x91, 0x1f, 0x7e, 0x93, 0xeb, 0xc3, 0x52, 0xd3, 0x7f, 0xd0, 0xdf, 0x94, 0x68, 0x1e, 0xb1,
	0xfa, 0xdd, 0xd2, 0x11, 0xb8, 0x75, 0x98, 0x67, 0xac, 0x13, 0x90, 0xc8, 0xa7, 0x63, 0xf0, 0x96,
	0x22, 0x7d, 0x5f, 0xf2, 0x02, 0x59, 0xb7, 0x29, 0x03, 0xe8, 0x6e, 0x4c, 0xce, 0x9c, 0xbf, 0xf3,
	0x82, 0x97, 0x36, 0x11, 0x05, 0xb2, 0x5e, 0x40, 0x22, 0x4a, 0x27, 0xe0, 0x4f, 0x0f, 0xdc, 0xec,
	0x71, 0x79, 0x2a, 0x98, 0x1b, 0x90, 0x68, 0x52, 0x8f, 0x76, 0x9f, 0xcf, 0x1a, 0xe5, 0x7c, 0xc6,
	0xfa, 0xcd, 0xdb, 0x7f, 0x18, 0x6c, 0x85, 0x9e, 0x21, 0xcf, 0x16, 0x42, 0x22, 0xf3, 0x9a, 0xcf,
	0x21, 0x38, 0xaf, 0x5b, 0xa1, 0x99, 0x1f, 0x90, 0xc8, 0xa3, 0x14, 0x60, 0xc5, 0xab, 0xa9, 0x92,
	0xd6, 0xa8, 0x9c, 0x41, 0xdd, 0xc2, 0x31, 0x8c, 0xe2, 0xe4, 0x01, 0xed, 0x53, 0x35, 0x97, 0x6f,
	0x6a, 0x8d, 0xc7, 0xf0, 0x8b, 0x5c, 0xa6, 0x52, 0x9f, 0x09, 0x48, 0x4b, 0xd0, 0x39, 0x17, 0x74,
	0x5b, 0x02, 0xa7, 0x2d, 0xe8, 0xb5, 0x04, 0xee, 0x35, 0x41, 0xff, 0x42, 0x50, 0x7b, 0xbc, 0x9f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x92, 0x7c, 0x6f, 0xe2, 0xb9, 0x01, 0x00, 0x00,
}
