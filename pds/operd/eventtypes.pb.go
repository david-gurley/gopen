// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: eventtypes.proto

package operd

import (
	fmt "fmt"
	math "math"

	proto "github.com/gogo/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
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

type EventType int32

const (
	EventType_DSC_EVENT_TYPE_NONE EventType = 0
	// ----------------------------- SYSTEM events --------------------------- //
	EventType_DSC_SERVICE_STARTED                  EventType = 1
	EventType_DSC_SERVICE_STOPPED                  EventType = 2
	EventType_SYSTEM_COLDBOOT                      EventType = 3
	EventType_DSC_MEM_TEMP_ABOVE_THRESHOLD         EventType = 4
	EventType_DSC_MEM_TEMP_BELOW_THRESHOLD         EventType = 5
	EventType_DSC_CATTRIP_INTERRUPT                EventType = 6
	EventType_DSC_PANIC_EVENT                      EventType = 7
	EventType_DSC_POST_DIAG_FAILURE_EVENT          EventType = 8
	EventType_DSC_INFO_PCIEHEALTH_EVENT            EventType = 9
	EventType_DSC_WARN_PCIEHEALTH_EVENT            EventType = 10
	EventType_DSC_ERR_PCIEHEALTH_EVENT             EventType = 11
	EventType_DSC_FILESYSTEM_USAGE_ABOVE_THRESHOLD EventType = 12
	EventType_DSC_FILESYSTEM_USAGE_BELOW_THRESHOLD EventType = 13
	EventType_DSC_SERVICE_RESTARTED                EventType = 14
	// ----------------------------- Network events -------------------------- //
	EventType_LINK_UP                   EventType = 1001
	EventType_LINK_DOWN                 EventType = 1002
	EventType_BGP_SESSION_ESTABLISHED   EventType = 1003
	EventType_BGP_SESSION_DOWN          EventType = 1004
	EventType_BGP_EVPN_DUP_MAC_IP       EventType = 1005
	EventType_BGP_EVPN_DUP_MAC_IP_CLEAR EventType = 1006
	// ----------------------------- Resource events ------------------------- //
	EventType_VNIC_SESSION_LIMIT_EXCEEDED     EventType = 2001
	EventType_VNIC_SESSION_THRESHOLD_EXCEEDED EventType = 2002
	EventType_VNIC_SESSION_WITHIN_THRESHOLD   EventType = 2003
	// ----------------------------- Learn events ---------------------------- //
	EventType_LEARN_PKT EventType = 3001
	// ----------------------------- Interrupt events ------------------------ //
	EventType_DSC_HW_RMA_INTERRUPT EventType = 4001
	EventType_DSC_FATAL_INTERRUPT  EventType = 4002
	EventType_DSC_ERROR_INTERRUPT  EventType = 4003
	// ----------------------------- NAT events ---------------------------- //
	EventType_NAT_PORT_USAGE_THRESHOLD_EXCEEDED EventType = 5001
	// ----------------------------- HA events ----------------------------- //
	EventType_HA_FLOW_SYNC_VALIDATION_FAILED    EventType = 6001
	EventType_HA_FLOW_SYNC_PEER_INIT_FAILED     EventType = 6002
	EventType_HA_FLOW_SYNC_PEER_CONNECTION_DOWN EventType = 6003
	EventType_HA_FLOW_SYNC_PEER_CONNECTION_UP   EventType = 6004
	EventType_HA_FLOW_SYNC_FSM_RESUME_FAILED    EventType = 6005
)

var EventType_name = map[int32]string{
	0:    "DSC_EVENT_TYPE_NONE",
	1:    "DSC_SERVICE_STARTED",
	2:    "DSC_SERVICE_STOPPED",
	3:    "SYSTEM_COLDBOOT",
	4:    "DSC_MEM_TEMP_ABOVE_THRESHOLD",
	5:    "DSC_MEM_TEMP_BELOW_THRESHOLD",
	6:    "DSC_CATTRIP_INTERRUPT",
	7:    "DSC_PANIC_EVENT",
	8:    "DSC_POST_DIAG_FAILURE_EVENT",
	9:    "DSC_INFO_PCIEHEALTH_EVENT",
	10:   "DSC_WARN_PCIEHEALTH_EVENT",
	11:   "DSC_ERR_PCIEHEALTH_EVENT",
	12:   "DSC_FILESYSTEM_USAGE_ABOVE_THRESHOLD",
	13:   "DSC_FILESYSTEM_USAGE_BELOW_THRESHOLD",
	14:   "DSC_SERVICE_RESTARTED",
	1001: "LINK_UP",
	1002: "LINK_DOWN",
	1003: "BGP_SESSION_ESTABLISHED",
	1004: "BGP_SESSION_DOWN",
	1005: "BGP_EVPN_DUP_MAC_IP",
	1006: "BGP_EVPN_DUP_MAC_IP_CLEAR",
	2001: "VNIC_SESSION_LIMIT_EXCEEDED",
	2002: "VNIC_SESSION_THRESHOLD_EXCEEDED",
	2003: "VNIC_SESSION_WITHIN_THRESHOLD",
	3001: "LEARN_PKT",
	4001: "DSC_HW_RMA_INTERRUPT",
	4002: "DSC_FATAL_INTERRUPT",
	4003: "DSC_ERROR_INTERRUPT",
	5001: "NAT_PORT_USAGE_THRESHOLD_EXCEEDED",
	6001: "HA_FLOW_SYNC_VALIDATION_FAILED",
	6002: "HA_FLOW_SYNC_PEER_INIT_FAILED",
	6003: "HA_FLOW_SYNC_PEER_CONNECTION_DOWN",
	6004: "HA_FLOW_SYNC_PEER_CONNECTION_UP",
	6005: "HA_FLOW_SYNC_FSM_RESUME_FAILED",
}

var EventType_value = map[string]int32{
	"DSC_EVENT_TYPE_NONE":                  0,
	"DSC_SERVICE_STARTED":                  1,
	"DSC_SERVICE_STOPPED":                  2,
	"SYSTEM_COLDBOOT":                      3,
	"DSC_MEM_TEMP_ABOVE_THRESHOLD":         4,
	"DSC_MEM_TEMP_BELOW_THRESHOLD":         5,
	"DSC_CATTRIP_INTERRUPT":                6,
	"DSC_PANIC_EVENT":                      7,
	"DSC_POST_DIAG_FAILURE_EVENT":          8,
	"DSC_INFO_PCIEHEALTH_EVENT":            9,
	"DSC_WARN_PCIEHEALTH_EVENT":            10,
	"DSC_ERR_PCIEHEALTH_EVENT":             11,
	"DSC_FILESYSTEM_USAGE_ABOVE_THRESHOLD": 12,
	"DSC_FILESYSTEM_USAGE_BELOW_THRESHOLD": 13,
	"DSC_SERVICE_RESTARTED":                14,
	"LINK_UP":                              1001,
	"LINK_DOWN":                            1002,
	"BGP_SESSION_ESTABLISHED":              1003,
	"BGP_SESSION_DOWN":                     1004,
	"BGP_EVPN_DUP_MAC_IP":                  1005,
	"BGP_EVPN_DUP_MAC_IP_CLEAR":            1006,
	"VNIC_SESSION_LIMIT_EXCEEDED":          2001,
	"VNIC_SESSION_THRESHOLD_EXCEEDED":      2002,
	"VNIC_SESSION_WITHIN_THRESHOLD":        2003,
	"LEARN_PKT":                            3001,
	"DSC_HW_RMA_INTERRUPT":                 4001,
	"DSC_FATAL_INTERRUPT":                  4002,
	"DSC_ERROR_INTERRUPT":                  4003,
	"NAT_PORT_USAGE_THRESHOLD_EXCEEDED":    5001,
	"HA_FLOW_SYNC_VALIDATION_FAILED":       6001,
	"HA_FLOW_SYNC_PEER_INIT_FAILED":        6002,
	"HA_FLOW_SYNC_PEER_CONNECTION_DOWN":    6003,
	"HA_FLOW_SYNC_PEER_CONNECTION_UP":      6004,
	"HA_FLOW_SYNC_FSM_RESUME_FAILED":       6005,
}

func (x EventType) String() string {
	return proto.EnumName(EventType_name, int32(x))
}

func (EventType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a92369539d307339, []int{0}
}

var E_Category = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.EnumValueOptions)(nil),
	ExtensionType: (*EventCategory)(nil),
	Field:         100001,
	Name:          "operd.Category",
	Tag:           "varint,100001,opt,name=Category,enum=operd.EventCategory",
	Filename:      "eventtypes.proto",
}

var E_Severity = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.EnumValueOptions)(nil),
	ExtensionType: (*EventSeverity)(nil),
	Field:         100002,
	Name:          "operd.Severity",
	Tag:           "varint,100002,opt,name=Severity,enum=operd.EventSeverity",
	Filename:      "eventtypes.proto",
}

var E_Description = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.EnumValueOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         100003,
	Name:          "operd.Description",
	Tag:           "bytes,100003,opt,name=Description",
	Filename:      "eventtypes.proto",
}

func init() {
	proto.RegisterEnum("operd.EventType", EventType_name, EventType_value)
	proto.RegisterExtension(E_Category)
	proto.RegisterExtension(E_Severity)
	proto.RegisterExtension(E_Description)
}

func init() { proto.RegisterFile("eventtypes.proto", fileDescriptor_a92369539d307339) }

var fileDescriptor_a92369539d307339 = []byte{
	// 1720 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x97, 0xcb, 0x6f, 0xdc, 0xc6,
	0x1d, 0xc7, 0xbd, 0x76, 0x62, 0x45, 0xcc, 0x43, 0x2c, 0x13, 0x23, 0x31, 0x1d, 0x4b, 0x63, 0x20,
	0x8d, 0xe3, 0x34, 0x59, 0x6f, 0x9d, 0x3e, 0x83, 0xb6, 0x28, 0xb5, 0x3b, 0xb2, 0xd8, 0xec, 0x92,
	0x04, 0x97, 0x2b, 0xd5, 0x4d, 0x01, 0x62, 0x96, 0xfc, 0x69, 0x77, 0x2a, 0xee, 0x0c, 0x3b, 0x9c,
	0x95, 0xaa, 0x9e, 0x0a, 0xe4, 0xa2, 0x02, 0x6d, 0x91, 0x53, 0x81, 0xfa, 0x96, 0xe4, 0xda, 0x3f,
	0xa0, 0xfd, 0x0f, 0x7a, 0xec, 0xe3, 0xda, 0x43, 0xe1, 0x5b, 0xb6, 0x0f, 0xa0, 0x4d, 0x7b, 0x2f,
	0x66, 0x96, 0x5c, 0x51, 0x2f, 0x47, 0x07, 0x5f, 0xc9, 0xf9, 0x7d, 0x7e, 0xdf, 0xf9, 0x3d, 0x49,
	0xc3, 0x84, 0x3d, 0x60, 0x52, 0x1e, 0xe4, 0x50, 0x34, 0x73, 0xc1, 0x25, 0xb7, 0x9e, 0xe6, 0x39,
	0x88, 0xd4, 0x46, 0x23, 0xce, 0x47, 0x19, 0xdc, 0xd5, 0x0f, 0x87, 0xd3, 0x9d, 0xbb, 0x29, 0x14,
	0x89, 0xa0, 0xb9, 0xe4, 0x62, 0x7e, 0xd0, 0xbe, 0xa6, 0x4d, 0x89, 0x94, 0x82, 0x0e, 0xa7, 0xb2,
	0xb2, 0x7f, 0xf3, 0xb3, 0xeb, 0xc6, 0x32, 0x56, 0x6f, 0xa2, 0x83, 0x1c, 0xac, 0x7b, 0xc6, 0x8b,
	0x9d, 0x7e, 0x3b, 0xc6, 0x5b, 0xd8, 0x8b, 0xe2, 0xe8, 0x41, 0x80, 0x63, 0xcf, 0xf7, 0xb0, 0x79,
	0xc9, 0xbe, 0x7e, 0x38, 0x6b, 0x5d, 0xfa, 0x70, 0xd6, 0xba, 0xf4, 0x70, 0xd6, 0x7a, 0xde, 0x65,
	0x7b, 0x24, 0xa3, 0x29, 0xd2, 0x44, 0xeb, 0x1b, 0x73, 0x9b, 0x3e, 0x0e, 0xb7, 0xdc, 0x36, 0x8e,
	0xfb, 0x91, 0x13, 0x46, 0xb8, 0x63, 0x36, 0xec, 0xb5, 0xc3, 0x59, 0xab, 0x51, 0xda, 0xa8, 0x23,
	0xa8, 0x00, 0xb1, 0x47, 0x13, 0x40, 0x85, 0x24, 0x42, 0x42, 0x6a, 0x7d, 0x70, 0xf9, 0xa4, 0xa9,
	0x1f, 0x04, 0xb8, 0x63, 0x5e, 0xb6, 0x3f, 0x6d, 0x94, 0xb6, 0x57, 0x1e, 0xce, 0x5a, 0x7f, 0x6d,
	0x28, 0x63, 0x01, 0x19, 0x91, 0x90, 0xa2, 0x5c, 0xf0, 0x04, 0x8a, 0x02, 0x8d, 0x49, 0x81, 0x86,
	0x00, 0x0c, 0x15, 0x92, 0xe7, 0x39, 0xa4, 0x4d, 0xd4, 0x3f, 0x28, 0x24, 0x4c, 0x50, 0x0e, 0x62,
	0x87, 0x8b, 0x09, 0x61, 0x09, 0xa0, 0x09, 0x39, 0x40, 0x43, 0x40, 0x74, 0x92, 0x93, 0x44, 0xaa,
	0x33, 0x1b, 0x53, 0x21, 0xc7, 0x20, 0x50, 0xce, 0x33, 0x9a, 0x1c, 0xa0, 0x64, 0x4c, 0xd8, 0x08,
	0x0a, 0xb4, 0x4f, 0xb3, 0x0c, 0x31, 0x2e, 0xd5, 0x61, 0x92, 0xe7, 0x19, 0x85, 0x14, 0x49, 0x8e,
	0xe4, 0x98, 0x16, 0x48, 0x39, 0x97, 0xea, 0x3d, 0x95, 0x88, 0x16, 0x4a, 0x07, 0x27, 0xa9, 0x62,
	0x85, 0x90, 0xf0, 0xc9, 0x04, 0x58, 0x4a, 0x24, 0xe5, 0x4c, 0xbd, 0x93, 0x1c, 0x25, 0x3c, 0xcb,
	0x20, 0x91, 0x48, 0x42, 0x32, 0x2e, 0xa6, 0x79, 0xce, 0x85, 0x44, 0x84, 0xa5, 0x28, 0xe1, 0x4c,
	0x92, 0x44, 0xa2, 0x00, 0x58, 0x41, 0x58, 0xca, 0x51, 0xf9, 0xb2, 0x69, 0xb5, 0x8c, 0x95, 0xfe,
	0x83, 0x7e, 0x84, 0x7b, 0x71, 0xdb, 0xef, 0x76, 0xd6, 0x7d, 0x3f, 0x32, 0xaf, 0xd8, 0x37, 0xca,
	0xfb, 0x5f, 0x7e, 0x38, 0x6b, 0xad, 0x28, 0x05, 0x09, 0xcf, 0x52, 0x34, 0xe4, 0x5c, 0xc5, 0xed,
	0xb0, 0x61, 0xbc, 0xaa, 0xe2, 0xd6, 0xc3, 0xbd, 0x38, 0xc2, 0xbd, 0x20, 0x76, 0xd6, 0xfd, 0x2d,
	0x1c, 0x47, 0x9b, 0x21, 0xee, 0x6f, 0xfa, 0xdd, 0x8e, 0xf9, 0x94, 0xbd, 0x53, 0x8b, 0xdf, 0x0f,
	0x94, 0xfd, 0x04, 0x26, 0x5c, 0x1c, 0x20, 0x09, 0x93, 0x1c, 0x04, 0x91, 0x53, 0x01, 0x4a, 0x30,
	0x19, 0xf2, 0x3d, 0x40, 0x72, 0x0c, 0x28, 0x11, 0x54, 0xd2, 0x84, 0x64, 0x48, 0x8e, 0x05, 0x14,
	0x63, 0x9e, 0xa9, 0x80, 0x6a, 0x59, 0x67, 0x05, 0x34, 0x85, 0x91, 0x50, 0x41, 0xb0, 0xb2, 0x13,
	0x4a, 0xd6, 0x71, 0xd7, 0xdf, 0xae, 0x29, 0x79, 0xda, 0xfe, 0x5e, 0xa9, 0xa4, 0xf1, 0x70, 0xd6,
	0xfa, 0xce, 0x39, 0x4a, 0x54, 0x2e, 0x77, 0x48, 0x96, 0x01, 0x43, 0x43, 0xc8, 0xf8, 0xfe, 0x39,
	0x8a, 0xac, 0xdf, 0x36, 0x8c, 0x6b, 0xca, 0x5d, 0xdb, 0x89, 0xa2, 0xd0, 0x0d, 0x62, 0xd7, 0x8b,
	0x70, 0x18, 0x0e, 0x82, 0xc8, 0xbc, 0x6a, 0x7f, 0x58, 0x2f, 0x99, 0x0f, 0x74, 0xc9, 0x9c, 0x74,
	0x91, 0x08, 0x5e, 0x14, 0x90, 0x22, 0x82, 0x76, 0x88, 0xac, 0x83, 0xdf, 0xd2, 0x29, 0x5a, 0x64,
	0x7b, 0x51, 0x5a, 0x9f, 0x97, 0xeb, 0x31, 0x24, 0xbb, 0x08, 0xd8, 0x1e, 0x15, 0x9c, 0x4d, 0x54,
	0x77, 0x65, 0x2a, 0xcf, 0x29, 0x55, 0x67, 0x8a, 0xa6, 0xf5, 0xeb, 0x86, 0xa1, 0x72, 0x17, 0x07,
	0x8e, 0xe7, 0x96, 0x4d, 0x65, 0x2e, 0xd9, 0x3f, 0xab, 0x0b, 0x95, 0xca, 0x1f, 0xb0, 0x84, 0x4f,
	0x99, 0x04, 0xa1, 0xb5, 0xe5, 0x84, 0xd1, 0x04, 0xa5, 0x53, 0x41, 0xd9, 0x48, 0x07, 0x22, 0x23,
	0x85, 0xd4, 0xa9, 0x7f, 0x62, 0x25, 0xb7, 0x61, 0xdc, 0xd0, 0xba, 0xfc, 0x7e, 0x14, 0x77, 0x5c,
	0xe7, 0x7e, 0xbc, 0xe1, 0xb8, 0xdd, 0x41, 0x88, 0x4b, 0x8d, 0xcf, 0xd8, 0x5f, 0xac, 0x95, 0xdf,
	0x75, 0x25, 0x31, 0xe7, 0x85, 0x44, 0x29, 0x25, 0x23, 0x24, 0xa1, 0x90, 0x68, 0x87, 0xd0, 0x0c,
	0x52, 0x2b, 0x32, 0xd4, 0xcb, 0xd8, 0xf5, 0x36, 0xfc, 0x38, 0x68, 0xbb, 0x78, 0x13, 0x3b, 0xdd,
	0x68, 0xb3, 0xa4, 0x2c, 0xdb, 0x5f, 0xad, 0xa5, 0xfe, 0x4e, 0x15, 0xd8, 0x14, 0x24, 0xa8, 0x5e,
	0x54, 0x37, 0x4d, 0x28, 0xa0, 0x8c, 0xb2, 0x5d, 0x34, 0x06, 0x92, 0xc9, 0x71, 0x39, 0x50, 0xb6,
	0xe6, 0xd4, 0x6d, 0x27, 0xf4, 0x4e, 0x53, 0x0d, 0xfb, 0xeb, 0x35, 0x6d, 0x5f, 0xba, 0x08, 0x75,
	0x9f, 0x08, 0x46, 0xd9, 0xc8, 0xfa, 0x5d, 0xc3, 0x78, 0x45, 0x4f, 0xb7, 0x30, 0x3c, 0xcd, 0x7d,
	0xd6, 0xfe, 0x4d, 0x3d, 0x2f, 0xbf, 0x68, 0x9c, 0x22, 0x07, 0x6d, 0xb7, 0xe4, 0x82, 0x10, 0x5c,
	0x5c, 0x6c, 0xde, 0x3c, 0xa1, 0x84, 0xfd, 0xc8, 0x78, 0x4d, 0x29, 0xdf, 0x70, 0xbb, 0xb8, 0x9c,
	0x15, 0x83, 0xbe, 0x73, 0x1f, 0x9f, 0x6a, 0xfc, 0xe7, 0xec, 0xef, 0xd6, 0xa2, 0xf3, 0x95, 0x41,
	0x41, 0x46, 0x80, 0xf8, 0x8e, 0x2e, 0xa3, 0x1d, 0x9a, 0x41, 0x51, 0x4a, 0x26, 0x42, 0xd2, 0x4a,
	0x50, 0x35, 0x03, 0xaa, 0x26, 0x3b, 0xcf, 0xd7, 0xc9, 0xd6, 0x7e, 0x7e, 0xe1, 0xab, 0x71, 0x61,
	0x5f, 0x55, 0x77, 0x57, 0xbe, 0xbe, 0x35, 0xef, 0xe7, 0x6a, 0x01, 0x84, 0xb8, 0xda, 0x1e, 0x2f,
	0xd8, 0xb7, 0x6a, 0x17, 0xb9, 0x56, 0xdf, 0x1e, 0x02, 0xaa, 0xfd, 0x71, 0xc7, 0x58, 0xea, 0xba,
	0xde, 0x7b, 0xf1, 0x20, 0x30, 0x3f, 0x5d, 0xd2, 0x23, 0xf3, 0x72, 0xa9, 0x66, 0x25, 0x50, 0x31,
	0xd5, 0xc9, 0xa2, 0x05, 0x9a, 0xe6, 0xd6, 0xdb, 0xc6, 0xb2, 0x3e, 0xda, 0xf1, 0xb7, 0x3d, 0x73,
	0xb6, 0x64, 0xdf, 0x2c, 0x0f, 0x2b, 0xfa, 0x17, 0x8e, 0x1d, 0x4e, 0xf9, 0x3e, 0xb3, 0x3a, 0xc6,
	0xcb, 0xeb, 0xf7, 0x83, 0xb8, 0x8f, 0xfb, 0x7d, 0xd7, 0xf7, 0x62, 0x25, 0x6b, 0xbd, 0xeb, 0xf6,
	0x37, 0x71, 0xc7, 0xfc, 0xfb, 0x92, 0xfd, 0x7a, 0xcd, 0x93, 0xbd, 0x7e, 0x3f, 0x40, 0x05, 0x14,
	0x45, 0x79, 0x43, 0x25, 0x6e, 0x98, 0xd1, 0x62, 0x0c, 0xa9, 0xf5, 0x35, 0xc3, 0xac, 0x53, 0xb4,
	0xef, 0x7f, 0x2c, 0xe9, 0xbd, 0x58, 0xf9, 0x7e, 0xf1, 0x84, 0xb9, 0xf6, 0xee, 0x1b, 0xea, 0x71,
	0x8c, 0xb7, 0x02, 0x2f, 0xee, 0x0c, 0x82, 0xb8, 0xe7, 0xb4, 0x63, 0x37, 0x30, 0xff, 0xb9, 0xa4,
	0x3b, 0xaa, 0x32, 0xbd, 0xa3, 0x4c, 0xd5, 0x99, 0xa3, 0x12, 0x4d, 0xa7, 0x79, 0x46, 0x13, 0x22,
	0x01, 0x49, 0x60, 0x84, 0x49, 0xd4, 0x73, 0xda, 0x77, 0xdd, 0xc0, 0x1a, 0x18, 0xd7, 0xcf, 0x00,
	0xc6, 0xed, 0x2e, 0x76, 0x42, 0xf3, 0x5f, 0x47, 0xd8, 0xc6, 0x71, 0xec, 0xd9, 0x34, 0x95, 0x00,
	0x9e, 0xed, 0x41, 0x6a, 0xe5, 0xc6, 0x8d, 0x2d, 0x35, 0xd9, 0xaa, 0x0b, 0x76, 0xdd, 0x9e, 0x1b,
	0xc5, 0xf8, 0xfb, 0x6d, 0x8c, 0x3b, 0xb8, 0x63, 0xfe, 0x69, 0xc5, 0xf6, 0x0e, 0x67, 0xad, 0x2b,
	0x65, 0x4b, 0xad, 0xf7, 0xcb, 0x6b, 0xea, 0x61, 0x87, 0x76, 0xb8, 0x40, 0x7b, 0x6a, 0xd2, 0x09,
	0x20, 0xc9, 0x18, 0x52, 0x94, 0xd1, 0x09, 0x95, 0x6f, 0x21, 0x06, 0xfb, 0x55, 0x40, 0xca, 0xbd,
	0xac, 0xf6, 0x8d, 0xd0, 0x3b, 0xde, 0xfa, 0x55, 0xc3, 0x58, 0x3b, 0xe6, 0x72, 0x51, 0x8e, 0x47,
	0x6e, 0xff, 0xbc, 0x62, 0x8f, 0x4b, 0xb7, 0x2a, 0x4c, 0x3f, 0x3c, 0xc7, 0x2d, 0xfc, 0x24, 0x01,
	0x48, 0x8b, 0xfa, 0x12, 0x78, 0x9c, 0x6b, 0xc4, 0x55, 0x7f, 0x6b, 0x8d, 0xf3, 0xaf, 0x00, 0x2d,
	0xda, 0x22, 0xc6, 0xcd, 0x63, 0x7a, 0xb6, 0xdd, 0x68, 0xd3, 0xad, 0xc9, 0x32, 0xff, 0xb2, 0x62,
	0x7f, 0xbb, 0x54, 0xa3, 0xa2, 0xfb, 0xe5, 0x73, 0xd4, 0xa8, 0xed, 0xa7, 0x7c, 0xca, 0x31, 0x65,
	0xba, 0x83, 0x8e, 0x7a, 0x64, 0xd7, 0x58, 0x56, 0x79, 0xf2, 0xe2, 0xe0, 0xbd, 0xc8, 0xfc, 0xfd,
	0xcb, 0xf6, 0xfb, 0x87, 0xb3, 0xd6, 0x53, 0x25, 0xce, 0x0f, 0x48, 0xb2, 0x0b, 0x12, 0x09, 0x48,
	0x80, 0xee, 0x41, 0xaa, 0x72, 0x33, 0xcd, 0xa4, 0xda, 0x1c, 0x94, 0xe9, 0xfb, 0x00, 0x4b, 0x73,
	0x4e, 0x99, 0x44, 0x19, 0xcc, 0xc7, 0x20, 0xe2, 0x02, 0x4d, 0xf3, 0x94, 0x48, 0xd0, 0xb3, 0x68,
	0x97, 0xf1, 0x7d, 0xb6, 0x38, 0x65, 0x45, 0xc6, 0x4b, 0xaa, 0x21, 0x37, 0xb7, 0xe3, 0xb0, 0xe7,
	0xd4, 0xf6, 0xeb, 0x47, 0x6b, 0xf6, 0x37, 0x0f, 0x67, 0xad, 0xab, 0x65, 0x2e, 0xdf, 0x9e, 0x4f,
	0x47, 0x91, 0xee, 0x13, 0x01, 0x28, 0xec, 0x39, 0x88, 0xaa, 0xfd, 0x25, 0xa6, 0xb9, 0x3c, 0xda,
	0x9e, 0x52, 0xd0, 0xd1, 0x48, 0x2d, 0x35, 0xab, 0x3b, 0xff, 0xce, 0xdb, 0x70, 0x22, 0xa7, 0x5b,
	0x83, 0x7e, 0xbc, 0x66, 0xdf, 0xab, 0x41, 0x5f, 0x57, 0xd0, 0xf9, 0x66, 0xbe, 0x08, 0x0d, 0x87,
	0xa1, 0x1f, 0xd6, 0x68, 0x9f, 0x9c, 0x41, 0xd3, 0x93, 0xfa, 0xf1, 0xb4, 0xf7, 0x8d, 0x5b, 0x9e,
	0x13, 0xc5, 0x81, 0x1f, 0x46, 0xe5, 0xa0, 0x3b, 0xa3, 0xa6, 0x7e, 0x7e, 0xdb, 0x7e, 0xa7, 0x56,
	0x53, 0xb7, 0x43, 0xf8, 0xf1, 0x54, 0xed, 0x41, 0xc9, 0x11, 0xc9, 0x32, 0xae, 0x9b, 0x64, 0xc2,
	0x05, 0x20, 0xcf, 0x89, 0x90, 0x1a, 0xd9, 0xc3, 0x8c, 0x27, 0xbb, 0x85, 0xd5, 0x35, 0x56, 0x37,
	0x9d, 0x78, 0x43, 0x0d, 0xce, 0xfe, 0x03, 0xaf, 0x1d, 0x6f, 0x39, 0x5d, 0xb7, 0xe3, 0x44, 0xaa,
	0x52, 0xd4, 0xca, 0xc5, 0x1d, 0xf3, 0xdf, 0x4d, 0xfb, 0xf6, 0xe1, 0xac, 0xb5, 0x54, 0xaa, 0xbe,
	0xb1, 0xa1, 0x06, 0x64, 0x71, 0xc0, 0x12, 0xa4, 0xbf, 0xb1, 0xe7, 0x6b, 0x63, 0xb1, 0x6e, 0x6f,
	0x1e, 0xa3, 0x05, 0x18, 0xab, 0x00, 0xb8, 0x51, 0x05, 0xfb, 0x4f, 0xd3, 0x6e, 0xd5, 0x60, 0xaf,
	0x1d, 0xc1, 0x72, 0x00, 0x15, 0x07, 0x2a, 0x29, 0xc9, 0xe8, 0x4f, 0x8f, 0x51, 0x43, 0xe3, 0xd6,
	0x69, 0x6a, 0xdb, 0xf7, 0x3c, 0xdc, 0x8e, 0x16, 0x63, 0xeb, 0xb3, 0xa6, 0xfd, 0x66, 0x8d, 0xbc,
	0x7a, 0x82, 0x9c, 0x70, 0xc6, 0x20, 0xd1, 0x54, 0x3d, 0xc1, 0x3c, 0x63, 0xed, 0xb1, 0xcc, 0x41,
	0x60, 0xfe, 0xb7, 0x69, 0xbf, 0x51, 0x12, 0x55, 0x25, 0xbf, 0x7a, 0x3e, 0x71, 0x9a, 0x5b, 0xde,
	0x89, 0x38, 0x6e, 0xf4, 0x7b, 0x6a, 0x59, 0x0c, 0x7a, 0xb8, 0xba, 0xfa, 0xff, 0x9a, 0xf6, 0x9d,
	0x9a, 0xc0, 0x9b, 0x47, 0xb8, 0x42, 0xaa, 0xc5, 0xa4, 0x1a, 0x63, 0x02, 0xe5, 0x9d, 0xdf, 0x1d,
	0x18, 0xcf, 0xb4, 0x89, 0x84, 0x11, 0x17, 0x07, 0xd6, 0xad, 0xe6, 0xfc, 0xdf, 0xa9, 0x59, 0xfd,
	0x3b, 0x35, 0x31, 0x9b, 0x4e, 0xb6, 0x48, 0x36, 0x05, 0x3f, 0xd7, 0x5f, 0x72, 0xaf, 0x7c, 0xf4,
	0xcb, 0xab, 0xa8, 0xf1, 0xc6, 0x0b, 0xf7, 0x5e, 0x6a, 0xea, 0xbf, 0xad, 0xa6, 0xfe, 0x61, 0xaa,
	0x00, 0xe1, 0x02, 0xa5, 0xb0, 0x7d, 0xd8, 0x03, 0x41, 0xe5, 0x85, 0xb0, 0x1f, 0x9f, 0x85, 0xad,
	0x00, 0xe1, 0x02, 0xf5, 0x2e, 0x36, 0x9e, 0xed, 0x94, 0xbf, 0x73, 0x94, 0xb3, 0x8b, 0x90, 0x3f,
	0xd1, 0xe4, 0xe5, 0xb0, 0x6e, 0xb7, 0xfe, 0xdc, 0x1f, 0x1e, 0xad, 0x36, 0xfe, 0xf8, 0x68, 0xb5,
	0xf1, 0xb7, 0x47, 0xab, 0x8d, 0xe1, 0x55, 0x6d, 0xfd, 0xce, 0xff, 0x03, 0x00, 0x00, 0xff, 0xff,
	0xe4, 0x68, 0x58, 0x31, 0x53, 0x0e, 0x00, 0x00,
}
