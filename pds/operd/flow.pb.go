// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: flow.proto

package operd

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	pds "github.com/david-gurley/gopen/pds"
	proto "github.com/gogo/protobuf/proto"
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

// flow log type indicates whether log is generated when flow is
// created or closed
type FlowLogType int32

const (
	FlowLogType_FLOW_LOG_TYPE_NONE FlowLogType = 0
	// flow log raised at the time of flow/session creation
	FlowLogType_FLOW_LOG_TYPE_OPEN FlowLogType = 1
	// flow log raised at the time of flow/session deletion
	FlowLogType_FLOW_LOG_TYPE_CLOSE FlowLogType = 2
	// flow log raised for long lived flows while flow is active
	FlowLogType_FLOW_LOG_TYPE_ACTIVE FlowLogType = 3
)

var FlowLogType_name = map[int32]string{
	0: "FLOW_LOG_TYPE_NONE",
	1: "FLOW_LOG_TYPE_OPEN",
	2: "FLOW_LOG_TYPE_CLOSE",
	3: "FLOW_LOG_TYPE_ACTIVE",
}

var FlowLogType_value = map[string]int32{
	"FLOW_LOG_TYPE_NONE":   0,
	"FLOW_LOG_TYPE_OPEN":   1,
	"FLOW_LOG_TYPE_CLOSE":  2,
	"FLOW_LOG_TYPE_ACTIVE": 3,
}

func (x FlowLogType) String() string {
	return proto.EnumName(FlowLogType_name, int32(x))
}

func (FlowLogType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_bb3fc33c49933823, []int{0}
}

// security policy type due to which the flow action was chosen
type SecurityPolicyType int32

const (
	SecurityPolicyType_SECURITY_POLICY_TYPE_NONE SecurityPolicyType = 0
	// flow action is the result of a security group rule
	SecurityPolicyType_SECURITY_POLICY_TYPE_SG SecurityPolicyType = 1
	// flow action is the result of a stateless NACL rule (subnet/vpc level)
	SecurityPolicyType_SECURITY_POLICY_TYPE_NACL SecurityPolicyType = 2
)

var SecurityPolicyType_name = map[int32]string{
	0: "SECURITY_POLICY_TYPE_NONE",
	1: "SECURITY_POLICY_TYPE_SG",
	2: "SECURITY_POLICY_TYPE_NACL",
}

var SecurityPolicyType_value = map[string]int32{
	"SECURITY_POLICY_TYPE_NONE": 0,
	"SECURITY_POLICY_TYPE_SG":   1,
	"SECURITY_POLICY_TYPE_NACL": 2,
}

func (x SecurityPolicyType) String() string {
	return proto.EnumName(SecurityPolicyType_name, int32(x))
}

func (SecurityPolicyType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_bb3fc33c49933823, []int{1}
}

// protobuf format for flow logs (this is the format consumers will see flow
// logs in). The producer (e.g., VPP) will do just a binary dump to shared
// memory and a decoder (invoked by operd) will transform the contents into
// this format
type FlowLog struct {
	// timestamp indicating when the flow is created/closed
	Timestamp uint64 `protobuf:"varint,1,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	// flow open log or flow close log
	LogType FlowLogType `protobuf:"varint,2,opt,name=LogType,proto3,enum=operd.FlowLogType" json:"LogType,omitempty"`
	// flow key
	Key *pds.FlowKey `protobuf:"bytes,3,opt,name=Key,proto3" json:"Key,omitempty"`
	// SessionId identifies unique id assigned for the the flow pair
	// for the lifetime of the session
	SessionId uint64 `protobuf:"varint,4,opt,name=SessionId,proto3" json:"SessionId,omitempty"`
	// source interface of first packet in the session
	// NOTE: this is source vnic's PF/VF/Tunnel uuid if the session is initiated
	//       locally (irrespective of whether destination is local or remote) and
	//       if the session is created because first packet came on uplink, this
	//       field is left empty
	SrcInterface []byte `protobuf:"bytes,5,opt,name=SrcInterface,proto3" json:"SrcInterface,omitempty"`
	// source vnic's uuid if the session is originated from local workload,
	// empty otherwise
	SrcVnic []byte `protobuf:"bytes,6,opt,name=SrcVnic,proto3" json:"SrcVnic,omitempty"`
	// vpc of the client/initiator of the session
	SrcVPC []byte `protobuf:"bytes,7,opt,name=SrcVPC,proto3" json:"SrcVPC,omitempty"`
	// vpc of the server/destination of the session
	DstVPC []byte `protobuf:"bytes,8,opt,name=DstVPC,proto3" json:"DstVPC,omitempty"`
	// security policy id, if security policy was evaluated
	SecurityPolicy []byte `protobuf:"bytes,9,opt,name=SecurityPolicy,proto3" json:"SecurityPolicy,omitempty"`
	// security policy rule id, if security policy was evaluated
	SecurityPolicyRule []byte `protobuf:"bytes,10,opt,name=SecurityPolicyRule,proto3" json:"SecurityPolicyRule,omitempty"`
	// flow action taken
	FlowAction pds.SecurityRuleAction `protobuf:"varint,11,opt,name=FlowAction,proto3,enum=types.SecurityRuleAction" json:"FlowAction,omitempty"`
	// NAT-ed source IP after the rewrite, if any
	SrcNATIP *pds.IPAddress `protobuf:"bytes,12,opt,name=SrcNATIP,proto3" json:"SrcNATIP,omitempty"`
	// source NAT port, if any
	SrcNATPort uint32 `protobuf:"varint,13,opt,name=SrcNATPort,proto3" json:"SrcNATPort,omitempty"`
	// NAT-ed destination IP after the rewrite, if any
	DstNATIP *pds.IPAddress `protobuf:"bytes,14,opt,name=DstNATIP,proto3" json:"DstNATIP,omitempty"`
	// destination NAT port, if any
	DstNATPort uint32 `protobuf:"varint,15,opt,name=DstNATPort,proto3" json:"DstNATPort,omitempty"`
	// initiator flow octet count
	IflowBytes uint64 `protobuf:"varint,16,opt,name=IflowBytes,proto3" json:"IflowBytes,omitempty"`
	// initator flow packet count
	IflowPackets uint64 `protobuf:"varint,17,opt,name=IflowPackets,proto3" json:"IflowPackets,omitempty"`
	// responder flow octet count
	RflowBytes uint64 `protobuf:"varint,18,opt,name=RflowBytes,proto3" json:"RflowBytes,omitempty"`
	// responder flow packet count
	RflowPackets uint64 `protobuf:"varint,19,opt,name=RflowPackets,proto3" json:"RflowPackets,omitempty"`
	// vlan id in the received packet
	SrcVlan uint32 `protobuf:"varint,20,opt,name=SrcVlan,proto3" json:"SrcVlan,omitempty"`
	// security policy rule hw id, if security policy was evaluated
	SecurityPloicyRuleId uint32   `protobuf:"varint,21,opt,name=SecurityPloicyRuleId,proto3" json:"SecurityPloicyRuleId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FlowLog) Reset()         { *m = FlowLog{} }
func (m *FlowLog) String() string { return proto.CompactTextString(m) }
func (*FlowLog) ProtoMessage()    {}
func (*FlowLog) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb3fc33c49933823, []int{0}
}
func (m *FlowLog) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FlowLog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FlowLog.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FlowLog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlowLog.Merge(m, src)
}
func (m *FlowLog) XXX_Size() int {
	return m.Size()
}
func (m *FlowLog) XXX_DiscardUnknown() {
	xxx_messageInfo_FlowLog.DiscardUnknown(m)
}

var xxx_messageInfo_FlowLog proto.InternalMessageInfo

func (m *FlowLog) GetTimestamp() uint64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *FlowLog) GetLogType() FlowLogType {
	if m != nil {
		return m.LogType
	}
	return FlowLogType_FLOW_LOG_TYPE_NONE
}

func (m *FlowLog) GetKey() *pds.FlowKey {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *FlowLog) GetSessionId() uint64 {
	if m != nil {
		return m.SessionId
	}
	return 0
}

func (m *FlowLog) GetSrcInterface() []byte {
	if m != nil {
		return m.SrcInterface
	}
	return nil
}

func (m *FlowLog) GetSrcVnic() []byte {
	if m != nil {
		return m.SrcVnic
	}
	return nil
}

func (m *FlowLog) GetSrcVPC() []byte {
	if m != nil {
		return m.SrcVPC
	}
	return nil
}

func (m *FlowLog) GetDstVPC() []byte {
	if m != nil {
		return m.DstVPC
	}
	return nil
}

func (m *FlowLog) GetSecurityPolicy() []byte {
	if m != nil {
		return m.SecurityPolicy
	}
	return nil
}

func (m *FlowLog) GetSecurityPolicyRule() []byte {
	if m != nil {
		return m.SecurityPolicyRule
	}
	return nil
}

func (m *FlowLog) GetFlowAction() pds.SecurityRuleAction {
	if m != nil {
		return m.FlowAction
	}
	return pds.SecurityRuleAction_SECURITY_RULE_ACTION_NONE
}

func (m *FlowLog) GetSrcNATIP() *pds.IPAddress {
	if m != nil {
		return m.SrcNATIP
	}
	return nil
}

func (m *FlowLog) GetSrcNATPort() uint32 {
	if m != nil {
		return m.SrcNATPort
	}
	return 0
}

func (m *FlowLog) GetDstNATIP() *pds.IPAddress {
	if m != nil {
		return m.DstNATIP
	}
	return nil
}

func (m *FlowLog) GetDstNATPort() uint32 {
	if m != nil {
		return m.DstNATPort
	}
	return 0
}

func (m *FlowLog) GetIflowBytes() uint64 {
	if m != nil {
		return m.IflowBytes
	}
	return 0
}

func (m *FlowLog) GetIflowPackets() uint64 {
	if m != nil {
		return m.IflowPackets
	}
	return 0
}

func (m *FlowLog) GetRflowBytes() uint64 {
	if m != nil {
		return m.RflowBytes
	}
	return 0
}

func (m *FlowLog) GetRflowPackets() uint64 {
	if m != nil {
		return m.RflowPackets
	}
	return 0
}

func (m *FlowLog) GetSrcVlan() uint32 {
	if m != nil {
		return m.SrcVlan
	}
	return 0
}

func (m *FlowLog) GetSecurityPloicyRuleId() uint32 {
	if m != nil {
		return m.SecurityPloicyRuleId
	}
	return 0
}

func init() {
	proto.RegisterEnum("operd.FlowLogType", FlowLogType_name, FlowLogType_value)
	proto.RegisterEnum("operd.SecurityPolicyType", SecurityPolicyType_name, SecurityPolicyType_value)
	proto.RegisterType((*FlowLog)(nil), "operd.FlowLog")
}

func init() { proto.RegisterFile("flow.proto", fileDescriptor_bb3fc33c49933823) }

var fileDescriptor_bb3fc33c49933823 = []byte{
	// 538 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x94, 0xdf, 0x4e, 0xdb, 0x30,
	0x14, 0xc6, 0x97, 0x16, 0x5a, 0x38, 0x2d, 0x5d, 0x66, 0x3a, 0x30, 0xfb, 0x53, 0x45, 0x5c, 0x4c,
	0x15, 0x42, 0xbd, 0xe8, 0xae, 0x76, 0x59, 0x42, 0x40, 0x11, 0x51, 0x13, 0x39, 0x19, 0x53, 0xaf,
	0xaa, 0x2e, 0x35, 0x28, 0x5a, 0xa8, 0xa3, 0xc4, 0x08, 0xe5, 0x0d, 0xf6, 0x68, 0xbb, 0xdc, 0x23,
	0x4c, 0x7d, 0x92, 0xc9, 0x76, 0xd3, 0x26, 0x50, 0xee, 0x7c, 0x7e, 0xdf, 0xe7, 0xef, 0xb4, 0x3e,
	0x47, 0x01, 0xb8, 0x8b, 0xd9, 0xd3, 0x20, 0x49, 0x19, 0x67, 0x68, 0x97, 0x25, 0x34, 0x9d, 0x7f,
	0x68, 0xf1, 0x3c, 0xa1, 0x99, 0x62, 0xa7, 0xbf, 0x1b, 0xd0, 0xbc, 0x8a, 0xd9, 0x93, 0xc3, 0xee,
	0xd1, 0x27, 0xd8, 0x0f, 0xa2, 0x07, 0x9a, 0xf1, 0xd9, 0x43, 0x82, 0x35, 0x43, 0xeb, 0xef, 0x90,
	0x0d, 0x40, 0xe7, 0xd0, 0x74, 0xd8, 0x7d, 0x90, 0x27, 0x14, 0xd7, 0x0c, 0xad, 0xdf, 0x19, 0xa2,
	0x81, 0xcc, 0x1b, 0xac, 0xae, 0x0b, 0x85, 0x14, 0x16, 0x64, 0x40, 0xfd, 0x86, 0xe6, 0xb8, 0x6e,
	0x68, 0xfd, 0xd6, 0xb0, 0x33, 0x50, 0x2d, 0x85, 0xf3, 0x86, 0xe6, 0x44, 0x48, 0xa2, 0x9b, 0x4f,
	0xb3, 0x2c, 0x62, 0x0b, 0x7b, 0x8e, 0x77, 0x54, 0xb7, 0x35, 0x40, 0xa7, 0xd0, 0xf6, 0xd3, 0xd0,
	0x5e, 0x70, 0x9a, 0xde, 0xcd, 0x42, 0x8a, 0x77, 0x0d, 0xad, 0xdf, 0x26, 0x15, 0x86, 0x30, 0x34,
	0xfd, 0x34, 0xbc, 0x5d, 0x44, 0x21, 0x6e, 0x48, 0xb9, 0x28, 0xd1, 0x11, 0x34, 0xc4, 0xd1, 0x33,
	0x71, 0x53, 0x0a, 0xab, 0x4a, 0xf0, 0xcb, 0x8c, 0x0b, 0xbe, 0xa7, 0xb8, 0xaa, 0xd0, 0x17, 0xe8,
	0xf8, 0x34, 0x7c, 0x4c, 0x23, 0x9e, 0x7b, 0x2c, 0x8e, 0xc2, 0x1c, 0xef, 0x4b, 0xfd, 0x19, 0x45,
	0x03, 0x40, 0x55, 0x42, 0x1e, 0x63, 0x8a, 0x41, 0x7a, 0xb7, 0x28, 0xe8, 0x1b, 0x80, 0xf8, 0xcf,
	0xa3, 0x90, 0x47, 0x6c, 0x81, 0x5b, 0xf2, 0xd9, 0x4e, 0x56, 0x8f, 0x51, 0xd8, 0x85, 0x51, 0x19,
	0x48, 0xc9, 0x8c, 0xce, 0x61, 0xcf, 0x4f, 0xc3, 0xf1, 0x28, 0xb0, 0x3d, 0xdc, 0x96, 0xaf, 0xa8,
	0xaf, 0x2e, 0xda, 0xde, 0x68, 0x3e, 0x4f, 0x69, 0x96, 0x91, 0xb5, 0x03, 0xf5, 0x00, 0xd4, 0xd9,
	0x63, 0x29, 0xc7, 0x07, 0x86, 0xd6, 0x3f, 0x20, 0x25, 0x22, 0xd2, 0x2e, 0x33, 0xae, 0xd2, 0x3a,
	0xaf, 0xa5, 0x15, 0x0e, 0x91, 0xa6, 0xce, 0x32, 0xed, 0xad, 0x4a, 0xdb, 0x10, 0xa1, 0xdb, 0x62,
	0xaf, 0x2e, 0x72, 0x4e, 0x33, 0xac, 0xcb, 0xd9, 0x95, 0x88, 0x18, 0x9e, 0xac, 0xbc, 0x59, 0xf8,
	0x8b, 0xf2, 0x0c, 0xbf, 0x93, 0x8e, 0x0a, 0x13, 0x19, 0x64, 0x93, 0x81, 0x54, 0x06, 0xa9, 0x64,
	0x90, 0x72, 0xc6, 0xa1, 0xca, 0x28, 0xb3, 0x62, 0x01, 0xe2, 0xd9, 0x02, 0x77, 0xe5, 0x8f, 0x2c,
	0x4a, 0x34, 0x84, 0xee, 0x7a, 0x1c, 0x31, 0x5b, 0x8d, 0xc3, 0x9e, 0xe3, 0xf7, 0xd2, 0xb6, 0x55,
	0x3b, 0x4b, 0xa0, 0x55, 0x5a, 0x65, 0x74, 0x04, 0xe8, 0xca, 0x71, 0x7f, 0x4c, 0x1d, 0xf7, 0x7a,
	0x1a, 0x4c, 0x3c, 0x6b, 0x3a, 0x76, 0xc7, 0x96, 0xfe, 0xe6, 0x25, 0x77, 0x3d, 0x6b, 0xac, 0x6b,
	0xe8, 0x18, 0x0e, 0xab, 0xdc, 0x74, 0x5c, 0xdf, 0xd2, 0x6b, 0x08, 0x43, 0xb7, 0x2a, 0x8c, 0xcc,
	0xc0, 0xbe, 0xb5, 0xf4, 0xfa, 0x19, 0x7b, 0xbe, 0x4e, 0xb2, 0xf1, 0x67, 0x38, 0xf1, 0x2d, 0xf3,
	0x3b, 0xb1, 0x83, 0xc9, 0xd4, 0x73, 0x1d, 0xdb, 0x9c, 0x54, 0xfa, 0x7f, 0x84, 0xe3, 0xad, 0xb2,
	0x7f, 0xad, 0x6b, 0xaf, 0xdf, 0x1d, 0x99, 0x8e, 0x5e, 0xbb, 0x68, 0xff, 0x59, 0xf6, 0xb4, 0xbf,
	0xcb, 0x9e, 0xf6, 0x6f, 0xd9, 0xd3, 0x7e, 0x36, 0xe4, 0x27, 0xe0, 0xeb, 0xff, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x70, 0x35, 0x82, 0xe3, 0x24, 0x04, 0x00, 0x00,
}

func (m *FlowLog) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FlowLog) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FlowLog) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.SecurityPloicyRuleId != 0 {
		i = encodeVarintFlow(dAtA, i, uint64(m.SecurityPloicyRuleId))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xa8
	}
	if m.SrcVlan != 0 {
		i = encodeVarintFlow(dAtA, i, uint64(m.SrcVlan))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xa0
	}
	if m.RflowPackets != 0 {
		i = encodeVarintFlow(dAtA, i, uint64(m.RflowPackets))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x98
	}
	if m.RflowBytes != 0 {
		i = encodeVarintFlow(dAtA, i, uint64(m.RflowBytes))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x90
	}
	if m.IflowPackets != 0 {
		i = encodeVarintFlow(dAtA, i, uint64(m.IflowPackets))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x88
	}
	if m.IflowBytes != 0 {
		i = encodeVarintFlow(dAtA, i, uint64(m.IflowBytes))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x80
	}
	if m.DstNATPort != 0 {
		i = encodeVarintFlow(dAtA, i, uint64(m.DstNATPort))
		i--
		dAtA[i] = 0x78
	}
	if m.DstNATIP != nil {
		{
			size, err := m.DstNATIP.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintFlow(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x72
	}
	if m.SrcNATPort != 0 {
		i = encodeVarintFlow(dAtA, i, uint64(m.SrcNATPort))
		i--
		dAtA[i] = 0x68
	}
	if m.SrcNATIP != nil {
		{
			size, err := m.SrcNATIP.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintFlow(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x62
	}
	if m.FlowAction != 0 {
		i = encodeVarintFlow(dAtA, i, uint64(m.FlowAction))
		i--
		dAtA[i] = 0x58
	}
	if len(m.SecurityPolicyRule) > 0 {
		i -= len(m.SecurityPolicyRule)
		copy(dAtA[i:], m.SecurityPolicyRule)
		i = encodeVarintFlow(dAtA, i, uint64(len(m.SecurityPolicyRule)))
		i--
		dAtA[i] = 0x52
	}
	if len(m.SecurityPolicy) > 0 {
		i -= len(m.SecurityPolicy)
		copy(dAtA[i:], m.SecurityPolicy)
		i = encodeVarintFlow(dAtA, i, uint64(len(m.SecurityPolicy)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.DstVPC) > 0 {
		i -= len(m.DstVPC)
		copy(dAtA[i:], m.DstVPC)
		i = encodeVarintFlow(dAtA, i, uint64(len(m.DstVPC)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.SrcVPC) > 0 {
		i -= len(m.SrcVPC)
		copy(dAtA[i:], m.SrcVPC)
		i = encodeVarintFlow(dAtA, i, uint64(len(m.SrcVPC)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.SrcVnic) > 0 {
		i -= len(m.SrcVnic)
		copy(dAtA[i:], m.SrcVnic)
		i = encodeVarintFlow(dAtA, i, uint64(len(m.SrcVnic)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.SrcInterface) > 0 {
		i -= len(m.SrcInterface)
		copy(dAtA[i:], m.SrcInterface)
		i = encodeVarintFlow(dAtA, i, uint64(len(m.SrcInterface)))
		i--
		dAtA[i] = 0x2a
	}
	if m.SessionId != 0 {
		i = encodeVarintFlow(dAtA, i, uint64(m.SessionId))
		i--
		dAtA[i] = 0x20
	}
	if m.Key != nil {
		{
			size, err := m.Key.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintFlow(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.LogType != 0 {
		i = encodeVarintFlow(dAtA, i, uint64(m.LogType))
		i--
		dAtA[i] = 0x10
	}
	if m.Timestamp != 0 {
		i = encodeVarintFlow(dAtA, i, uint64(m.Timestamp))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintFlow(dAtA []byte, offset int, v uint64) int {
	offset -= sovFlow(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *FlowLog) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Timestamp != 0 {
		n += 1 + sovFlow(uint64(m.Timestamp))
	}
	if m.LogType != 0 {
		n += 1 + sovFlow(uint64(m.LogType))
	}
	if m.Key != nil {
		l = m.Key.Size()
		n += 1 + l + sovFlow(uint64(l))
	}
	if m.SessionId != 0 {
		n += 1 + sovFlow(uint64(m.SessionId))
	}
	l = len(m.SrcInterface)
	if l > 0 {
		n += 1 + l + sovFlow(uint64(l))
	}
	l = len(m.SrcVnic)
	if l > 0 {
		n += 1 + l + sovFlow(uint64(l))
	}
	l = len(m.SrcVPC)
	if l > 0 {
		n += 1 + l + sovFlow(uint64(l))
	}
	l = len(m.DstVPC)
	if l > 0 {
		n += 1 + l + sovFlow(uint64(l))
	}
	l = len(m.SecurityPolicy)
	if l > 0 {
		n += 1 + l + sovFlow(uint64(l))
	}
	l = len(m.SecurityPolicyRule)
	if l > 0 {
		n += 1 + l + sovFlow(uint64(l))
	}
	if m.FlowAction != 0 {
		n += 1 + sovFlow(uint64(m.FlowAction))
	}
	if m.SrcNATIP != nil {
		l = m.SrcNATIP.Size()
		n += 1 + l + sovFlow(uint64(l))
	}
	if m.SrcNATPort != 0 {
		n += 1 + sovFlow(uint64(m.SrcNATPort))
	}
	if m.DstNATIP != nil {
		l = m.DstNATIP.Size()
		n += 1 + l + sovFlow(uint64(l))
	}
	if m.DstNATPort != 0 {
		n += 1 + sovFlow(uint64(m.DstNATPort))
	}
	if m.IflowBytes != 0 {
		n += 2 + sovFlow(uint64(m.IflowBytes))
	}
	if m.IflowPackets != 0 {
		n += 2 + sovFlow(uint64(m.IflowPackets))
	}
	if m.RflowBytes != 0 {
		n += 2 + sovFlow(uint64(m.RflowBytes))
	}
	if m.RflowPackets != 0 {
		n += 2 + sovFlow(uint64(m.RflowPackets))
	}
	if m.SrcVlan != 0 {
		n += 2 + sovFlow(uint64(m.SrcVlan))
	}
	if m.SecurityPloicyRuleId != 0 {
		n += 2 + sovFlow(uint64(m.SecurityPloicyRuleId))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovFlow(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozFlow(x uint64) (n int) {
	return sovFlow(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *FlowLog) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFlow
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
			return fmt.Errorf("proto: FlowLog: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FlowLog: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			m.Timestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFlow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timestamp |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LogType", wireType)
			}
			m.LogType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFlow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LogType |= FlowLogType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFlow
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
				return ErrInvalidLengthFlow
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFlow
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Key == nil {
				m.Key = &pds.FlowKey{}
			}
			if err := m.Key.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SessionId", wireType)
			}
			m.SessionId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFlow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SessionId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SrcInterface", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFlow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthFlow
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthFlow
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SrcInterface = append(m.SrcInterface[:0], dAtA[iNdEx:postIndex]...)
			if m.SrcInterface == nil {
				m.SrcInterface = []byte{}
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SrcVnic", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFlow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthFlow
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthFlow
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SrcVnic = append(m.SrcVnic[:0], dAtA[iNdEx:postIndex]...)
			if m.SrcVnic == nil {
				m.SrcVnic = []byte{}
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SrcVPC", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFlow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthFlow
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthFlow
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SrcVPC = append(m.SrcVPC[:0], dAtA[iNdEx:postIndex]...)
			if m.SrcVPC == nil {
				m.SrcVPC = []byte{}
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DstVPC", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFlow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthFlow
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthFlow
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DstVPC = append(m.DstVPC[:0], dAtA[iNdEx:postIndex]...)
			if m.DstVPC == nil {
				m.DstVPC = []byte{}
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SecurityPolicy", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFlow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthFlow
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthFlow
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SecurityPolicy = append(m.SecurityPolicy[:0], dAtA[iNdEx:postIndex]...)
			if m.SecurityPolicy == nil {
				m.SecurityPolicy = []byte{}
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SecurityPolicyRule", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFlow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthFlow
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthFlow
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SecurityPolicyRule = append(m.SecurityPolicyRule[:0], dAtA[iNdEx:postIndex]...)
			if m.SecurityPolicyRule == nil {
				m.SecurityPolicyRule = []byte{}
			}
			iNdEx = postIndex
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FlowAction", wireType)
			}
			m.FlowAction = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFlow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FlowAction |= pds.SecurityRuleAction(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SrcNATIP", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFlow
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
				return ErrInvalidLengthFlow
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFlow
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SrcNATIP == nil {
				m.SrcNATIP = &pds.IPAddress{}
			}
			if err := m.SrcNATIP.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SrcNATPort", wireType)
			}
			m.SrcNATPort = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFlow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SrcNATPort |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DstNATIP", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFlow
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
				return ErrInvalidLengthFlow
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFlow
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.DstNATIP == nil {
				m.DstNATIP = &pds.IPAddress{}
			}
			if err := m.DstNATIP.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 15:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DstNATPort", wireType)
			}
			m.DstNATPort = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFlow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DstNATPort |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 16:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IflowBytes", wireType)
			}
			m.IflowBytes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFlow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.IflowBytes |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 17:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IflowPackets", wireType)
			}
			m.IflowPackets = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFlow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.IflowPackets |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 18:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RflowBytes", wireType)
			}
			m.RflowBytes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFlow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RflowBytes |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 19:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RflowPackets", wireType)
			}
			m.RflowPackets = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFlow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RflowPackets |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 20:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SrcVlan", wireType)
			}
			m.SrcVlan = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFlow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SrcVlan |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 21:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SecurityPloicyRuleId", wireType)
			}
			m.SecurityPloicyRuleId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFlow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SecurityPloicyRuleId |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipFlow(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFlow
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthFlow
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
func skipFlow(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFlow
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
					return 0, ErrIntOverflowFlow
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
					return 0, ErrIntOverflowFlow
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
				return 0, ErrInvalidLengthFlow
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthFlow
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowFlow
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
				next, err := skipFlow(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthFlow
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
	ErrInvalidLengthFlow = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFlow   = fmt.Errorf("proto: integer overflow")
)