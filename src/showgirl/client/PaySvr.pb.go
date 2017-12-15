// Code generated by protoc-gen-go. DO NOT EDIT.
// source: PaySvr.proto

package client

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// 创建微信支付订单请求
type STCreateTransactionReq struct {
	GoodsDesc        *string `protobuf:"bytes,1,opt,name=GoodsDesc" json:"GoodsDesc,omitempty"`
	GoodsDetail      *string `protobuf:"bytes,2,opt,name=GoodsDetail" json:"GoodsDetail,omitempty"`
	FeeAmount        *int32  `protobuf:"zigzag32,3,opt,name=FeeAmount" json:"FeeAmount,omitempty"`
	OpenID           *string `protobuf:"bytes,4,opt,name=OpenID" json:"OpenID,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *STCreateTransactionReq) Reset()                    { *m = STCreateTransactionReq{} }
func (m *STCreateTransactionReq) String() string            { return proto.CompactTextString(m) }
func (*STCreateTransactionReq) ProtoMessage()               {}
func (*STCreateTransactionReq) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *STCreateTransactionReq) GetGoodsDesc() string {
	if m != nil && m.GoodsDesc != nil {
		return *m.GoodsDesc
	}
	return ""
}

func (m *STCreateTransactionReq) GetGoodsDetail() string {
	if m != nil && m.GoodsDetail != nil {
		return *m.GoodsDetail
	}
	return ""
}

func (m *STCreateTransactionReq) GetFeeAmount() int32 {
	if m != nil && m.FeeAmount != nil {
		return *m.FeeAmount
	}
	return 0
}

func (m *STCreateTransactionReq) GetOpenID() string {
	if m != nil && m.OpenID != nil {
		return *m.OpenID
	}
	return ""
}

// 创建微信支付订单响应
type STCreateTransactionRsp struct {
	AppId            *string `protobuf:"bytes,1,opt,name=AppId" json:"AppId,omitempty"`
	PartnerId        *string `protobuf:"bytes,2,opt,name=PartnerId" json:"PartnerId,omitempty"`
	DeviceInfo       *string `protobuf:"bytes,3,opt,name=DeviceInfo" json:"DeviceInfo,omitempty"`
	NonceStr         *string `protobuf:"bytes,4,opt,name=NonceStr" json:"NonceStr,omitempty"`
	Sign             *string `protobuf:"bytes,5,opt,name=Sign" json:"Sign,omitempty"`
	ResultCode       *string `protobuf:"bytes,6,opt,name=ResultCode" json:"ResultCode,omitempty"`
	ErrCode          *string `protobuf:"bytes,7,opt,name=ErrCode" json:"ErrCode,omitempty"`
	ErrCodeDes       *string `protobuf:"bytes,8,opt,name=ErrCodeDes" json:"ErrCodeDes,omitempty"`
	TradeType        *string `protobuf:"bytes,9,opt,name=TradeType" json:"TradeType,omitempty"`
	PrepayID         *string `protobuf:"bytes,10,opt,name=PrepayID" json:"PrepayID,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *STCreateTransactionRsp) Reset()                    { *m = STCreateTransactionRsp{} }
func (m *STCreateTransactionRsp) String() string            { return proto.CompactTextString(m) }
func (*STCreateTransactionRsp) ProtoMessage()               {}
func (*STCreateTransactionRsp) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *STCreateTransactionRsp) GetAppId() string {
	if m != nil && m.AppId != nil {
		return *m.AppId
	}
	return ""
}

func (m *STCreateTransactionRsp) GetPartnerId() string {
	if m != nil && m.PartnerId != nil {
		return *m.PartnerId
	}
	return ""
}

func (m *STCreateTransactionRsp) GetDeviceInfo() string {
	if m != nil && m.DeviceInfo != nil {
		return *m.DeviceInfo
	}
	return ""
}

func (m *STCreateTransactionRsp) GetNonceStr() string {
	if m != nil && m.NonceStr != nil {
		return *m.NonceStr
	}
	return ""
}

func (m *STCreateTransactionRsp) GetSign() string {
	if m != nil && m.Sign != nil {
		return *m.Sign
	}
	return ""
}

func (m *STCreateTransactionRsp) GetResultCode() string {
	if m != nil && m.ResultCode != nil {
		return *m.ResultCode
	}
	return ""
}

func (m *STCreateTransactionRsp) GetErrCode() string {
	if m != nil && m.ErrCode != nil {
		return *m.ErrCode
	}
	return ""
}

func (m *STCreateTransactionRsp) GetErrCodeDes() string {
	if m != nil && m.ErrCodeDes != nil {
		return *m.ErrCodeDes
	}
	return ""
}

func (m *STCreateTransactionRsp) GetTradeType() string {
	if m != nil && m.TradeType != nil {
		return *m.TradeType
	}
	return ""
}

func (m *STCreateTransactionRsp) GetPrepayID() string {
	if m != nil && m.PrepayID != nil {
		return *m.PrepayID
	}
	return ""
}

func init() {
	proto.RegisterType((*STCreateTransactionReq)(nil), "client.STCreateTransactionReq")
	proto.RegisterType((*STCreateTransactionRsp)(nil), "client.STCreateTransactionRsp")
}

func init() { proto.RegisterFile("PaySvr.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 246 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0xcd, 0x4d, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x61, 0x05, 0xda, 0xb4, 0x1e, 0xca, 0x4f, 0x8c, 0x84, 0xbc, 0xac, 0xba, 0xea, 0x8a,
	0x3b, 0x54, 0x35, 0x20, 0x6f, 0x20, 0x6a, 0x72, 0x01, 0xcb, 0x1e, 0x90, 0xa5, 0x60, 0x9b, 0xb1,
	0x5b, 0x29, 0x77, 0xe4, 0x50, 0xa8, 0x21, 0xd9, 0x75, 0xfb, 0xea, 0x9b, 0x79, 0x60, 0x55, 0xeb,
	0xbe, 0x39, 0xd1, 0x73, 0xa4, 0x90, 0x03, 0x2f, 0x4d, 0xe7, 0xd0, 0xe7, 0x8d, 0x81, 0xa7, 0xa6,
	0xdd, 0x13, 0xea, 0x8c, 0x2d, 0x69, 0x9f, 0xb4, 0xc9, 0x2e, 0xf8, 0x03, 0xfe, 0xf0, 0x0a, 0xd8,
	0x5b, 0x08, 0x36, 0x49, 0x4c, 0x46, 0x14, 0xeb, 0x62, 0xcb, 0xf8, 0x23, 0xdc, 0x8c, 0x29, 0x6b,
	0xd7, 0x89, 0xab, 0x21, 0x56, 0xc0, 0x5e, 0x11, 0x77, 0xdf, 0xe1, 0xe8, 0xb3, 0xb8, 0x5e, 0x17,
	0xdb, 0x8a, 0xdf, 0x41, 0xf9, 0x11, 0xd1, 0x2b, 0x29, 0x66, 0xe7, 0xc9, 0xe6, 0xb7, 0xb8, 0xac,
	0xa4, 0xc8, 0x6f, 0x61, 0xbe, 0x8b, 0x51, 0xd9, 0x51, 0xa8, 0x80, 0xd5, 0x9a, 0xb2, 0x47, 0x52,
	0x76, 0xfc, 0xcf, 0x01, 0x24, 0x9e, 0x9c, 0x41, 0xe5, 0x3f, 0xc3, 0x00, 0x30, 0xfe, 0x00, 0xcb,
	0xf7, 0xe0, 0x0d, 0x36, 0x99, 0xfe, 0x09, 0xbe, 0x82, 0x59, 0xe3, 0xbe, 0xbc, 0x98, 0x4f, 0x37,
	0x07, 0x4c, 0xc7, 0x2e, 0xef, 0x83, 0x45, 0x51, 0x0e, 0xed, 0x1e, 0x16, 0x2f, 0x44, 0x43, 0x58,
	0x4c, 0xa3, 0x31, 0x48, 0x4c, 0x62, 0x39, 0xf9, 0x2d, 0x69, 0x8b, 0x6d, 0x1f, 0x51, 0xb0, 0xc9,
	0xaa, 0x09, 0xa3, 0xee, 0x95, 0x14, 0x70, 0x2e, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa2, 0xfa,
	0xae, 0x9a, 0x4a, 0x01, 0x00, 0x00,
}