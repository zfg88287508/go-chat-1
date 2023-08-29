// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.9
// source: web/v1/contact.proto

package web

import (
	_ "github.com/srikrsna/protoc-gen-gotag/tagger"
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

// 联系人列表接口请求参数
type ContactListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ContactListRequest) Reset() {
	*x = ContactListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_v1_contact_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContactListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContactListRequest) ProtoMessage() {}

func (x *ContactListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_web_v1_contact_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContactListRequest.ProtoReflect.Descriptor instead.
func (*ContactListRequest) Descriptor() ([]byte, []int) {
	return file_web_v1_contact_proto_rawDescGZIP(), []int{0}
}

// 联系人列表接口响应参数
type ContactListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*ContactListResponse_Item `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ContactListResponse) Reset() {
	*x = ContactListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_v1_contact_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContactListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContactListResponse) ProtoMessage() {}

func (x *ContactListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_web_v1_contact_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContactListResponse.ProtoReflect.Descriptor instead.
func (*ContactListResponse) Descriptor() ([]byte, []int) {
	return file_web_v1_contact_proto_rawDescGZIP(), []int{1}
}

func (x *ContactListResponse) GetItems() []*ContactListResponse_Item {
	if x != nil {
		return x.Items
	}
	return nil
}

// 联系人删除接口请求参数
type ContactDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FriendId int32 `protobuf:"varint,1,opt,name=friend_id,json=friendId,proto3" json:"friend_id,omitempty" binding:"required"`
}

func (x *ContactDeleteRequest) Reset() {
	*x = ContactDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_v1_contact_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContactDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContactDeleteRequest) ProtoMessage() {}

func (x *ContactDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_web_v1_contact_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContactDeleteRequest.ProtoReflect.Descriptor instead.
func (*ContactDeleteRequest) Descriptor() ([]byte, []int) {
	return file_web_v1_contact_proto_rawDescGZIP(), []int{2}
}

func (x *ContactDeleteRequest) GetFriendId() int32 {
	if x != nil {
		return x.FriendId
	}
	return 0
}

// 联系人删除接口响应参数
type ContactDeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ContactDeleteResponse) Reset() {
	*x = ContactDeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_v1_contact_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContactDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContactDeleteResponse) ProtoMessage() {}

func (x *ContactDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_web_v1_contact_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContactDeleteResponse.ProtoReflect.Descriptor instead.
func (*ContactDeleteResponse) Descriptor() ([]byte, []int) {
	return file_web_v1_contact_proto_rawDescGZIP(), []int{3}
}

// 联系人备注修改接口请求参数
type ContactEditRemarkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FriendId int32  `protobuf:"varint,1,opt,name=friend_id,json=friendId,proto3" json:"friend_id,omitempty" binding:"required"`
	Remark   string `protobuf:"bytes,2,opt,name=remark,proto3" json:"remark,omitempty"`
}

func (x *ContactEditRemarkRequest) Reset() {
	*x = ContactEditRemarkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_v1_contact_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContactEditRemarkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContactEditRemarkRequest) ProtoMessage() {}

func (x *ContactEditRemarkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_web_v1_contact_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContactEditRemarkRequest.ProtoReflect.Descriptor instead.
func (*ContactEditRemarkRequest) Descriptor() ([]byte, []int) {
	return file_web_v1_contact_proto_rawDescGZIP(), []int{4}
}

func (x *ContactEditRemarkRequest) GetFriendId() int32 {
	if x != nil {
		return x.FriendId
	}
	return 0
}

func (x *ContactEditRemarkRequest) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

// 联系人备注修改接口响应参数
type ContactEditRemarkResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ContactEditRemarkResponse) Reset() {
	*x = ContactEditRemarkResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_v1_contact_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContactEditRemarkResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContactEditRemarkResponse) ProtoMessage() {}

func (x *ContactEditRemarkResponse) ProtoReflect() protoreflect.Message {
	mi := &file_web_v1_contact_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContactEditRemarkResponse.ProtoReflect.Descriptor instead.
func (*ContactEditRemarkResponse) Descriptor() ([]byte, []int) {
	return file_web_v1_contact_proto_rawDescGZIP(), []int{5}
}

// 联系人详情接口请求参数
type ContactDetailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty" form:"user_id" binding:"required"`
}

func (x *ContactDetailRequest) Reset() {
	*x = ContactDetailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_v1_contact_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContactDetailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContactDetailRequest) ProtoMessage() {}

func (x *ContactDetailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_web_v1_contact_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContactDetailRequest.ProtoReflect.Descriptor instead.
func (*ContactDetailRequest) Descriptor() ([]byte, []int) {
	return file_web_v1_contact_proto_rawDescGZIP(), []int{6}
}

func (x *ContactDetailRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

// 联系人详情接口响应参数
type ContactDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Mobile       string `protobuf:"bytes,2,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Nickname     string `protobuf:"bytes,3,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Remark       string `protobuf:"bytes,4,opt,name=remark,proto3" json:"remark,omitempty"`
	Avatar       string `protobuf:"bytes,5,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Gender       int32  `protobuf:"varint,6,opt,name=gender,proto3" json:"gender,omitempty"`
	Motto        string `protobuf:"bytes,7,opt,name=motto,proto3" json:"motto,omitempty"`
	FriendApply  int32  `protobuf:"varint,8,opt,name=friend_apply,json=friendApply,proto3" json:"friend_apply,omitempty"`
	FriendStatus int32  `protobuf:"varint,9,opt,name=friend_status,json=friendStatus,proto3" json:"friend_status,omitempty"`
	GroupId      int32  `protobuf:"varint,10,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	Email        int32  `protobuf:"varint,11,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *ContactDetailResponse) Reset() {
	*x = ContactDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_v1_contact_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContactDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContactDetailResponse) ProtoMessage() {}

func (x *ContactDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_web_v1_contact_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContactDetailResponse.ProtoReflect.Descriptor instead.
func (*ContactDetailResponse) Descriptor() ([]byte, []int) {
	return file_web_v1_contact_proto_rawDescGZIP(), []int{7}
}

func (x *ContactDetailResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ContactDetailResponse) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *ContactDetailResponse) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *ContactDetailResponse) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *ContactDetailResponse) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *ContactDetailResponse) GetGender() int32 {
	if x != nil {
		return x.Gender
	}
	return 0
}

func (x *ContactDetailResponse) GetMotto() string {
	if x != nil {
		return x.Motto
	}
	return ""
}

func (x *ContactDetailResponse) GetFriendApply() int32 {
	if x != nil {
		return x.FriendApply
	}
	return 0
}

func (x *ContactDetailResponse) GetFriendStatus() int32 {
	if x != nil {
		return x.FriendStatus
	}
	return 0
}

func (x *ContactDetailResponse) GetGroupId() int32 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

func (x *ContactDetailResponse) GetEmail() int32 {
	if x != nil {
		return x.Email
	}
	return 0
}

// 联系人搜索接口请求参数
type ContactSearchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mobile string `protobuf:"bytes,2,opt,name=mobile,proto3" json:"mobile,omitempty" form:"mobile" binding:"required"`
}

func (x *ContactSearchRequest) Reset() {
	*x = ContactSearchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_v1_contact_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContactSearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContactSearchRequest) ProtoMessage() {}

func (x *ContactSearchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_web_v1_contact_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContactSearchRequest.ProtoReflect.Descriptor instead.
func (*ContactSearchRequest) Descriptor() ([]byte, []int) {
	return file_web_v1_contact_proto_rawDescGZIP(), []int{8}
}

func (x *ContactSearchRequest) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

// 联系人搜索接口响应参数
type ContactSearchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Mobile   string `protobuf:"bytes,2,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Nickname string `protobuf:"bytes,3,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Avatar   string `protobuf:"bytes,5,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Gender   int32  `protobuf:"varint,6,opt,name=gender,proto3" json:"gender,omitempty"`
	Motto    string `protobuf:"bytes,7,opt,name=motto,proto3" json:"motto,omitempty"`
}

func (x *ContactSearchResponse) Reset() {
	*x = ContactSearchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_v1_contact_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContactSearchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContactSearchResponse) ProtoMessage() {}

func (x *ContactSearchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_web_v1_contact_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContactSearchResponse.ProtoReflect.Descriptor instead.
func (*ContactSearchResponse) Descriptor() ([]byte, []int) {
	return file_web_v1_contact_proto_rawDescGZIP(), []int{9}
}

func (x *ContactSearchResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ContactSearchResponse) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *ContactSearchResponse) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *ContactSearchResponse) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *ContactSearchResponse) GetGender() int32 {
	if x != nil {
		return x.Gender
	}
	return 0
}

func (x *ContactSearchResponse) GetMotto() string {
	if x != nil {
		return x.Motto
	}
	return ""
}

// 修改联系人分组接口请求参数
type ContactChangeGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  int32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty" form:"user_id" binding:"required"`
	GroupId int32 `protobuf:"varint,2,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty" form:"group_id"`
}

func (x *ContactChangeGroupRequest) Reset() {
	*x = ContactChangeGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_v1_contact_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContactChangeGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContactChangeGroupRequest) ProtoMessage() {}

func (x *ContactChangeGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_web_v1_contact_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContactChangeGroupRequest.ProtoReflect.Descriptor instead.
func (*ContactChangeGroupRequest) Descriptor() ([]byte, []int) {
	return file_web_v1_contact_proto_rawDescGZIP(), []int{10}
}

func (x *ContactChangeGroupRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ContactChangeGroupRequest) GetGroupId() int32 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

// 修改联系人分组接口响应参数
type ContactChangeGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ContactChangeGroupResponse) Reset() {
	*x = ContactChangeGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_v1_contact_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContactChangeGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContactChangeGroupResponse) ProtoMessage() {}

func (x *ContactChangeGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_web_v1_contact_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContactChangeGroupResponse.ProtoReflect.Descriptor instead.
func (*ContactChangeGroupResponse) Descriptor() ([]byte, []int) {
	return file_web_v1_contact_proto_rawDescGZIP(), []int{11}
}

type ContactListResponse_Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 用户ID
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// 昵称
	Nickname string `protobuf:"bytes,2,opt,name=nickname,proto3" json:"nickname,omitempty"`
	// 性别[0:未知;1:男;2:女;]
	Gender int32 `protobuf:"varint,3,opt,name=gender,proto3" json:"gender,omitempty"`
	// 座右铭
	Motto string `protobuf:"bytes,4,opt,name=motto,proto3" json:"motto,omitempty"`
	// 头像
	Avatar string `protobuf:"bytes,5,opt,name=avatar,proto3" json:"avatar,omitempty"`
	// 备注
	Remark string `protobuf:"bytes,6,opt,name=remark,proto3" json:"remark,omitempty"`
	// 是否在线
	IsOnline int32 `protobuf:"varint,7,opt,name=is_online,json=isOnline,proto3" json:"is_online,omitempty"`
	// 联系人分组ID
	GroupId int32 `protobuf:"varint,8,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
}

func (x *ContactListResponse_Item) Reset() {
	*x = ContactListResponse_Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_v1_contact_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContactListResponse_Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContactListResponse_Item) ProtoMessage() {}

func (x *ContactListResponse_Item) ProtoReflect() protoreflect.Message {
	mi := &file_web_v1_contact_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContactListResponse_Item.ProtoReflect.Descriptor instead.
func (*ContactListResponse_Item) Descriptor() ([]byte, []int) {
	return file_web_v1_contact_proto_rawDescGZIP(), []int{1, 0}
}

func (x *ContactListResponse_Item) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ContactListResponse_Item) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *ContactListResponse_Item) GetGender() int32 {
	if x != nil {
		return x.Gender
	}
	return 0
}

func (x *ContactListResponse_Item) GetMotto() string {
	if x != nil {
		return x.Motto
	}
	return ""
}

func (x *ContactListResponse_Item) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *ContactListResponse_Item) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *ContactListResponse_Item) GetIsOnline() int32 {
	if x != nil {
		return x.IsOnline
	}
	return 0
}

func (x *ContactListResponse_Item) GetGroupId() int32 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

var File_web_v1_contact_proto protoreflect.FileDescriptor

var file_web_v1_contact_proto_rawDesc = []byte{
	0x0a, 0x14, 0x77, 0x65, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x77, 0x65, 0x62, 0x1a, 0x13, 0x74, 0x61, 0x67,
	0x67, 0x65, 0x72, 0x2f, 0x74, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x14, 0x0a, 0x12, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x95, 0x02, 0x0a, 0x13, 0x43, 0x6f, 0x6e, 0x74, 0x61,
	0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33,
	0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x77, 0x65, 0x62, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x1a, 0xc8, 0x01, 0x0a, 0x04, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x74, 0x74, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x6d, 0x6f, 0x74, 0x74, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x16,
	0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x6f, 0x6e, 0x6c,
	0x69, 0x6e, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x69, 0x73, 0x4f, 0x6e, 0x6c,
	0x69, 0x6e, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x22, 0x4c,
	0x0a, 0x14, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x09, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x17, 0x9a, 0x84, 0x9e, 0x03, 0x12,
	0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x3a, 0x22, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65,
	0x64, 0x22, 0x52, 0x08, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x49, 0x64, 0x22, 0x17, 0x0a, 0x15,
	0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x68, 0x0a, 0x18, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74,
	0x45, 0x64, 0x69, 0x74, 0x52, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x34, 0x0a, 0x09, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x42, 0x17, 0x9a, 0x84, 0x9e, 0x03, 0x12, 0x62, 0x69, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x3a, 0x22, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x22, 0x52, 0x08, 0x66,
	0x72, 0x69, 0x65, 0x6e, 0x64, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72,
	0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x22,
	0x1b, 0x0a, 0x19, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x45, 0x64, 0x69, 0x74, 0x52, 0x65,
	0x6d, 0x61, 0x72, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x57, 0x0a, 0x14,
	0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x3f, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x26, 0x9a, 0x84, 0x9e, 0x03, 0x21, 0x66, 0x6f, 0x72, 0x6d,
	0x3a, 0x22, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x22, 0x20, 0x62, 0x69, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x3a, 0x22, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x22, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0xb2, 0x02, 0x0a, 0x15, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63,
	0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6d,
	0x6f, 0x74, 0x74, 0x6f, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f, 0x74, 0x74,
	0x6f, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x5f, 0x61, 0x70, 0x70, 0x6c,
	0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x41,
	0x70, 0x70, 0x6c, 0x79, 0x12, 0x23, 0x0a, 0x0d, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x66, 0x72, 0x69,
	0x65, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x55, 0x0a, 0x14, 0x43, 0x6f,
	0x6e, 0x74, 0x61, 0x63, 0x74, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x3d, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x25, 0x9a, 0x84, 0x9e, 0x03, 0x20, 0x66, 0x6f, 0x72, 0x6d, 0x3a, 0x22, 0x6d,
	0x6f, 0x62, 0x69, 0x6c, 0x65, 0x22, 0x20, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x3a, 0x22,
	0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x22, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c,
	0x65, 0x22, 0xa1, 0x01, 0x0a, 0x15, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6d,
	0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62,
	0x69, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12,
	0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x74, 0x74, 0x6f, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x6d, 0x6f, 0x74, 0x74, 0x6f, 0x22, 0x8d, 0x01, 0x0a, 0x19, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63,
	0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x3f, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x42, 0x26, 0x9a, 0x84, 0x9e, 0x03, 0x21, 0x66, 0x6f, 0x72, 0x6d, 0x3a,
	0x22, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x22, 0x20, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x3a, 0x22, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x22, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x2f, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x42, 0x14, 0x9a, 0x84, 0x9e, 0x03, 0x0f, 0x66, 0x6f, 0x72,
	0x6d, 0x3a, 0x22, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x22, 0x52, 0x07, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x49, 0x64, 0x22, 0x1c, 0x0a, 0x1a, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x0c, 0x5a, 0x0a, 0x77, 0x65, 0x62, 0x2f, 0x76, 0x31, 0x3b, 0x77, 0x65,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_web_v1_contact_proto_rawDescOnce sync.Once
	file_web_v1_contact_proto_rawDescData = file_web_v1_contact_proto_rawDesc
)

func file_web_v1_contact_proto_rawDescGZIP() []byte {
	file_web_v1_contact_proto_rawDescOnce.Do(func() {
		file_web_v1_contact_proto_rawDescData = protoimpl.X.CompressGZIP(file_web_v1_contact_proto_rawDescData)
	})
	return file_web_v1_contact_proto_rawDescData
}

var file_web_v1_contact_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_web_v1_contact_proto_goTypes = []interface{}{
	(*ContactListRequest)(nil),         // 0: web.ContactListRequest
	(*ContactListResponse)(nil),        // 1: web.ContactListResponse
	(*ContactDeleteRequest)(nil),       // 2: web.ContactDeleteRequest
	(*ContactDeleteResponse)(nil),      // 3: web.ContactDeleteResponse
	(*ContactEditRemarkRequest)(nil),   // 4: web.ContactEditRemarkRequest
	(*ContactEditRemarkResponse)(nil),  // 5: web.ContactEditRemarkResponse
	(*ContactDetailRequest)(nil),       // 6: web.ContactDetailRequest
	(*ContactDetailResponse)(nil),      // 7: web.ContactDetailResponse
	(*ContactSearchRequest)(nil),       // 8: web.ContactSearchRequest
	(*ContactSearchResponse)(nil),      // 9: web.ContactSearchResponse
	(*ContactChangeGroupRequest)(nil),  // 10: web.ContactChangeGroupRequest
	(*ContactChangeGroupResponse)(nil), // 11: web.ContactChangeGroupResponse
	(*ContactListResponse_Item)(nil),   // 12: web.ContactListResponse.Item
}
var file_web_v1_contact_proto_depIdxs = []int32{
	12, // 0: web.ContactListResponse.items:type_name -> web.ContactListResponse.Item
	1,  // [1:1] is the sub-list for method output_type
	1,  // [1:1] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_web_v1_contact_proto_init() }
func file_web_v1_contact_proto_init() {
	if File_web_v1_contact_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_web_v1_contact_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContactListRequest); i {
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
		file_web_v1_contact_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContactListResponse); i {
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
		file_web_v1_contact_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContactDeleteRequest); i {
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
		file_web_v1_contact_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContactDeleteResponse); i {
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
		file_web_v1_contact_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContactEditRemarkRequest); i {
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
		file_web_v1_contact_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContactEditRemarkResponse); i {
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
		file_web_v1_contact_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContactDetailRequest); i {
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
		file_web_v1_contact_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContactDetailResponse); i {
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
		file_web_v1_contact_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContactSearchRequest); i {
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
		file_web_v1_contact_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContactSearchResponse); i {
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
		file_web_v1_contact_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContactChangeGroupRequest); i {
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
		file_web_v1_contact_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContactChangeGroupResponse); i {
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
		file_web_v1_contact_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContactListResponse_Item); i {
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
			RawDescriptor: file_web_v1_contact_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_web_v1_contact_proto_goTypes,
		DependencyIndexes: file_web_v1_contact_proto_depIdxs,
		MessageInfos:      file_web_v1_contact_proto_msgTypes,
	}.Build()
	File_web_v1_contact_proto = out.File
	file_web_v1_contact_proto_rawDesc = nil
	file_web_v1_contact_proto_goTypes = nil
	file_web_v1_contact_proto_depIdxs = nil
}
