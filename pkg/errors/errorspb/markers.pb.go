// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/cockroachdb/cockroach/pkg/errors/errorspb/markers.proto

package errorspb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type MarkPayload struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	Types                []string `protobuf:"bytes,2,rep,name=types,proto3" json:"types,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MarkPayload) Reset()         { *m = MarkPayload{} }
func (m *MarkPayload) String() string { return proto.CompactTextString(m) }
func (*MarkPayload) ProtoMessage()    {}
func (*MarkPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_markers_e4318b485555d408, []int{0}
}
func (m *MarkPayload) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MarkPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalTo(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (dst *MarkPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MarkPayload.Merge(dst, src)
}
func (m *MarkPayload) XXX_Size() int {
	return m.Size()
}
func (m *MarkPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_MarkPayload.DiscardUnknown(m)
}

var xxx_messageInfo_MarkPayload proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MarkPayload)(nil), "cockroach.errorspb.MarkPayload")
}
func (m *MarkPayload) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MarkPayload) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Msg) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMarkers(dAtA, i, uint64(len(m.Msg)))
		i += copy(dAtA[i:], m.Msg)
	}
	if len(m.Types) > 0 {
		for _, s := range m.Types {
			dAtA[i] = 0x12
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	return i, nil
}

func encodeVarintMarkers(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *MarkPayload) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Msg)
	if l > 0 {
		n += 1 + l + sovMarkers(uint64(l))
	}
	if len(m.Types) > 0 {
		for _, s := range m.Types {
			l = len(s)
			n += 1 + l + sovMarkers(uint64(l))
		}
	}
	return n
}

func sovMarkers(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozMarkers(x uint64) (n int) {
	return sovMarkers(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MarkPayload) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMarkers
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MarkPayload: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MarkPayload: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msg", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMarkers
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMarkers
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Msg = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Types", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMarkers
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMarkers
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Types = append(m.Types, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMarkers(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMarkers
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
func skipMarkers(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMarkers
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
					return 0, ErrIntOverflowMarkers
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
					return 0, ErrIntOverflowMarkers
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthMarkers
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowMarkers
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
				next, err := skipMarkers(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
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
	ErrInvalidLengthMarkers = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMarkers   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("github.com/cockroachdb/cockroach/pkg/errors/errorspb/markers.proto", fileDescriptor_markers_e4318b485555d408)
}

var fileDescriptor_markers_e4318b485555d408 = []byte{
	// 150 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4d, 0x2d, 0x2a, 0xca,
	0x2f, 0x2a, 0xd6, 0x87, 0x50, 0x05, 0x49, 0xfa, 0xb9, 0x89, 0x45, 0xd9, 0xa9, 0x45, 0xc5, 0x7a,
	0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x42, 0xc9, 0xf9, 0xc9, 0xd9, 0x45, 0xf9, 0x89, 0xc9, 0x19,
	0x7a, 0x30, 0x15, 0x4a, 0xa6, 0x5c, 0xdc, 0xbe, 0x89, 0x45, 0xd9, 0x01, 0x89, 0x95, 0x39, 0xf9,
	0x89, 0x29, 0x42, 0x02, 0x5c, 0xcc, 0xb9, 0xc5, 0xe9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41,
	0x20, 0xa6, 0x90, 0x08, 0x17, 0x6b, 0x49, 0x65, 0x41, 0x6a, 0xb1, 0x04, 0x93, 0x02, 0xb3, 0x06,
	0x67, 0x10, 0x84, 0xe3, 0xa4, 0x74, 0xe2, 0xa1, 0x1c, 0xc3, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e,
	0xc9, 0x31, 0xde, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x51,
	0x1c, 0x30, 0xa3, 0x93, 0xd8, 0xc0, 0xb6, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x68, 0xdb,
	0xee, 0x90, 0x96, 0x00, 0x00, 0x00,
}
