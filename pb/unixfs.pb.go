// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: unixfs.proto

package unixfs_pb

import (
	fmt "fmt"
	math "math"

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

type Data_DataType int32

const (
	Data_Raw       Data_DataType = 0
	Data_Directory Data_DataType = 1
	Data_File      Data_DataType = 2
	Data_Metadata  Data_DataType = 3
	Data_Symlink   Data_DataType = 4
	Data_HAMTShard Data_DataType = 5
)

var Data_DataType_name = map[int32]string{
	0: "Raw",
	1: "Directory",
	2: "File",
	3: "Metadata",
	4: "Symlink",
	5: "HAMTShard",
}

var Data_DataType_value = map[string]int32{
	"Raw":       0,
	"Directory": 1,
	"File":      2,
	"Metadata":  3,
	"Symlink":   4,
	"HAMTShard": 5,
}

func (x Data_DataType) Enum() *Data_DataType {
	p := new(Data_DataType)
	*p = x
	return p
}

func (x Data_DataType) String() string {
	return proto.EnumName(Data_DataType_name, int32(x))
}

func (x *Data_DataType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Data_DataType_value, data, "Data_DataType")
	if err != nil {
		return err
	}
	*x = Data_DataType(value)
	return nil
}

func (Data_DataType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e2fd76cc44dfc7c3, []int{0, 0}
}

type Data struct {
	Type                 *Data_DataType `protobuf:"varint,1,req,name=Type,enum=unixfs.pb.Data_DataType" json:"Type,omitempty"`
	Data                 []byte         `protobuf:"bytes,2,opt,name=Data" json:"Data,omitempty"`
	Filesize             *uint64        `protobuf:"varint,3,opt,name=filesize" json:"filesize,omitempty"`
	Blocksizes           []uint64       `protobuf:"varint,4,rep,name=blocksizes" json:"blocksizes,omitempty"`
	HashType             *uint64        `protobuf:"varint,5,opt,name=hashType" json:"hashType,omitempty"`
	Fanout               *uint64        `protobuf:"varint,6,opt,name=fanout" json:"fanout,omitempty"`
	Mode                 *uint32        `protobuf:"varint,7,opt,name=mode" json:"mode,omitempty"`
	Mtime                *IPFSTimestamp `protobuf:"bytes,8,opt,name=mtime" json:"mtime,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Data) Reset()         { *m = Data{} }
func (m *Data) String() string { return proto.CompactTextString(m) }
func (*Data) ProtoMessage()    {}
func (*Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_e2fd76cc44dfc7c3, []int{0}
}
func (m *Data) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Data.Unmarshal(m, b)
}
func (m *Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Data.Marshal(b, m, deterministic)
}
func (m *Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Data.Merge(m, src)
}
func (m *Data) XXX_Size() int {
	return xxx_messageInfo_Data.Size(m)
}
func (m *Data) XXX_DiscardUnknown() {
	xxx_messageInfo_Data.DiscardUnknown(m)
}

var xxx_messageInfo_Data proto.InternalMessageInfo

func (m *Data) GetType() Data_DataType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return Data_Raw
}

func (m *Data) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Data) GetFilesize() uint64 {
	if m != nil && m.Filesize != nil {
		return *m.Filesize
	}
	return 0
}

func (m *Data) GetBlocksizes() []uint64 {
	if m != nil {
		return m.Blocksizes
	}
	return nil
}

func (m *Data) GetHashType() uint64 {
	if m != nil && m.HashType != nil {
		return *m.HashType
	}
	return 0
}

func (m *Data) GetFanout() uint64 {
	if m != nil && m.Fanout != nil {
		return *m.Fanout
	}
	return 0
}

func (m *Data) GetMode() uint32 {
	if m != nil && m.Mode != nil {
		return *m.Mode
	}
	return 0
}

func (m *Data) GetMtime() *IPFSTimestamp {
	if m != nil {
		return m.Mtime
	}
	return nil
}

type Metadata struct {
	MimeType             *string  `protobuf:"bytes,1,opt,name=MimeType" json:"MimeType,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Metadata) Reset()         { *m = Metadata{} }
func (m *Metadata) String() string { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()    {}
func (*Metadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_e2fd76cc44dfc7c3, []int{1}
}
func (m *Metadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Metadata.Unmarshal(m, b)
}
func (m *Metadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Metadata.Marshal(b, m, deterministic)
}
func (m *Metadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metadata.Merge(m, src)
}
func (m *Metadata) XXX_Size() int {
	return xxx_messageInfo_Metadata.Size(m)
}
func (m *Metadata) XXX_DiscardUnknown() {
	xxx_messageInfo_Metadata.DiscardUnknown(m)
}

var xxx_messageInfo_Metadata proto.InternalMessageInfo

func (m *Metadata) GetMimeType() string {
	if m != nil && m.MimeType != nil {
		return *m.MimeType
	}
	return ""
}

// mostly copied from proto 3 - with int32 nanos changed to fixed32 for js-ipfs compatibility
// https://github.com/protocolbuffers/protobuf/blob/master/src/google/protobuf/timestamp.proto
type IPFSTimestamp struct {
	// Represents seconds of UTC time since Unix epoch
	// 1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to
	// 9999-12-31T23:59:59Z inclusive.
	Seconds *int64 `protobuf:"varint,1,req,name=seconds" json:"seconds,omitempty"`
	// Non-negative fractions of a second at nanosecond resolution. Negative
	// second values with fractions must still have non-negative nanos values
	// that count forward in time. Must be from 0 to 999,999,999
	// inclusive.
	Nanos                *uint32  `protobuf:"fixed32,2,opt,name=nanos" json:"nanos,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IPFSTimestamp) Reset()         { *m = IPFSTimestamp{} }
func (m *IPFSTimestamp) String() string { return proto.CompactTextString(m) }
func (*IPFSTimestamp) ProtoMessage()    {}
func (*IPFSTimestamp) Descriptor() ([]byte, []int) {
	return fileDescriptor_e2fd76cc44dfc7c3, []int{2}
}
func (m *IPFSTimestamp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IPFSTimestamp.Unmarshal(m, b)
}
func (m *IPFSTimestamp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IPFSTimestamp.Marshal(b, m, deterministic)
}
func (m *IPFSTimestamp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IPFSTimestamp.Merge(m, src)
}
func (m *IPFSTimestamp) XXX_Size() int {
	return xxx_messageInfo_IPFSTimestamp.Size(m)
}
func (m *IPFSTimestamp) XXX_DiscardUnknown() {
	xxx_messageInfo_IPFSTimestamp.DiscardUnknown(m)
}

var xxx_messageInfo_IPFSTimestamp proto.InternalMessageInfo

func (m *IPFSTimestamp) GetSeconds() int64 {
	if m != nil && m.Seconds != nil {
		return *m.Seconds
	}
	return 0
}

func (m *IPFSTimestamp) GetNanos() uint32 {
	if m != nil && m.Nanos != nil {
		return *m.Nanos
	}
	return 0
}

func init() {
	proto.RegisterEnum("unixfs.pb.Data_DataType", Data_DataType_name, Data_DataType_value)
	proto.RegisterType((*Data)(nil), "unixfs.pb.Data")
	proto.RegisterType((*Metadata)(nil), "unixfs.pb.Metadata")
	proto.RegisterType((*IPFSTimestamp)(nil), "unixfs.pb.IPFSTimestamp")
}

func init() { proto.RegisterFile("unixfs.proto", fileDescriptor_e2fd76cc44dfc7c3) }

var fileDescriptor_e2fd76cc44dfc7c3 = []byte{
	// 322 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0x5f, 0x4b, 0xc3, 0x30,
	0x14, 0xc5, 0xed, 0xbf, 0xb5, 0xbb, 0xdb, 0xa4, 0x5c, 0x44, 0x82, 0x0f, 0x52, 0xfa, 0x20, 0x7d,
	0x90, 0x3e, 0xf8, 0x05, 0x44, 0x18, 0x43, 0x1f, 0x06, 0x92, 0x0d, 0xdf, 0xb3, 0x35, 0x63, 0x61,
	0x4d, 0x33, 0x9a, 0x0c, 0x9d, 0x9f, 0xd3, 0x0f, 0x24, 0x49, 0xd7, 0xe9, 0x5e, 0x4a, 0x7f, 0xb9,
	0xe7, 0x84, 0x73, 0x6e, 0x60, 0x7c, 0x68, 0xc4, 0xd7, 0x46, 0x97, 0xfb, 0x56, 0x19, 0x85, 0xc3,
	0x9e, 0x56, 0xf9, 0x8f, 0x0f, 0xe1, 0x94, 0x19, 0x86, 0x8f, 0x10, 0x2e, 0x8f, 0x7b, 0x4e, 0xbc,
	0xcc, 0x2f, 0xae, 0x9f, 0x48, 0x79, 0x96, 0x94, 0x76, 0xec, 0x3e, 0x76, 0x4e, 0x9d, 0x0a, 0xb1,
	0x73, 0x11, 0x3f, 0xf3, 0x8a, 0x31, 0xed, 0x6e, 0xb8, 0x83, 0x64, 0x23, 0x6a, 0xae, 0xc5, 0x37,
	0x27, 0x41, 0xe6, 0x15, 0x21, 0x3d, 0x33, 0xde, 0x03, 0xac, 0x6a, 0xb5, 0xde, 0x59, 0xd0, 0x24,
	0xcc, 0x82, 0x22, 0xa4, 0xff, 0x4e, 0xac, 0x77, 0xcb, 0xf4, 0xd6, 0x25, 0x88, 0x3a, 0x6f, 0xcf,
	0x78, 0x0b, 0x83, 0x0d, 0x6b, 0xd4, 0xc1, 0x90, 0x81, 0x9b, 0x9c, 0xc8, 0x66, 0x90, 0xaa, 0xe2,
	0x24, 0xce, 0xbc, 0x62, 0x42, 0xdd, 0x3f, 0x96, 0x10, 0x49, 0x23, 0x24, 0x27, 0x49, 0xe6, 0x15,
	0xa3, 0x8b, 0x1a, 0x6f, 0xef, 0xb3, 0xc5, 0x52, 0x48, 0xae, 0x0d, 0x93, 0x7b, 0xda, 0xc9, 0xf2,
	0x0f, 0x48, 0xfa, 0x66, 0x18, 0x43, 0x40, 0xd9, 0x67, 0x7a, 0x85, 0x13, 0x18, 0x4e, 0x45, 0xcb,
	0xd7, 0x46, 0xb5, 0xc7, 0xd4, 0xc3, 0x04, 0xc2, 0x99, 0xa8, 0x79, 0xea, 0xe3, 0x18, 0x92, 0x39,
	0x37, 0xac, 0x62, 0x86, 0xa5, 0x01, 0x8e, 0x20, 0x5e, 0x1c, 0x65, 0x2d, 0x9a, 0x5d, 0x1a, 0x5a,
	0xcf, 0xeb, 0xcb, 0x7c, 0xb9, 0xd8, 0xb2, 0xb6, 0x4a, 0xa3, 0xfc, 0xe1, 0x4f, 0x69, 0xbb, 0xcd,
	0x85, 0xe4, 0xa7, 0xed, 0x7a, 0xc5, 0x90, 0x9e, 0x39, 0x7f, 0x86, 0xc9, 0x45, 0x2e, 0x24, 0x10,
	0x6b, 0xbe, 0x56, 0x4d, 0xa5, 0xdd, 0x4b, 0x04, 0xb4, 0x47, 0xbc, 0x81, 0xa8, 0x61, 0x8d, 0xd2,
	0x6e, 0xe7, 0x31, 0xed, 0xe0, 0x37, 0x00, 0x00, 0xff, 0xff, 0x36, 0xaf, 0xfa, 0x7c, 0xd9, 0x01,
	0x00, 0x00,
}
