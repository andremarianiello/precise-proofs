// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proofs/proto/proof.proto

package proofspb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type MerkleHash struct {
	Left                 []byte   `protobuf:"bytes,1,opt,name=left,proto3" json:"left,omitempty"`
	Right                []byte   `protobuf:"bytes,2,opt,name=right,proto3" json:"right,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MerkleHash) Reset()         { *m = MerkleHash{} }
func (m *MerkleHash) String() string { return proto.CompactTextString(m) }
func (*MerkleHash) ProtoMessage()    {}
func (*MerkleHash) Descriptor() ([]byte, []int) {
	return fileDescriptor_proof_efe8e9a45ef150c9, []int{0}
}
func (m *MerkleHash) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MerkleHash.Unmarshal(m, b)
}
func (m *MerkleHash) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MerkleHash.Marshal(b, m, deterministic)
}
func (dst *MerkleHash) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MerkleHash.Merge(dst, src)
}
func (m *MerkleHash) XXX_Size() int {
	return xxx_messageInfo_MerkleHash.Size(m)
}
func (m *MerkleHash) XXX_DiscardUnknown() {
	xxx_messageInfo_MerkleHash.DiscardUnknown(m)
}

var xxx_messageInfo_MerkleHash proto.InternalMessageInfo

func (m *MerkleHash) GetLeft() []byte {
	if m != nil {
		return m.Left
	}
	return nil
}

func (m *MerkleHash) GetRight() []byte {
	if m != nil {
		return m.Right
	}
	return nil
}

type Proof struct {
	Property []byte `protobuf:"bytes,1,opt,name=property,proto3" json:"property,omitempty"`
	Value    string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	Salt     []byte `protobuf:"bytes,3,opt,name=salt,proto3" json:"salt,omitempty"`
	// hash is filled if value & salt are not available
	Hash []byte `protobuf:"bytes,6,opt,name=hash,proto3" json:"hash,omitempty"`
	// Fills either 'hashes' for standard Merkle trees or 'sorted_hashes' for a lexicograhical ordered of a node hash
	// not both
	Hashes               []*MerkleHash `protobuf:"bytes,4,rep,name=hashes" json:"hashes,omitempty"`
	SortedHashes         [][]byte      `protobuf:"bytes,5,rep,name=sorted_hashes,json=sortedHashes,proto3" json:"sorted_hashes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Proof) Reset()         { *m = Proof{} }
func (m *Proof) String() string { return proto.CompactTextString(m) }
func (*Proof) ProtoMessage()    {}
func (*Proof) Descriptor() ([]byte, []int) {
	return fileDescriptor_proof_efe8e9a45ef150c9, []int{1}
}
func (m *Proof) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Proof.Unmarshal(m, b)
}
func (m *Proof) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Proof.Marshal(b, m, deterministic)
}
func (dst *Proof) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Proof.Merge(dst, src)
}
func (m *Proof) XXX_Size() int {
	return xxx_messageInfo_Proof.Size(m)
}
func (m *Proof) XXX_DiscardUnknown() {
	xxx_messageInfo_Proof.DiscardUnknown(m)
}

var xxx_messageInfo_Proof proto.InternalMessageInfo

func (m *Proof) GetProperty() []byte {
	if m != nil {
		return m.Property
	}
	return nil
}

func (m *Proof) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *Proof) GetSalt() []byte {
	if m != nil {
		return m.Salt
	}
	return nil
}

func (m *Proof) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *Proof) GetHashes() []*MerkleHash {
	if m != nil {
		return m.Hashes
	}
	return nil
}

func (m *Proof) GetSortedHashes() [][]byte {
	if m != nil {
		return m.SortedHashes
	}
	return nil
}

var E_ExcludeFromTree = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         2862100,
	Name:          "proofs.exclude_from_tree",
	Tag:           "varint,2862100,opt,name=exclude_from_tree,json=excludeFromTree",
	Filename:      "proofs/proto/proof.proto",
}

var E_HashedField = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         2862101,
	Name:          "proofs.hashed_field",
	Tag:           "varint,2862101,opt,name=hashed_field,json=hashedField",
	Filename:      "proofs/proto/proof.proto",
}

func init() {
	proto.RegisterType((*MerkleHash)(nil), "proofs.MerkleHash")
	proto.RegisterType((*Proof)(nil), "proofs.Proof")
	proto.RegisterExtension(E_ExcludeFromTree)
	proto.RegisterExtension(E_HashedField)
}

func init() { proto.RegisterFile("proofs/proto/proof.proto", fileDescriptor_proof_efe8e9a45ef150c9) }

var fileDescriptor_proof_efe8e9a45ef150c9 = []byte{
	// 327 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0xc1, 0x4a, 0xfb, 0x40,
	0x10, 0xc6, 0xc9, 0xbf, 0x6d, 0xe8, 0x7f, 0x1a, 0x11, 0x17, 0x0f, 0xa1, 0x20, 0x84, 0xea, 0xa1,
	0x78, 0x48, 0x41, 0xc1, 0x43, 0x8f, 0x15, 0x4a, 0x0f, 0x8a, 0x25, 0x78, 0xf2, 0x12, 0xd2, 0x66,
	0xd2, 0x04, 0xb7, 0xce, 0x32, 0xbb, 0x15, 0x7d, 0x0f, 0x7d, 0x0d, 0x0f, 0x3e, 0x8d, 0x8f, 0x23,
	0xbb, 0x9b, 0xea, 0xd1, 0x53, 0xe6, 0xfb, 0x76, 0x7e, 0xf3, 0x65, 0x06, 0x62, 0xc5, 0x44, 0x95,
	0x9e, 0x28, 0x26, 0x43, 0x13, 0x27, 0x52, 0x57, 0x8b, 0xd0, 0xbf, 0x0c, 0x93, 0x0d, 0xd1, 0x46,
	0xa2, 0xef, 0x58, 0xed, 0xaa, 0x49, 0x89, 0x7a, 0xcd, 0x8d, 0x32, 0xc4, 0xbe, 0x73, 0x74, 0x05,
	0x70, 0x8b, 0xfc, 0x28, 0x71, 0x51, 0xe8, 0x5a, 0x08, 0xe8, 0x4a, 0xac, 0x4c, 0x1c, 0x24, 0xc1,
	0x38, 0xca, 0x5c, 0x2d, 0x8e, 0xa1, 0xc7, 0xcd, 0xa6, 0x36, 0xf1, 0x3f, 0x67, 0x7a, 0x31, 0xfa,
	0x0c, 0xa0, 0xb7, 0xb4, 0x21, 0x62, 0x08, 0x7d, 0xc5, 0xa4, 0x90, 0xcd, 0x6b, 0xcb, 0xfd, 0x68,
	0xcb, 0x3e, 0x17, 0x72, 0x87, 0x8e, 0xfd, 0x9f, 0x79, 0x61, 0x53, 0x74, 0x21, 0x4d, 0xdc, 0xf1,
	0x29, 0xb6, 0xb6, 0x5e, 0x5d, 0xe8, 0x3a, 0x0e, 0xbd, 0x67, 0x6b, 0x71, 0x0e, 0xa1, 0xfd, 0xa2,
	0x8e, 0xbb, 0x49, 0x67, 0x3c, 0xb8, 0x10, 0xa9, 0x5f, 0x2b, 0xfd, 0xfd, 0xe3, 0xac, 0xed, 0x10,
	0xa7, 0x70, 0xa0, 0x89, 0x0d, 0x96, 0x79, 0x8b, 0xf4, 0x92, 0xce, 0x38, 0xca, 0x22, 0x6f, 0x2e,
	0x9c, 0x37, 0xbd, 0x81, 0x23, 0x7c, 0x59, 0xcb, 0x5d, 0x89, 0x79, 0xc5, 0xb4, 0xcd, 0x0d, 0x23,
	0x8a, 0x93, 0xd4, 0x1f, 0x29, 0xdd, 0x1f, 0x29, 0x9d, 0x37, 0x28, 0xcb, 0x3b, 0x65, 0x1a, 0x7a,
	0xd2, 0xf1, 0xdb, 0xd7, 0x87, 0xdd, 0xaa, 0x9f, 0x1d, 0xb6, 0xe8, 0x9c, 0x69, 0x7b, 0xcf, 0x88,
	0xd3, 0x6b, 0x88, 0x5c, 0x56, 0x99, 0x57, 0x16, 0xf8, 0x6b, 0xd0, 0xfb, 0x7e, 0xd0, 0xc0, 0x53,
	0xee, 0x71, 0x76, 0x06, 0xb0, 0xa6, 0x6d, 0xbb, 0xd8, 0x0c, 0xdc, 0x49, 0x97, 0x96, 0x5f, 0x06,
	0x0f, 0x7d, 0xef, 0xaa, 0xd5, 0x2a, 0x74, 0x23, 0x2f, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x9e,
	0xd0, 0xf6, 0xac, 0xf2, 0x01, 0x00, 0x00,
}
