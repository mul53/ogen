// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: wallet.proto

package proto

import (
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type SendTransactionInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account string `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Amount  string `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *SendTransactionInfo) Reset() {
	*x = SendTransactionInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendTransactionInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendTransactionInfo) ProtoMessage() {}

func (x *SendTransactionInfo) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendTransactionInfo.ProtoReflect.Descriptor instead.
func (*SendTransactionInfo) Descriptor() ([]byte, []int) {
	return file_wallet_proto_rawDescGZIP(), []int{0}
}

func (x *SendTransactionInfo) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *SendTransactionInfo) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

type Wallets struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Wallets []string `protobuf:"bytes,1,rep,name=wallets,proto3" json:"wallets,omitempty"`
}

func (x *Wallets) Reset() {
	*x = Wallets{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Wallets) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Wallets) ProtoMessage() {}

func (x *Wallets) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Wallets.ProtoReflect.Descriptor instead.
func (*Wallets) Descriptor() ([]byte, []int) {
	return file_wallet_proto_rawDescGZIP(), []int{1}
}

func (x *Wallets) GetWallets() []string {
	if x != nil {
		return x.Wallets
	}
	return nil
}

type WalletReference struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *WalletReference) Reset() {
	*x = WalletReference{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WalletReference) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WalletReference) ProtoMessage() {}

func (x *WalletReference) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WalletReference.ProtoReflect.Descriptor instead.
func (*WalletReference) Descriptor() ([]byte, []int) {
	return file_wallet_proto_rawDescGZIP(), []int{2}
}

func (x *WalletReference) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *WalletReference) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type NewWalletInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Account  string `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"`
	Mnemonic string `protobuf:"bytes,3,opt,name=mnemonic,proto3" json:"mnemonic,omitempty"`
}

func (x *NewWalletInfo) Reset() {
	*x = NewWalletInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewWalletInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewWalletInfo) ProtoMessage() {}

func (x *NewWalletInfo) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewWalletInfo.ProtoReflect.Descriptor instead.
func (*NewWalletInfo) Descriptor() ([]byte, []int) {
	return file_wallet_proto_rawDescGZIP(), []int{3}
}

func (x *NewWalletInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NewWalletInfo) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *NewWalletInfo) GetMnemonic() string {
	if x != nil {
		return x.Mnemonic
	}
	return ""
}

type ImportWalletData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Mnemonic string `protobuf:"bytes,2,opt,name=mnemonic,proto3" json:"mnemonic,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *ImportWalletData) Reset() {
	*x = ImportWalletData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImportWalletData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImportWalletData) ProtoMessage() {}

func (x *ImportWalletData) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImportWalletData.ProtoReflect.Descriptor instead.
func (*ImportWalletData) Descriptor() ([]byte, []int) {
	return file_wallet_proto_rawDescGZIP(), []int{4}
}

func (x *ImportWalletData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ImportWalletData) GetMnemonic() string {
	if x != nil {
		return x.Mnemonic
	}
	return ""
}

func (x *ImportWalletData) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type DumpHDWalletInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mnemonic string `protobuf:"bytes,1,opt,name=mnemonic,proto3" json:"mnemonic,omitempty"`
}

func (x *DumpHDWalletInfo) Reset() {
	*x = DumpHDWalletInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DumpHDWalletInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DumpHDWalletInfo) ProtoMessage() {}

func (x *DumpHDWalletInfo) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DumpHDWalletInfo.ProtoReflect.Descriptor instead.
func (*DumpHDWalletInfo) Descriptor() ([]byte, []int) {
	return file_wallet_proto_rawDescGZIP(), []int{5}
}

func (x *DumpHDWalletInfo) GetMnemonic() string {
	if x != nil {
		return x.Mnemonic
	}
	return ""
}

var File_wallet_proto protoreflect.FileDescriptor

var file_wallet_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x47, 0x0a, 0x13, 0x53, 0x65,
	0x6e, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x22, 0x23, 0x0a, 0x07, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x73, 0x12, 0x18,
	0x0a, 0x07, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x07, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x73, 0x22, 0x41, 0x0a, 0x0f, 0x57, 0x61, 0x6c, 0x6c,
	0x65, 0x74, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x59, 0x0a, 0x0d, 0x4e,
	0x65, 0x77, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x6e,
	0x65, 0x6d, 0x6f, 0x6e, 0x69, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x6e,
	0x65, 0x6d, 0x6f, 0x6e, 0x69, 0x63, 0x22, 0x5e, 0x0a, 0x10, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74,
	0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x6d, 0x6e, 0x65, 0x6d, 0x6f, 0x6e, 0x69, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6d, 0x6e, 0x65, 0x6d, 0x6f, 0x6e, 0x69, 0x63, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x2e, 0x0a, 0x10, 0x44, 0x75, 0x6d, 0x70, 0x48, 0x44,
	0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x6e,
	0x65, 0x6d, 0x6f, 0x6e, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x6e,
	0x65, 0x6d, 0x6f, 0x6e, 0x69, 0x63, 0x32, 0x97, 0x08, 0x0a, 0x06, 0x57, 0x61, 0x6c, 0x6c, 0x65,
	0x74, 0x12, 0x35, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x73,
	0x12, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x08, 0x2e, 0x57, 0x61, 0x6c, 0x6c, 0x65,
	0x74, 0x73, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x12, 0x0c, 0x2f, 0x77, 0x61, 0x6c,
	0x6c, 0x65, 0x74, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x4b, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x12, 0x10, 0x2e, 0x57, 0x61, 0x6c, 0x6c, 0x65,
	0x74, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x1a, 0x0e, 0x2e, 0x4e, 0x65, 0x77,
	0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x13, 0x22, 0x0e, 0x2f, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2f, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x41, 0x0a, 0x0a, 0x4f, 0x70, 0x65, 0x6e, 0x57, 0x61, 0x6c,
	0x6c, 0x65, 0x74, 0x12, 0x10, 0x2e, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x65, 0x66, 0x65,
	0x72, 0x65, 0x6e, 0x63, 0x65, 0x1a, 0x08, 0x2e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22,
	0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x22, 0x0c, 0x2f, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74,
	0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x3a, 0x01, 0x2a, 0x12, 0x46, 0x0a, 0x0c, 0x49, 0x6d, 0x70, 0x6f,
	0x72, 0x74, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x12, 0x11, 0x2e, 0x49, 0x6d, 0x70, 0x6f, 0x72,
	0x74, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x08, 0x2e, 0x4b, 0x65,
	0x79, 0x50, 0x61, 0x69, 0x72, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x22, 0x0e, 0x2f,
	0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2f, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x3a, 0x01, 0x2a,
	0x12, 0x34, 0x0a, 0x0a, 0x44, 0x75, 0x6d, 0x70, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x12, 0x06,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x08, 0x2e, 0x4b, 0x65, 0x79, 0x50, 0x61, 0x69, 0x72,
	0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x12, 0x0c, 0x2f, 0x77, 0x61, 0x6c, 0x6c, 0x65,
	0x74, 0x2f, 0x64, 0x75, 0x6d, 0x70, 0x12, 0x41, 0x0a, 0x0c, 0x44, 0x75, 0x6d, 0x70, 0x48, 0x44,
	0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x12, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x11,
	0x2e, 0x44, 0x75, 0x6d, 0x70, 0x48, 0x44, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x12, 0x0e, 0x2f, 0x77, 0x61, 0x6c, 0x6c,
	0x65, 0x74, 0x2f, 0x64, 0x75, 0x6d, 0x70, 0x68, 0x64, 0x12, 0x36, 0x0a, 0x0b, 0x43, 0x6c, 0x6f,
	0x73, 0x65, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x12, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x08, 0x2e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x0f, 0x12, 0x0d, 0x2f, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2f, 0x63, 0x6c, 0x6f, 0x73,
	0x65, 0x12, 0x37, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12,
	0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x08, 0x2e, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x12, 0x0f, 0x2f, 0x77, 0x61, 0x6c, 0x6c,
	0x65, 0x74, 0x2f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x48, 0x0a, 0x0d, 0x47, 0x65,
	0x74, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x06, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x13, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14,
	0x12, 0x12, 0x2f, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x6f, 0x72, 0x73, 0x12, 0x37, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x08, 0x2e, 0x4b, 0x65, 0x79,
	0x50, 0x61, 0x69, 0x72, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x12, 0x0f, 0x2f, 0x77,
	0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x52, 0x0a,
	0x0f, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x14, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x05, 0x2e, 0x48, 0x61, 0x73, 0x68, 0x22, 0x22, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x22, 0x17, 0x2f, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2f, 0x73,
	0x65, 0x6e, 0x64, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x01,
	0x2a, 0x12, 0x4e, 0x0a, 0x0e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x6f, 0x72, 0x12, 0x08, 0x2e, 0x4b, 0x65, 0x79, 0x50, 0x61, 0x69, 0x72, 0x1a, 0x08, 0x2e,
	0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x28, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x12,
	0x20, 0x2f, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x7b, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65,
	0x7d, 0x12, 0x50, 0x0a, 0x12, 0x53, 0x74, 0x61, 0x72, 0x74, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x6f, 0x72, 0x42, 0x75, 0x6c, 0x6b, 0x12, 0x09, 0x2e, 0x4b, 0x65, 0x79, 0x50, 0x61, 0x69,
	0x72, 0x73, 0x1a, 0x08, 0x2e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x25, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x1f, 0x22, 0x1a, 0x2f, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2f, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x62, 0x75, 0x6c, 0x6b,
	0x3a, 0x01, 0x2a, 0x12, 0x4b, 0x0a, 0x0d, 0x45, 0x78, 0x69, 0x74, 0x56, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x6f, 0x72, 0x12, 0x08, 0x2e, 0x4b, 0x65, 0x79, 0x50, 0x61, 0x69, 0x72, 0x1a, 0x08,
	0x2e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20,
	0x12, 0x1e, 0x2f, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2f, 0x65, 0x78, 0x69, 0x74, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x7b, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x7d,
	0x12, 0x4e, 0x0a, 0x11, 0x45, 0x78, 0x69, 0x74, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f,
	0x72, 0x42, 0x75, 0x6c, 0x6b, 0x12, 0x09, 0x2e, 0x4b, 0x65, 0x79, 0x50, 0x61, 0x69, 0x72, 0x73,
	0x1a, 0x08, 0x2e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x1e, 0x22, 0x19, 0x2f, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2f, 0x65, 0x78, 0x69, 0x74,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x62, 0x75, 0x6c, 0x6b, 0x3a, 0x01, 0x2a,
	0x42, 0x09, 0x5a, 0x07, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_wallet_proto_rawDescOnce sync.Once
	file_wallet_proto_rawDescData = file_wallet_proto_rawDesc
)

func file_wallet_proto_rawDescGZIP() []byte {
	file_wallet_proto_rawDescOnce.Do(func() {
		file_wallet_proto_rawDescData = protoimpl.X.CompressGZIP(file_wallet_proto_rawDescData)
	})
	return file_wallet_proto_rawDescData
}

var file_wallet_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_wallet_proto_goTypes = []interface{}{
	(*SendTransactionInfo)(nil), // 0: SendTransactionInfo
	(*Wallets)(nil),             // 1: Wallets
	(*WalletReference)(nil),     // 2: WalletReference
	(*NewWalletInfo)(nil),       // 3: NewWalletInfo
	(*ImportWalletData)(nil),    // 4: ImportWalletData
	(*DumpHDWalletInfo)(nil),    // 5: DumpHDWalletInfo
	(*Empty)(nil),               // 6: Empty
	(*KeyPair)(nil),             // 7: KeyPair
	(*KeyPairs)(nil),            // 8: KeyPairs
	(*Success)(nil),             // 9: Success
	(*Balance)(nil),             // 10: Balance
	(*ValidatorsRegistry)(nil),  // 11: ValidatorsRegistry
	(*Hash)(nil),                // 12: Hash
}
var file_wallet_proto_depIdxs = []int32{
	6,  // 0: Wallet.ListWallets:input_type -> Empty
	2,  // 1: Wallet.CreateWallet:input_type -> WalletReference
	2,  // 2: Wallet.OpenWallet:input_type -> WalletReference
	4,  // 3: Wallet.ImportWallet:input_type -> ImportWalletData
	6,  // 4: Wallet.DumpWallet:input_type -> Empty
	6,  // 5: Wallet.DumpHDWallet:input_type -> Empty
	6,  // 6: Wallet.CloseWallet:input_type -> Empty
	6,  // 7: Wallet.GetBalance:input_type -> Empty
	6,  // 8: Wallet.GetValidators:input_type -> Empty
	6,  // 9: Wallet.GetAccount:input_type -> Empty
	0,  // 10: Wallet.SendTransaction:input_type -> SendTransactionInfo
	7,  // 11: Wallet.StartValidator:input_type -> KeyPair
	8,  // 12: Wallet.StartValidatorBulk:input_type -> KeyPairs
	7,  // 13: Wallet.ExitValidator:input_type -> KeyPair
	8,  // 14: Wallet.ExitValidatorBulk:input_type -> KeyPairs
	1,  // 15: Wallet.ListWallets:output_type -> Wallets
	3,  // 16: Wallet.CreateWallet:output_type -> NewWalletInfo
	9,  // 17: Wallet.OpenWallet:output_type -> Success
	7,  // 18: Wallet.ImportWallet:output_type -> KeyPair
	7,  // 19: Wallet.DumpWallet:output_type -> KeyPair
	5,  // 20: Wallet.DumpHDWallet:output_type -> DumpHDWalletInfo
	9,  // 21: Wallet.CloseWallet:output_type -> Success
	10, // 22: Wallet.GetBalance:output_type -> Balance
	11, // 23: Wallet.GetValidators:output_type -> ValidatorsRegistry
	7,  // 24: Wallet.GetAccount:output_type -> KeyPair
	12, // 25: Wallet.SendTransaction:output_type -> Hash
	9,  // 26: Wallet.StartValidator:output_type -> Success
	9,  // 27: Wallet.StartValidatorBulk:output_type -> Success
	9,  // 28: Wallet.ExitValidator:output_type -> Success
	9,  // 29: Wallet.ExitValidatorBulk:output_type -> Success
	15, // [15:30] is the sub-list for method output_type
	0,  // [0:15] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_wallet_proto_init() }
func file_wallet_proto_init() {
	if File_wallet_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_wallet_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendTransactionInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wallet_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Wallets); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wallet_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WalletReference); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wallet_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewWalletInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wallet_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImportWalletData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wallet_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DumpHDWalletInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_wallet_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_wallet_proto_goTypes,
		DependencyIndexes: file_wallet_proto_depIdxs,
		MessageInfos:      file_wallet_proto_msgTypes,
	}.Build()
	File_wallet_proto = out.File
	file_wallet_proto_rawDesc = nil
	file_wallet_proto_goTypes = nil
	file_wallet_proto_depIdxs = nil
}
