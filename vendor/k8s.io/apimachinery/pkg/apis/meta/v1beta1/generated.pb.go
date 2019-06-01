/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by protoc-gen-gogo.
// source: k8s.io/kubernetes/vendor/k8s.io/apimachinery/pkg/apis/meta/v1beta1/generated.proto
// DO NOT EDIT!

/*
	Package v1beta1 is a generated protocol buffer package.

	It is generated from these files:
		k8s.io/kubernetes/vendor/k8s.io/apimachinery/pkg/apis/meta/v1beta1/generated.proto

	It has these top-level messages:
		PartialObjectMetadata
		PartialObjectMetadataList
		TableOptions
*/
package v1beta1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import strings "strings"
import reflect "reflect"

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

func (m *PartialObjectMetadata) Reset()                    { *m = PartialObjectMetadata{} }
func (*PartialObjectMetadata) ProtoMessage()               {}
func (*PartialObjectMetadata) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{0} }

func (m *PartialObjectMetadataList) Reset()      { *m = PartialObjectMetadataList{} }
func (*PartialObjectMetadataList) ProtoMessage() {}
func (*PartialObjectMetadataList) Descriptor() ([]byte, []int) {
	return fileDescriptorGenerated, []int{1}
}

func (m *TableOptions) Reset()                    { *m = TableOptions{} }
func (*TableOptions) ProtoMessage()               {}
func (*TableOptions) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{2} }

func init() {
	proto.RegisterType((*PartialObjectMetadata)(nil), "k8s.io.apimachinery.pkg.apis.meta.v1beta1.PartialObjectMetadata")
	proto.RegisterType((*PartialObjectMetadataList)(nil), "k8s.io.apimachinery.pkg.apis.meta.v1beta1.PartialObjectMetadataList")
	proto.RegisterType((*TableOptions)(nil), "k8s.io.apimachinery.pkg.apis.meta.v1beta1.TableOptions")
}
func (m *PartialObjectMetadata) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PartialObjectMetadata) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.ObjectMeta.Size()))
	n1, err := m.ObjectMeta.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	return i, nil
}

func (m *PartialObjectMetadataList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PartialObjectMetadataList) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Items) > 0 {
		for _, msg := range m.Items {
			dAtA[i] = 0xa
			i++
			i = encodeVarintGenerated(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *TableOptions) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TableOptions) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.IncludeObject)))
	i += copy(dAtA[i:], m.IncludeObject)
	return i, nil
}

func encodeFixed64Generated(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Generated(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintGenerated(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *PartialObjectMetadata) Size() (n int) {
	var l int
	_ = l
	l = m.ObjectMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func (m *PartialObjectMetadataList) Size() (n int) {
	var l int
	_ = l
	if len(m.Items) > 0 {
		for _, e := range m.Items {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	return n
}

func (m *TableOptions) Size() (n int) {
	var l int
	_ = l
	l = len(m.IncludeObject)
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func sovGenerated(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozGenerated(x uint64) (n int) {
	return sovGenerated(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *PartialObjectMetadata) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&PartialObjectMetadata{`,
		`ObjectMeta:` + strings.Replace(strings.Replace(this.ObjectMeta.String(), "ObjectMeta", "k8s_io_apimachinery_pkg_apis_meta_v1.ObjectMeta", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *PartialObjectMetadataList) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&PartialObjectMetadataList{`,
		`Items:` + strings.Replace(fmt.Sprintf("%v", this.Items), "PartialObjectMetadata", "PartialObjectMetadata", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *TableOptions) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&TableOptions{`,
		`IncludeObject:` + fmt.Sprintf("%v", this.IncludeObject) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringGenerated(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *PartialObjectMetadata) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: PartialObjectMetadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PartialObjectMetadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObjectMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ObjectMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
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
func (m *PartialObjectMetadataList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: PartialObjectMetadataList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PartialObjectMetadataList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Items", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Items = append(m.Items, &PartialObjectMetadata{})
			if err := m.Items[len(m.Items)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
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
func (m *TableOptions) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: TableOptions: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TableOptions: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IncludeObject", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IncludeObject = IncludeObjectPolicy(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
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
func skipGenerated(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenerated
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
					return 0, ErrIntOverflowGenerated
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
					return 0, ErrIntOverflowGenerated
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
				return 0, ErrInvalidLengthGenerated
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowGenerated
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
				next, err := skipGenerated(dAtA[start:])
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
	ErrInvalidLengthGenerated = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenerated   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("k8s.io/kubernetes/vendor/k8s.io/apimachinery/pkg/apis/meta/v1beta1/generated.proto", fileDescriptorGenerated)
}

var fileDescriptorGenerated = []byte{
	// 391 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xbd, 0x6e, 0xd4, 0x40,
	0x10, 0xc7, 0xbd, 0x42, 0x11, 0x64, 0x43, 0x1a, 0x23, 0xa4, 0x70, 0xc5, 0x3a, 0xba, 0x2a, 0x48,
	0x64, 0x97, 0x04, 0x84, 0x28, 0x91, 0xbb, 0x48, 0xa0, 0x44, 0x16, 0x15, 0x15, 0x6b, 0x7b, 0xf0,
	0x2d, 0xb6, 0x77, 0xad, 0xdd, 0x71, 0xa4, 0x6b, 0x10, 0x8f, 0xc0, 0x63, 0x5d, 0x99, 0x32, 0x95,
	0xc5, 0x99, 0xb7, 0xa0, 0x42, 0xfe, 0x10, 0xf9, 0xb8, 0x3b, 0xe5, 0xba, 0x99, 0xff, 0xe8, 0xf7,
	0xf3, 0x8c, 0x97, 0x46, 0xf9, 0x7b, 0xc7, 0x95, 0x11, 0x79, 0x1d, 0x83, 0xd5, 0x80, 0xe0, 0xc4,
	0x25, 0xe8, 0xd4, 0x58, 0x31, 0x0e, 0x64, 0xa5, 0x4a, 0x99, 0xcc, 0x94, 0x06, 0x3b, 0x17, 0x55,
	0x9e, 0x75, 0x81, 0x13, 0x25, 0xa0, 0x14, 0x97, 0x27, 0x31, 0xa0, 0x3c, 0x11, 0x19, 0x68, 0xb0,
	0x12, 0x21, 0xe5, 0x95, 0x35, 0x68, 0xfc, 0x97, 0x03, 0xca, 0x6f, 0xa3, 0xbc, 0xca, 0xb3, 0x2e,
	0x70, 0xbc, 0x43, 0xf9, 0x88, 0x4e, 0x8e, 0x33, 0x85, 0xb3, 0x3a, 0xe6, 0x89, 0x29, 0x45, 0x66,
	0x32, 0x23, 0x7a, 0x43, 0x5c, 0x7f, 0xeb, 0xbb, 0xbe, 0xe9, 0xab, 0xc1, 0x3c, 0x79, 0xbb, 0xcd,
	0x52, 0xf7, 0xf7, 0x99, 0x6c, 0x3c, 0xc5, 0xd6, 0x1a, 0x55, 0x09, 0x2b, 0xc0, 0xbb, 0x87, 0x00,
	0x97, 0xcc, 0xa0, 0x94, 0x2b, 0xdc, 0x9b, 0x4d, 0x5c, 0x8d, 0xaa, 0x10, 0x4a, 0xa3, 0x43, 0x7b,
	0x1f, 0x9a, 0xce, 0xe9, 0xf3, 0x0b, 0x69, 0x51, 0xc9, 0xe2, 0x3c, 0xfe, 0x0e, 0x09, 0x7e, 0x02,
	0x94, 0xa9, 0x44, 0xe9, 0x7f, 0xa5, 0x4f, 0xca, 0xb1, 0x3e, 0x20, 0x87, 0xe4, 0x68, 0xef, 0xf4,
	0x35, 0xdf, 0xe6, 0xcf, 0xf2, 0x1b, 0x4f, 0xe8, 0x2f, 0x9a, 0xc0, 0x6b, 0x9b, 0x80, 0xde, 0x64,
	0xd1, 0x7f, 0xeb, 0xf4, 0x07, 0x7d, 0xb1, 0xf6, 0xd3, 0x1f, 0x95, 0x43, 0x5f, 0xd2, 0x1d, 0x85,
	0x50, 0xba, 0x03, 0x72, 0xf8, 0xe8, 0x68, 0xef, 0xf4, 0x03, 0xdf, 0xfa, 0x55, 0xf9, 0x5a, 0x69,
	0xb8, 0xdb, 0x36, 0xc1, 0xce, 0x59, 0xa7, 0x8c, 0x06, 0xf3, 0x34, 0xa6, 0x4f, 0x3f, 0xcb, 0xb8,
	0x80, 0xf3, 0x0a, 0x95, 0xd1, 0xce, 0x8f, 0xe8, 0xbe, 0xd2, 0x49, 0x51, 0xa7, 0x30, 0xa0, 0xfd,
	0xd9, 0xbb, 0xe1, 0xab, 0xf1, 0x88, 0xfd, 0xb3, 0xdb, 0xc3, 0xbf, 0x4d, 0xf0, 0xec, 0x4e, 0x70,
	0x61, 0x0a, 0x95, 0xcc, 0xa3, 0xbb, 0x8a, 0xf0, 0x78, 0xb1, 0x64, 0xde, 0xd5, 0x92, 0x79, 0xd7,
	0x4b, 0xe6, 0xfd, 0x6c, 0x19, 0x59, 0xb4, 0x8c, 0x5c, 0xb5, 0x8c, 0x5c, 0xb7, 0x8c, 0xfc, 0x6e,
	0x19, 0xf9, 0xf5, 0x87, 0x79, 0x5f, 0x1e, 0x8f, 0xab, 0xff, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x73,
	0xdf, 0x3a, 0x0c, 0x10, 0x03, 0x00, 0x00,
}
