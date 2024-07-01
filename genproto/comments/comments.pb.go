// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: comments.proto

package comments

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

type CreateComment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PodcastId string `protobuf:"bytes,1,opt,name=podcast_id,json=podcastId,proto3" json:"podcast_id,omitempty"`
	UserId    string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Content   string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *CreateComment) Reset() {
	*x = CreateComment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comments_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateComment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateComment) ProtoMessage() {}

func (x *CreateComment) ProtoReflect() protoreflect.Message {
	mi := &file_comments_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateComment.ProtoReflect.Descriptor instead.
func (*CreateComment) Descriptor() ([]byte, []int) {
	return file_comments_proto_rawDescGZIP(), []int{0}
}

func (x *CreateComment) GetPodcastId() string {
	if x != nil {
		return x.PodcastId
	}
	return ""
}

func (x *CreateComment) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreateComment) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type EpisodeComment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EpisodeId string `protobuf:"bytes,1,opt,name=episode_id,json=episodeId,proto3" json:"episode_id,omitempty"`
	UserId    string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Content   string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *EpisodeComment) Reset() {
	*x = EpisodeComment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comments_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EpisodeComment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EpisodeComment) ProtoMessage() {}

func (x *EpisodeComment) ProtoReflect() protoreflect.Message {
	mi := &file_comments_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EpisodeComment.ProtoReflect.Descriptor instead.
func (*EpisodeComment) Descriptor() ([]byte, []int) {
	return file_comments_proto_rawDescGZIP(), []int{1}
}

func (x *EpisodeComment) GetEpisodeId() string {
	if x != nil {
		return x.EpisodeId
	}
	return ""
}

func (x *EpisodeComment) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *EpisodeComment) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type ID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ID) Reset() {
	*x = ID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comments_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ID) ProtoMessage() {}

func (x *ID) ProtoReflect() protoreflect.Message {
	mi := &file_comments_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ID.ProtoReflect.Descriptor instead.
func (*ID) Descriptor() ([]byte, []int) {
	return file_comments_proto_rawDescGZIP(), []int{2}
}

func (x *ID) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CommentInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    string `protobuf:"bytes,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	Content   string `protobuf:"bytes,2,opt,name=Content,proto3" json:"Content,omitempty"`
	CreatedAt string `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt string `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *CommentInfo) Reset() {
	*x = CommentInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comments_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentInfo) ProtoMessage() {}

func (x *CommentInfo) ProtoReflect() protoreflect.Message {
	mi := &file_comments_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentInfo.ProtoReflect.Descriptor instead.
func (*CommentInfo) Descriptor() ([]byte, []int) {
	return file_comments_proto_rawDescGZIP(), []int{3}
}

func (x *CommentInfo) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CommentInfo) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *CommentInfo) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *CommentInfo) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type Comment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username  string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Content   string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	CreatedAt string `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt string `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Comment) Reset() {
	*x = Comment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comments_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Comment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Comment) ProtoMessage() {}

func (x *Comment) ProtoReflect() protoreflect.Message {
	mi := &file_comments_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Comment.ProtoReflect.Descriptor instead.
func (*Comment) Descriptor() ([]byte, []int) {
	return file_comments_proto_rawDescGZIP(), []int{4}
}

func (x *Comment) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Comment) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Comment) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Comment) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type AllComments struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Comments []*Comment `protobuf:"bytes,1,rep,name=comments,proto3" json:"comments,omitempty"`
}

func (x *AllComments) Reset() {
	*x = AllComments{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comments_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllComments) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllComments) ProtoMessage() {}

func (x *AllComments) ProtoReflect() protoreflect.Message {
	mi := &file_comments_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllComments.ProtoReflect.Descriptor instead.
func (*AllComments) Descriptor() ([]byte, []int) {
	return file_comments_proto_rawDescGZIP(), []int{5}
}

func (x *AllComments) GetComments() []*Comment {
	if x != nil {
		return x.Comments
	}
	return nil
}

type CommentCount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int64 `protobuf:"varint,1,opt,name=Count,proto3" json:"Count,omitempty"`
}

func (x *CommentCount) Reset() {
	*x = CommentCount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comments_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentCount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentCount) ProtoMessage() {}

func (x *CommentCount) ProtoReflect() protoreflect.Message {
	mi := &file_comments_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentCount.ProtoReflect.Descriptor instead.
func (*CommentCount) Descriptor() ([]byte, []int) {
	return file_comments_proto_rawDescGZIP(), []int{6}
}

func (x *CommentCount) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type CountFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EpisodeId string `protobuf:"bytes,1,opt,name=EpisodeId,proto3" json:"EpisodeId,omitempty"`
	PodcastId string `protobuf:"bytes,2,opt,name=PodcastId,proto3" json:"PodcastId,omitempty"`
}

func (x *CountFilter) Reset() {
	*x = CountFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comments_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountFilter) ProtoMessage() {}

func (x *CountFilter) ProtoReflect() protoreflect.Message {
	mi := &file_comments_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountFilter.ProtoReflect.Descriptor instead.
func (*CountFilter) Descriptor() ([]byte, []int) {
	return file_comments_proto_rawDescGZIP(), []int{7}
}

func (x *CountFilter) GetEpisodeId() string {
	if x != nil {
		return x.EpisodeId
	}
	return ""
}

func (x *CountFilter) GetPodcastId() string {
	if x != nil {
		return x.PodcastId
	}
	return ""
}

var File_comments_proto protoreflect.FileDescriptor

var file_comments_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x61, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x6f, 0x64, 0x63, 0x61, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x6f, 0x64, 0x63, 0x61, 0x73, 0x74, 0x49, 0x64,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x22, 0x62, 0x0a, 0x0e, 0x45, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x70, 0x69, 0x73, 0x6f,
	0x64, 0x65, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x14, 0x0a, 0x02, 0x49, 0x44, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x7d, 0x0a,
	0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x0a, 0x06,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x7d, 0x0a, 0x07,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x33, 0x0a, 0x0b, 0x41,
	0x6c, 0x6c, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x24, 0x0a, 0x08, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x43,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x22, 0x24, 0x0a, 0x0c, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x49, 0x0a, 0x0b, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x46,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x45, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x45, 0x70, 0x69, 0x73, 0x6f, 0x64,
	0x65, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x6f, 0x64, 0x63, 0x61, 0x73, 0x74, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x50, 0x6f, 0x64, 0x63, 0x61, 0x73, 0x74, 0x49,
	0x64, 0x32, 0xf5, 0x01, 0x0a, 0x08, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x2f,
	0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x42,
	0x79, 0x50, 0x6f, 0x64, 0x63, 0x61, 0x73, 0x74, 0x49, 0x64, 0x12, 0x0e, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x03, 0x2e, 0x49, 0x44, 0x12,
	0x2b, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x79,
	0x50, 0x6f, 0x64, 0x63, 0x61, 0x73, 0x74, 0x49, 0x64, 0x12, 0x03, 0x2e, 0x49, 0x44, 0x1a, 0x0c,
	0x2e, 0x41, 0x6c, 0x6c, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x30, 0x0a, 0x18,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x45,
	0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x0f, 0x2e, 0x45, 0x70, 0x69, 0x73, 0x6f,
	0x64, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x03, 0x2e, 0x49, 0x44, 0x12, 0x2b,
	0x0a, 0x16, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x79, 0x45,
	0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x03, 0x2e, 0x49, 0x44, 0x1a, 0x0c, 0x2e,
	0x41, 0x6c, 0x6c, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x2c, 0x0a, 0x0d, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x0c, 0x2e, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x1a, 0x0d, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x13, 0x5a, 0x11, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_comments_proto_rawDescOnce sync.Once
	file_comments_proto_rawDescData = file_comments_proto_rawDesc
)

func file_comments_proto_rawDescGZIP() []byte {
	file_comments_proto_rawDescOnce.Do(func() {
		file_comments_proto_rawDescData = protoimpl.X.CompressGZIP(file_comments_proto_rawDescData)
	})
	return file_comments_proto_rawDescData
}

var file_comments_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_comments_proto_goTypes = []interface{}{
	(*CreateComment)(nil),  // 0: CreateComment
	(*EpisodeComment)(nil), // 1: EpisodeComment
	(*ID)(nil),             // 2: ID
	(*CommentInfo)(nil),    // 3: CommentInfo
	(*Comment)(nil),        // 4: Comment
	(*AllComments)(nil),    // 5: AllComments
	(*CommentCount)(nil),   // 6: CommentCount
	(*CountFilter)(nil),    // 7: CountFilter
}
var file_comments_proto_depIdxs = []int32{
	4, // 0: AllComments.comments:type_name -> Comment
	0, // 1: Comments.CreateCommentByPodcastId:input_type -> CreateComment
	2, // 2: Comments.GetCommentsByPodcastId:input_type -> ID
	1, // 3: Comments.CreateCommentByEpisodeId:input_type -> EpisodeComment
	2, // 4: Comments.GetCommentsByEpisodeId:input_type -> ID
	7, // 5: Comments.CountComments:input_type -> CountFilter
	2, // 6: Comments.CreateCommentByPodcastId:output_type -> ID
	5, // 7: Comments.GetCommentsByPodcastId:output_type -> AllComments
	2, // 8: Comments.CreateCommentByEpisodeId:output_type -> ID
	5, // 9: Comments.GetCommentsByEpisodeId:output_type -> AllComments
	6, // 10: Comments.CountComments:output_type -> CommentCount
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_comments_proto_init() }
func file_comments_proto_init() {
	if File_comments_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_comments_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateComment); i {
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
		file_comments_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EpisodeComment); i {
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
		file_comments_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ID); i {
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
		file_comments_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentInfo); i {
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
		file_comments_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Comment); i {
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
		file_comments_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllComments); i {
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
		file_comments_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentCount); i {
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
		file_comments_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountFilter); i {
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
			RawDescriptor: file_comments_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_comments_proto_goTypes,
		DependencyIndexes: file_comments_proto_depIdxs,
		MessageInfos:      file_comments_proto_msgTypes,
	}.Build()
	File_comments_proto = out.File
	file_comments_proto_rawDesc = nil
	file_comments_proto_goTypes = nil
	file_comments_proto_depIdxs = nil
}
