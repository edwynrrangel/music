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

type Content struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Order int32  `protobuf:"varint,2,opt,name=order,proto3" json:"order,omitempty"`
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

func (x *Content) GetOrder() int32 {
	if x != nil {
		return x.Order
	}
	return 0
}

type Playlist struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string     `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Contents []*Content `protobuf:"bytes,3,rep,name=contents,proto3" json:"contents,omitempty"`
}

func (x *Playlist) Reset() {
	*x = Playlist{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_playlist_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Playlist) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Playlist) ProtoMessage() {}

func (x *Playlist) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Playlist.ProtoReflect.Descriptor instead.
func (*Playlist) Descriptor() ([]byte, []int) {
	return file_proto_playlist_proto_rawDescGZIP(), []int{1}
}

func (x *Playlist) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Playlist) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Playlist) GetContents() []*Content {
	if x != nil {
		return x.Contents
	}
	return nil
}

type CreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Mode   string `protobuf:"bytes,3,opt,name=mode,proto3" json:"mode,omitempty"`
}

func (x *CreateRequest) Reset() {
	*x = CreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_playlist_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest) ProtoMessage() {}

func (x *CreateRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use CreateRequest.ProtoReflect.Descriptor instead.
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return file_proto_playlist_proto_rawDescGZIP(), []int{2}
}

func (x *CreateRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateRequest) GetMode() string {
	if x != nil {
		return x.Mode
	}
	return ""
}

type ListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Query  string `protobuf:"bytes,2,opt,name=query,proto3" json:"query,omitempty"`
	Page   int32  `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	Limit  int32  `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *ListRequest) Reset() {
	*x = ListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_playlist_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest) ProtoMessage() {}

func (x *ListRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ListRequest.ProtoReflect.Descriptor instead.
func (*ListRequest) Descriptor() ([]byte, []int) {
	return file_proto_playlist_proto_rawDescGZIP(), []int{3}
}

func (x *ListRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ListRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *ListRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type ListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page  int32       `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Limit int32       `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Total int64       `protobuf:"varint,3,opt,name=total,proto3" json:"total,omitempty"`
	Data  []*Playlist `protobuf:"bytes,4,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *ListResponse) Reset() {
	*x = ListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_playlist_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse) ProtoMessage() {}

func (x *ListResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ListResponse.ProtoReflect.Descriptor instead.
func (*ListResponse) Descriptor() ([]byte, []int) {
	return file_proto_playlist_proto_rawDescGZIP(), []int{4}
}

func (x *ListResponse) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListResponse) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ListResponse) GetData() []*Playlist {
	if x != nil {
		return x.Data
	}
	return nil
}

type PlaylistRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	PlaylistId string `protobuf:"bytes,2,opt,name=playlist_id,json=playlistId,proto3" json:"playlist_id,omitempty"`
}

func (x *PlaylistRequest) Reset() {
	*x = PlaylistRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_playlist_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlaylistRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlaylistRequest) ProtoMessage() {}

func (x *PlaylistRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use PlaylistRequest.ProtoReflect.Descriptor instead.
func (*PlaylistRequest) Descriptor() ([]byte, []int) {
	return file_proto_playlist_proto_rawDescGZIP(), []int{5}
}

func (x *PlaylistRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *PlaylistRequest) GetPlaylistId() string {
	if x != nil {
		return x.PlaylistId
	}
	return ""
}

type OperationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	PlaylistId string `protobuf:"bytes,2,opt,name=playlist_id,json=playlistId,proto3" json:"playlist_id,omitempty"`
	ContentId  string `protobuf:"bytes,3,opt,name=content_id,json=contentId,proto3" json:"content_id,omitempty"`
	Order      int32  `protobuf:"varint,4,opt,name=order,proto3" json:"order,omitempty"`
}

func (x *OperationRequest) Reset() {
	*x = OperationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_playlist_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperationRequest) ProtoMessage() {}

func (x *OperationRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use OperationRequest.ProtoReflect.Descriptor instead.
func (*OperationRequest) Descriptor() ([]byte, []int) {
	return file_proto_playlist_proto_rawDescGZIP(), []int{6}
}

func (x *OperationRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *OperationRequest) GetPlaylistId() string {
	if x != nil {
		return x.PlaylistId
	}
	return ""
}

func (x *OperationRequest) GetContentId() string {
	if x != nil {
		return x.ContentId
	}
	return ""
}

func (x *OperationRequest) GetOrder() int32 {
	if x != nil {
		return x.Order
	}
	return 0
}

type AddContentEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContentId string `protobuf:"bytes,1,opt,name=content_id,json=contentId,proto3" json:"content_id,omitempty"`
	Order     int32  `protobuf:"varint,2,opt,name=order,proto3" json:"order,omitempty"`
}

func (x *AddContentEvent) Reset() {
	*x = AddContentEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_playlist_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddContentEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddContentEvent) ProtoMessage() {}

func (x *AddContentEvent) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use AddContentEvent.ProtoReflect.Descriptor instead.
func (*AddContentEvent) Descriptor() ([]byte, []int) {
	return file_proto_playlist_proto_rawDescGZIP(), []int{7}
}

func (x *AddContentEvent) GetContentId() string {
	if x != nil {
		return x.ContentId
	}
	return ""
}

func (x *AddContentEvent) GetOrder() int32 {
	if x != nil {
		return x.Order
	}
	return 0
}

type SkipContentEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ForceSkip bool `protobuf:"varint,1,opt,name=force_skip,json=forceSkip,proto3" json:"force_skip,omitempty"`
}

func (x *SkipContentEvent) Reset() {
	*x = SkipContentEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_playlist_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SkipContentEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SkipContentEvent) ProtoMessage() {}

func (x *SkipContentEvent) ProtoReflect() protoreflect.Message {
	mi := &file_proto_playlist_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SkipContentEvent.ProtoReflect.Descriptor instead.
func (*SkipContentEvent) Descriptor() ([]byte, []int) {
	return file_proto_playlist_proto_rawDescGZIP(), []int{8}
}

func (x *SkipContentEvent) GetForceSkip() bool {
	if x != nil {
		return x.ForceSkip
	}
	return false
}

type PartyModeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlaylistId string `protobuf:"bytes,1,opt,name=playlist_id,json=playlistId,proto3" json:"playlist_id,omitempty"`
	// Types that are assignable to Action:
	//
	//	*PartyModeRequest_AddContent
	//	*PartyModeRequest_SkipContent
	Action isPartyModeRequest_Action `protobuf_oneof:"action"`
}

func (x *PartyModeRequest) Reset() {
	*x = PartyModeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_playlist_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PartyModeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartyModeRequest) ProtoMessage() {}

func (x *PartyModeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_playlist_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PartyModeRequest.ProtoReflect.Descriptor instead.
func (*PartyModeRequest) Descriptor() ([]byte, []int) {
	return file_proto_playlist_proto_rawDescGZIP(), []int{9}
}

func (x *PartyModeRequest) GetPlaylistId() string {
	if x != nil {
		return x.PlaylistId
	}
	return ""
}

func (m *PartyModeRequest) GetAction() isPartyModeRequest_Action {
	if m != nil {
		return m.Action
	}
	return nil
}

func (x *PartyModeRequest) GetAddContent() *AddContentEvent {
	if x, ok := x.GetAction().(*PartyModeRequest_AddContent); ok {
		return x.AddContent
	}
	return nil
}

func (x *PartyModeRequest) GetSkipContent() *SkipContentEvent {
	if x, ok := x.GetAction().(*PartyModeRequest_SkipContent); ok {
		return x.SkipContent
	}
	return nil
}

type isPartyModeRequest_Action interface {
	isPartyModeRequest_Action()
}

type PartyModeRequest_AddContent struct {
	AddContent *AddContentEvent `protobuf:"bytes,2,opt,name=add_content,json=addContent,proto3,oneof"`
}

type PartyModeRequest_SkipContent struct {
	SkipContent *SkipContentEvent `protobuf:"bytes,3,opt,name=skip_content,json=skipContent,proto3,oneof"`
}

func (*PartyModeRequest_AddContent) isPartyModeRequest_Action() {}

func (*PartyModeRequest_SkipContent) isPartyModeRequest_Action() {}

type PartyModeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message  string    `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Playlist *Playlist `protobuf:"bytes,2,opt,name=playlist,proto3" json:"playlist,omitempty"`
}

func (x *PartyModeResponse) Reset() {
	*x = PartyModeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_playlist_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PartyModeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartyModeResponse) ProtoMessage() {}

func (x *PartyModeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_playlist_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PartyModeResponse.ProtoReflect.Descriptor instead.
func (*PartyModeResponse) Descriptor() ([]byte, []int) {
	return file_proto_playlist_proto_rawDescGZIP(), []int{10}
}

func (x *PartyModeResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *PartyModeResponse) GetPlaylist() *Playlist {
	if x != nil {
		return x.Playlist
	}
	return nil
}

var File_proto_playlist_proto protoreflect.FileDescriptor

var file_proto_playlist_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74,
	0x22, 0x2f, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x22, 0x5d, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x2d, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73,
	0x22, 0x50, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x6f,
	0x64, 0x65, 0x22, 0x66, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x76, 0x0a, 0x0c, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x26, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x6c,
	0x69, 0x73, 0x74, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x22, 0x4b, 0x0a, 0x0f, 0x50, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f,
	0x0a, 0x0b, 0x70, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x49, 0x64, 0x22,
	0x81, 0x01, 0x0a, 0x10, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a,
	0x0b, 0x70, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x70, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x49, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x22, 0x46, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x22, 0x31, 0x0a, 0x10, 0x53,
	0x6b, 0x69, 0x70, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x5f, 0x73, 0x6b, 0x69, 0x70, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x09, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x53, 0x6b, 0x69, 0x70, 0x22, 0xbc,
	0x01, 0x0a, 0x10, 0x50, 0x61, 0x72, 0x74, 0x79, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6c, 0x61, 0x79, 0x6c, 0x69,
	0x73, 0x74, 0x49, 0x64, 0x12, 0x3c, 0x0a, 0x0b, 0x61, 0x64, 0x64, 0x5f, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x6c, 0x61, 0x79,
	0x6c, 0x69, 0x73, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x0a, 0x61, 0x64, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x12, 0x3f, 0x0a, 0x0c, 0x73, 0x6b, 0x69, 0x70, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x6c,
	0x69, 0x73, 0x74, 0x2e, 0x53, 0x6b, 0x69, 0x70, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x0b, 0x73, 0x6b, 0x69, 0x70, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x42, 0x08, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x5d, 0x0a,
	0x11, 0x50, 0x61, 0x72, 0x74, 0x79, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2e, 0x0a, 0x08,
	0x70, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x70, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x6c, 0x69,
	0x73, 0x74, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x32, 0x82, 0x03, 0x0a,
	0x0f, 0x50, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x35, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x17, 0x2e, 0x70, 0x6c, 0x61,
	0x79, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x50,
	0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x15, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73,
	0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34,
	0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x19, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74,
	0x2e, 0x50, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x12, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x50, 0x6c, 0x61, 0x79,
	0x6c, 0x69, 0x73, 0x74, 0x12, 0x3e, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x12, 0x1a, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12,
	0x2e, 0x70, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x6c, 0x69,
	0x73, 0x74, 0x28, 0x01, 0x12, 0x41, 0x0a, 0x0d, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74,
	0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x12, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x50, 0x6c, 0x61,
	0x79, 0x6c, 0x69, 0x73, 0x74, 0x28, 0x01, 0x12, 0x48, 0x0a, 0x09, 0x50, 0x61, 0x72, 0x74, 0x79,
	0x4d, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x2e,
	0x50, 0x61, 0x72, 0x74, 0x79, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1b, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x50, 0x61, 0x72, 0x74,
	0x79, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30,
	0x01, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x70, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_proto_playlist_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_proto_playlist_proto_goTypes = []interface{}{
	(*Content)(nil),           // 0: playlist.Content
	(*Playlist)(nil),          // 1: playlist.Playlist
	(*CreateRequest)(nil),     // 2: playlist.CreateRequest
	(*ListRequest)(nil),       // 3: playlist.ListRequest
	(*ListResponse)(nil),      // 4: playlist.ListResponse
	(*PlaylistRequest)(nil),   // 5: playlist.PlaylistRequest
	(*OperationRequest)(nil),  // 6: playlist.OperationRequest
	(*AddContentEvent)(nil),   // 7: playlist.AddContentEvent
	(*SkipContentEvent)(nil),  // 8: playlist.SkipContentEvent
	(*PartyModeRequest)(nil),  // 9: playlist.PartyModeRequest
	(*PartyModeResponse)(nil), // 10: playlist.PartyModeResponse
}
var file_proto_playlist_proto_depIdxs = []int32{
	0,  // 0: playlist.Playlist.contents:type_name -> playlist.Content
	1,  // 1: playlist.ListResponse.data:type_name -> playlist.Playlist
	7,  // 2: playlist.PartyModeRequest.add_content:type_name -> playlist.AddContentEvent
	8,  // 3: playlist.PartyModeRequest.skip_content:type_name -> playlist.SkipContentEvent
	1,  // 4: playlist.PartyModeResponse.playlist:type_name -> playlist.Playlist
	2,  // 5: playlist.PlaylistService.Create:input_type -> playlist.CreateRequest
	3,  // 6: playlist.PlaylistService.List:input_type -> playlist.ListRequest
	5,  // 7: playlist.PlaylistService.Get:input_type -> playlist.PlaylistRequest
	6,  // 8: playlist.PlaylistService.AddContent:input_type -> playlist.OperationRequest
	6,  // 9: playlist.PlaylistService.RemoveContent:input_type -> playlist.OperationRequest
	9,  // 10: playlist.PlaylistService.PartyMode:input_type -> playlist.PartyModeRequest
	1,  // 11: playlist.PlaylistService.Create:output_type -> playlist.Playlist
	4,  // 12: playlist.PlaylistService.List:output_type -> playlist.ListResponse
	1,  // 13: playlist.PlaylistService.Get:output_type -> playlist.Playlist
	1,  // 14: playlist.PlaylistService.AddContent:output_type -> playlist.Playlist
	1,  // 15: playlist.PlaylistService.RemoveContent:output_type -> playlist.Playlist
	10, // 16: playlist.PlaylistService.PartyMode:output_type -> playlist.PartyModeResponse
	11, // [11:17] is the sub-list for method output_type
	5,  // [5:11] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
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
			switch v := v.(*Playlist); i {
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
			switch v := v.(*CreateRequest); i {
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
			switch v := v.(*ListRequest); i {
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
			switch v := v.(*ListResponse); i {
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
		file_proto_playlist_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperationRequest); i {
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
			switch v := v.(*AddContentEvent); i {
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
		file_proto_playlist_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SkipContentEvent); i {
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
		file_proto_playlist_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PartyModeRequest); i {
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
		file_proto_playlist_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PartyModeResponse); i {
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
	file_proto_playlist_proto_msgTypes[9].OneofWrappers = []interface{}{
		(*PartyModeRequest_AddContent)(nil),
		(*PartyModeRequest_SkipContent)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_playlist_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_playlist_proto_goTypes,
		DependencyIndexes: file_proto_playlist_proto_depIdxs,
		MessageInfos:      file_proto_playlist_proto_msgTypes,
	}.Build()
	File_proto_playlist_proto = out.File
	file_proto_playlist_proto_rawDesc = nil
	file_proto_playlist_proto_goTypes = nil
	file_proto_playlist_proto_depIdxs = nil
}
