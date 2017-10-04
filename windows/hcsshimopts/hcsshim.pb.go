// Code generated by protoc-gen-gogo.
// source: github.com/containerd/containerd/windows/hcsshimopts/hcsshim.proto
// DO NOT EDIT!

/*
	Package hcsshimopts is a generated protocol buffer package.

	It is generated from these files:
		github.com/containerd/containerd/windows/hcsshimopts/hcsshim.proto

	It has these top-level messages:
		CreateOptions
		ProcessDetails
*/
package hcsshimopts

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "github.com/gogo/protobuf/types"
import google_protobuf2 "github.com/gogo/protobuf/types"

import time "time"

import github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type CreateOptions struct {
	TerminateDuration time.Duration `protobuf:"bytes,1,opt,name=terminate_duration,json=terminateDuration,stdduration" json:"terminate_duration"`
}

func (m *CreateOptions) Reset()                    { *m = CreateOptions{} }
func (*CreateOptions) ProtoMessage()               {}
func (*CreateOptions) Descriptor() ([]byte, []int) { return fileDescriptorHcsshim, []int{0} }

// ProcessDetails contains additional information about a process
// ProcessDetails is made of the same fields as found in hcsshim.ProcessListItem
type ProcessDetails struct {
	ImageName                    string                      `protobuf:"bytes,1,opt,name=image_name,json=imageName,proto3" json:"image_name,omitempty"`
	CreatedAt                    *google_protobuf2.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	KernelTime_100Ns             uint64                      `protobuf:"varint,3,opt,name=kernel_time_100_ns,json=kernelTime100Ns,proto3" json:"kernel_time_100_ns,omitempty"`
	MemoryCommitBytes            uint64                      `protobuf:"varint,4,opt,name=memory_commit_bytes,json=memoryCommitBytes,proto3" json:"memory_commit_bytes,omitempty"`
	MemoryWorkingSetPrivateBytes uint64                      `protobuf:"varint,5,opt,name=memory_working_set_private_bytes,json=memoryWorkingSetPrivateBytes,proto3" json:"memory_working_set_private_bytes,omitempty"`
	MemoryWorkingSetSharedBytes  uint64                      `protobuf:"varint,6,opt,name=memory_working_set_shared_bytes,json=memoryWorkingSetSharedBytes,proto3" json:"memory_working_set_shared_bytes,omitempty"`
	ProcessID                    uint32                      `protobuf:"varint,7,opt,name=process_id,json=processId,proto3" json:"process_id,omitempty"`
	UserTime_100Ns               uint64                      `protobuf:"varint,8,opt,name=user_time_100_ns,json=userTime100Ns,proto3" json:"user_time_100_ns,omitempty"`
}

func (m *ProcessDetails) Reset()                    { *m = ProcessDetails{} }
func (*ProcessDetails) ProtoMessage()               {}
func (*ProcessDetails) Descriptor() ([]byte, []int) { return fileDescriptorHcsshim, []int{1} }

func init() {
	proto.RegisterType((*CreateOptions)(nil), "containerd.windows.hcsshim.CreateOptions")
	proto.RegisterType((*ProcessDetails)(nil), "containerd.windows.hcsshim.ProcessDetails")
}
func (m *CreateOptions) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CreateOptions) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintHcsshim(dAtA, i, uint64(github_com_gogo_protobuf_types.SizeOfStdDuration(m.TerminateDuration)))
	n1, err := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.TerminateDuration, dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	return i, nil
}

func (m *ProcessDetails) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProcessDetails) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ImageName) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintHcsshim(dAtA, i, uint64(len(m.ImageName)))
		i += copy(dAtA[i:], m.ImageName)
	}
	if m.CreatedAt != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintHcsshim(dAtA, i, uint64(m.CreatedAt.Size()))
		n2, err := m.CreatedAt.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.KernelTime_100Ns != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintHcsshim(dAtA, i, uint64(m.KernelTime_100Ns))
	}
	if m.MemoryCommitBytes != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintHcsshim(dAtA, i, uint64(m.MemoryCommitBytes))
	}
	if m.MemoryWorkingSetPrivateBytes != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintHcsshim(dAtA, i, uint64(m.MemoryWorkingSetPrivateBytes))
	}
	if m.MemoryWorkingSetSharedBytes != 0 {
		dAtA[i] = 0x30
		i++
		i = encodeVarintHcsshim(dAtA, i, uint64(m.MemoryWorkingSetSharedBytes))
	}
	if m.ProcessID != 0 {
		dAtA[i] = 0x38
		i++
		i = encodeVarintHcsshim(dAtA, i, uint64(m.ProcessID))
	}
	if m.UserTime_100Ns != 0 {
		dAtA[i] = 0x40
		i++
		i = encodeVarintHcsshim(dAtA, i, uint64(m.UserTime_100Ns))
	}
	return i, nil
}

func encodeFixed64Hcsshim(dAtA []byte, offset int, v uint64) int {
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
func encodeFixed32Hcsshim(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintHcsshim(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *CreateOptions) Size() (n int) {
	var l int
	_ = l
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.TerminateDuration)
	n += 1 + l + sovHcsshim(uint64(l))
	return n
}

func (m *ProcessDetails) Size() (n int) {
	var l int
	_ = l
	l = len(m.ImageName)
	if l > 0 {
		n += 1 + l + sovHcsshim(uint64(l))
	}
	if m.CreatedAt != nil {
		l = m.CreatedAt.Size()
		n += 1 + l + sovHcsshim(uint64(l))
	}
	if m.KernelTime_100Ns != 0 {
		n += 1 + sovHcsshim(uint64(m.KernelTime_100Ns))
	}
	if m.MemoryCommitBytes != 0 {
		n += 1 + sovHcsshim(uint64(m.MemoryCommitBytes))
	}
	if m.MemoryWorkingSetPrivateBytes != 0 {
		n += 1 + sovHcsshim(uint64(m.MemoryWorkingSetPrivateBytes))
	}
	if m.MemoryWorkingSetSharedBytes != 0 {
		n += 1 + sovHcsshim(uint64(m.MemoryWorkingSetSharedBytes))
	}
	if m.ProcessID != 0 {
		n += 1 + sovHcsshim(uint64(m.ProcessID))
	}
	if m.UserTime_100Ns != 0 {
		n += 1 + sovHcsshim(uint64(m.UserTime_100Ns))
	}
	return n
}

func sovHcsshim(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozHcsshim(x uint64) (n int) {
	return sovHcsshim(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *CreateOptions) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&CreateOptions{`,
		`TerminateDuration:` + strings.Replace(strings.Replace(this.TerminateDuration.String(), "Duration", "google_protobuf1.Duration", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ProcessDetails) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ProcessDetails{`,
		`ImageName:` + fmt.Sprintf("%v", this.ImageName) + `,`,
		`CreatedAt:` + strings.Replace(fmt.Sprintf("%v", this.CreatedAt), "Timestamp", "google_protobuf2.Timestamp", 1) + `,`,
		`KernelTime_100Ns:` + fmt.Sprintf("%v", this.KernelTime_100Ns) + `,`,
		`MemoryCommitBytes:` + fmt.Sprintf("%v", this.MemoryCommitBytes) + `,`,
		`MemoryWorkingSetPrivateBytes:` + fmt.Sprintf("%v", this.MemoryWorkingSetPrivateBytes) + `,`,
		`MemoryWorkingSetSharedBytes:` + fmt.Sprintf("%v", this.MemoryWorkingSetSharedBytes) + `,`,
		`ProcessID:` + fmt.Sprintf("%v", this.ProcessID) + `,`,
		`UserTime_100Ns:` + fmt.Sprintf("%v", this.UserTime_100Ns) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringHcsshim(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *CreateOptions) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHcsshim
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
			return fmt.Errorf("proto: CreateOptions: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CreateOptions: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TerminateDuration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHcsshim
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
				return ErrInvalidLengthHcsshim
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.TerminateDuration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHcsshim(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHcsshim
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
func (m *ProcessDetails) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHcsshim
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
			return fmt.Errorf("proto: ProcessDetails: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProcessDetails: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ImageName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHcsshim
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
				return ErrInvalidLengthHcsshim
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ImageName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHcsshim
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
				return ErrInvalidLengthHcsshim
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CreatedAt == nil {
				m.CreatedAt = &google_protobuf2.Timestamp{}
			}
			if err := m.CreatedAt.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field KernelTime_100Ns", wireType)
			}
			m.KernelTime_100Ns = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHcsshim
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.KernelTime_100Ns |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MemoryCommitBytes", wireType)
			}
			m.MemoryCommitBytes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHcsshim
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MemoryCommitBytes |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MemoryWorkingSetPrivateBytes", wireType)
			}
			m.MemoryWorkingSetPrivateBytes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHcsshim
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MemoryWorkingSetPrivateBytes |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MemoryWorkingSetSharedBytes", wireType)
			}
			m.MemoryWorkingSetSharedBytes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHcsshim
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MemoryWorkingSetSharedBytes |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProcessID", wireType)
			}
			m.ProcessID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHcsshim
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ProcessID |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserTime_100Ns", wireType)
			}
			m.UserTime_100Ns = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHcsshim
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UserTime_100Ns |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipHcsshim(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHcsshim
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
func skipHcsshim(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHcsshim
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
					return 0, ErrIntOverflowHcsshim
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
					return 0, ErrIntOverflowHcsshim
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
				return 0, ErrInvalidLengthHcsshim
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowHcsshim
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
				next, err := skipHcsshim(dAtA[start:])
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
	ErrInvalidLengthHcsshim = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHcsshim   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("github.com/containerd/containerd/windows/hcsshimopts/hcsshim.proto", fileDescriptorHcsshim)
}

var fileDescriptorHcsshim = []byte{
	// 473 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xc1, 0x6e, 0xd3, 0x30,
	0x1c, 0xc6, 0x1b, 0x36, 0xc6, 0x6a, 0x54, 0x60, 0x86, 0x43, 0x28, 0x90, 0x54, 0xbb, 0x50, 0x09,
	0x94, 0x74, 0x70, 0x42, 0x5c, 0x20, 0xab, 0x90, 0x76, 0x19, 0x53, 0x86, 0x84, 0xb4, 0x8b, 0xe5,
	0x26, 0x7f, 0x52, 0x6b, 0xb5, 0x1d, 0xd9, 0x2e, 0xd5, 0x6e, 0x3c, 0x06, 0x0f, 0xc1, 0x83, 0xf4,
	0xc8, 0x91, 0xd3, 0x60, 0x79, 0x12, 0x14, 0xdb, 0x2d, 0x63, 0x70, 0xda, 0xed, 0x1f, 0x7f, 0xbf,
	0xef, 0xfb, 0xc7, 0x5f, 0x82, 0xb2, 0x8a, 0x99, 0xe9, 0x7c, 0x92, 0x14, 0x92, 0xa7, 0x85, 0x14,
	0x86, 0x32, 0x01, 0xaa, 0xbc, 0x3c, 0x2e, 0x98, 0x28, 0xe5, 0x42, 0xa7, 0xd3, 0x42, 0xeb, 0x29,
	0xe3, 0xb2, 0x36, 0xeb, 0x39, 0xa9, 0x95, 0x34, 0x12, 0xf7, 0xff, 0xd0, 0x89, 0xa7, 0x13, 0x4f,
	0xf4, 0x1f, 0x54, 0xb2, 0x92, 0x16, 0x4b, 0xdb, 0xc9, 0x39, 0xfa, 0x51, 0x25, 0x65, 0x35, 0x83,
	0xd4, 0x3e, 0x4d, 0xe6, 0x9f, 0xd2, 0x72, 0xae, 0xa8, 0x61, 0x52, 0x78, 0x3d, 0xbe, 0xaa, 0x1b,
	0xc6, 0x41, 0x1b, 0xca, 0x6b, 0x07, 0xec, 0x16, 0xa8, 0xb7, 0xaf, 0x80, 0x1a, 0x78, 0x5f, 0xb7,
	0x36, 0x8d, 0x73, 0x84, 0x0d, 0x28, 0xce, 0x04, 0x35, 0x40, 0x56, 0x69, 0x61, 0x30, 0x08, 0x86,
	0xb7, 0x5f, 0x3c, 0x4c, 0x5c, 0x5c, 0xb2, 0x8a, 0x4b, 0xc6, 0x1e, 0xc8, 0xb6, 0x97, 0xe7, 0x71,
	0xe7, 0xeb, 0xcf, 0x38, 0xc8, 0x77, 0xd6, 0xf6, 0x95, 0xb8, 0xfb, 0x6d, 0x03, 0xdd, 0x39, 0x52,
	0xb2, 0x00, 0xad, 0xc7, 0x60, 0x28, 0x9b, 0x69, 0xfc, 0x04, 0x21, 0xc6, 0x69, 0x05, 0x44, 0x50,
	0x0e, 0x36, 0xbe, 0x9b, 0x77, 0xed, 0xc9, 0x21, 0xe5, 0x80, 0x5f, 0x21, 0x54, 0xd8, 0xd7, 0x2a,
	0x09, 0x35, 0xe1, 0x0d, 0xbb, 0xbd, 0xff, 0xcf, 0xf6, 0x0f, 0xab, 0xcb, 0xe4, 0x5d, 0x4f, 0xbf,
	0x35, 0xf8, 0x19, 0xc2, 0xa7, 0xa0, 0x04, 0xcc, 0x48, 0x7b, 0x57, 0xb2, 0x37, 0x1a, 0x11, 0xa1,
	0xc3, 0x8d, 0x41, 0x30, 0xdc, 0xcc, 0xef, 0x3a, 0xa5, 0xf5, 0xed, 0x8d, 0x46, 0x87, 0x1a, 0x27,
	0xe8, 0x3e, 0x07, 0x2e, 0xd5, 0x19, 0x29, 0x24, 0xe7, 0xcc, 0x90, 0xc9, 0x99, 0x01, 0x1d, 0x6e,
	0x5a, 0x7a, 0xc7, 0x49, 0xfb, 0x56, 0xc9, 0x5a, 0x01, 0xbf, 0x43, 0x03, 0xcf, 0x2f, 0xa4, 0x3a,
	0x65, 0xa2, 0x22, 0x1a, 0x0c, 0xa9, 0x15, 0xfb, 0xdc, 0xd6, 0xe5, 0xcc, 0x37, 0xad, 0xf9, 0xb1,
	0xe3, 0x3e, 0x3a, 0xec, 0x18, 0xcc, 0x91, 0x83, 0x5c, 0xce, 0x18, 0xc5, 0xff, 0xc9, 0xd1, 0x53,
	0xaa, 0xa0, 0xf4, 0x31, 0x5b, 0x36, 0xe6, 0xd1, 0xd5, 0x98, 0x63, 0xcb, 0xb8, 0x94, 0xe7, 0x08,
	0xd5, 0xae, 0x56, 0xc2, 0xca, 0xf0, 0xd6, 0x20, 0x18, 0xf6, 0xb2, 0x5e, 0x73, 0x1e, 0x77, 0x7d,
	0xd9, 0x07, 0xe3, 0xbc, 0xeb, 0x81, 0x83, 0x12, 0x3f, 0x45, 0xf7, 0xe6, 0x1a, 0xd4, 0x5f, 0xb5,
	0x6c, 0xdb, 0x25, 0xbd, 0xf6, 0x7c, 0x5d, 0x4a, 0x76, 0xb2, 0xbc, 0x88, 0x3a, 0x3f, 0x2e, 0xa2,
	0xce, 0x97, 0x26, 0x0a, 0x96, 0x4d, 0x14, 0x7c, 0x6f, 0xa2, 0xe0, 0x57, 0x13, 0x05, 0x27, 0x6f,
	0xae, 0xf3, 0x93, 0xbf, 0xbe, 0x34, 0x4f, 0xb6, 0xec, 0xc7, 0x7b, 0xf9, 0x3b, 0x00, 0x00, 0xff,
	0xff, 0x75, 0x5d, 0xc9, 0xa6, 0x2f, 0x03, 0x00, 0x00,
}
