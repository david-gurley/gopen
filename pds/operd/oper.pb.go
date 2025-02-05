// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: oper.proto

package operd

import (
	context "context"
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	pds "github.com/david-gurley/gopen/pds"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Operational Info exported by operd plugin through push / subscription model
type OperInfoType int32

const (
	OperInfoType_OPER_INFO_TYPE_NONE     OperInfoType = 0
	OperInfoType_OPER_INFO_TYPE_EVENT    OperInfoType = 1
	OperInfoType_OPER_INFO_TYPE_FLOW_LOG OperInfoType = 2
)

var OperInfoType_name = map[int32]string{
	0: "OPER_INFO_TYPE_NONE",
	1: "OPER_INFO_TYPE_EVENT",
	2: "OPER_INFO_TYPE_FLOW_LOG",
}

var OperInfoType_value = map[string]int32{
	"OPER_INFO_TYPE_NONE":     0,
	"OPER_INFO_TYPE_EVENT":    1,
	"OPER_INFO_TYPE_FLOW_LOG": 2,
}

func (x OperInfoType) String() string {
	return proto.EnumName(OperInfoType_name, int32(x))
}

func (OperInfoType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e02e71d32ba5bf6d, []int{0}
}

// OperInfoOp indicates whether client is subscribing or unsubscribing
type OperInfoOp int32

const (
	OperInfoOp_OPER_INFO_OP_NONE        OperInfoOp = 0
	OperInfoOp_OPER_INFO_OP_SUBSCRIBE   OperInfoOp = 1
	OperInfoOp_OPER_INFO_OP_UNSUBSCRIBE OperInfoOp = 2
)

var OperInfoOp_name = map[int32]string{
	0: "OPER_INFO_OP_NONE",
	1: "OPER_INFO_OP_SUBSCRIBE",
	2: "OPER_INFO_OP_UNSUBSCRIBE",
}

var OperInfoOp_value = map[string]int32{
	"OPER_INFO_OP_NONE":        0,
	"OPER_INFO_OP_SUBSCRIBE":   1,
	"OPER_INFO_OP_UNSUBSCRIBE": 2,
}

func (x OperInfoOp) String() string {
	return proto.EnumName(OperInfoOp_name, int32(x))
}

func (OperInfoOp) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e02e71d32ba5bf6d, []int{1}
}

type OperInfoSpec struct {
	// oper info of interest
	InfoType OperInfoType `protobuf:"varint,1,opt,name=InfoType,proto3,enum=operd.OperInfoType" json:"InfoType,omitempty"`
	// action is either subscribe or unsubscribe
	Action OperInfoOp `protobuf:"varint,2,opt,name=Action,proto3,enum=operd.OperInfoOp" json:"Action,omitempty"`
	// any additional InfoType specific filters
	//
	// Types that are valid to be assigned to InfoFilter:
	//	*OperInfoSpec_EventFilter
	InfoFilter           isOperInfoSpec_InfoFilter `protobuf_oneof:"info_filter"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *OperInfoSpec) Reset()         { *m = OperInfoSpec{} }
func (m *OperInfoSpec) String() string { return proto.CompactTextString(m) }
func (*OperInfoSpec) ProtoMessage()    {}
func (*OperInfoSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_e02e71d32ba5bf6d, []int{0}
}
func (m *OperInfoSpec) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OperInfoSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OperInfoSpec.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OperInfoSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OperInfoSpec.Merge(m, src)
}
func (m *OperInfoSpec) XXX_Size() int {
	return m.Size()
}
func (m *OperInfoSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_OperInfoSpec.DiscardUnknown(m)
}

var xxx_messageInfo_OperInfoSpec proto.InternalMessageInfo

type isOperInfoSpec_InfoFilter interface {
	isOperInfoSpec_InfoFilter()
	MarshalTo([]byte) (int, error)
	Size() int
}

type OperInfoSpec_EventFilter struct {
	EventFilter *EventFilter `protobuf:"bytes,3,opt,name=EventFilter,proto3,oneof"`
}

func (*OperInfoSpec_EventFilter) isOperInfoSpec_InfoFilter() {}

func (m *OperInfoSpec) GetInfoFilter() isOperInfoSpec_InfoFilter {
	if m != nil {
		return m.InfoFilter
	}
	return nil
}

func (m *OperInfoSpec) GetInfoType() OperInfoType {
	if m != nil {
		return m.InfoType
	}
	return OperInfoType_OPER_INFO_TYPE_NONE
}

func (m *OperInfoSpec) GetAction() OperInfoOp {
	if m != nil {
		return m.Action
	}
	return OperInfoOp_OPER_INFO_OP_NONE
}

func (m *OperInfoSpec) GetEventFilter() *EventFilter {
	if x, ok := m.GetInfoFilter().(*OperInfoSpec_EventFilter); ok {
		return x.EventFilter
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*OperInfoSpec) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*OperInfoSpec_EventFilter)(nil),
	}
}

// OperInfoRequest captures the operational information for subscription/unsubscription
type OperInfoRequest struct {
	Request              []*OperInfoSpec `protobuf:"bytes,1,rep,name=Request,proto3" json:"Request,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *OperInfoRequest) Reset()         { *m = OperInfoRequest{} }
func (m *OperInfoRequest) String() string { return proto.CompactTextString(m) }
func (*OperInfoRequest) ProtoMessage()    {}
func (*OperInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e02e71d32ba5bf6d, []int{1}
}
func (m *OperInfoRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OperInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OperInfoRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OperInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OperInfoRequest.Merge(m, src)
}
func (m *OperInfoRequest) XXX_Size() int {
	return m.Size()
}
func (m *OperInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_OperInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_OperInfoRequest proto.InternalMessageInfo

func (m *OperInfoRequest) GetRequest() []*OperInfoSpec {
	if m != nil {
		return m.Request
	}
	return nil
}

type OperInfoResponse struct {
	Status   pds.ApiStatus `protobuf:"varint,1,opt,name=Status,proto3,enum=types.ApiStatus" json:"Status,omitempty"`
	InfoType OperInfoType  `protobuf:"varint,2,opt,name=InfoType,proto3,enum=operd.OperInfoType" json:"InfoType,omitempty"`
	// operational information specific to InfoType
	//
	// Types that are valid to be assigned to OperInfo:
	//	*OperInfoResponse_EventInfo
	//	*OperInfoResponse_FlowLogInfo
	OperInfo             isOperInfoResponse_OperInfo `protobuf_oneof:"oper_info"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *OperInfoResponse) Reset()         { *m = OperInfoResponse{} }
func (m *OperInfoResponse) String() string { return proto.CompactTextString(m) }
func (*OperInfoResponse) ProtoMessage()    {}
func (*OperInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e02e71d32ba5bf6d, []int{2}
}
func (m *OperInfoResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OperInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OperInfoResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OperInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OperInfoResponse.Merge(m, src)
}
func (m *OperInfoResponse) XXX_Size() int {
	return m.Size()
}
func (m *OperInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_OperInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_OperInfoResponse proto.InternalMessageInfo

type isOperInfoResponse_OperInfo interface {
	isOperInfoResponse_OperInfo()
	MarshalTo([]byte) (int, error)
	Size() int
}

type OperInfoResponse_EventInfo struct {
	EventInfo *Event `protobuf:"bytes,3,opt,name=EventInfo,proto3,oneof"`
}
type OperInfoResponse_FlowLogInfo struct {
	FlowLogInfo *FlowLog `protobuf:"bytes,4,opt,name=FlowLogInfo,proto3,oneof"`
}

func (*OperInfoResponse_EventInfo) isOperInfoResponse_OperInfo()   {}
func (*OperInfoResponse_FlowLogInfo) isOperInfoResponse_OperInfo() {}

func (m *OperInfoResponse) GetOperInfo() isOperInfoResponse_OperInfo {
	if m != nil {
		return m.OperInfo
	}
	return nil
}

func (m *OperInfoResponse) GetStatus() pds.ApiStatus {
	if m != nil {
		return m.Status
	}
	return pds.ApiStatus_API_STATUS_OK
}

func (m *OperInfoResponse) GetInfoType() OperInfoType {
	if m != nil {
		return m.InfoType
	}
	return OperInfoType_OPER_INFO_TYPE_NONE
}

func (m *OperInfoResponse) GetEventInfo() *Event {
	if x, ok := m.GetOperInfo().(*OperInfoResponse_EventInfo); ok {
		return x.EventInfo
	}
	return nil
}

func (m *OperInfoResponse) GetFlowLogInfo() *FlowLog {
	if x, ok := m.GetOperInfo().(*OperInfoResponse_FlowLogInfo); ok {
		return x.FlowLogInfo
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*OperInfoResponse) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*OperInfoResponse_EventInfo)(nil),
		(*OperInfoResponse_FlowLogInfo)(nil),
	}
}

func init() {
	proto.RegisterEnum("operd.OperInfoType", OperInfoType_name, OperInfoType_value)
	proto.RegisterEnum("operd.OperInfoOp", OperInfoOp_name, OperInfoOp_value)
	proto.RegisterType((*OperInfoSpec)(nil), "operd.OperInfoSpec")
	proto.RegisterType((*OperInfoRequest)(nil), "operd.OperInfoRequest")
	proto.RegisterType((*OperInfoResponse)(nil), "operd.OperInfoResponse")
}

func init() { proto.RegisterFile("oper.proto", fileDescriptor_e02e71d32ba5bf6d) }

var fileDescriptor_e02e71d32ba5bf6d = []byte{
	// 451 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0xb3, 0x29, 0xa4, 0x74, 0x1c, 0x8a, 0xb3, 0x85, 0xc6, 0x0a, 0x28, 0x8a, 0x72, 0x32,
	0x15, 0x04, 0x14, 0x24, 0xce, 0xd4, 0xc8, 0xa1, 0x41, 0x91, 0x37, 0x5a, 0x27, 0x20, 0x0e, 0x60,
	0x91, 0xb0, 0x41, 0x96, 0x22, 0xef, 0x62, 0x6f, 0x5a, 0xf5, 0xb1, 0x78, 0x0b, 0x8e, 0x3c, 0x00,
	0x07, 0x94, 0x27, 0x41, 0xbb, 0x5e, 0xd7, 0x89, 0x7b, 0xe1, 0x36, 0x3b, 0xff, 0xa7, 0xf1, 0xcc,
	0xff, 0x1b, 0x80, 0x0b, 0x96, 0x0e, 0x44, 0xca, 0x25, 0xc7, 0x77, 0x55, 0xfd, 0xad, 0x63, 0xc9,
	0x6b, 0xc1, 0xb2, 0xbc, 0xd7, 0x69, 0xb2, 0x4b, 0x96, 0xc8, 0xe2, 0x05, 0xab, 0x35, 0xbf, 0xca,
	0xeb, 0xfe, 0x4f, 0x04, 0x4d, 0x22, 0x58, 0x3a, 0x4e, 0x56, 0x3c, 0x14, 0x6c, 0x89, 0x5f, 0xc0,
	0x3d, 0x55, 0xcf, 0xae, 0x05, 0x73, 0x50, 0x0f, 0xb9, 0xc7, 0xc3, 0x93, 0x81, 0x9e, 0x38, 0x28,
	0x30, 0x25, 0xd1, 0x1b, 0x08, 0x3f, 0x85, 0xc6, 0xf9, 0x52, 0xc6, 0x3c, 0x71, 0xea, 0x1a, 0x6f,
	0x55, 0x70, 0x22, 0xa8, 0x01, 0xf0, 0x6b, 0xb0, 0x7c, 0xb5, 0xc8, 0x28, 0x5e, 0x4b, 0x96, 0x3a,
	0x07, 0x3d, 0xe4, 0x5a, 0x43, 0x6c, 0xf8, 0x1d, 0xe5, 0xa2, 0x46, 0x77, 0x41, 0xef, 0x3e, 0x58,
	0x71, 0xb2, 0xe2, 0xd1, 0x4a, 0x3f, 0xfb, 0x6f, 0xe0, 0x41, 0x31, 0x9c, 0xb2, 0x1f, 0x1b, 0x96,
	0x49, 0xfc, 0x1c, 0x0e, 0x4d, 0xe9, 0xa0, 0xde, 0x81, 0x6b, 0xdd, 0x5a, 0x5a, 0xdd, 0x46, 0x0b,
	0xa6, 0xff, 0x07, 0x81, 0x5d, 0x8e, 0xc8, 0x04, 0x4f, 0x32, 0x86, 0x5d, 0x68, 0x84, 0xf2, 0xab,
	0xdc, 0x64, 0xe6, 0x6e, 0x7b, 0x90, 0x5b, 0x78, 0x2e, 0xe2, 0xbc, 0x4f, 0x8d, 0xbe, 0xe7, 0x51,
	0xfd, 0x7f, 0x3c, 0x7a, 0x06, 0x47, 0xfa, 0x1e, 0xd5, 0x30, 0x67, 0x37, 0x77, 0xcf, 0xbe, 0xa8,
	0xd1, 0x12, 0xc0, 0x43, 0xb0, 0x46, 0x6b, 0x7e, 0x35, 0xe1, 0xdf, 0x35, 0x7f, 0x47, 0xf3, 0xc7,
	0x86, 0x37, 0x8a, 0xb2, 0x68, 0x07, 0xf2, 0x2c, 0x38, 0x52, 0x7a, 0xa4, 0x7c, 0x3a, 0xfb, 0x52,
	0x66, 0xaa, 0x3f, 0xdf, 0x86, 0x13, 0x32, 0xf5, 0x69, 0x34, 0x0e, 0x46, 0x24, 0x9a, 0x7d, 0x9a,
	0xfa, 0x51, 0x40, 0x02, 0xdf, 0xae, 0x61, 0x07, 0x1e, 0x56, 0x04, 0xff, 0x83, 0x1f, 0xcc, 0x6c,
	0x84, 0x1f, 0x43, 0xbb, 0xa2, 0x8c, 0x26, 0xe4, 0x63, 0x34, 0x21, 0xef, 0xec, 0xfa, 0xd9, 0x67,
	0x80, 0x32, 0x5d, 0xfc, 0x08, 0x5a, 0x25, 0x4a, 0xa6, 0xc5, 0xec, 0x0e, 0x9c, 0xee, 0xb5, 0xc3,
	0xb9, 0x17, 0xbe, 0xa5, 0x63, 0xcf, 0xb7, 0x11, 0x7e, 0x02, 0xce, 0x9e, 0x36, 0x0f, 0x4a, 0xb5,
	0x3e, 0x9c, 0xc3, 0xa1, 0x1a, 0x1f, 0x5e, 0x2e, 0xf1, 0x7b, 0x68, 0xdd, 0x24, 0xb8, 0x59, 0x64,
	0xcb, 0x34, 0x5e, 0x30, 0x7c, 0x5a, 0x31, 0xdb, 0xa4, 0xda, 0x69, 0xdf, 0xea, 0xe7, 0xc9, 0xf6,
	0x6b, 0x2e, 0x7a, 0x89, 0xbc, 0xe6, 0xaf, 0x6d, 0x17, 0xfd, 0xde, 0x76, 0xd1, 0xdf, 0x6d, 0x17,
	0x2d, 0x1a, 0xfa, 0xff, 0x7f, 0xf5, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x3e, 0x24, 0xbd, 0x28, 0x3b,
	0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// OperSvcClient is the client API for OperSvc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OperSvcClient interface {
	OperInfoSubscribe(ctx context.Context, opts ...grpc.CallOption) (OperSvc_OperInfoSubscribeClient, error)
}

type operSvcClient struct {
	cc *grpc.ClientConn
}

func NewOperSvcClient(cc *grpc.ClientConn) OperSvcClient {
	return &operSvcClient{cc}
}

func (c *operSvcClient) OperInfoSubscribe(ctx context.Context, opts ...grpc.CallOption) (OperSvc_OperInfoSubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_OperSvc_serviceDesc.Streams[0], "/operd.OperSvc/OperInfoSubscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &operSvcOperInfoSubscribeClient{stream}
	return x, nil
}

type OperSvc_OperInfoSubscribeClient interface {
	Send(*OperInfoRequest) error
	Recv() (*OperInfoResponse, error)
	grpc.ClientStream
}

type operSvcOperInfoSubscribeClient struct {
	grpc.ClientStream
}

func (x *operSvcOperInfoSubscribeClient) Send(m *OperInfoRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *operSvcOperInfoSubscribeClient) Recv() (*OperInfoResponse, error) {
	m := new(OperInfoResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// OperSvcServer is the server API for OperSvc service.
type OperSvcServer interface {
	OperInfoSubscribe(OperSvc_OperInfoSubscribeServer) error
}

// UnimplementedOperSvcServer can be embedded to have forward compatible implementations.
type UnimplementedOperSvcServer struct {
}

func (*UnimplementedOperSvcServer) OperInfoSubscribe(srv OperSvc_OperInfoSubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method OperInfoSubscribe not implemented")
}

func RegisterOperSvcServer(s *grpc.Server, srv OperSvcServer) {
	s.RegisterService(&_OperSvc_serviceDesc, srv)
}

func _OperSvc_OperInfoSubscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(OperSvcServer).OperInfoSubscribe(&operSvcOperInfoSubscribeServer{stream})
}

type OperSvc_OperInfoSubscribeServer interface {
	Send(*OperInfoResponse) error
	Recv() (*OperInfoRequest, error)
	grpc.ServerStream
}

type operSvcOperInfoSubscribeServer struct {
	grpc.ServerStream
}

func (x *operSvcOperInfoSubscribeServer) Send(m *OperInfoResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *operSvcOperInfoSubscribeServer) Recv() (*OperInfoRequest, error) {
	m := new(OperInfoRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _OperSvc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "operd.OperSvc",
	HandlerType: (*OperSvcServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "OperInfoSubscribe",
			Handler:       _OperSvc_OperInfoSubscribe_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "oper.proto",
}

func (m *OperInfoSpec) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OperInfoSpec) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OperInfoSpec) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.InfoFilter != nil {
		{
			size := m.InfoFilter.Size()
			i -= size
			if _, err := m.InfoFilter.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	if m.Action != 0 {
		i = encodeVarintOper(dAtA, i, uint64(m.Action))
		i--
		dAtA[i] = 0x10
	}
	if m.InfoType != 0 {
		i = encodeVarintOper(dAtA, i, uint64(m.InfoType))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *OperInfoSpec_EventFilter) MarshalTo(dAtA []byte) (int, error) {
	return m.MarshalToSizedBuffer(dAtA[:m.Size()])
}

func (m *OperInfoSpec_EventFilter) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.EventFilter != nil {
		{
			size, err := m.EventFilter.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintOper(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	return len(dAtA) - i, nil
}
func (m *OperInfoRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OperInfoRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OperInfoRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Request) > 0 {
		for iNdEx := len(m.Request) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Request[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintOper(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *OperInfoResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OperInfoResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OperInfoResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.OperInfo != nil {
		{
			size := m.OperInfo.Size()
			i -= size
			if _, err := m.OperInfo.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	if m.InfoType != 0 {
		i = encodeVarintOper(dAtA, i, uint64(m.InfoType))
		i--
		dAtA[i] = 0x10
	}
	if m.Status != 0 {
		i = encodeVarintOper(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *OperInfoResponse_EventInfo) MarshalTo(dAtA []byte) (int, error) {
	return m.MarshalToSizedBuffer(dAtA[:m.Size()])
}

func (m *OperInfoResponse_EventInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.EventInfo != nil {
		{
			size, err := m.EventInfo.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintOper(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	return len(dAtA) - i, nil
}
func (m *OperInfoResponse_FlowLogInfo) MarshalTo(dAtA []byte) (int, error) {
	return m.MarshalToSizedBuffer(dAtA[:m.Size()])
}

func (m *OperInfoResponse_FlowLogInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.FlowLogInfo != nil {
		{
			size, err := m.FlowLogInfo.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintOper(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	return len(dAtA) - i, nil
}
func encodeVarintOper(dAtA []byte, offset int, v uint64) int {
	offset -= sovOper(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *OperInfoSpec) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.InfoType != 0 {
		n += 1 + sovOper(uint64(m.InfoType))
	}
	if m.Action != 0 {
		n += 1 + sovOper(uint64(m.Action))
	}
	if m.InfoFilter != nil {
		n += m.InfoFilter.Size()
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *OperInfoSpec_EventFilter) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.EventFilter != nil {
		l = m.EventFilter.Size()
		n += 1 + l + sovOper(uint64(l))
	}
	return n
}
func (m *OperInfoRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Request) > 0 {
		for _, e := range m.Request {
			l = e.Size()
			n += 1 + l + sovOper(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *OperInfoResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Status != 0 {
		n += 1 + sovOper(uint64(m.Status))
	}
	if m.InfoType != 0 {
		n += 1 + sovOper(uint64(m.InfoType))
	}
	if m.OperInfo != nil {
		n += m.OperInfo.Size()
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *OperInfoResponse_EventInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.EventInfo != nil {
		l = m.EventInfo.Size()
		n += 1 + l + sovOper(uint64(l))
	}
	return n
}
func (m *OperInfoResponse_FlowLogInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.FlowLogInfo != nil {
		l = m.FlowLogInfo.Size()
		n += 1 + l + sovOper(uint64(l))
	}
	return n
}

func sovOper(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozOper(x uint64) (n int) {
	return sovOper(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *OperInfoSpec) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOper
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: OperInfoSpec: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OperInfoSpec: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field InfoType", wireType)
			}
			m.InfoType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOper
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.InfoType |= OperInfoType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Action", wireType)
			}
			m.Action = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOper
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Action |= OperInfoOp(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EventFilter", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOper
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthOper
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOper
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &EventFilter{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.InfoFilter = &OperInfoSpec_EventFilter{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOper(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthOper
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthOper
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *OperInfoRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOper
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: OperInfoRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OperInfoRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Request", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOper
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthOper
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOper
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Request = append(m.Request, &OperInfoSpec{})
			if err := m.Request[len(m.Request)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOper(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthOper
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthOper
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *OperInfoResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOper
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: OperInfoResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OperInfoResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOper
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= pds.ApiStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field InfoType", wireType)
			}
			m.InfoType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOper
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.InfoType |= OperInfoType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EventInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOper
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthOper
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOper
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &Event{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.OperInfo = &OperInfoResponse_EventInfo{v}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FlowLogInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOper
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthOper
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOper
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &FlowLog{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.OperInfo = &OperInfoResponse_FlowLogInfo{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOper(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthOper
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthOper
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipOper(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOper
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowOper
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowOper
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthOper
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthOper
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowOper
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipOper(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthOper
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthOper = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOper   = fmt.Errorf("proto: integer overflow")
)
