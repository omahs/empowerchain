// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: empowerchain/accesscontrol/genesis.proto

package accesscontrol

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

// GenesisState defines the accesscontrol module's genesis state.
type GenesisState struct {
	PermStores []ModulePermStore `protobuf:"bytes,1,rep,name=perm_stores,json=permStores,proto3" json:"perm_stores"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_12aa109cac8f21f4, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetPermStores() []ModulePermStore {
	if m != nil {
		return m.PermStores
	}
	return nil
}

// All accesses for a given module
type ModulePermStore struct {
	// Name - will be used as a name for a PermStore
	ModuleName string `protobuf:"bytes,1,opt,name=module_name,json=moduleName,proto3" json:"module_name,omitempty"`
	// List of accesses
	Accesses []Access `protobuf:"bytes,2,rep,name=accesses,proto3" json:"accesses"`
}

func (m *ModulePermStore) Reset()         { *m = ModulePermStore{} }
func (m *ModulePermStore) String() string { return proto.CompactTextString(m) }
func (*ModulePermStore) ProtoMessage()    {}
func (*ModulePermStore) Descriptor() ([]byte, []int) {
	return fileDescriptor_12aa109cac8f21f4, []int{1}
}
func (m *ModulePermStore) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ModulePermStore) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ModulePermStore.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ModulePermStore) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModulePermStore.Merge(m, src)
}
func (m *ModulePermStore) XXX_Size() int {
	return m.Size()
}
func (m *ModulePermStore) XXX_DiscardUnknown() {
	xxx_messageInfo_ModulePermStore.DiscardUnknown(m)
}

var xxx_messageInfo_ModulePermStore proto.InternalMessageInfo

func (m *ModulePermStore) GetModuleName() string {
	if m != nil {
		return m.ModuleName
	}
	return ""
}

func (m *ModulePermStore) GetAccesses() []Access {
	if m != nil {
		return m.Accesses
	}
	return nil
}

// Single access consisting of permissioned address and msgType
// of the message it can invoke
type Access struct {
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	MsgType string `protobuf:"bytes,2,opt,name=msg_type,json=msgType,proto3" json:"msg_type,omitempty"`
}

func (m *Access) Reset()         { *m = Access{} }
func (m *Access) String() string { return proto.CompactTextString(m) }
func (*Access) ProtoMessage()    {}
func (*Access) Descriptor() ([]byte, []int) {
	return fileDescriptor_12aa109cac8f21f4, []int{2}
}
func (m *Access) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Access) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Access.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Access) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Access.Merge(m, src)
}
func (m *Access) XXX_Size() int {
	return m.Size()
}
func (m *Access) XXX_DiscardUnknown() {
	xxx_messageInfo_Access.DiscardUnknown(m)
}

var xxx_messageInfo_Access proto.InternalMessageInfo

func (m *Access) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Access) GetMsgType() string {
	if m != nil {
		return m.MsgType
	}
	return ""
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "empowerchain.accesscontrol.GenesisState")
	proto.RegisterType((*ModulePermStore)(nil), "empowerchain.accesscontrol.ModulePermStore")
	proto.RegisterType((*Access)(nil), "empowerchain.accesscontrol.Access")
}

func init() {
	proto.RegisterFile("empowerchain/accesscontrol/genesis.proto", fileDescriptor_12aa109cac8f21f4)
}

var fileDescriptor_12aa109cac8f21f4 = []byte{
	// 329 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0xc1, 0x6a, 0xf2, 0x40,
	0x18, 0x4c, 0xfc, 0x7f, 0xd4, 0xae, 0x85, 0x42, 0xf0, 0x10, 0x3d, 0x44, 0xc9, 0x49, 0x28, 0x4d,
	0xc0, 0xf6, 0x05, 0x94, 0x42, 0x4f, 0x95, 0x12, 0x0b, 0x85, 0x5e, 0x42, 0x8c, 0x1f, 0x6b, 0xc0,
	0xcd, 0x17, 0xf6, 0x5b, 0xa9, 0xbe, 0x45, 0x1f, 0xa6, 0x0f, 0xe1, 0x51, 0x7a, 0xea, 0xa9, 0x14,
	0x7d, 0x91, 0xe2, 0xae, 0x96, 0xa6, 0xd0, 0xde, 0x32, 0xf3, 0xcd, 0x64, 0x86, 0x1d, 0xd6, 0x03,
	0x51, 0xe0, 0x13, 0xc8, 0x74, 0x96, 0x64, 0x79, 0x98, 0xa4, 0x29, 0x10, 0xa5, 0x98, 0x2b, 0x89,
	0xf3, 0x90, 0x43, 0x0e, 0x94, 0x51, 0x50, 0x48, 0x54, 0xe8, 0xb4, 0xbf, 0x2b, 0x83, 0x92, 0xb2,
	0xdd, 0x4a, 0x91, 0x04, 0x52, 0xac, 0x95, 0xa1, 0x01, 0xc6, 0xd6, 0x6e, 0x72, 0xe4, 0x68, 0xf8,
	0xfd, 0x97, 0x61, 0xfd, 0x09, 0x3b, 0xbd, 0x31, 0x7f, 0x1f, 0xab, 0x44, 0x81, 0x13, 0xb1, 0x46,
	0x01, 0x52, 0xc4, 0xa4, 0x50, 0x02, 0xb9, 0x76, 0xf7, 0x5f, 0xaf, 0xd1, 0x3f, 0x0f, 0x7e, 0x8f,
	0x0c, 0x6e, 0x71, 0xba, 0x98, 0xc3, 0x1d, 0x48, 0x31, 0xde, 0x7b, 0x86, 0xff, 0xd7, 0xef, 0x1d,
	0x2b, 0x62, 0xc5, 0x91, 0x20, 0x7f, 0xc9, 0xce, 0x7e, 0x88, 0x9c, 0x0e, 0x6b, 0x08, 0x4d, 0xc5,
	0x79, 0x22, 0xc0, 0xb5, 0xbb, 0x76, 0xef, 0x24, 0x62, 0x86, 0x1a, 0x25, 0x02, 0x9c, 0x6b, 0x56,
	0x37, 0x31, 0x40, 0x6e, 0x45, 0x97, 0xf0, 0xff, 0x2a, 0x31, 0xd0, 0xe8, 0x90, 0xfd, 0xe5, 0xf4,
	0x1f, 0x58, 0xd5, 0x5c, 0x9c, 0x3e, 0xab, 0x25, 0xd3, 0xa9, 0x04, 0x22, 0x13, 0x36, 0x74, 0x5f,
	0x5f, 0x2e, 0x9a, 0x87, 0x07, 0x1a, 0x98, 0xcb, 0x58, 0xc9, 0x2c, 0xe7, 0xd1, 0x51, 0xe8, 0xb4,
	0x58, 0x5d, 0x10, 0x8f, 0xd5, 0xaa, 0x00, 0xb7, 0xa2, 0x1b, 0xd6, 0x04, 0xf1, 0xfb, 0x55, 0x01,
	0xc3, 0xd1, 0x7a, 0xeb, 0xd9, 0x9b, 0xad, 0x67, 0x7f, 0x6c, 0x3d, 0xfb, 0x79, 0xe7, 0x59, 0x9b,
	0x9d, 0x67, 0xbd, 0xed, 0x3c, 0xeb, 0xf1, 0x8a, 0x67, 0x6a, 0xb6, 0x98, 0x04, 0x29, 0x8a, 0xb0,
	0x34, 0x69, 0x09, 0x2c, 0xcb, 0x0b, 0x4f, 0xaa, 0x7a, 0x8d, 0xcb, 0xcf, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xd9, 0x4d, 0xcf, 0x41, 0x06, 0x02, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PermStores) > 0 {
		for iNdEx := len(m.PermStores) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PermStores[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *ModulePermStore) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ModulePermStore) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ModulePermStore) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Accesses) > 0 {
		for iNdEx := len(m.Accesses) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Accesses[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.ModuleName) > 0 {
		i -= len(m.ModuleName)
		copy(dAtA[i:], m.ModuleName)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.ModuleName)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Access) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Access) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Access) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.MsgType) > 0 {
		i -= len(m.MsgType)
		copy(dAtA[i:], m.MsgType)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.MsgType)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.PermStores) > 0 {
		for _, e := range m.PermStores {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *ModulePermStore) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ModuleName)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.Accesses) > 0 {
		for _, e := range m.Accesses {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *Access) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.MsgType)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PermStores", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PermStores = append(m.PermStores, ModulePermStore{})
			if err := m.PermStores[len(m.PermStores)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *ModulePermStore) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: ModulePermStore: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ModulePermStore: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ModuleName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ModuleName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Accesses", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Accesses = append(m.Accesses, Access{})
			if err := m.Accesses[len(m.Accesses)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *Access) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: Access: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Access: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MsgType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MsgType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)