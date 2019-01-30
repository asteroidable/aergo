// Code generated by protoc-gen-go. DO NOT EDIT.
// source: metric.proto

package types // import "github.com/aergoio/aergo/types"

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type MetricType int32

const (
	// NOTHING should not be used.
	MetricType_NOTHING MetricType = 0
	// Metric for p2p network transfer
	MetricType_P2P_NETWORK MetricType = 1
)

var MetricType_name = map[int32]string{
	0: "NOTHING",
	1: "P2P_NETWORK",
}
var MetricType_value = map[string]int32{
	"NOTHING":     0,
	"P2P_NETWORK": 1,
}

func (x MetricType) String() string {
	return proto.EnumName(MetricType_name, int32(x))
}
func (MetricType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_metric_295631cc212b5411, []int{0}
}

type MetricsRequest struct {
	Types                []MetricType `protobuf:"varint,1,rep,packed,name=types,enum=types.MetricType" json:"types,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *MetricsRequest) Reset()         { *m = MetricsRequest{} }
func (m *MetricsRequest) String() string { return proto.CompactTextString(m) }
func (*MetricsRequest) ProtoMessage()    {}
func (*MetricsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_metric_295631cc212b5411, []int{0}
}

func (m *MetricsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetricsRequest.Unmarshal(m, b)
}
func (m *MetricsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetricsRequest.Marshal(b, m, deterministic)
}
func (dst *MetricsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetricsRequest.Merge(dst, src)
}
func (m *MetricsRequest) XXX_Size() int {
	return xxx_messageInfo_MetricsRequest.Size(m)
}
func (m *MetricsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MetricsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MetricsRequest proto.InternalMessageInfo

func (m *MetricsRequest) GetTypes() []MetricType {
	if m != nil {
		return m.Types
	}
	return nil
}

type Metrics struct {
	Peers                []*PeerMetric `protobuf:"bytes,1,rep,name=peers" json:"peers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Metrics) Reset()         { *m = Metrics{} }
func (m *Metrics) String() string { return proto.CompactTextString(m) }
func (*Metrics) ProtoMessage()    {}
func (*Metrics) Descriptor() ([]byte, []int) {
	return fileDescriptor_metric_295631cc212b5411, []int{1}
}

func (m *Metrics) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Metrics.Unmarshal(m, b)
}
func (m *Metrics) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Metrics.Marshal(b, m, deterministic)
}
func (dst *Metrics) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metrics.Merge(dst, src)
}
func (m *Metrics) XXX_Size() int {
	return xxx_messageInfo_Metrics.Size(m)
}
func (m *Metrics) XXX_DiscardUnknown() {
	xxx_messageInfo_Metrics.DiscardUnknown(m)
}

var xxx_messageInfo_Metrics proto.InternalMessageInfo

func (m *Metrics) GetPeers() []*PeerMetric {
	if m != nil {
		return m.Peers
	}
	return nil
}

type PeerMetric struct {
	PeerID               []byte   `protobuf:"bytes,1,opt,name=peerID,proto3" json:"peerID,omitempty"`
	SumIn                int64    `protobuf:"varint,2,opt,name=sumIn" json:"sumIn,omitempty"`
	AvrIn                int64    `protobuf:"varint,3,opt,name=avrIn" json:"avrIn,omitempty"`
	SumOut               int64    `protobuf:"varint,4,opt,name=sumOut" json:"sumOut,omitempty"`
	AvrOut               int64    `protobuf:"varint,5,opt,name=avrOut" json:"avrOut,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PeerMetric) Reset()         { *m = PeerMetric{} }
func (m *PeerMetric) String() string { return proto.CompactTextString(m) }
func (*PeerMetric) ProtoMessage()    {}
func (*PeerMetric) Descriptor() ([]byte, []int) {
	return fileDescriptor_metric_295631cc212b5411, []int{2}
}

func (m *PeerMetric) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PeerMetric.Unmarshal(m, b)
}
func (m *PeerMetric) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PeerMetric.Marshal(b, m, deterministic)
}
func (dst *PeerMetric) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PeerMetric.Merge(dst, src)
}
func (m *PeerMetric) XXX_Size() int {
	return xxx_messageInfo_PeerMetric.Size(m)
}
func (m *PeerMetric) XXX_DiscardUnknown() {
	xxx_messageInfo_PeerMetric.DiscardUnknown(m)
}

var xxx_messageInfo_PeerMetric proto.InternalMessageInfo

func (m *PeerMetric) GetPeerID() []byte {
	if m != nil {
		return m.PeerID
	}
	return nil
}

func (m *PeerMetric) GetSumIn() int64 {
	if m != nil {
		return m.SumIn
	}
	return 0
}

func (m *PeerMetric) GetAvrIn() int64 {
	if m != nil {
		return m.AvrIn
	}
	return 0
}

func (m *PeerMetric) GetSumOut() int64 {
	if m != nil {
		return m.SumOut
	}
	return 0
}

func (m *PeerMetric) GetAvrOut() int64 {
	if m != nil {
		return m.AvrOut
	}
	return 0
}

func init() {
	proto.RegisterEnum("types.MetricType", MetricType_name, MetricType_value)
	proto.RegisterType((*MetricsRequest)(nil), "types.MetricsRequest")
	proto.RegisterType((*Metrics)(nil), "types.Metrics")
	proto.RegisterType((*PeerMetric)(nil), "types.PeerMetric")
}

func init() { proto.RegisterFile("metric.proto", fileDescriptor_metric_295631cc212b5411) }

var fileDescriptor_metric_295631cc212b5411 = []byte{
	// 251 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x90, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0x40, 0x5d, 0x63, 0x5a, 0x98, 0x94, 0x5a, 0x17, 0x91, 0x9c, 0x24, 0xf4, 0x62, 0xe8, 0x21,
	0x85, 0x78, 0xf2, 0x2a, 0x8a, 0x06, 0x31, 0x09, 0x4b, 0x40, 0xf0, 0x22, 0x69, 0x19, 0x6a, 0x0e,
	0xe9, 0xc6, 0xfd, 0x28, 0xf4, 0xe6, 0x4f, 0x97, 0xdd, 0x09, 0xf6, 0x14, 0xde, 0x9b, 0x79, 0x81,
	0x59, 0x98, 0xf5, 0x68, 0x54, 0xb7, 0xcd, 0x06, 0x25, 0x8d, 0xe4, 0xa1, 0x39, 0x0e, 0xa8, 0x97,
	0x0f, 0x30, 0x7f, 0xf7, 0x5a, 0x0b, 0xfc, 0xb1, 0xa8, 0x0d, 0xbf, 0x03, 0x1a, 0xc5, 0x2c, 0x09,
	0xd2, 0x79, 0x7e, 0x95, 0x79, 0xca, 0x68, 0xab, 0x39, 0x0e, 0x28, 0xc6, 0x34, 0x87, 0xe9, 0x98,
	0xba, 0x66, 0x40, 0x54, 0xd4, 0x44, 0xff, 0x4d, 0x8d, 0xa8, 0x68, 0x45, 0xd0, 0x7c, 0xf9, 0xcb,
	0x00, 0x4e, 0x96, 0xdf, 0xc0, 0xc4, 0xf9, 0xe2, 0x29, 0x66, 0x09, 0x4b, 0x67, 0x62, 0x24, 0x7e,
	0x0d, 0xa1, 0xb6, 0x7d, 0xb1, 0x8f, 0xcf, 0x13, 0x96, 0x06, 0x82, 0xc0, 0xd9, 0xf6, 0xa0, 0x8a,
	0x7d, 0x1c, 0x90, 0xf5, 0xe0, 0xfe, 0xa1, 0x6d, 0x5f, 0x59, 0x13, 0x5f, 0x78, 0x3d, 0x92, 0xf3,
	0xed, 0x41, 0x39, 0x1f, 0x92, 0x27, 0x5a, 0xad, 0x00, 0x4e, 0xb7, 0xf0, 0x08, 0xa6, 0x65, 0xd5,
	0xbc, 0x16, 0xe5, 0xcb, 0xe2, 0x8c, 0x5f, 0x42, 0x54, 0xe7, 0xf5, 0x57, 0xf9, 0xdc, 0x7c, 0x54,
	0xe2, 0x6d, 0xc1, 0x1e, 0x93, 0xcf, 0xdb, 0x5d, 0x67, 0xbe, 0xed, 0x26, 0xdb, 0xca, 0x7e, 0xdd,
	0xa2, 0xda, 0xc9, 0x4e, 0xd2, 0x77, 0xed, 0x4f, 0xdc, 0x4c, 0xfc, 0x6b, 0xde, 0xff, 0x05, 0x00,
	0x00, 0xff, 0xff, 0xb5, 0xeb, 0x41, 0xa0, 0x5d, 0x01, 0x00, 0x00,
}
