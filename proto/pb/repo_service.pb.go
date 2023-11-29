// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: repo_service.proto

package pb

import (
	empty "github.com/golang/protobuf/ptypes/empty"
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

type PrivateRepoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	CreatedAt   string `protobuf:"bytes,4,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt   string `protobuf:"bytes,5,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	OwnerId     string `protobuf:"bytes,6,opt,name=ownerId,proto3" json:"ownerId,omitempty"`
}

func (x *PrivateRepoRequest) Reset() {
	*x = PrivateRepoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_repo_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrivateRepoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrivateRepoRequest) ProtoMessage() {}

func (x *PrivateRepoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_repo_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrivateRepoRequest.ProtoReflect.Descriptor instead.
func (*PrivateRepoRequest) Descriptor() ([]byte, []int) {
	return file_repo_service_proto_rawDescGZIP(), []int{0}
}

func (x *PrivateRepoRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PrivateRepoRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PrivateRepoRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *PrivateRepoRequest) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *PrivateRepoRequest) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *PrivateRepoRequest) GetOwnerId() string {
	if x != nil {
		return x.OwnerId
	}
	return ""
}

type PrivateReposResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Repos []*PrivateRepoResponse `protobuf:"bytes,1,rep,name=repos,proto3" json:"repos,omitempty"`
}

func (x *PrivateReposResponse) Reset() {
	*x = PrivateReposResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_repo_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrivateReposResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrivateReposResponse) ProtoMessage() {}

func (x *PrivateReposResponse) ProtoReflect() protoreflect.Message {
	mi := &file_repo_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrivateReposResponse.ProtoReflect.Descriptor instead.
func (*PrivateReposResponse) Descriptor() ([]byte, []int) {
	return file_repo_service_proto_rawDescGZIP(), []int{1}
}

func (x *PrivateReposResponse) GetRepos() []*PrivateRepoResponse {
	if x != nil {
		return x.Repos
	}
	return nil
}

type PrivateRepoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	CreatedAt   string `protobuf:"bytes,4,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt   string `protobuf:"bytes,5,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	OwnerId     string `protobuf:"bytes,6,opt,name=ownerId,proto3" json:"ownerId,omitempty"`
}

func (x *PrivateRepoResponse) Reset() {
	*x = PrivateRepoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_repo_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrivateRepoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrivateRepoResponse) ProtoMessage() {}

func (x *PrivateRepoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_repo_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrivateRepoResponse.ProtoReflect.Descriptor instead.
func (*PrivateRepoResponse) Descriptor() ([]byte, []int) {
	return file_repo_service_proto_rawDescGZIP(), []int{2}
}

func (x *PrivateRepoResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PrivateRepoResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PrivateRepoResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *PrivateRepoResponse) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *PrivateRepoResponse) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *PrivateRepoResponse) GetOwnerId() string {
	if x != nil {
		return x.OwnerId
	}
	return ""
}

type PublicReposResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Repos []*PublicRepoResponse `protobuf:"bytes,1,rep,name=repos,proto3" json:"repos,omitempty"`
}

func (x *PublicReposResponse) Reset() {
	*x = PublicReposResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_repo_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublicReposResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicReposResponse) ProtoMessage() {}

func (x *PublicReposResponse) ProtoReflect() protoreflect.Message {
	mi := &file_repo_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublicReposResponse.ProtoReflect.Descriptor instead.
func (*PublicReposResponse) Descriptor() ([]byte, []int) {
	return file_repo_service_proto_rawDescGZIP(), []int{3}
}

func (x *PublicReposResponse) GetRepos() []*PublicRepoResponse {
	if x != nil {
		return x.Repos
	}
	return nil
}

type PublicRepoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	CreatedAt   string `protobuf:"bytes,4,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt   string `protobuf:"bytes,5,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
}

func (x *PublicRepoResponse) Reset() {
	*x = PublicRepoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_repo_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublicRepoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicRepoResponse) ProtoMessage() {}

func (x *PublicRepoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_repo_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublicRepoResponse.ProtoReflect.Descriptor instead.
func (*PublicRepoResponse) Descriptor() ([]byte, []int) {
	return file_repo_service_proto_rawDescGZIP(), []int{4}
}

func (x *PublicRepoResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PublicRepoResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PublicRepoResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *PublicRepoResponse) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *PublicRepoResponse) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type CreateRepoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OwnerId     string `protobuf:"bytes,1,opt,name=ownerId,proto3" json:"ownerId,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *CreateRepoRequest) Reset() {
	*x = CreateRepoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_repo_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRepoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRepoRequest) ProtoMessage() {}

func (x *CreateRepoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_repo_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRepoRequest.ProtoReflect.Descriptor instead.
func (*CreateRepoRequest) Descriptor() ([]byte, []int) {
	return file_repo_service_proto_rawDescGZIP(), []int{5}
}

func (x *CreateRepoRequest) GetOwnerId() string {
	if x != nil {
		return x.OwnerId
	}
	return ""
}

func (x *CreateRepoRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateRepoRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type IdentifierRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RepoIdentifier string `protobuf:"bytes,1,opt,name=repoIdentifier,proto3" json:"repoIdentifier,omitempty"`
}

func (x *IdentifierRequest) Reset() {
	*x = IdentifierRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_repo_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdentifierRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdentifierRequest) ProtoMessage() {}

func (x *IdentifierRequest) ProtoReflect() protoreflect.Message {
	mi := &file_repo_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdentifierRequest.ProtoReflect.Descriptor instead.
func (*IdentifierRequest) Descriptor() ([]byte, []int) {
	return file_repo_service_proto_rawDescGZIP(), []int{6}
}

func (x *IdentifierRequest) GetRepoIdentifier() string {
	if x != nil {
		return x.RepoIdentifier
	}
	return ""
}

var File_repo_service_proto protoreflect.FileDescriptor

var file_repo_service_proto_rawDesc = []byte{
	0x0a, 0x12, 0x72, 0x65, 0x70, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xb0, 0x01, 0x0a, 0x12, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70,
	0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c,
	0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x77,
	0x6e, 0x65, 0x72, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x77, 0x6e,
	0x65, 0x72, 0x49, 0x64, 0x22, 0x42, 0x0a, 0x14, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x70, 0x6f, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x05,
	0x72, 0x65, 0x70, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x50, 0x72,
	0x69, 0x76, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x52, 0x05, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x22, 0xb1, 0x01, 0x0a, 0x13, 0x50, 0x72, 0x69,
	0x76, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x22, 0x40, 0x0a, 0x13,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x05, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x52, 0x65, 0x70, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x05, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x22, 0x96,
	0x01, 0x0a, 0x12, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x52, 0x65, 0x70, 0x6f, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x63, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x70, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3b, 0x0a, 0x11,
	0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x26, 0x0a, 0x0e, 0x72, 0x65, 0x70, 0x6f, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66,
	0x69, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x72, 0x65, 0x70, 0x6f, 0x49,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x32, 0xd0, 0x03, 0x0a, 0x0b, 0x52, 0x65,
	0x70, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x0a, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x12, 0x12, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x70, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x50, 0x72,
	0x69, 0x76, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x70, 0x6f, 0x42, 0x79, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x12, 0x12, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x70, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x46, 0x0a,
	0x19, 0x47, 0x65, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x52, 0x65, 0x70, 0x6f, 0x42, 0x79,
	0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x12, 0x2e, 0x49, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13,
	0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x52, 0x65, 0x70, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x50, 0x72, 0x69, 0x76,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x12, 0x12, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x50,
	0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x50, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x12, 0x12, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x66, 0x69, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x50, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70,
	0x6f, 0x12, 0x13, 0x2e, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x70, 0x6f, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x3b, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x12, 0x13, 0x2e,
	0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x11, 0x5a, 0x0f,
	0x72, 0x65, 0x70, 0x6f, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_repo_service_proto_rawDescOnce sync.Once
	file_repo_service_proto_rawDescData = file_repo_service_proto_rawDesc
)

func file_repo_service_proto_rawDescGZIP() []byte {
	file_repo_service_proto_rawDescOnce.Do(func() {
		file_repo_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_repo_service_proto_rawDescData)
	})
	return file_repo_service_proto_rawDescData
}

var file_repo_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_repo_service_proto_goTypes = []interface{}{
	(*PrivateRepoRequest)(nil),   // 0: PrivateRepoRequest
	(*PrivateReposResponse)(nil), // 1: PrivateReposResponse
	(*PrivateRepoResponse)(nil),  // 2: PrivateRepoResponse
	(*PublicReposResponse)(nil),  // 3: PublicReposResponse
	(*PublicRepoResponse)(nil),   // 4: PublicRepoResponse
	(*CreateRepoRequest)(nil),    // 5: CreateRepoRequest
	(*IdentifierRequest)(nil),    // 6: IdentifierRequest
	(*empty.Empty)(nil),          // 7: google.protobuf.Empty
}
var file_repo_service_proto_depIdxs = []int32{
	2, // 0: PrivateReposResponse.repos:type_name -> PrivateRepoResponse
	4, // 1: PublicReposResponse.repos:type_name -> PublicRepoResponse
	5, // 2: RepoService.CreateRepo:input_type -> CreateRepoRequest
	6, // 3: RepoService.GetPrivateRepoByIdentifier:input_type -> IdentifierRequest
	6, // 4: RepoService.GetPublicRepoByIdentifier:input_type -> IdentifierRequest
	6, // 5: RepoService.GetPrivateRepos:input_type -> IdentifierRequest
	6, // 6: RepoService.GetPublicRepos:input_type -> IdentifierRequest
	0, // 7: RepoService.UpdateRepo:input_type -> PrivateRepoRequest
	0, // 8: RepoService.DeleteRepo:input_type -> PrivateRepoRequest
	2, // 9: RepoService.CreateRepo:output_type -> PrivateRepoResponse
	2, // 10: RepoService.GetPrivateRepoByIdentifier:output_type -> PrivateRepoResponse
	4, // 11: RepoService.GetPublicRepoByIdentifier:output_type -> PublicRepoResponse
	1, // 12: RepoService.GetPrivateRepos:output_type -> PrivateReposResponse
	3, // 13: RepoService.GetPublicRepos:output_type -> PublicReposResponse
	1, // 14: RepoService.UpdateRepo:output_type -> PrivateReposResponse
	7, // 15: RepoService.DeleteRepo:output_type -> google.protobuf.Empty
	9, // [9:16] is the sub-list for method output_type
	2, // [2:9] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_repo_service_proto_init() }
func file_repo_service_proto_init() {
	if File_repo_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_repo_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrivateRepoRequest); i {
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
		file_repo_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrivateReposResponse); i {
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
		file_repo_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrivateRepoResponse); i {
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
		file_repo_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublicReposResponse); i {
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
		file_repo_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublicRepoResponse); i {
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
		file_repo_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRepoRequest); i {
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
		file_repo_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdentifierRequest); i {
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
			RawDescriptor: file_repo_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_repo_service_proto_goTypes,
		DependencyIndexes: file_repo_service_proto_depIdxs,
		MessageInfos:      file_repo_service_proto_msgTypes,
	}.Build()
	File_repo_service_proto = out.File
	file_repo_service_proto_rawDesc = nil
	file_repo_service_proto_goTypes = nil
	file_repo_service_proto_depIdxs = nil
}