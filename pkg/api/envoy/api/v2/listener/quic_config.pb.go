// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/api/v2/listener/quic_config.proto

package envoy_api_v2_listener

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
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

// Configuration specific to the QUIC protocol.
// Next id: 4
type QuicProtocolOptions struct {
	// Maximum number of streams that the client can negotiate per connection. 100
	// if not specified.
	MaxConcurrentStreams *types.UInt32Value `protobuf:"bytes,1,opt,name=max_concurrent_streams,json=maxConcurrentStreams,proto3" json:"max_concurrent_streams,omitempty"`
	// Maximum number of milliseconds that connection will be alive when there is
	// no network activity. 300000ms if not specified.
	IdleTimeout *types.Duration `protobuf:"bytes,2,opt,name=idle_timeout,json=idleTimeout,proto3" json:"idle_timeout,omitempty"`
	// Connection timeout in milliseconds before the crypto handshake is finished.
	// 20000ms if not specified.
	CryptoHandshakeTimeout *types.Duration `protobuf:"bytes,3,opt,name=crypto_handshake_timeout,json=cryptoHandshakeTimeout,proto3" json:"crypto_handshake_timeout,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}        `json:"-"`
	XXX_unrecognized       []byte          `json:"-"`
	XXX_sizecache          int32           `json:"-"`
}

func (m *QuicProtocolOptions) Reset()         { *m = QuicProtocolOptions{} }
func (m *QuicProtocolOptions) String() string { return proto.CompactTextString(m) }
func (*QuicProtocolOptions) ProtoMessage()    {}
func (*QuicProtocolOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f6a4a402e708e40, []int{0}
}
func (m *QuicProtocolOptions) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QuicProtocolOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QuicProtocolOptions.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QuicProtocolOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuicProtocolOptions.Merge(m, src)
}
func (m *QuicProtocolOptions) XXX_Size() int {
	return m.Size()
}
func (m *QuicProtocolOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_QuicProtocolOptions.DiscardUnknown(m)
}

var xxx_messageInfo_QuicProtocolOptions proto.InternalMessageInfo

func (m *QuicProtocolOptions) GetMaxConcurrentStreams() *types.UInt32Value {
	if m != nil {
		return m.MaxConcurrentStreams
	}
	return nil
}

func (m *QuicProtocolOptions) GetIdleTimeout() *types.Duration {
	if m != nil {
		return m.IdleTimeout
	}
	return nil
}

func (m *QuicProtocolOptions) GetCryptoHandshakeTimeout() *types.Duration {
	if m != nil {
		return m.CryptoHandshakeTimeout
	}
	return nil
}

func init() {
	proto.RegisterType((*QuicProtocolOptions)(nil), "envoy.api.v2.listener.QuicProtocolOptions")
}

func init() {
	proto.RegisterFile("envoy/api/v2/listener/quic_config.proto", fileDescriptor_1f6a4a402e708e40)
}

var fileDescriptor_1f6a4a402e708e40 = []byte{
	// 385 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xb1, 0x4e, 0xdb, 0x40,
	0x18, 0xc7, 0x7b, 0xa9, 0x9a, 0xc1, 0xa9, 0xd4, 0xca, 0x6d, 0x53, 0x37, 0x6a, 0xad, 0xaa, 0x1d,
	0x60, 0xba, 0x93, 0x9c, 0x95, 0x85, 0x04, 0x24, 0x90, 0x10, 0x84, 0x04, 0xb2, 0x5a, 0x17, 0xfb,
	0xe2, 0x9c, 0xb0, 0xef, 0x8e, 0xf3, 0x9d, 0x49, 0x36, 0xde, 0x80, 0x95, 0x91, 0x99, 0x47, 0xe0,
	0x09, 0x32, 0xc2, 0x1b, 0xa0, 0x8c, 0x8c, 0xcc, 0x08, 0x21, 0x9f, 0x6d, 0x18, 0x02, 0x62, 0x3c,
	0xfd, 0x7f, 0xbf, 0xbf, 0xee, 0xfb, 0x3e, 0x6b, 0x85, 0xb0, 0x8c, 0xcf, 0x10, 0x16, 0x14, 0x65,
	0x1e, 0x8a, 0x69, 0xaa, 0x08, 0x23, 0x12, 0x1d, 0x6b, 0x1a, 0xf8, 0x01, 0x67, 0x63, 0x1a, 0x41,
	0x21, 0xb9, 0xe2, 0xf6, 0x0f, 0x03, 0x42, 0x2c, 0x28, 0xcc, 0x3c, 0x58, 0x81, 0x2d, 0x37, 0xe2,
	0x3c, 0x8a, 0x09, 0x32, 0xd0, 0x48, 0x8f, 0x51, 0xa8, 0x25, 0x56, 0x94, 0xb3, 0x42, 0x5b, 0xce,
	0x4f, 0x24, 0x16, 0x82, 0xc8, 0xb4, 0xca, 0x75, 0x28, 0x30, 0xc2, 0x8c, 0x71, 0x65, 0xb4, 0x14,
	0x25, 0x34, 0x92, 0x58, 0x91, 0x32, 0xff, 0xb3, 0x94, 0xa7, 0x0a, 0x2b, 0x5d, 0xea, 0xff, 0x1e,
	0x80, 0xf5, 0x6d, 0x5f, 0xd3, 0xa0, 0x97, 0xbf, 0x02, 0x1e, 0xef, 0x09, 0x03, 0xd9, 0x7d, 0xab,
	0x99, 0xe0, 0x69, 0x3e, 0x41, 0xa0, 0xa5, 0x24, 0x4c, 0xf9, 0xa9, 0x92, 0x04, 0x27, 0xa9, 0x03,
	0xfe, 0x82, 0xd5, 0x86, 0xf7, 0x1b, 0x16, 0xff, 0x82, 0xd5, 0xbf, 0xe0, 0xe1, 0x36, 0x53, 0x6d,
	0x6f, 0x88, 0x63, 0x4d, 0xfa, 0xdf, 0x13, 0x3c, 0xed, 0x3e, 0xab, 0x83, 0xc2, 0xb4, 0xd7, 0xac,
	0xcf, 0x34, 0x8c, 0x89, 0xaf, 0x68, 0x42, 0xb8, 0x56, 0x4e, 0xcd, 0x34, 0xfd, 0x5a, 0x6a, 0xda,
	0x28, 0x37, 0xd0, 0x6f, 0xe4, 0xf8, 0x41, 0x41, 0xdb, 0x03, 0xcb, 0x09, 0xe4, 0x4c, 0x28, 0xee,
	0x4f, 0x30, 0x0b, 0xd3, 0x09, 0x3e, 0x7a, 0x69, 0xfa, 0xf8, 0x5e, 0x53, 0xb3, 0x50, 0xb7, 0x2a,
	0xb3, 0x2c, 0xed, 0x5c, 0x80, 0xf9, 0xc2, 0x05, 0xd7, 0x0b, 0x17, 0xdc, 0x2e, 0x5c, 0x70, 0x7f,
	0xfe, 0x78, 0xf6, 0xa9, 0x65, 0x3b, 0xc5, 0xa5, 0xca, 0xeb, 0x55, 0x97, 0x82, 0x59, 0xfb, 0xea,
	0x74, 0x7e, 0x53, 0xaf, 0x7d, 0xfd, 0x60, 0xfd, 0xa7, 0x1c, 0x1a, 0x48, 0x48, 0x3e, 0x9d, 0xc1,
	0x57, 0x2f, 0xdb, 0xf9, 0x92, 0xef, 0xb5, 0x6b, 0x4a, 0xcc, 0x76, 0x7b, 0xe0, 0xb2, 0xf6, 0x73,
	0xd3, 0xa0, 0xeb, 0x82, 0xc2, 0xa1, 0x07, 0x77, 0x4a, 0x74, 0x77, 0x70, 0xf7, 0x66, 0x32, 0xaa,
	0x9b, 0x69, 0xda, 0x4f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xf9, 0x80, 0x76, 0x86, 0x69, 0x02, 0x00,
	0x00,
}

func (m *QuicProtocolOptions) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QuicProtocolOptions) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QuicProtocolOptions) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.CryptoHandshakeTimeout != nil {
		{
			size, err := m.CryptoHandshakeTimeout.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuicConfig(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.IdleTimeout != nil {
		{
			size, err := m.IdleTimeout.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuicConfig(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.MaxConcurrentStreams != nil {
		{
			size, err := m.MaxConcurrentStreams.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuicConfig(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuicConfig(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuicConfig(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QuicProtocolOptions) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MaxConcurrentStreams != nil {
		l = m.MaxConcurrentStreams.Size()
		n += 1 + l + sovQuicConfig(uint64(l))
	}
	if m.IdleTimeout != nil {
		l = m.IdleTimeout.Size()
		n += 1 + l + sovQuicConfig(uint64(l))
	}
	if m.CryptoHandshakeTimeout != nil {
		l = m.CryptoHandshakeTimeout.Size()
		n += 1 + l + sovQuicConfig(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovQuicConfig(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuicConfig(x uint64) (n int) {
	return sovQuicConfig(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QuicProtocolOptions) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuicConfig
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
			return fmt.Errorf("proto: QuicProtocolOptions: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuicProtocolOptions: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxConcurrentStreams", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuicConfig
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
				return ErrInvalidLengthQuicConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuicConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MaxConcurrentStreams == nil {
				m.MaxConcurrentStreams = &types.UInt32Value{}
			}
			if err := m.MaxConcurrentStreams.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IdleTimeout", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuicConfig
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
				return ErrInvalidLengthQuicConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuicConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.IdleTimeout == nil {
				m.IdleTimeout = &types.Duration{}
			}
			if err := m.IdleTimeout.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CryptoHandshakeTimeout", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuicConfig
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
				return ErrInvalidLengthQuicConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuicConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CryptoHandshakeTimeout == nil {
				m.CryptoHandshakeTimeout = &types.Duration{}
			}
			if err := m.CryptoHandshakeTimeout.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuicConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuicConfig
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQuicConfig
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
func skipQuicConfig(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuicConfig
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
					return 0, ErrIntOverflowQuicConfig
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
					return 0, ErrIntOverflowQuicConfig
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
				return 0, ErrInvalidLengthQuicConfig
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuicConfig
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuicConfig
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuicConfig        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuicConfig          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuicConfig = fmt.Errorf("proto: unexpected end of group")
)
