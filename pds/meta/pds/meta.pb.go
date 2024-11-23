// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: meta.proto

package pds

import (
	fmt "fmt"
	gogoproto "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	io "io"
	math "math"
	math_bits "math/bits"
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

var E_GoprotoEnumPrefix = gogoproto.E_GoprotoEnumPrefix

var E_GoprotoEnumStringer = gogoproto.E_GoprotoEnumStringer

var E_EnumStringer = gogoproto.E_EnumStringer

var E_EnumCustomname = gogoproto.E_EnumCustomname

var E_Enumdecl = gogoproto.E_Enumdecl

var E_EnumvalueCustomname = gogoproto.E_EnumvalueCustomname

var E_GoprotoGettersAll = gogoproto.E_GoprotoGettersAll

var E_GoprotoEnumPrefixAll = gogoproto.E_GoprotoEnumPrefixAll

var E_GoprotoStringerAll = gogoproto.E_GoprotoStringerAll

var E_VerboseEqualAll = gogoproto.E_VerboseEqualAll

var E_FaceAll = gogoproto.E_FaceAll

var E_GostringAll = gogoproto.E_GostringAll

var E_PopulateAll = gogoproto.E_PopulateAll

var E_StringerAll = gogoproto.E_StringerAll

var E_OnlyoneAll = gogoproto.E_OnlyoneAll

var E_EqualAll = gogoproto.E_EqualAll

var E_DescriptionAll = gogoproto.E_DescriptionAll

var E_TestgenAll = gogoproto.E_TestgenAll

var E_BenchgenAll = gogoproto.E_BenchgenAll

var E_MarshalerAll = gogoproto.E_MarshalerAll

var E_UnmarshalerAll = gogoproto.E_UnmarshalerAll

var E_StableMarshalerAll = gogoproto.E_StableMarshalerAll

var E_SizerAll = gogoproto.E_SizerAll

var E_GoprotoEnumStringerAll = gogoproto.E_GoprotoEnumStringerAll

var E_EnumStringerAll = gogoproto.E_EnumStringerAll

var E_UnsafeMarshalerAll = gogoproto.E_UnsafeMarshalerAll

var E_UnsafeUnmarshalerAll = gogoproto.E_UnsafeUnmarshalerAll

var E_GoprotoExtensionsMapAll = gogoproto.E_GoprotoExtensionsMapAll

var E_GoprotoUnrecognizedAll = gogoproto.E_GoprotoUnrecognizedAll

var E_GogoprotoImport = gogoproto.E_GogoprotoImport

var E_ProtosizerAll = gogoproto.E_ProtosizerAll

var E_CompareAll = gogoproto.E_CompareAll

var E_TypedeclAll = gogoproto.E_TypedeclAll

var E_EnumdeclAll = gogoproto.E_EnumdeclAll

var E_GoprotoRegistration = gogoproto.E_GoprotoRegistration

var E_MessagenameAll = gogoproto.E_MessagenameAll

var E_GoprotoSizecacheAll = gogoproto.E_GoprotoSizecacheAll

var E_GoprotoUnkeyedAll = gogoproto.E_GoprotoUnkeyedAll

var E_GoprotoGetters = gogoproto.E_GoprotoGetters

var E_GoprotoStringer = gogoproto.E_GoprotoStringer

var E_VerboseEqual = gogoproto.E_VerboseEqual

var E_Face = gogoproto.E_Face

var E_Gostring = gogoproto.E_Gostring

var E_Populate = gogoproto.E_Populate

var E_Stringer = gogoproto.E_Stringer

var E_Onlyone = gogoproto.E_Onlyone

var E_Equal = gogoproto.E_Equal

var E_Description = gogoproto.E_Description

var E_Testgen = gogoproto.E_Testgen

var E_Benchgen = gogoproto.E_Benchgen

var E_Marshaler = gogoproto.E_Marshaler

var E_Unmarshaler = gogoproto.E_Unmarshaler

var E_StableMarshaler = gogoproto.E_StableMarshaler

var E_Sizer = gogoproto.E_Sizer

var E_UnsafeMarshaler = gogoproto.E_UnsafeMarshaler

var E_UnsafeUnmarshaler = gogoproto.E_UnsafeUnmarshaler

var E_GoprotoExtensionsMap = gogoproto.E_GoprotoExtensionsMap

var E_GoprotoUnrecognized = gogoproto.E_GoprotoUnrecognized

var E_Protosizer = gogoproto.E_Protosizer

var E_Compare = gogoproto.E_Compare

var E_Typedecl = gogoproto.E_Typedecl

var E_Messagename = gogoproto.E_Messagename

var E_GoprotoSizecache = gogoproto.E_GoprotoSizecache

var E_GoprotoUnkeyed = gogoproto.E_GoprotoUnkeyed

var E_Nullable = gogoproto.E_Nullable

var E_Embed = gogoproto.E_Embed

var E_Customtype = gogoproto.E_Customtype

var E_Customname = gogoproto.E_Customname

var E_Jsontag = gogoproto.E_Jsontag

var E_Moretags = gogoproto.E_Moretags

var E_Casttype = gogoproto.E_Casttype

var E_Castkey = gogoproto.E_Castkey

var E_Castvalue = gogoproto.E_Castvalue

var E_Stdtime = gogoproto.E_Stdtime

var E_Stdduration = gogoproto.E_Stdduration

var E_Wktpointer = gogoproto.E_Wktpointer

// TypeMeta contains the metadata about kind and version for all API objects
type TypeMeta struct {
	// Kind represents the type of the API object.
	Kind                 string   `protobuf:"bytes,1,opt,name=Kind,proto3" json:"kind"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TypeMeta) Reset()         { *m = TypeMeta{} }
func (m *TypeMeta) String() string { return proto.CompactTextString(m) }
func (*TypeMeta) ProtoMessage()    {}
func (*TypeMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b5ea8fe65782bcc, []int{0}
}
func (m *TypeMeta) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TypeMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TypeMeta.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TypeMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TypeMeta.Merge(m, src)
}
func (m *TypeMeta) XXX_Size() int {
	return m.Size()
}
func (m *TypeMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_TypeMeta.DiscardUnknown(m)
}

var xxx_messageInfo_TypeMeta proto.InternalMessageInfo

func (m *TypeMeta) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

type Timestamp struct {
	*types.Timestamp     `protobuf:"bytes,1,opt,name=Time,proto3,embedded=Time" json:"Time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Timestamp) Reset()         { *m = Timestamp{} }
func (m *Timestamp) String() string { return proto.CompactTextString(m) }
func (*Timestamp) ProtoMessage()    {}
func (*Timestamp) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b5ea8fe65782bcc, []int{1}
}
func (m *Timestamp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Timestamp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Timestamp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Timestamp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Timestamp.Merge(m, src)
}
func (m *Timestamp) XXX_Size() int {
	return m.Size()
}
func (m *Timestamp) XXX_DiscardUnknown() {
	xxx_messageInfo_Timestamp.DiscardUnknown(m)
}

var xxx_messageInfo_Timestamp proto.InternalMessageInfo

// ObjectMeta contains metadata that all objects stored in kvstore must have
type ObjMeta struct {
	// Name of the object, unique within a Namespace for scoped objects.
	// Should start and end in an alphanumeric character and can contain alphanumeric or ._-: characters
	// minimum length is 2 and maximum length is 64 characters
	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"name"`
	// UUID is the unique identifier for the object. This is generated on creation of the object.
	// System generated, not updatable by user.
	UUID string `protobuf:"bytes,2,opt,name=UUID,proto3" json:"uuid,omitempty"`
	// CreationTime is the creation time of the object
	//  System generated and updated, not updatable by user.
	CreationTime *Timestamp `protobuf:"bytes,3,opt,name=CreationTime,proto3" json:"creation-time,omitempty"`
	// ModTime is the Last Modification time of the object
	// System generated and updated, not updatable by user.
	ModTime              *Timestamp `protobuf:"bytes,4,opt,name=ModTime,proto3" json:"mod-time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ObjMeta) Reset()         { *m = ObjMeta{} }
func (m *ObjMeta) String() string { return proto.CompactTextString(m) }
func (*ObjMeta) ProtoMessage()    {}
func (*ObjMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b5ea8fe65782bcc, []int{2}
}
func (m *ObjMeta) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ObjMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ObjMeta.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ObjMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ObjMeta.Merge(m, src)
}
func (m *ObjMeta) XXX_Size() int {
	return m.Size()
}
func (m *ObjMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_ObjMeta.DiscardUnknown(m)
}

var xxx_messageInfo_ObjMeta proto.InternalMessageInfo

func (m *ObjMeta) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ObjMeta) GetUUID() string {
	if m != nil {
		return m.UUID
	}
	return ""
}

func (m *ObjMeta) GetCreationTime() *Timestamp {
	if m != nil {
		return m.CreationTime
	}
	return nil
}

func (m *ObjMeta) GetModTime() *Timestamp {
	if m != nil {
		return m.ModTime
	}
	return nil
}

func init() {
	proto.RegisterType((*TypeMeta)(nil), "meta.TypeMeta")
	proto.RegisterType((*Timestamp)(nil), "meta.Timestamp")
	proto.RegisterType((*ObjMeta)(nil), "meta.ObjMeta")
}

func init() { proto.RegisterFile("meta.proto", fileDescriptor_3b5ea8fe65782bcc) }

var fileDescriptor_3b5ea8fe65782bcc = []byte{
	// 299 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xbd, 0x4e, 0xc3, 0x30,
	0x14, 0x85, 0x31, 0x58, 0xb4, 0x35, 0x08, 0x24, 0x0f, 0xa5, 0x2a, 0x28, 0x46, 0x19, 0x50, 0x07,
	0x48, 0x24, 0xe0, 0x05, 0x1a, 0x58, 0x10, 0x2a, 0x54, 0x51, 0xbb, 0xb0, 0x39, 0xd8, 0x44, 0x06,
	0x1c, 0x47, 0xad, 0x33, 0xf4, 0x0d, 0x19, 0xfb, 0x04, 0x11, 0x8a, 0x98, 0xf2, 0x14, 0x28, 0x37,
	0x09, 0x7f, 0x62, 0xb3, 0xef, 0x3d, 0xf7, 0x3b, 0x47, 0x87, 0x10, 0x2d, 0x2d, 0xf7, 0xd2, 0x85,
	0xb1, 0x86, 0xe2, 0xea, 0x3d, 0x24, 0xb1, 0x89, 0x4d, 0x3d, 0x19, 0xb2, 0xd8, 0x98, 0xf8, 0x55,
	0xfa, 0xf0, 0x8b, 0xb2, 0x27, 0xdf, 0x2a, 0x2d, 0x97, 0x96, 0xeb, 0xb4, 0x16, 0xb8, 0x23, 0xd2,
	0x9d, 0xad, 0x52, 0x39, 0x91, 0x96, 0xd3, 0x23, 0x82, 0x6f, 0x55, 0x22, 0x06, 0xe8, 0x18, 0x8d,
	0x7a, 0x41, 0xb7, 0xcc, 0x19, 0x7e, 0x51, 0x89, 0x08, 0x61, 0xea, 0x8e, 0x49, 0x6f, 0xd6, 0x1e,
	0xd3, 0x4b, 0x82, 0xab, 0x0f, 0x48, 0x77, 0xce, 0x87, 0x5e, 0x6d, 0xe3, 0xb5, 0x36, 0xde, 0x97,
	0x32, 0xc0, 0xeb, 0x9c, 0xa1, 0x10, 0xd4, 0xee, 0x07, 0x22, 0x9d, 0xfb, 0xe8, 0xb9, 0x35, 0xbb,
	0xe3, 0x0d, 0xa1, 0x31, 0x4b, 0xb8, 0x96, 0x21, 0x4c, 0xe9, 0x09, 0xc1, 0xf3, 0xf9, 0xcd, 0xf5,
	0x60, 0x13, 0xb6, 0xb4, 0xcc, 0xd9, 0x5e, 0x96, 0x29, 0x71, 0x6a, 0xb4, 0xb2, 0x52, 0xa7, 0x76,
	0x15, 0xc2, 0x9e, 0x4e, 0xc9, 0xee, 0xd5, 0x42, 0x72, 0xab, 0x4c, 0x02, 0x79, 0xb6, 0x20, 0xcf,
	0xbe, 0x07, 0xa5, 0x7c, 0x87, 0x38, 0x2c, 0x73, 0x76, 0xf0, 0xd8, 0x08, 0xcf, 0xaa, 0x0e, 0x7e,
	0x90, 0x7e, 0x11, 0xe8, 0x98, 0x74, 0x26, 0x46, 0x00, 0x0c, 0xff, 0x0f, 0xeb, 0x97, 0x39, 0xa3,
	0xda, 0x88, 0xbf, 0x9c, 0xf6, 0x2e, 0xe8, 0xbf, 0x15, 0x0e, 0x5a, 0x17, 0x0e, 0x7a, 0x2f, 0x1c,
	0xf4, 0xd0, 0xad, 0xce, 0xfd, 0x54, 0x2c, 0xa7, 0x1b, 0xd1, 0x36, 0x14, 0x74, 0xf1, 0x19, 0x00,
	0x00, 0xff, 0xff, 0x47, 0x3d, 0x7f, 0x5a, 0xb5, 0x01, 0x00, 0x00,
}

func (m *TypeMeta) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TypeMeta) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TypeMeta) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Kind) > 0 {
		i -= len(m.Kind)
		copy(dAtA[i:], m.Kind)
		i = encodeVarintMeta(dAtA, i, uint64(len(m.Kind)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Timestamp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Timestamp) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Timestamp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Timestamp != nil {
		{
			size, err := m.Timestamp.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMeta(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ObjMeta) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ObjMeta) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ObjMeta) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.ModTime != nil {
		{
			size, err := m.ModTime.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMeta(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.CreationTime != nil {
		{
			size, err := m.CreationTime.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMeta(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.UUID) > 0 {
		i -= len(m.UUID)
		copy(dAtA[i:], m.UUID)
		i = encodeVarintMeta(dAtA, i, uint64(len(m.UUID)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintMeta(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMeta(dAtA []byte, offset int, v uint64) int {
	offset -= sovMeta(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TypeMeta) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Kind)
	if l > 0 {
		n += 1 + l + sovMeta(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Timestamp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Timestamp != nil {
		l = m.Timestamp.Size()
		n += 1 + l + sovMeta(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ObjMeta) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovMeta(uint64(l))
	}
	l = len(m.UUID)
	if l > 0 {
		n += 1 + l + sovMeta(uint64(l))
	}
	if m.CreationTime != nil {
		l = m.CreationTime.Size()
		n += 1 + l + sovMeta(uint64(l))
	}
	if m.ModTime != nil {
		l = m.ModTime.Size()
		n += 1 + l + sovMeta(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovMeta(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMeta(x uint64) (n int) {
	return sovMeta(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TypeMeta) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMeta
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
			return fmt.Errorf("proto: TypeMeta: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TypeMeta: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Kind", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMeta
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMeta
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMeta
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Kind = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMeta(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMeta
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMeta
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
func (m *Timestamp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMeta
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
			return fmt.Errorf("proto: Timestamp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Timestamp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMeta
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
				return ErrInvalidLengthMeta
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMeta
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Timestamp == nil {
				m.Timestamp = &types.Timestamp{}
			}
			if err := m.Timestamp.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMeta(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMeta
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMeta
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
func (m *ObjMeta) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMeta
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
			return fmt.Errorf("proto: ObjMeta: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ObjMeta: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMeta
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMeta
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMeta
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UUID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMeta
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMeta
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMeta
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UUID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreationTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMeta
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
				return ErrInvalidLengthMeta
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMeta
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CreationTime == nil {
				m.CreationTime = &Timestamp{}
			}
			if err := m.CreationTime.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ModTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMeta
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
				return ErrInvalidLengthMeta
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMeta
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ModTime == nil {
				m.ModTime = &Timestamp{}
			}
			if err := m.ModTime.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMeta(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMeta
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMeta
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
func skipMeta(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMeta
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
					return 0, ErrIntOverflowMeta
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
					return 0, ErrIntOverflowMeta
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
				return 0, ErrInvalidLengthMeta
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthMeta
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowMeta
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
				next, err := skipMeta(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthMeta
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
	ErrInvalidLengthMeta = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMeta   = fmt.Errorf("proto: integer overflow")
)
