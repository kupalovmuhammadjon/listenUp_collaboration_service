// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: collaborations.proto

package collaborations

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

type CreateInvite struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PodcastId string `protobuf:"bytes,1,opt,name=podcast_id,json=podcastId,proto3" json:"podcast_id,omitempty"`
	InviterId string `protobuf:"bytes,2,opt,name=inviter_id,json=inviterId,proto3" json:"inviter_id,omitempty"`
	InviteeId string `protobuf:"bytes,3,opt,name=invitee_id,json=inviteeId,proto3" json:"invitee_id,omitempty"`
}

func (x *CreateInvite) Reset() {
	*x = CreateInvite{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collaborations_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateInvite) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateInvite) ProtoMessage() {}

func (x *CreateInvite) ProtoReflect() protoreflect.Message {
	mi := &file_collaborations_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateInvite.ProtoReflect.Descriptor instead.
func (*CreateInvite) Descriptor() ([]byte, []int) {
	return file_collaborations_proto_rawDescGZIP(), []int{0}
}

func (x *CreateInvite) GetPodcastId() string {
	if x != nil {
		return x.PodcastId
	}
	return ""
}

func (x *CreateInvite) GetInviterId() string {
	if x != nil {
		return x.InviterId
	}
	return ""
}

func (x *CreateInvite) GetInviteeId() string {
	if x != nil {
		return x.InviteeId
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
		mi := &file_collaborations_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ID) ProtoMessage() {}

func (x *ID) ProtoReflect() protoreflect.Message {
	mi := &file_collaborations_proto_msgTypes[1]
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
	return file_collaborations_proto_rawDescGZIP(), []int{1}
}

func (x *ID) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CreateCollaboration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Role       string        `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
	Invitation *CreateInvite `protobuf:"bytes,2,opt,name=invitation,proto3" json:"invitation,omitempty"`
}

func (x *CreateCollaboration) Reset() {
	*x = CreateCollaboration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collaborations_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCollaboration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCollaboration) ProtoMessage() {}

func (x *CreateCollaboration) ProtoReflect() protoreflect.Message {
	mi := &file_collaborations_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCollaboration.ProtoReflect.Descriptor instead.
func (*CreateCollaboration) Descriptor() ([]byte, []int) {
	return file_collaborations_proto_rawDescGZIP(), []int{2}
}

func (x *CreateCollaboration) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *CreateCollaboration) GetInvitation() *CreateInvite {
	if x != nil {
		return x.Invitation
	}
	return nil
}

type Collaborator struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Email    string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Role     string `protobuf:"bytes,3,opt,name=role,proto3" json:"role,omitempty"`
}

func (x *Collaborator) Reset() {
	*x = Collaborator{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collaborations_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Collaborator) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Collaborator) ProtoMessage() {}

func (x *Collaborator) ProtoReflect() protoreflect.Message {
	mi := &file_collaborations_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Collaborator.ProtoReflect.Descriptor instead.
func (*Collaborator) Descriptor() ([]byte, []int) {
	return file_collaborations_proto_rawDescGZIP(), []int{3}
}

func (x *Collaborator) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Collaborator) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Collaborator) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

type Void struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Void) Reset() {
	*x = Void{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collaborations_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Void) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Void) ProtoMessage() {}

func (x *Void) ProtoReflect() protoreflect.Message {
	mi := &file_collaborations_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Void.ProtoReflect.Descriptor instead.
func (*Void) Descriptor() ([]byte, []int) {
	return file_collaborations_proto_rawDescGZIP(), []int{4}
}

var File_collaborations_proto protoreflect.FileDescriptor

var file_collaborations_proto_rawDesc = []byte{
	0x0a, 0x14, 0x63, 0x6f, 0x6c, 0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6b, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x6f, 0x64, 0x63, 0x61, 0x73,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x6f, 0x64, 0x63,
	0x61, 0x73, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6e, 0x76, 0x69, 0x74,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65,
	0x65, 0x49, 0x64, 0x22, 0x14, 0x0a, 0x02, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x58, 0x0a, 0x13, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x72, 0x6f, 0x6c, 0x65, 0x12, 0x2d, 0x0a, 0x0a, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x52, 0x0a, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x54, 0x0a, 0x0c, 0x43, 0x6f, 0x6c, 0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x22, 0x06, 0x0a, 0x04, 0x56, 0x6f, 0x69,
	0x64, 0x32, 0xf5, 0x01, 0x0a, 0x0e, 0x43, 0x6f, 0x6c, 0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x26, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e,
	0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0d, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x1a, 0x03, 0x2e, 0x49, 0x44, 0x12, 0x2e, 0x0a, 0x11,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x14, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x61, 0x62,
	0x6f, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x03, 0x2e, 0x49, 0x44, 0x12, 0x31, 0x0a, 0x1b,
	0x47, 0x65, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x73,
	0x42, 0x79, 0x50, 0x6f, 0x64, 0x63, 0x61, 0x73, 0x74, 0x49, 0x64, 0x12, 0x03, 0x2e, 0x49, 0x44,
	0x1a, 0x0d, 0x2e, 0x43, 0x6f, 0x6c, 0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12,
	0x2b, 0x0a, 0x1d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x61, 0x62, 0x6f,
	0x72, 0x61, 0x74, 0x6f, 0x72, 0x42, 0x79, 0x50, 0x6f, 0x64, 0x63, 0x61, 0x73, 0x74, 0x49, 0x64,
	0x12, 0x03, 0x2e, 0x49, 0x44, 0x1a, 0x05, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12, 0x2b, 0x0a, 0x1d,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x42, 0x79, 0x50, 0x6f, 0x64, 0x63, 0x61, 0x73, 0x74, 0x49, 0x64, 0x12, 0x03, 0x2e,
	0x49, 0x44, 0x1a, 0x05, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x42, 0x19, 0x5a, 0x17, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6c, 0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_collaborations_proto_rawDescOnce sync.Once
	file_collaborations_proto_rawDescData = file_collaborations_proto_rawDesc
)

func file_collaborations_proto_rawDescGZIP() []byte {
	file_collaborations_proto_rawDescOnce.Do(func() {
		file_collaborations_proto_rawDescData = protoimpl.X.CompressGZIP(file_collaborations_proto_rawDescData)
	})
	return file_collaborations_proto_rawDescData
}

var file_collaborations_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_collaborations_proto_goTypes = []interface{}{
	(*CreateInvite)(nil),        // 0: CreateInvite
	(*ID)(nil),                  // 1: ID
	(*CreateCollaboration)(nil), // 2: CreateCollaboration
	(*Collaborator)(nil),        // 3: Collaborator
	(*Void)(nil),                // 4: Void
}
var file_collaborations_proto_depIdxs = []int32{
	0, // 0: CreateCollaboration.invitation:type_name -> CreateInvite
	0, // 1: Collaborations.CreateInvitation:input_type -> CreateInvite
	2, // 2: Collaborations.RespondInvitation:input_type -> CreateCollaboration
	1, // 3: Collaborations.GetCollaboratorsByPodcastId:input_type -> ID
	1, // 4: Collaborations.UpdateCollaboratorByPodcastId:input_type -> ID
	1, // 5: Collaborations.DeleteCollaboratorByPodcastId:input_type -> ID
	1, // 6: Collaborations.CreateInvitation:output_type -> ID
	1, // 7: Collaborations.RespondInvitation:output_type -> ID
	3, // 8: Collaborations.GetCollaboratorsByPodcastId:output_type -> Collaborator
	4, // 9: Collaborations.UpdateCollaboratorByPodcastId:output_type -> Void
	4, // 10: Collaborations.DeleteCollaboratorByPodcastId:output_type -> Void
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_collaborations_proto_init() }
func file_collaborations_proto_init() {
	if File_collaborations_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_collaborations_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateInvite); i {
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
		file_collaborations_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_collaborations_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCollaboration); i {
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
		file_collaborations_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Collaborator); i {
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
		file_collaborations_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Void); i {
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
			RawDescriptor: file_collaborations_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_collaborations_proto_goTypes,
		DependencyIndexes: file_collaborations_proto_depIdxs,
		MessageInfos:      file_collaborations_proto_msgTypes,
	}.Build()
	File_collaborations_proto = out.File
	file_collaborations_proto_rawDesc = nil
	file_collaborations_proto_goTypes = nil
	file_collaborations_proto_depIdxs = nil
}
