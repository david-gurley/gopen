// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: eventinfo.proto

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

// Endpoint info for learning
type EpLearnPkt struct {
	// host interface on which packet was received
	HostIf []byte `protobuf:"bytes,1,opt,name=HostIf,proto3" json:"HostIf,omitempty"`
	// packet as received upto a maximum size of 128 bytes
	Packet               []byte   `protobuf:"bytes,2,opt,name=Packet,proto3" json:"Packet,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EpLearnPkt) Reset()         { *m = EpLearnPkt{} }
func (m *EpLearnPkt) String() string { return proto.CompactTextString(m) }
func (*EpLearnPkt) ProtoMessage()    {}
func (*EpLearnPkt) Descriptor() ([]byte, []int) {
	return fileDescriptor_3038fa90fea99f0a, []int{0}
}
func (m *EpLearnPkt) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EpLearnPkt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EpLearnPkt.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EpLearnPkt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EpLearnPkt.Merge(m, src)
}
func (m *EpLearnPkt) XXX_Size() int {
	return m.Size()
}
func (m *EpLearnPkt) XXX_DiscardUnknown() {
	xxx_messageInfo_EpLearnPkt.DiscardUnknown(m)
}

var xxx_messageInfo_EpLearnPkt proto.InternalMessageInfo

func (m *EpLearnPkt) GetHostIf() []byte {
	if m != nil {
		return m.HostIf
	}
	return nil
}

func (m *EpLearnPkt) GetPacket() []byte {
	if m != nil {
		return m.Packet
	}
	return nil
}

// NAT port block event specification
type NatPortblockEvent struct {
	// unique key identifying the port block
	Id []byte `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	// protocol
	Protocol uint32 `protobuf:"varint,2,opt,name=Protocol,proto3" json:"Protocol,omitempty"`
	// source NAT IP address
	SnatIP *pds.IPAddress `protobuf:"bytes,3,opt,name=SnatIP,proto3" json:"SnatIP,omitempty"`
	// NAT address is from Internet or Cloud Infra space
	AddressType          pds.AddressType `protobuf:"varint,4,opt,name=AddressType,proto3,enum=types.AddressType" json:"AddressType,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *NatPortblockEvent) Reset()         { *m = NatPortblockEvent{} }
func (m *NatPortblockEvent) String() string { return proto.CompactTextString(m) }
func (*NatPortblockEvent) ProtoMessage()    {}
func (*NatPortblockEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_3038fa90fea99f0a, []int{1}
}
func (m *NatPortblockEvent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NatPortblockEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NatPortblockEvent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NatPortblockEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NatPortblockEvent.Merge(m, src)
}
func (m *NatPortblockEvent) XXX_Size() int {
	return m.Size()
}
func (m *NatPortblockEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_NatPortblockEvent.DiscardUnknown(m)
}

var xxx_messageInfo_NatPortblockEvent proto.InternalMessageInfo

func (m *NatPortblockEvent) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *NatPortblockEvent) GetProtocol() uint32 {
	if m != nil {
		return m.Protocol
	}
	return 0
}

func (m *NatPortblockEvent) GetSnatIP() *pds.IPAddress {
	if m != nil {
		return m.SnatIP
	}
	return nil
}

func (m *NatPortblockEvent) GetAddressType() pds.AddressType {
	if m != nil {
		return m.AddressType
	}
	return pds.AddressType_ADDR_TYPE_NONE
}

func init() {
	proto.RegisterType((*EpLearnPkt)(nil), "operd.EpLearnPkt")
	proto.RegisterType((*NatPortblockEvent)(nil), "operd.NatPortblockEvent")
}

func init() { proto.RegisterFile("eventinfo.proto", fileDescriptor_3038fa90fea99f0a) }

var fileDescriptor_3038fa90fea99f0a = []byte{
	// 229 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4f, 0x2d, 0x4b, 0xcd,
	0x2b, 0xc9, 0xcc, 0x4b, 0xcb, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xcd, 0x2f, 0x48,
	0x2d, 0x4a, 0x91, 0xe2, 0x2e, 0xa9, 0x2c, 0x48, 0x2d, 0x86, 0x88, 0x29, 0xd9, 0x70, 0x71, 0xb9,
	0x16, 0xf8, 0xa4, 0x26, 0x16, 0xe5, 0x05, 0x64, 0x97, 0x08, 0x89, 0x71, 0xb1, 0x79, 0xe4, 0x17,
	0x97, 0x78, 0xa6, 0x49, 0x30, 0x2a, 0x30, 0x6a, 0xf0, 0x04, 0x41, 0x79, 0x20, 0xf1, 0x80, 0xc4,
	0xe4, 0xec, 0xd4, 0x12, 0x09, 0x26, 0x88, 0x38, 0x84, 0xa7, 0x34, 0x9f, 0x91, 0x4b, 0xd0, 0x2f,
	0xb1, 0x24, 0x20, 0xbf, 0xa8, 0x24, 0x29, 0x27, 0x3f, 0x39, 0xdb, 0x15, 0x64, 0xa3, 0x10, 0x1f,
	0x17, 0x93, 0x67, 0x0a, 0xd4, 0x04, 0x26, 0xcf, 0x14, 0x21, 0x29, 0x2e, 0x8e, 0x00, 0x90, 0x65,
	0xc9, 0xf9, 0x39, 0x60, 0xfd, 0xbc, 0x41, 0x70, 0xbe, 0x90, 0x06, 0x17, 0x5b, 0x70, 0x5e, 0x62,
	0x89, 0x67, 0x80, 0x04, 0xb3, 0x02, 0xa3, 0x06, 0xb7, 0x91, 0x80, 0x1e, 0xc4, 0x75, 0x9e, 0x01,
	0x8e, 0x29, 0x29, 0x45, 0xa9, 0xc5, 0xc5, 0x41, 0x50, 0x79, 0x21, 0x13, 0x2e, 0x6e, 0xa8, 0x50,
	0x48, 0x65, 0x41, 0xaa, 0x04, 0x8b, 0x02, 0xa3, 0x06, 0x9f, 0x91, 0x10, 0x54, 0x39, 0x92, 0x4c,
	0x10, 0xb2, 0x32, 0x27, 0x9e, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48,
	0x8e, 0x31, 0x89, 0x0d, 0xec, 0x69, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x81, 0xb7, 0x16,
	0xd8, 0x1b, 0x01, 0x00, 0x00,
}

func (m *EpLearnPkt) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EpLearnPkt) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EpLearnPkt) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Packet) > 0 {
		i -= len(m.Packet)
		copy(dAtA[i:], m.Packet)
		i = encodeVarintEventinfo(dAtA, i, uint64(len(m.Packet)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.HostIf) > 0 {
		i -= len(m.HostIf)
		copy(dAtA[i:], m.HostIf)
		i = encodeVarintEventinfo(dAtA, i, uint64(len(m.HostIf)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *NatPortblockEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NatPortblockEvent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NatPortblockEvent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.AddressType != 0 {
		i = encodeVarintEventinfo(dAtA, i, uint64(m.AddressType))
		i--
		dAtA[i] = 0x20
	}
	if m.SnatIP != nil {
		{
			size, err := m.SnatIP.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintEventinfo(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.Protocol != 0 {
		i = encodeVarintEventinfo(dAtA, i, uint64(m.Protocol))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintEventinfo(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEventinfo(dAtA []byte, offset int, v uint64) int {
	offset -= sovEventinfo(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EpLearnPkt) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.HostIf)
	if l > 0 {
		n += 1 + l + sovEventinfo(uint64(l))
	}
	l = len(m.Packet)
	if l > 0 {
		n += 1 + l + sovEventinfo(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *NatPortblockEvent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovEventinfo(uint64(l))
	}
	if m.Protocol != 0 {
		n += 1 + sovEventinfo(uint64(m.Protocol))
	}
	if m.SnatIP != nil {
		l = m.SnatIP.Size()
		n += 1 + l + sovEventinfo(uint64(l))
	}
	if m.AddressType != 0 {
		n += 1 + sovEventinfo(uint64(m.AddressType))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovEventinfo(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEventinfo(x uint64) (n int) {
	return sovEventinfo(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EpLearnPkt) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEventinfo
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
			return fmt.Errorf("proto: EpLearnPkt: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EpLearnPkt: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HostIf", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEventinfo
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
				return ErrInvalidLengthEventinfo
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthEventinfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HostIf = append(m.HostIf[:0], dAtA[iNdEx:postIndex]...)
			if m.HostIf == nil {
				m.HostIf = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Packet", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEventinfo
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
				return ErrInvalidLengthEventinfo
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthEventinfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Packet = append(m.Packet[:0], dAtA[iNdEx:postIndex]...)
			if m.Packet == nil {
				m.Packet = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEventinfo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEventinfo
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthEventinfo
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
func (m *NatPortblockEvent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEventinfo
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
			return fmt.Errorf("proto: NatPortblockEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NatPortblockEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEventinfo
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
				return ErrInvalidLengthEventinfo
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthEventinfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = append(m.Id[:0], dAtA[iNdEx:postIndex]...)
			if m.Id == nil {
				m.Id = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Protocol", wireType)
			}
			m.Protocol = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEventinfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Protocol |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SnatIP", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEventinfo
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
				return ErrInvalidLengthEventinfo
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEventinfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SnatIP == nil {
				m.SnatIP = &pds.IPAddress{}
			}
			if err := m.SnatIP.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AddressType", wireType)
			}
			m.AddressType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEventinfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AddressType |= pds.AddressType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEventinfo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEventinfo
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthEventinfo
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
func skipEventinfo(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEventinfo
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
					return 0, ErrIntOverflowEventinfo
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
					return 0, ErrIntOverflowEventinfo
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
				return 0, ErrInvalidLengthEventinfo
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthEventinfo
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowEventinfo
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
				next, err := skipEventinfo(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthEventinfo
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
	ErrInvalidLengthEventinfo = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEventinfo   = fmt.Errorf("proto: integer overflow")
)
