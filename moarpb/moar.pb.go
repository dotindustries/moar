// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: moar.proto

package moarpb

import (
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

type GetUrlRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ModuleName string `protobuf:"bytes,1,opt,name=moduleName,proto3" json:"moduleName,omitempty"`
	// Types that are assignable to VersionSelector:
	//	*GetUrlRequest_VersionConstraint
	//	*GetUrlRequest_Version
	VersionSelector isGetUrlRequest_VersionSelector `protobuf_oneof:"versionSelector"`
}

func (x *GetUrlRequest) Reset() {
	*x = GetUrlRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moar_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUrlRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUrlRequest) ProtoMessage() {}

func (x *GetUrlRequest) ProtoReflect() protoreflect.Message {
	mi := &file_moar_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUrlRequest.ProtoReflect.Descriptor instead.
func (*GetUrlRequest) Descriptor() ([]byte, []int) {
	return file_moar_proto_rawDescGZIP(), []int{0}
}

func (x *GetUrlRequest) GetModuleName() string {
	if x != nil {
		return x.ModuleName
	}
	return ""
}

func (m *GetUrlRequest) GetVersionSelector() isGetUrlRequest_VersionSelector {
	if m != nil {
		return m.VersionSelector
	}
	return nil
}

func (x *GetUrlRequest) GetVersionConstraint() string {
	if x, ok := x.GetVersionSelector().(*GetUrlRequest_VersionConstraint); ok {
		return x.VersionConstraint
	}
	return ""
}

func (x *GetUrlRequest) GetVersion() string {
	if x, ok := x.GetVersionSelector().(*GetUrlRequest_Version); ok {
		return x.Version
	}
	return ""
}

type isGetUrlRequest_VersionSelector interface {
	isGetUrlRequest_VersionSelector()
}

type GetUrlRequest_VersionConstraint struct {
	VersionConstraint string `protobuf:"bytes,2,opt,name=versionConstraint,proto3,oneof"`
}

type GetUrlRequest_Version struct {
	Version string `protobuf:"bytes,3,opt,name=version,proto3,oneof"`
}

func (*GetUrlRequest_VersionConstraint) isGetUrlRequest_VersionSelector() {}

func (*GetUrlRequest_Version) isGetUrlRequest_VersionSelector() {}

type GetUrlResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uri             string  `protobuf:"bytes,1,opt,name=uri,proto3" json:"uri,omitempty"`
	Module          *Module `protobuf:"bytes,2,opt,name=module,proto3" json:"module,omitempty"`
	SelectedVersion string  `protobuf:"bytes,3,opt,name=selectedVersion,proto3" json:"selectedVersion,omitempty"`
}

func (x *GetUrlResponse) Reset() {
	*x = GetUrlResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moar_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUrlResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUrlResponse) ProtoMessage() {}

func (x *GetUrlResponse) ProtoReflect() protoreflect.Message {
	mi := &file_moar_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUrlResponse.ProtoReflect.Descriptor instead.
func (*GetUrlResponse) Descriptor() ([]byte, []int) {
	return file_moar_proto_rawDescGZIP(), []int{1}
}

func (x *GetUrlResponse) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

func (x *GetUrlResponse) GetModule() *Module {
	if x != nil {
		return x.Module
	}
	return nil
}

func (x *GetUrlResponse) GetSelectedVersion() string {
	if x != nil {
		return x.SelectedVersion
	}
	return ""
}

type CreateModuleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ModuleName string `protobuf:"bytes,1,opt,name=moduleName,proto3" json:"moduleName,omitempty"`
	Author     string `protobuf:"bytes,2,opt,name=author,proto3" json:"author,omitempty"`
}

func (x *CreateModuleRequest) Reset() {
	*x = CreateModuleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moar_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateModuleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateModuleRequest) ProtoMessage() {}

func (x *CreateModuleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_moar_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateModuleRequest.ProtoReflect.Descriptor instead.
func (*CreateModuleRequest) Descriptor() ([]byte, []int) {
	return file_moar_proto_rawDescGZIP(), []int{2}
}

func (x *CreateModuleRequest) GetModuleName() string {
	if x != nil {
		return x.ModuleName
	}
	return ""
}

func (x *CreateModuleRequest) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

type CreateModuleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateModuleResponse) Reset() {
	*x = CreateModuleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moar_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateModuleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateModuleResponse) ProtoMessage() {}

func (x *CreateModuleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_moar_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateModuleResponse.ProtoReflect.Descriptor instead.
func (*CreateModuleResponse) Descriptor() ([]byte, []int) {
	return file_moar_proto_rawDescGZIP(), []int{3}
}

type DeleteModuleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ModuleName string `protobuf:"bytes,1,opt,name=moduleName,proto3" json:"moduleName,omitempty"`
}

func (x *DeleteModuleRequest) Reset() {
	*x = DeleteModuleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moar_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteModuleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteModuleRequest) ProtoMessage() {}

func (x *DeleteModuleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_moar_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteModuleRequest.ProtoReflect.Descriptor instead.
func (*DeleteModuleRequest) Descriptor() ([]byte, []int) {
	return file_moar_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteModuleRequest) GetModuleName() string {
	if x != nil {
		return x.ModuleName
	}
	return ""
}

type DeleteModuleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteModuleResponse) Reset() {
	*x = DeleteModuleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moar_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteModuleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteModuleResponse) ProtoMessage() {}

func (x *DeleteModuleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_moar_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteModuleResponse.ProtoReflect.Descriptor instead.
func (*DeleteModuleResponse) Descriptor() ([]byte, []int) {
	return file_moar_proto_rawDescGZIP(), []int{5}
}

type UploadVersionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ModuleName string `protobuf:"bytes,1,opt,name=moduleName,proto3" json:"moduleName,omitempty"`
	Version    string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	FileData   []byte `protobuf:"bytes,3,opt,name=fileData,proto3" json:"fileData,omitempty"`
}

func (x *UploadVersionRequest) Reset() {
	*x = UploadVersionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moar_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadVersionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadVersionRequest) ProtoMessage() {}

func (x *UploadVersionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_moar_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadVersionRequest.ProtoReflect.Descriptor instead.
func (*UploadVersionRequest) Descriptor() ([]byte, []int) {
	return file_moar_proto_rawDescGZIP(), []int{6}
}

func (x *UploadVersionRequest) GetModuleName() string {
	if x != nil {
		return x.ModuleName
	}
	return ""
}

func (x *UploadVersionRequest) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *UploadVersionRequest) GetFileData() []byte {
	if x != nil {
		return x.FileData
	}
	return nil
}

type UploadVersionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UploadVersionResponse) Reset() {
	*x = UploadVersionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moar_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadVersionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadVersionResponse) ProtoMessage() {}

func (x *UploadVersionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_moar_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadVersionResponse.ProtoReflect.Descriptor instead.
func (*UploadVersionResponse) Descriptor() ([]byte, []int) {
	return file_moar_proto_rawDescGZIP(), []int{7}
}

type DeleteVersionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ModuleName string `protobuf:"bytes,1,opt,name=moduleName,proto3" json:"moduleName,omitempty"`
	Version    string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *DeleteVersionRequest) Reset() {
	*x = DeleteVersionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moar_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteVersionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteVersionRequest) ProtoMessage() {}

func (x *DeleteVersionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_moar_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteVersionRequest.ProtoReflect.Descriptor instead.
func (*DeleteVersionRequest) Descriptor() ([]byte, []int) {
	return file_moar_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteVersionRequest) GetModuleName() string {
	if x != nil {
		return x.ModuleName
	}
	return ""
}

func (x *DeleteVersionRequest) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type DeleteVersionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteVersionResponse) Reset() {
	*x = DeleteVersionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moar_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteVersionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteVersionResponse) ProtoMessage() {}

func (x *DeleteVersionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_moar_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteVersionResponse.ProtoReflect.Descriptor instead.
func (*DeleteVersionResponse) Descriptor() ([]byte, []int) {
	return file_moar_proto_rawDescGZIP(), []int{9}
}

type Module struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Versions []string `protobuf:"bytes,2,rep,name=versions,proto3" json:"versions,omitempty"`
	Author   string   `protobuf:"bytes,3,opt,name=author,proto3" json:"author,omitempty"`
	Language string   `protobuf:"bytes,4,opt,name=language,proto3" json:"language,omitempty"`
}

func (x *Module) Reset() {
	*x = Module{}
	if protoimpl.UnsafeEnabled {
		mi := &file_moar_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Module) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Module) ProtoMessage() {}

func (x *Module) ProtoReflect() protoreflect.Message {
	mi := &file_moar_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Module.ProtoReflect.Descriptor instead.
func (*Module) Descriptor() ([]byte, []int) {
	return file_moar_proto_rawDescGZIP(), []int{10}
}

func (x *Module) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Module) GetVersions() []string {
	if x != nil {
		return x.Versions
	}
	return nil
}

func (x *Module) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *Module) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

var File_moar_proto protoreflect.FileDescriptor

var file_moar_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6d, 0x6f, 0x61, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6d, 0x6f,
	0x61, 0x72, 0x70, 0x62, 0x22, 0x8e, 0x01, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x55, 0x72, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x6f, 0x64, 0x75,
	0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x11, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x11, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x73,
	0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x42, 0x11, 0x0a, 0x0f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x6c,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x22, 0x74, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55, 0x72, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x69, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x69, 0x12, 0x26, 0x0a, 0x06, 0x6d, 0x6f, 0x64,
	0x75, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6d, 0x6f, 0x61, 0x72,
	0x70, 0x62, 0x2e, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x06, 0x6d, 0x6f, 0x64, 0x75, 0x6c,
	0x65, 0x12, 0x28, 0x0a, 0x0f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x65, 0x6c, 0x65,
	0x63, 0x74, 0x65, 0x64, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x4d, 0x0a, 0x13, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x22, 0x16, 0x0a, 0x14, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x35, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x6f, 0x64, 0x75,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x6f, 0x64,
	0x75, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d,
	0x6f, 0x64, 0x75, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x16, 0x0a, 0x14, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x6c, 0x0a, 0x14, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x6f, 0x64,
	0x75, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d,
	0x6f, 0x64, 0x75, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x22,
	0x17, 0x0a, 0x15, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x50, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x17, 0x0a, 0x15, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x6c, 0x0a, 0x06, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x08, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67,
	0x65, 0x32, 0xfb, 0x02, 0x0a, 0x0e, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x79, 0x12, 0x37, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x55, 0x72, 0x6c, 0x12, 0x15,
	0x2e, 0x6d, 0x6f, 0x61, 0x72, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x72, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x6d, 0x6f, 0x61, 0x72, 0x70, 0x62, 0x2e, 0x47,
	0x65, 0x74, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a,
	0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x1b, 0x2e,
	0x6d, 0x6f, 0x61, 0x72, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x64,
	0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x6d, 0x6f, 0x61,
	0x72, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x1b, 0x2e, 0x6d, 0x6f, 0x61, 0x72, 0x70,
	0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x6d, 0x6f, 0x61, 0x72, 0x70, 0x62, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a, 0x0d, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x2e, 0x6d, 0x6f, 0x61, 0x72, 0x70, 0x62, 0x2e, 0x55, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x6d, 0x6f, 0x61, 0x72, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x4c, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x1c, 0x2e, 0x6d, 0x6f, 0x61, 0x72, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1d, 0x2e, 0x6d, 0x6f, 0x61, 0x72, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x20, 0x5a, 0x1e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x61,
	0x64, 0x69, 0x6c, 0x61, 0x73, 0x2f, 0x6d, 0x6f, 0x61, 0x72, 0x2f, 0x6d, 0x6f, 0x61, 0x72, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_moar_proto_rawDescOnce sync.Once
	file_moar_proto_rawDescData = file_moar_proto_rawDesc
)

func file_moar_proto_rawDescGZIP() []byte {
	file_moar_proto_rawDescOnce.Do(func() {
		file_moar_proto_rawDescData = protoimpl.X.CompressGZIP(file_moar_proto_rawDescData)
	})
	return file_moar_proto_rawDescData
}

var file_moar_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_moar_proto_goTypes = []interface{}{
	(*GetUrlRequest)(nil),         // 0: moarpb.GetUrlRequest
	(*GetUrlResponse)(nil),        // 1: moarpb.GetUrlResponse
	(*CreateModuleRequest)(nil),   // 2: moarpb.CreateModuleRequest
	(*CreateModuleResponse)(nil),  // 3: moarpb.CreateModuleResponse
	(*DeleteModuleRequest)(nil),   // 4: moarpb.DeleteModuleRequest
	(*DeleteModuleResponse)(nil),  // 5: moarpb.DeleteModuleResponse
	(*UploadVersionRequest)(nil),  // 6: moarpb.UploadVersionRequest
	(*UploadVersionResponse)(nil), // 7: moarpb.UploadVersionResponse
	(*DeleteVersionRequest)(nil),  // 8: moarpb.DeleteVersionRequest
	(*DeleteVersionResponse)(nil), // 9: moarpb.DeleteVersionResponse
	(*Module)(nil),                // 10: moarpb.Module
}
var file_moar_proto_depIdxs = []int32{
	10, // 0: moarpb.GetUrlResponse.module:type_name -> moarpb.Module
	0,  // 1: moarpb.ModuleRegistry.GetUrl:input_type -> moarpb.GetUrlRequest
	2,  // 2: moarpb.ModuleRegistry.CreateModule:input_type -> moarpb.CreateModuleRequest
	4,  // 3: moarpb.ModuleRegistry.DeleteModule:input_type -> moarpb.DeleteModuleRequest
	6,  // 4: moarpb.ModuleRegistry.UploadVersion:input_type -> moarpb.UploadVersionRequest
	8,  // 5: moarpb.ModuleRegistry.DeleteVersion:input_type -> moarpb.DeleteVersionRequest
	1,  // 6: moarpb.ModuleRegistry.GetUrl:output_type -> moarpb.GetUrlResponse
	3,  // 7: moarpb.ModuleRegistry.CreateModule:output_type -> moarpb.CreateModuleResponse
	5,  // 8: moarpb.ModuleRegistry.DeleteModule:output_type -> moarpb.DeleteModuleResponse
	7,  // 9: moarpb.ModuleRegistry.UploadVersion:output_type -> moarpb.UploadVersionResponse
	9,  // 10: moarpb.ModuleRegistry.DeleteVersion:output_type -> moarpb.DeleteVersionResponse
	6,  // [6:11] is the sub-list for method output_type
	1,  // [1:6] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_moar_proto_init() }
func file_moar_proto_init() {
	if File_moar_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_moar_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUrlRequest); i {
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
		file_moar_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUrlResponse); i {
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
		file_moar_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateModuleRequest); i {
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
		file_moar_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateModuleResponse); i {
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
		file_moar_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteModuleRequest); i {
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
		file_moar_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteModuleResponse); i {
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
		file_moar_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadVersionRequest); i {
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
		file_moar_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadVersionResponse); i {
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
		file_moar_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteVersionRequest); i {
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
		file_moar_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteVersionResponse); i {
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
		file_moar_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Module); i {
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
	file_moar_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*GetUrlRequest_VersionConstraint)(nil),
		(*GetUrlRequest_Version)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_moar_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_moar_proto_goTypes,
		DependencyIndexes: file_moar_proto_depIdxs,
		MessageInfos:      file_moar_proto_msgTypes,
	}.Build()
	File_moar_proto = out.File
	file_moar_proto_rawDesc = nil
	file_moar_proto_goTypes = nil
	file_moar_proto_depIdxs = nil
}
