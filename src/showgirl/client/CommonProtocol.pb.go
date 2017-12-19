// Code generated by protoc-gen-go. DO NOT EDIT.
// source: CommonProtocol.proto

package client

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type EErrorTypeDef int32

const (
	EErrorTypeDef_RESULT_OK               EErrorTypeDef = 200
	EErrorTypeDef_RESULT_NOT_IMPLEMENTED  EErrorTypeDef = 501
	EErrorTypeDef_CHECK_CONTENT_ERROR     EErrorTypeDef = 1000
	EErrorTypeDef_CHECK_COOKIE_ERROR      EErrorTypeDef = 1010
	EErrorTypeDef_SYS_INTERNAL_ERROR      EErrorTypeDef = 1020
	EErrorTypeDef_PROGRAM_EXCEPTION_ERROR EErrorTypeDef = 1021
	EErrorTypeDef_RPC_FAILED_ERROR        EErrorTypeDef = 1080
	EErrorTypeDef_RPC_INTERFACE_ABNORMAL  EErrorTypeDef = 1081
	EErrorTypeDef_RPC_CLIENT_TIMEOUT      EErrorTypeDef = 1082
	EErrorTypeDef_GENERATE_CONTENT_ERROR  EErrorTypeDef = 1090
	EErrorTypeDef_CHECK_PARAM_ERROR       EErrorTypeDef = 1100
	EErrorTypeDef_CHECK_PERMISSION_ERROR  EErrorTypeDef = 1110
)

var EErrorTypeDef_name = map[int32]string{
	200:  "RESULT_OK",
	501:  "RESULT_NOT_IMPLEMENTED",
	1000: "CHECK_CONTENT_ERROR",
	1010: "CHECK_COOKIE_ERROR",
	1020: "SYS_INTERNAL_ERROR",
	1021: "PROGRAM_EXCEPTION_ERROR",
	1080: "RPC_FAILED_ERROR",
	1081: "RPC_INTERFACE_ABNORMAL",
	1082: "RPC_CLIENT_TIMEOUT",
	1090: "GENERATE_CONTENT_ERROR",
	1100: "CHECK_PARAM_ERROR",
	1110: "CHECK_PERMISSION_ERROR",
}
var EErrorTypeDef_value = map[string]int32{
	"RESULT_OK":               200,
	"RESULT_NOT_IMPLEMENTED":  501,
	"CHECK_CONTENT_ERROR":     1000,
	"CHECK_COOKIE_ERROR":      1010,
	"SYS_INTERNAL_ERROR":      1020,
	"PROGRAM_EXCEPTION_ERROR": 1021,
	"RPC_FAILED_ERROR":        1080,
	"RPC_INTERFACE_ABNORMAL":  1081,
	"RPC_CLIENT_TIMEOUT":      1082,
	"GENERATE_CONTENT_ERROR":  1090,
	"CHECK_PARAM_ERROR":       1100,
	"CHECK_PERMISSION_ERROR":  1110,
}

func (x EErrorTypeDef) Enum() *EErrorTypeDef {
	p := new(EErrorTypeDef)
	*p = x
	return p
}
func (x EErrorTypeDef) String() string {
	return proto.EnumName(EErrorTypeDef_name, int32(x))
}
func (x *EErrorTypeDef) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(EErrorTypeDef_value, data, "EErrorTypeDef")
	if err != nil {
		return err
	}
	*x = EErrorTypeDef(value)
	return nil
}
func (EErrorTypeDef) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type EFindAccountDef int32

const (
	EFindAccountDef_FOUND_ACCOUNT_SUCCESS EFindAccountDef = 0
	EFindAccountDef_NOT_FOUND_ACCOUNT     EFindAccountDef = 1
	EFindAccountDef_ACCOUNT_SYSTEM_ERROR  EFindAccountDef = 2
	EFindAccountDef_ACCOUNT_PARAM_ERROR   EFindAccountDef = 3
)

var EFindAccountDef_name = map[int32]string{
	0: "FOUND_ACCOUNT_SUCCESS",
	1: "NOT_FOUND_ACCOUNT",
	2: "ACCOUNT_SYSTEM_ERROR",
	3: "ACCOUNT_PARAM_ERROR",
}
var EFindAccountDef_value = map[string]int32{
	"FOUND_ACCOUNT_SUCCESS": 0,
	"NOT_FOUND_ACCOUNT":     1,
	"ACCOUNT_SYSTEM_ERROR":  2,
	"ACCOUNT_PARAM_ERROR":   3,
}

func (x EFindAccountDef) Enum() *EFindAccountDef {
	p := new(EFindAccountDef)
	*p = x
	return p
}
func (x EFindAccountDef) String() string {
	return proto.EnumName(EFindAccountDef_name, int32(x))
}
func (x *EFindAccountDef) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(EFindAccountDef_value, data, "EFindAccountDef")
	if err != nil {
		return err
	}
	*x = EFindAccountDef(value)
	return nil
}
func (EFindAccountDef) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

// 可信字段
type STUserTrustInfo struct {
	UserID           *string `protobuf:"bytes,1,opt,name=UserID" json:"UserID,omitempty"`
	Url              *string `protobuf:"bytes,2,opt,name=Url" json:"Url,omitempty"`
	Name             *string `protobuf:"bytes,3,opt,name=Name" json:"Name,omitempty"`
	FlowId           *int64  `protobuf:"zigzag64,4,opt,name=FlowId" json:"FlowId,omitempty"`
	RealIP           *string `protobuf:"bytes,5,opt,name=RealIP" json:"RealIP,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *STUserTrustInfo) Reset()                    { *m = STUserTrustInfo{} }
func (m *STUserTrustInfo) String() string            { return proto.CompactTextString(m) }
func (*STUserTrustInfo) ProtoMessage()               {}
func (*STUserTrustInfo) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *STUserTrustInfo) GetUserID() string {
	if m != nil && m.UserID != nil {
		return *m.UserID
	}
	return ""
}

func (m *STUserTrustInfo) GetUrl() string {
	if m != nil && m.Url != nil {
		return *m.Url
	}
	return ""
}

func (m *STUserTrustInfo) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *STUserTrustInfo) GetFlowId() int64 {
	if m != nil && m.FlowId != nil {
		return *m.FlowId
	}
	return 0
}

func (m *STUserTrustInfo) GetRealIP() string {
	if m != nil && m.RealIP != nil {
		return *m.RealIP
	}
	return ""
}

type STCookieInfo struct {
	UserTrustInfo    *STUserTrustInfo `protobuf:"bytes,1,opt,name=UserTrustInfo" json:"UserTrustInfo,omitempty"`
	CurTime          *int64           `protobuf:"zigzag64,2,opt,name=CurTime" json:"CurTime,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *STCookieInfo) Reset()                    { *m = STCookieInfo{} }
func (m *STCookieInfo) String() string            { return proto.CompactTextString(m) }
func (*STCookieInfo) ProtoMessage()               {}
func (*STCookieInfo) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *STCookieInfo) GetUserTrustInfo() *STUserTrustInfo {
	if m != nil {
		return m.UserTrustInfo
	}
	return nil
}

func (m *STCookieInfo) GetCurTime() int64 {
	if m != nil && m.CurTime != nil {
		return *m.CurTime
	}
	return 0
}

// 回包header
type STRspHeader struct {
	ErrNo            *EErrorTypeDef `protobuf:"varint,1,req,name=ErrNo,enum=client.EErrorTypeDef" json:"ErrNo,omitempty"`
	ErrMsg           *string        `protobuf:"bytes,2,opt,name=ErrMsg" json:"ErrMsg,omitempty"`
	FlowId           *int64         `protobuf:"zigzag64,3,opt,name=FlowId" json:"FlowId,omitempty"`
	ErrDetail        *string        `protobuf:"bytes,4,opt,name=ErrDetail" json:"ErrDetail,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *STRspHeader) Reset()                    { *m = STRspHeader{} }
func (m *STRspHeader) String() string            { return proto.CompactTextString(m) }
func (*STRspHeader) ProtoMessage()               {}
func (*STRspHeader) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *STRspHeader) GetErrNo() EErrorTypeDef {
	if m != nil && m.ErrNo != nil {
		return *m.ErrNo
	}
	return EErrorTypeDef_RESULT_OK
}

func (m *STRspHeader) GetErrMsg() string {
	if m != nil && m.ErrMsg != nil {
		return *m.ErrMsg
	}
	return ""
}

func (m *STRspHeader) GetFlowId() int64 {
	if m != nil && m.FlowId != nil {
		return *m.FlowId
	}
	return 0
}

func (m *STRspHeader) GetErrDetail() string {
	if m != nil && m.ErrDetail != nil {
		return *m.ErrDetail
	}
	return ""
}

type CommonReq struct {
	UserTrustInfo    *STUserTrustInfo `protobuf:"bytes,1,req,name=UserTrustInfo" json:"UserTrustInfo,omitempty"`
	ReqInfo          []byte           `protobuf:"bytes,2,opt,name=ReqInfo" json:"ReqInfo,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *CommonReq) Reset()                    { *m = CommonReq{} }
func (m *CommonReq) String() string            { return proto.CompactTextString(m) }
func (*CommonReq) ProtoMessage()               {}
func (*CommonReq) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *CommonReq) GetUserTrustInfo() *STUserTrustInfo {
	if m != nil {
		return m.UserTrustInfo
	}
	return nil
}

func (m *CommonReq) GetReqInfo() []byte {
	if m != nil {
		return m.ReqInfo
	}
	return nil
}

type CommonRsp struct {
	RspHeader        *STRspHeader     `protobuf:"bytes,1,req,name=RspHeader" json:"RspHeader,omitempty"`
	UserTrustInfo    *STUserTrustInfo `protobuf:"bytes,2,req,name=UserTrustInfo" json:"UserTrustInfo,omitempty"`
	RspInfo          []byte           `protobuf:"bytes,3,opt,name=RspInfo" json:"RspInfo,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *CommonRsp) Reset()                    { *m = CommonRsp{} }
func (m *CommonRsp) String() string            { return proto.CompactTextString(m) }
func (*CommonRsp) ProtoMessage()               {}
func (*CommonRsp) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *CommonRsp) GetRspHeader() *STRspHeader {
	if m != nil {
		return m.RspHeader
	}
	return nil
}

func (m *CommonRsp) GetUserTrustInfo() *STUserTrustInfo {
	if m != nil {
		return m.UserTrustInfo
	}
	return nil
}

func (m *CommonRsp) GetRspInfo() []byte {
	if m != nil {
		return m.RspInfo
	}
	return nil
}

type CommonClientRsp struct {
	RspHeader        *STRspHeader `protobuf:"bytes,1,req,name=RspHeader" json:"RspHeader,omitempty"`
	RspJson          *string      `protobuf:"bytes,2,opt,name=RspJson" json:"RspJson,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *CommonClientRsp) Reset()                    { *m = CommonClientRsp{} }
func (m *CommonClientRsp) String() string            { return proto.CompactTextString(m) }
func (*CommonClientRsp) ProtoMessage()               {}
func (*CommonClientRsp) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

func (m *CommonClientRsp) GetRspHeader() *STRspHeader {
	if m != nil {
		return m.RspHeader
	}
	return nil
}

func (m *CommonClientRsp) GetRspJson() string {
	if m != nil && m.RspJson != nil {
		return *m.RspJson
	}
	return ""
}

func init() {
	proto.RegisterType((*STUserTrustInfo)(nil), "client.STUserTrustInfo")
	proto.RegisterType((*STCookieInfo)(nil), "client.STCookieInfo")
	proto.RegisterType((*STRspHeader)(nil), "client.STRspHeader")
	proto.RegisterType((*CommonReq)(nil), "client.CommonReq")
	proto.RegisterType((*CommonRsp)(nil), "client.CommonRsp")
	proto.RegisterType((*CommonClientRsp)(nil), "client.CommonClientRsp")
	proto.RegisterEnum("client.EErrorTypeDef", EErrorTypeDef_name, EErrorTypeDef_value)
	proto.RegisterEnum("client.EFindAccountDef", EFindAccountDef_name, EFindAccountDef_value)
}

func init() { proto.RegisterFile("CommonProtocol.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 580 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x90, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0x49, 0xd2, 0x36, 0xc9, 0xf4, 0x4f, 0xdc, 0x6d, 0xd3, 0x18, 0xc1, 0xa1, 0x8a, 0x10,
	0xaa, 0x7a, 0xc8, 0xa1, 0x6f, 0x60, 0x36, 0x93, 0xd6, 0xad, 0xbd, 0xb6, 0xd6, 0x6b, 0x89, 0x4a,
	0x48, 0x56, 0x94, 0x6e, 0x51, 0x44, 0x12, 0xa7, 0x6b, 0x17, 0xc4, 0x9b, 0x01, 0x47, 0x24, 0x24,
	0x0e, 0x9c, 0x79, 0x06, 0xce, 0x48, 0xdc, 0x40, 0x42, 0xbb, 0x76, 0x1a, 0x5a, 0x90, 0x2a, 0x6e,
	0x3b, 0xf3, 0xed, 0x7c, 0xdf, 0x6f, 0x06, 0x76, 0x69, 0x3a, 0x9d, 0xa6, 0xb3, 0x50, 0xa5, 0x79,
	0x3a, 0x4a, 0x27, 0xbd, 0xb9, 0x7e, 0x90, 0xb5, 0xd1, 0x64, 0x2c, 0x67, 0x79, 0xf7, 0x05, 0xb4,
	0x22, 0x11, 0x67, 0x52, 0x09, 0x75, 0x9d, 0xe5, 0xee, 0xec, 0x32, 0x25, 0x5b, 0xb0, 0xa6, 0x1b,
	0x6e, 0xdf, 0xae, 0xec, 0x57, 0x0e, 0x9a, 0x64, 0x1d, 0x6a, 0xb1, 0x9a, 0xd8, 0x55, 0x53, 0x6c,
	0xc0, 0x0a, 0x1b, 0x4e, 0xa5, 0x5d, 0x33, 0xd5, 0x16, 0xac, 0x0d, 0x26, 0xe9, 0x1b, 0xf7, 0xc2,
	0x5e, 0xd9, 0xaf, 0x1c, 0x10, 0x5d, 0x73, 0x39, 0x9c, 0xb8, 0xa1, 0xbd, 0xaa, 0xf5, 0x6e, 0x00,
	0x1b, 0x91, 0xa0, 0x69, 0xfa, 0x6a, 0x2c, 0x8d, 0x75, 0x0f, 0x36, 0x6f, 0x65, 0x99, 0x84, 0xf5,
	0xa3, 0x4e, 0xaf, 0xa0, 0xe9, 0xdd, 0x45, 0x69, 0x41, 0x9d, 0x5e, 0x2b, 0x31, 0x9e, 0x4a, 0x13,
	0x4f, 0xba, 0x97, 0xb0, 0x1e, 0x09, 0x9e, 0xcd, 0x4f, 0xe4, 0xf0, 0x42, 0x2a, 0xf2, 0x04, 0x56,
	0x51, 0x29, 0xa6, 0x7d, 0xaa, 0x07, 0x5b, 0x47, 0xed, 0x85, 0x0f, 0xa2, 0x52, 0xa9, 0x12, 0x6f,
	0xe7, 0xb2, 0x2f, 0x2f, 0x35, 0x15, 0x2a, 0xe5, 0x67, 0x2f, 0xcb, 0x1d, 0x96, 0xd4, 0x35, 0x43,
	0xbd, 0x0d, 0x4d, 0x54, 0xaa, 0x2f, 0xf3, 0xe1, 0x78, 0x62, 0x16, 0x69, 0x76, 0x3d, 0x68, 0x16,
	0x67, 0xe3, 0xf2, 0xea, 0x5f, 0xd4, 0xd5, 0x7b, 0xa8, 0xb9, 0xbc, 0x32, 0x3f, 0x75, 0xe0, 0x46,
	0x37, 0xbf, 0x71, 0xcb, 0xe6, 0xe4, 0x29, 0x34, 0x6f, 0x16, 0x28, 0x9d, 0x76, 0x96, 0x4e, 0xcb,
	0xdd, 0xfe, 0x4a, 0xad, 0xde, 0x9f, 0x9a, 0xcd, 0xcd, 0xcf, 0x9a, 0x49, 0x3d, 0x85, 0x56, 0x91,
	0x4a, 0xcd, 0xc0, 0xff, 0x64, 0x17, 0x5e, 0xa7, 0x59, 0x3a, 0x2b, 0x4e, 0x76, 0xf8, 0xa9, 0x0a,
	0x9b, 0x77, 0x8f, 0xda, 0xe4, 0x18, 0xc5, 0x9e, 0x48, 0x82, 0x33, 0xeb, 0x73, 0x85, 0x3c, 0x82,
	0xbd, 0xb2, 0x66, 0x81, 0x48, 0x5c, 0x3f, 0xf4, 0xd0, 0x47, 0x26, 0xb0, 0x6f, 0xfd, 0xa8, 0x11,
	0x1b, 0x76, 0xe8, 0x09, 0xd2, 0xb3, 0x84, 0x06, 0x4c, 0x20, 0x13, 0x09, 0x72, 0x1e, 0x70, 0xeb,
	0x5b, 0x9d, 0x74, 0x80, 0x2c, 0x94, 0xe0, 0xcc, 0xc5, 0x52, 0xf8, 0x6e, 0x84, 0xe8, 0x3c, 0x4a,
	0x5c, 0x26, 0x90, 0x33, 0xc7, 0x2b, 0x85, 0x9f, 0x75, 0xf2, 0x18, 0x3a, 0x21, 0x0f, 0x8e, 0xb9,
	0xe3, 0x27, 0xf8, 0x9c, 0x62, 0x28, 0xdc, 0x80, 0x95, 0xea, 0xaf, 0x3a, 0x69, 0x83, 0xc5, 0x43,
	0x9a, 0x0c, 0x1c, 0xd7, 0xc3, 0x7e, 0xd9, 0x7e, 0xd7, 0x30, 0x74, 0x21, 0x2d, 0xdc, 0x06, 0x0e,
	0xc5, 0xc4, 0x79, 0xc6, 0x02, 0xee, 0x3b, 0x9e, 0xf5, 0xbe, 0xa1, 0xa3, 0xb4, 0x48, 0x3d, 0x57,
	0xa3, 0x09, 0xd7, 0xc7, 0x20, 0x16, 0xd6, 0x07, 0x33, 0x75, 0x8c, 0x0c, 0xb9, 0x23, 0xf0, 0x0e,
	0xf9, 0xc7, 0x06, 0xd9, 0x83, 0xed, 0x82, 0x3c, 0x74, 0x0c, 0x8b, 0xe9, 0x7f, 0x31, 0x43, 0x65,
	0x1f, 0xb9, 0xef, 0x46, 0xd1, 0x12, 0xef, 0x6b, 0xe3, 0xf0, 0x35, 0xb4, 0x70, 0x30, 0x9e, 0x5d,
	0x38, 0xa3, 0x51, 0x7a, 0x3d, 0xcb, 0xf5, 0x21, 0x1f, 0x42, 0x7b, 0x10, 0xc4, 0xac, 0x9f, 0x38,
	0x94, 0x06, 0x31, 0x13, 0x49, 0x14, 0x53, 0x8a, 0x51, 0x64, 0x3d, 0x20, 0x6d, 0xd8, 0xd6, 0xc7,
	0xbc, 0x25, 0x5b, 0x15, 0x62, 0xc3, 0xee, 0xcd, 0xdf, 0xf3, 0x48, 0xe0, 0x22, 0xbc, 0x4a, 0x3a,
	0xb0, 0xb3, 0x50, 0xfe, 0xa4, 0xaa, 0xfd, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x0a, 0x59, 0xab, 0xb6,
	0x05, 0x04, 0x00, 0x00,
}
