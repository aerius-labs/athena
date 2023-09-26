// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: athena/qcrescent/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type MsgMsgSendQueryAllBalances struct {
	Creator    string             `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	ChannelId  string             `protobuf:"bytes,2,opt,name=channelId,proto3" json:"channelId,omitempty"`
	Address    string             `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	Pagination *query.PageRequest `protobuf:"bytes,4,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *MsgMsgSendQueryAllBalances) Reset()         { *m = MsgMsgSendQueryAllBalances{} }
func (m *MsgMsgSendQueryAllBalances) String() string { return proto.CompactTextString(m) }
func (*MsgMsgSendQueryAllBalances) ProtoMessage()    {}
func (*MsgMsgSendQueryAllBalances) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d362d6dcba2c18e, []int{0}
}
func (m *MsgMsgSendQueryAllBalances) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgMsgSendQueryAllBalances) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgMsgSendQueryAllBalances.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgMsgSendQueryAllBalances) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgMsgSendQueryAllBalances.Merge(m, src)
}
func (m *MsgMsgSendQueryAllBalances) XXX_Size() int {
	return m.Size()
}
func (m *MsgMsgSendQueryAllBalances) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgMsgSendQueryAllBalances.DiscardUnknown(m)
}

var xxx_messageInfo_MsgMsgSendQueryAllBalances proto.InternalMessageInfo

func (m *MsgMsgSendQueryAllBalances) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgMsgSendQueryAllBalances) GetChannelId() string {
	if m != nil {
		return m.ChannelId
	}
	return ""
}

func (m *MsgMsgSendQueryAllBalances) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *MsgMsgSendQueryAllBalances) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type MsgMsgSendQueryAllBalancesResponse struct {
	Sequence uint64 `protobuf:"varint,1,opt,name=sequence,proto3" json:"sequence,omitempty"`
}

func (m *MsgMsgSendQueryAllBalancesResponse) Reset()         { *m = MsgMsgSendQueryAllBalancesResponse{} }
func (m *MsgMsgSendQueryAllBalancesResponse) String() string { return proto.CompactTextString(m) }
func (*MsgMsgSendQueryAllBalancesResponse) ProtoMessage()    {}
func (*MsgMsgSendQueryAllBalancesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d362d6dcba2c18e, []int{1}
}
func (m *MsgMsgSendQueryAllBalancesResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgMsgSendQueryAllBalancesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgMsgSendQueryAllBalancesResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgMsgSendQueryAllBalancesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgMsgSendQueryAllBalancesResponse.Merge(m, src)
}
func (m *MsgMsgSendQueryAllBalancesResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgMsgSendQueryAllBalancesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgMsgSendQueryAllBalancesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgMsgSendQueryAllBalancesResponse proto.InternalMessageInfo

func (m *MsgMsgSendQueryAllBalancesResponse) GetSequence() uint64 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func init() {
	proto.RegisterType((*MsgMsgSendQueryAllBalances)(nil), "athena.qcrescent.MsgMsgSendQueryAllBalances")
	proto.RegisterType((*MsgMsgSendQueryAllBalancesResponse)(nil), "athena.qcrescent.MsgMsgSendQueryAllBalancesResponse")
}

func init() { proto.RegisterFile("athena/qcrescent/tx.proto", fileDescriptor_6d362d6dcba2c18e) }

var fileDescriptor_6d362d6dcba2c18e = []byte{
	// 338 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xbf, 0x4e, 0xfb, 0x30,
	0x10, 0xc7, 0xeb, 0x5f, 0xab, 0x1f, 0xd4, 0x2c, 0x28, 0x0b, 0x21, 0x42, 0x51, 0xd5, 0x01, 0x55,
	0x08, 0x6c, 0xb5, 0xf0, 0x00, 0xd0, 0x01, 0xa9, 0x43, 0x25, 0x08, 0x1b, 0x9b, 0xe3, 0x9c, 0xd2,
	0x48, 0xa9, 0x9d, 0xfa, 0x1c, 0xd4, 0x0e, 0x6c, 0x3c, 0x00, 0x4f, 0xc3, 0x33, 0x30, 0x76, 0x64,
	0x44, 0xed, 0x8b, 0xa0, 0x24, 0xfd, 0x27, 0xa4, 0x0e, 0x8c, 0xe7, 0xef, 0x7d, 0xe4, 0xcf, 0xd9,
	0x47, 0x4f, 0x85, 0x1d, 0x81, 0x12, 0x7c, 0x22, 0x0d, 0xa0, 0x04, 0x65, 0xb9, 0x9d, 0xb2, 0xcc,
	0x68, 0xab, 0x9d, 0xe3, 0x2a, 0x62, 0x9b, 0xc8, 0xbb, 0x90, 0x1a, 0xc7, 0x1a, 0x79, 0x28, 0x10,
	0xf8, 0x24, 0x07, 0x33, 0xe3, 0x2f, 0xdd, 0x10, 0xac, 0xe8, 0xf2, 0x4c, 0xc4, 0x89, 0x12, 0x36,
	0xd1, 0xaa, 0xa2, 0xdb, 0x1f, 0x84, 0x7a, 0x43, 0x8c, 0x87, 0x18, 0x3f, 0x81, 0x8a, 0x1e, 0x8b,
	0xe6, 0xbb, 0x34, 0xed, 0x8b, 0x54, 0x28, 0x09, 0xe8, 0xb8, 0xf4, 0x40, 0x1a, 0x10, 0x56, 0x1b,
	0x97, 0xb4, 0x48, 0xa7, 0x19, 0xac, 0x4b, 0xe7, 0x8c, 0x36, 0xe5, 0x48, 0x28, 0x05, 0xe9, 0x20,
	0x72, 0xff, 0x95, 0xd9, 0xf6, 0xa0, 0xe0, 0x44, 0x14, 0x19, 0x40, 0x74, 0xeb, 0x15, 0xb7, 0x2a,
	0x9d, 0x7b, 0x4a, 0xb7, 0x12, 0x6e, 0xa3, 0x45, 0x3a, 0x47, 0xbd, 0x73, 0x56, 0x19, 0xb3, 0xc2,
	0x98, 0x95, 0xc6, 0x6c, 0x65, 0xcc, 0x1e, 0x44, 0x0c, 0x01, 0x4c, 0x72, 0x40, 0x1b, 0xec, 0x90,
	0xed, 0x5b, 0xda, 0xde, 0xef, 0x1d, 0x00, 0x66, 0x5a, 0x21, 0x38, 0x1e, 0x3d, 0xc4, 0x02, 0x56,
	0x12, 0xca, 0x01, 0x1a, 0xc1, 0xa6, 0xee, 0xbd, 0x11, 0x5a, 0x1f, 0x62, 0xec, 0xbc, 0xd2, 0x93,
	0x7d, 0xe3, 0x5f, 0xb2, 0xdf, 0x8f, 0xcb, 0xf6, 0x5f, 0xea, 0xdd, 0xfc, 0xa5, 0x7b, 0xad, 0xd8,
	0x1f, 0x7c, 0x2e, 0x7c, 0x32, 0x5f, 0xf8, 0xe4, 0x7b, 0xe1, 0x93, 0xf7, 0xa5, 0x5f, 0x9b, 0x2f,
	0xfd, 0xda, 0xd7, 0xd2, 0xaf, 0x3d, 0xf3, 0x38, 0xb1, 0xa3, 0x3c, 0x64, 0x52, 0x8f, 0xb9, 0x00,
	0x93, 0xe4, 0x78, 0x95, 0x8a, 0x10, 0xf9, 0x6a, 0x17, 0xa6, 0xbb, 0xdb, 0x30, 0xcb, 0x00, 0xc3,
	0xff, 0xe5, 0x9f, 0x5e, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x31, 0x58, 0x18, 0xc1, 0x2e, 0x02,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	MsgSendQueryAllBalances(ctx context.Context, in *MsgMsgSendQueryAllBalances, opts ...grpc.CallOption) (*MsgMsgSendQueryAllBalancesResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) MsgSendQueryAllBalances(ctx context.Context, in *MsgMsgSendQueryAllBalances, opts ...grpc.CallOption) (*MsgMsgSendQueryAllBalancesResponse, error) {
	out := new(MsgMsgSendQueryAllBalancesResponse)
	err := c.cc.Invoke(ctx, "/athena.qcrescent.Msg/MsgSendQueryAllBalances", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	MsgSendQueryAllBalances(context.Context, *MsgMsgSendQueryAllBalances) (*MsgMsgSendQueryAllBalancesResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) MsgSendQueryAllBalances(ctx context.Context, req *MsgMsgSendQueryAllBalances) (*MsgMsgSendQueryAllBalancesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MsgSendQueryAllBalances not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_MsgSendQueryAllBalances_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgMsgSendQueryAllBalances)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).MsgSendQueryAllBalances(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/athena.qcrescent.Msg/MsgSendQueryAllBalances",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).MsgSendQueryAllBalances(ctx, req.(*MsgMsgSendQueryAllBalances))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "athena.qcrescent.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MsgSendQueryAllBalances",
			Handler:    _Msg_MsgSendQueryAllBalances_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "athena/qcrescent/tx.proto",
}

func (m *MsgMsgSendQueryAllBalances) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgMsgSendQueryAllBalances) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgMsgSendQueryAllBalances) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ChannelId) > 0 {
		i -= len(m.ChannelId)
		copy(dAtA[i:], m.ChannelId)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ChannelId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgMsgSendQueryAllBalancesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgMsgSendQueryAllBalancesResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgMsgSendQueryAllBalancesResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Sequence != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Sequence))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgMsgSendQueryAllBalances) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.ChannelId)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgMsgSendQueryAllBalancesResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Sequence != 0 {
		n += 1 + sovTx(uint64(m.Sequence))
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgMsgSendQueryAllBalances) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgMsgSendQueryAllBalances: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgMsgSendQueryAllBalances: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChannelId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChannelId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgMsgSendQueryAllBalancesResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgMsgSendQueryAllBalancesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgMsgSendQueryAllBalancesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sequence", wireType)
			}
			m.Sequence = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Sequence |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
