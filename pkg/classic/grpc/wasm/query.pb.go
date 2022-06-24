// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: terra/wasm/v1beta1/query.proto

package wasm

import (
	context "context"
	encoding_json "encoding/json"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// QueryContractInfoRequest is the request type for the Query/ContractInfo RPC method.
type QueryContractInfoRequest struct {
	ContractAddress string `protobuf:"bytes,1,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
}

func (m *QueryContractInfoRequest) Reset()         { *m = QueryContractInfoRequest{} }
func (m *QueryContractInfoRequest) String() string { return proto.CompactTextString(m) }
func (*QueryContractInfoRequest) ProtoMessage()    {}
func (*QueryContractInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7601576355e80c46, []int{0}
}
func (m *QueryContractInfoRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryContractInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryContractInfoRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryContractInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryContractInfoRequest.Merge(m, src)
}
func (m *QueryContractInfoRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryContractInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryContractInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryContractInfoRequest proto.InternalMessageInfo

// QueryContractStoreRequest is the request type for the Query/ContractStore RPC method.
type QueryContractStoreRequest struct {
	ContractAddress string                   `protobuf:"bytes,1,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
	QueryMsg        encoding_json.RawMessage `protobuf:"bytes,2,opt,name=query_msg,json=queryMsg,proto3,casttype=encoding/json.RawMessage" json:"query_msg,omitempty"`
}

func (m *QueryContractStoreRequest) Reset()         { *m = QueryContractStoreRequest{} }
func (m *QueryContractStoreRequest) String() string { return proto.CompactTextString(m) }
func (*QueryContractStoreRequest) ProtoMessage()    {}
func (*QueryContractStoreRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7601576355e80c46, []int{1}
}
func (m *QueryContractStoreRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryContractStoreRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryContractStoreRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryContractStoreRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryContractStoreRequest.Merge(m, src)
}
func (m *QueryContractStoreRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryContractStoreRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryContractStoreRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryContractStoreRequest proto.InternalMessageInfo

// QueryContractStoreResponse is response type for the
// Query/ContractStore RPC method.
type QueryContractStoreResponse struct {
	QueryResult encoding_json.RawMessage `protobuf:"bytes,1,opt,name=query_result,json=queryResult,proto3,casttype=encoding/json.RawMessage" json:"query_result,omitempty"`
}

func (m *QueryContractStoreResponse) Reset()         { *m = QueryContractStoreResponse{} }
func (m *QueryContractStoreResponse) String() string { return proto.CompactTextString(m) }
func (*QueryContractStoreResponse) ProtoMessage()    {}
func (*QueryContractStoreResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7601576355e80c46, []int{2}
}
func (m *QueryContractStoreResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryContractStoreResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryContractStoreResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryContractStoreResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryContractStoreResponse.Merge(m, src)
}
func (m *QueryContractStoreResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryContractStoreResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryContractStoreResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryContractStoreResponse proto.InternalMessageInfo

func (m *QueryContractStoreResponse) GetQueryResult() encoding_json.RawMessage {
	if m != nil {
		return m.QueryResult
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryContractInfoRequest)(nil), "terra.wasm.v1beta1.QueryContractInfoRequest")
	proto.RegisterType((*QueryContractStoreRequest)(nil), "terra.wasm.v1beta1.QueryContractStoreRequest")
	proto.RegisterType((*QueryContractStoreResponse)(nil), "terra.wasm.v1beta1.QueryContractStoreResponse")
}

func init() { proto.RegisterFile("terra/wasm/v1beta1/query.proto", fileDescriptor_7601576355e80c46) }

var fileDescriptor_7601576355e80c46 = []byte{
	// 419 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x31, 0x6b, 0x14, 0x41,
	0x14, 0xc7, 0x77, 0x02, 0x4a, 0x32, 0x46, 0x94, 0xc1, 0xe2, 0x3c, 0xe2, 0x5c, 0xb8, 0xea, 0x2c,
	0x32, 0x43, 0x14, 0x44, 0x2d, 0x14, 0xcf, 0xca, 0x22, 0x88, 0x6b, 0x27, 0x48, 0x98, 0xdb, 0x7b,
	0x8e, 0x2b, 0xd9, 0x79, 0x9b, 0x79, 0xb3, 0xc6, 0x43, 0x6c, 0xac, 0x04, 0x1b, 0xc1, 0x2f, 0x90,
	0x4f, 0x60, 0xed, 0x47, 0xb0, 0x0c, 0xd8, 0x58, 0x89, 0xdc, 0x59, 0xf8, 0x19, 0xac, 0x64, 0x67,
	0x2e, 0x62, 0xf4, 0x44, 0xd3, 0x2d, 0xef, 0xf7, 0xe6, 0xff, 0xff, 0xf3, 0xdf, 0xc7, 0x65, 0x00,
	0xef, 0x8d, 0xde, 0x33, 0x54, 0xe9, 0xa7, 0x9b, 0x23, 0x08, 0x66, 0x53, 0xef, 0x36, 0xe0, 0x27,
	0xaa, 0xf6, 0x18, 0x50, 0x88, 0xc8, 0x55, 0xcb, 0xd5, 0x9c, 0x77, 0xcf, 0x59, 0xb4, 0x18, 0xb1,
	0x6e, 0xbf, 0xd2, 0x66, 0x77, 0xcd, 0x22, 0xda, 0x1d, 0xd0, 0xa6, 0x2e, 0xb5, 0x71, 0x0e, 0x83,
	0x09, 0x25, 0x3a, 0x9a, 0xd3, 0x0b, 0x0b, 0x7c, 0xa2, 0x68, 0xc2, 0xb2, 0x40, 0xaa, 0x90, 0xf4,
	0xc8, 0x10, 0xfc, 0xe4, 0x05, 0x96, 0x2e, 0xf1, 0xfe, 0x5d, 0xde, 0xb9, 0xd7, 0xa6, 0xba, 0x8d,
	0x2e, 0x78, 0x53, 0x84, 0x3b, 0xee, 0x11, 0xe6, 0xb0, 0xdb, 0x00, 0x05, 0x71, 0x91, 0x9f, 0x2d,
	0xe6, 0xe3, 0x6d, 0x33, 0x1e, 0x7b, 0x20, 0xea, 0xb0, 0x75, 0x36, 0x58, 0xc9, 0xcf, 0x1c, 0xce,
	0x6f, 0xa5, 0xf1, 0xf5, 0xe5, 0x57, 0xfb, 0xbd, 0xec, 0xdb, 0x7e, 0x2f, 0xeb, 0xbf, 0x66, 0xfc,
	0xfc, 0x11, 0xc5, 0xfb, 0x01, 0x3d, 0x1c, 0x5f, 0x52, 0x5c, 0xe3, 0x2b, 0xb1, 0xaf, 0xed, 0x8a,
	0x6c, 0x67, 0x69, 0x9d, 0x0d, 0x56, 0x87, 0x6b, 0xdf, 0x3f, 0xf7, 0x3a, 0xe0, 0x0a, 0x1c, 0x97,
	0xce, 0xea, 0x27, 0x84, 0x4e, 0xe5, 0x66, 0x6f, 0x0b, 0x88, 0x8c, 0x85, 0x7c, 0x39, 0xae, 0x6f,
	0x91, 0xfd, 0x25, 0xcd, 0x43, 0xde, 0x5d, 0x14, 0x86, 0x6a, 0x74, 0x04, 0xe2, 0x26, 0x5f, 0x4d,
	0x16, 0x1e, 0xa8, 0xd9, 0x09, 0x31, 0xc9, 0xbf, 0x5c, 0x4e, 0xc5, 0x17, 0x79, 0x7c, 0x70, 0xe9,
	0x3d, 0xe3, 0x27, 0xa2, 0xbe, 0x78, 0xc7, 0xf8, 0xe9, 0x23, 0x26, 0x62, 0x43, 0xfd, 0xf9, 0x87,
	0xd5, 0x5f, 0x9b, 0xe9, 0xaa, 0xff, 0x5d, 0x4f, 0xd9, 0xfb, 0x37, 0x5e, 0x7e, 0xfc, 0xfa, 0x76,
	0xe9, 0xaa, 0xb8, 0xa2, 0x17, 0x1c, 0xc0, 0x61, 0x97, 0xa4, 0x9f, 0xff, 0x5e, 0xf7, 0x0b, 0x4d,
	0xad, 0xce, 0x70, 0xf8, 0x61, 0x2a, 0xd9, 0xc1, 0x54, 0xb2, 0x2f, 0x53, 0xc9, 0xde, 0xcc, 0x64,
	0x76, 0x30, 0x93, 0xd9, 0xa7, 0x99, 0xcc, 0x1e, 0x0c, 0x6c, 0x19, 0x1e, 0x37, 0x23, 0x55, 0x60,
	0x95, 0xb4, 0x37, 0x2a, 0x74, 0x30, 0xd1, 0x05, 0x7a, 0xd0, 0xcf, 0x92, 0x51, 0x98, 0xd4, 0x40,
	0xa3, 0x93, 0xf1, 0x86, 0x2e, 0xff, 0x08, 0x00, 0x00, 0xff, 0xff, 0xd8, 0x82, 0x30, 0x1a, 0xec,
	0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// ContractStore return smart query result from the contract
	ContractStore(ctx context.Context, in *QueryContractStoreRequest, opts ...grpc.CallOption) (*QueryContractStoreResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) ContractStore(ctx context.Context, in *QueryContractStoreRequest, opts ...grpc.CallOption) (*QueryContractStoreResponse, error) {
	out := new(QueryContractStoreResponse)
	err := c.cc.Invoke(ctx, "/terra.wasm.v1beta1.Query/ContractStore", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// ContractStore return smart query result from the contract
	ContractStore(context.Context, *QueryContractStoreRequest) (*QueryContractStoreResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) ContractStore(ctx context.Context, req *QueryContractStoreRequest) (*QueryContractStoreResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ContractStore not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_ContractStore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryContractStoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ContractStore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/terra.wasm.v1beta1.Query/ContractStore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ContractStore(ctx, req.(*QueryContractStoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "terra.wasm.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ContractStore",
			Handler:    _Query_ContractStore_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "terra/wasm/v1beta1/query.proto",
}

func (m *QueryContractInfoRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryContractInfoRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryContractInfoRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ContractAddress) > 0 {
		i -= len(m.ContractAddress)
		copy(dAtA[i:], m.ContractAddress)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.ContractAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryContractStoreRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryContractStoreRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryContractStoreRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.QueryMsg) > 0 {
		i -= len(m.QueryMsg)
		copy(dAtA[i:], m.QueryMsg)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.QueryMsg)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ContractAddress) > 0 {
		i -= len(m.ContractAddress)
		copy(dAtA[i:], m.ContractAddress)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.ContractAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryContractStoreResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryContractStoreResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryContractStoreResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.QueryResult) > 0 {
		i -= len(m.QueryResult)
		copy(dAtA[i:], m.QueryResult)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.QueryResult)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryContractInfoRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ContractAddress)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryContractStoreRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ContractAddress)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.QueryMsg)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryContractStoreResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.QueryResult)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryContractInfoRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryContractInfoRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryContractInfoRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryContractStoreRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryContractStoreRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryContractStoreRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field QueryMsg", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.QueryMsg = append(m.QueryMsg[:0], dAtA[iNdEx:postIndex]...)
			if m.QueryMsg == nil {
				m.QueryMsg = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryContractStoreResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryContractStoreResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryContractStoreResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field QueryResult", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.QueryResult = append(m.QueryResult[:0], dAtA[iNdEx:postIndex]...)
			if m.QueryResult == nil {
				m.QueryResult = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)