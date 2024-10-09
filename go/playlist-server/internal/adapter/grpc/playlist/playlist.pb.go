// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: proto/playlist.proto

package playlist

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

type Operation int32

const (
	Operation_ADD    Operation = 0
	Operation_REMOVE Operation = 1
	Operation_GET    Operation = 2
)

// Enum value maps for Operation.
var (
	Operation_name = map[int32]string{
		0: "ADD",
		1: "REMOVE",
		2: "GET",
	}
	Operation_value = map[string]int32{
		"ADD":    0,
		"REMOVE": 1,
		"GET":    2,
	}
)

func (x Operation) Enum() *Operation {
	p := new(Operation)
	*p = x
	return p
}

func (x Operation) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Operation) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_playlist_proto_enumTypes[0].Descriptor()
}

func (Operation) Type() protoreflect.EnumType {
	return &file_proto_playlist_proto_enumTypes[0]
}

func (x Operation) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Operation.Descriptor instead.
func (Operation) EnumDescriptor() ([]byte, []int) {
	return file_proto_playlist_proto_rawDescGZIP(), []int{0}
}

// Content message
type Content struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title    string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Genre    string `protobuf:"bytes,3,opt,name=genre,proto3" json:"genre,omitempty"`
	Creator  string `protobuf:"bytes,4,opt,name=creator,proto3" json:"creator,omitempty"` // Artist or Director
	Duration string `protobuf:"bytes,5,opt,name=duration,proto3" json:"duration,omitempty"`
}

func (x *Content) Reset() {
	*x = Content{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_playlist_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Content) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Content) ProtoMessage() {}

func (x *Content) ProtoReflect() protoreflect.Message {
	mi := &file_proto_playlist_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Content.ProtoReflect.Descriptor instead.
func (*Content) Descriptor() ([]byte, []int) {
	return file_proto_playlist_proto_rawDescGZIP(), []int{0}
}

func (x *Content) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Content) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Content) GetGenre() string {
	if x != nil {
		return x.Genre
	}
	return ""
}

func (x *Content) GetCreator() string {
	if x != nil {
		return x.Creator
	}
	return ""
}

func (x *Content) GetDuration() string {
	if x != nil {
		return x.Duration
	}
	return ""
}

// Playlist request message
type PlaylistRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserID    string     `protobuf:"bytes,2,opt,name=userID,json=user_id,proto3" json:"userID,omitempty"`
	Name      string     `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Contents  []*Content `protobuf:"bytes,4,rep,name=contents,proto3" json:"contents,omitempty"`
	Operation Operation  `protobuf:"varint,5,opt,name=operation,proto3,enum=Operation" json:"operation,omitempty"`
}

func (x *PlaylistRequest) Reset() {
	*x = PlaylistRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_playlist_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlaylistRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlaylistRequest) ProtoMessage() {}

func (x *PlaylistRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_playlist_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlaylistRequest.ProtoReflect.Descriptor instead.
func (*PlaylistRequest) Descriptor() ([]byte, []int) {
	return file_proto_playlist_proto_rawDescGZIP(), []int{1}
}

func (x *PlaylistRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PlaylistRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *PlaylistRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PlaylistRequest) GetContents() []*Content {
	if x != nil {
		return x.Contents
	}
	return nil
}

func (x *PlaylistRequest) GetOperation() Operation {
	if x != nil {
		return x.Operation
	}
	return Operation_ADD
}

// List playlist request message
type ListPlaylistRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID string `protobuf:"bytes,1,opt,name=userID,json=user_id,proto3" json:"userID,omitempty"`
}

func (x *ListPlaylistRequest) Reset() {
	*x = ListPlaylistRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_playlist_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPlaylistRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPlaylistRequest) ProtoMessage() {}

func (x *ListPlaylistRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_playlist_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPlaylistRequest.ProtoReflect.Descriptor instead.
func (*ListPlaylistRequest) Descriptor() ([]byte, []int) {
	return file_proto_playlist_proto_rawDescGZIP(), []int{2}
}

func (x *ListPlaylistRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

// Remove content request message
type RemovePlaylistRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserID string `protobuf:"bytes,2,opt,name=userID,json=user_id,proto3" json:"userID,omitempty"`
}

func (x *RemovePlaylistRequest) Reset() {
	*x = RemovePlaylistRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_playlist_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemovePlaylistRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemovePlaylistRequest) ProtoMessage() {}

func (x *RemovePlaylistRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_playlist_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemovePlaylistRequest.ProtoReflect.Descriptor instead.
func (*RemovePlaylistRequest) Descriptor() ([]byte, []int) {
	return file_proto_playlist_proto_rawDescGZIP(), []int{3}
}

func (x *RemovePlaylistRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RemovePlaylistRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

// Playlist response message
type PlaylistResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string     `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Contents []*Content `protobuf:"bytes,3,rep,name=contents,proto3" json:"contents,omitempty"`
}

func (x *PlaylistResponse) Reset() {
	*x = PlaylistResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_playlist_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlaylistResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlaylistResponse) ProtoMessage() {}

func (x *PlaylistResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_playlist_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlaylistResponse.ProtoReflect.Descriptor instead.
func (*PlaylistResponse) Descriptor() ([]byte, []int) {
	return file_proto_playlist_proto_rawDescGZIP(), []int{4}
}

func (x *PlaylistResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PlaylistResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PlaylistResponse) GetContents() []*Content {
	if x != nil {
		return x.Contents
	}
	return nil
}

type ListPlaylistResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*ListPlaylistResponse_Playlist `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *ListPlaylistResponse) Reset() {
	*x = ListPlaylistResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_playlist_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPlaylistResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPlaylistResponse) ProtoMessage() {}

func (x *ListPlaylistResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_playlist_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPlaylistResponse.ProtoReflect.Descriptor instead.
func (*ListPlaylistResponse) Descriptor() ([]byte, []int) {
	return file_proto_playlist_proto_rawDescGZIP(), []int{5}
}

func (x *ListPlaylistResponse) GetData() []*ListPlaylistResponse_Playlist {
	if x != nil {
		return x.Data
	}
	return nil
}

type RemovePlaylistResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *RemovePlaylistResponse) Reset() {
	*x = RemovePlaylistResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_playlist_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemovePlaylistResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemovePlaylistResponse) ProtoMessage() {}

func (x *RemovePlaylistResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_playlist_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemovePlaylistResponse.ProtoReflect.Descriptor instead.
func (*RemovePlaylistResponse) Descriptor() ([]byte, []int) {
	return file_proto_playlist_proto_rawDescGZIP(), []int{6}
}

func (x *RemovePlaylistResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ListPlaylistResponse_Playlist struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ListPlaylistResponse_Playlist) Reset() {
	*x = ListPlaylistResponse_Playlist{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_playlist_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPlaylistResponse_Playlist) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPlaylistResponse_Playlist) ProtoMessage() {}

func (x *ListPlaylistResponse_Playlist) ProtoReflect() protoreflect.Message {
	mi := &file_proto_playlist_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPlaylistResponse_Playlist.ProtoReflect.Descriptor instead.
func (*ListPlaylistResponse_Playlist) Descriptor() ([]byte, []int) {
	return file_proto_playlist_proto_rawDescGZIP(), []int{5, 0}
}

func (x *ListPlaylistResponse_Playlist) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ListPlaylistResponse_Playlist) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_proto_playlist_proto protoreflect.FileDescriptor

var file_proto_playlist_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7b, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x65, 0x6e, 0x72, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x9e, 0x01, 0x0a, 0x0f, 0x50, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x52, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x28, 0x0a, 0x09, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0a, 0x2e,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x2e, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6c, 0x61, 0x79,
	0x6c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x22, 0x40, 0x0a, 0x15, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x6c,
	0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x22, 0x5c, 0x0a, 0x10, 0x50, 0x6c, 0x61, 0x79, 0x6c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x24,
	0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x08, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x73, 0x22, 0x7a, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6c, 0x61, 0x79,
	0x6c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x50, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x1a, 0x2e, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x32, 0x0a, 0x16, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x6c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2a, 0x29, 0x0a, 0x09, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x44, 0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x52, 0x45,
	0x4d, 0x4f, 0x56, 0x45, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x47, 0x45, 0x54, 0x10, 0x02, 0x32,
	0xb4, 0x01, 0x0a, 0x0f, 0x50, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x31, 0x0a, 0x06, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x12, 0x10, 0x2e,
	0x50, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x11, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x28, 0x01, 0x30, 0x01, 0x12, 0x33, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x14,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x6c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x06, 0x52,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x12, 0x16, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x6c,
	0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e,
	0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x70, 0x6c, 0x61, 0x79,
	0x6c, 0x69, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_playlist_proto_rawDescOnce sync.Once
	file_proto_playlist_proto_rawDescData = file_proto_playlist_proto_rawDesc
)

func file_proto_playlist_proto_rawDescGZIP() []byte {
	file_proto_playlist_proto_rawDescOnce.Do(func() {
		file_proto_playlist_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_playlist_proto_rawDescData)
	})
	return file_proto_playlist_proto_rawDescData
}

var file_proto_playlist_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_playlist_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_playlist_proto_goTypes = []interface{}{
	(Operation)(0),                        // 0: Operation
	(*Content)(nil),                       // 1: Content
	(*PlaylistRequest)(nil),               // 2: PlaylistRequest
	(*ListPlaylistRequest)(nil),           // 3: ListPlaylistRequest
	(*RemovePlaylistRequest)(nil),         // 4: RemovePlaylistRequest
	(*PlaylistResponse)(nil),              // 5: PlaylistResponse
	(*ListPlaylistResponse)(nil),          // 6: ListPlaylistResponse
	(*RemovePlaylistResponse)(nil),        // 7: RemovePlaylistResponse
	(*ListPlaylistResponse_Playlist)(nil), // 8: ListPlaylistResponse.Playlist
}
var file_proto_playlist_proto_depIdxs = []int32{
	1, // 0: PlaylistRequest.contents:type_name -> Content
	0, // 1: PlaylistRequest.operation:type_name -> Operation
	1, // 2: PlaylistResponse.contents:type_name -> Content
	8, // 3: ListPlaylistResponse.data:type_name -> ListPlaylistResponse.Playlist
	2, // 4: PlaylistService.Manage:input_type -> PlaylistRequest
	3, // 5: PlaylistService.List:input_type -> ListPlaylistRequest
	4, // 6: PlaylistService.Remove:input_type -> RemovePlaylistRequest
	5, // 7: PlaylistService.Manage:output_type -> PlaylistResponse
	6, // 8: PlaylistService.List:output_type -> ListPlaylistResponse
	7, // 9: PlaylistService.Remove:output_type -> RemovePlaylistResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_playlist_proto_init() }
func file_proto_playlist_proto_init() {
	if File_proto_playlist_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_playlist_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Content); i {
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
		file_proto_playlist_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlaylistRequest); i {
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
		file_proto_playlist_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPlaylistRequest); i {
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
		file_proto_playlist_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemovePlaylistRequest); i {
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
		file_proto_playlist_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlaylistResponse); i {
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
		file_proto_playlist_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPlaylistResponse); i {
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
		file_proto_playlist_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemovePlaylistResponse); i {
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
		file_proto_playlist_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPlaylistResponse_Playlist); i {
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
			RawDescriptor: file_proto_playlist_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_playlist_proto_goTypes,
		DependencyIndexes: file_proto_playlist_proto_depIdxs,
		EnumInfos:         file_proto_playlist_proto_enumTypes,
		MessageInfos:      file_proto_playlist_proto_msgTypes,
	}.Build()
	File_proto_playlist_proto = out.File
	file_proto_playlist_proto_rawDesc = nil
	file_proto_playlist_proto_goTypes = nil
	file_proto_playlist_proto_depIdxs = nil
}
