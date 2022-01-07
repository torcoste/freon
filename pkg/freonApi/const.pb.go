// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: const.proto

package freonApi

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

type UserStatus int32

const (
	UserStatus_USER_ACTIVE     UserStatus = 0
	UserStatus_USER_NOT_ACTIVE UserStatus = 1
	UserStatus_USER_IS_BANNED  UserStatus = 2
)

// Enum value maps for UserStatus.
var (
	UserStatus_name = map[int32]string{
		0: "USER_ACTIVE",
		1: "USER_NOT_ACTIVE",
		2: "USER_IS_BANNED",
	}
	UserStatus_value = map[string]int32{
		"USER_ACTIVE":     0,
		"USER_NOT_ACTIVE": 1,
		"USER_IS_BANNED":  2,
	}
)

func (x UserStatus) Enum() *UserStatus {
	p := new(UserStatus)
	*p = x
	return p
}

func (x UserStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_const_proto_enumTypes[0].Descriptor()
}

func (UserStatus) Type() protoreflect.EnumType {
	return &file_const_proto_enumTypes[0]
}

func (x UserStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserStatus.Descriptor instead.
func (UserStatus) EnumDescriptor() ([]byte, []int) {
	return file_const_proto_rawDescGZIP(), []int{0}
}

type UserRole int32

const (
	UserRole_USER_ROLE_ADMIN      UserRole = 0
	UserRole_USER_ROLE_TRANSLATOR UserRole = 1
	UserRole_USER_ROLE_MODERATOR  UserRole = 2
	UserRole_USER_ROLE_USER       UserRole = 3
)

// Enum value maps for UserRole.
var (
	UserRole_name = map[int32]string{
		0: "USER_ROLE_ADMIN",
		1: "USER_ROLE_TRANSLATOR",
		2: "USER_ROLE_MODERATOR",
		3: "USER_ROLE_USER",
	}
	UserRole_value = map[string]int32{
		"USER_ROLE_ADMIN":      0,
		"USER_ROLE_TRANSLATOR": 1,
		"USER_ROLE_MODERATOR":  2,
		"USER_ROLE_USER":       3,
	}
)

func (x UserRole) Enum() *UserRole {
	p := new(UserRole)
	*p = x
	return p
}

func (x UserRole) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserRole) Descriptor() protoreflect.EnumDescriptor {
	return file_const_proto_enumTypes[1].Descriptor()
}

func (UserRole) Type() protoreflect.EnumType {
	return &file_const_proto_enumTypes[1]
}

func (x UserRole) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserRole.Descriptor instead.
func (UserRole) EnumDescriptor() ([]byte, []int) {
	return file_const_proto_rawDescGZIP(), []int{1}
}

type Status int32

const (
	Status_ACTIVE     Status = 0
	Status_NOT_ACTIVE Status = 1
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "ACTIVE",
		1: "NOT_ACTIVE",
	}
	Status_value = map[string]int32{
		"ACTIVE":     0,
		"NOT_ACTIVE": 1,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_const_proto_enumTypes[2].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_const_proto_enumTypes[2]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_const_proto_rawDescGZIP(), []int{2}
}

type StatusTranslation int32

const (
	StatusTranslation_HIDDEN  StatusTranslation = 0
	StatusTranslation_DRAFT   StatusTranslation = 1
	StatusTranslation_RELEASE StatusTranslation = 2
)

// Enum value maps for StatusTranslation.
var (
	StatusTranslation_name = map[int32]string{
		0: "HIDDEN",
		1: "DRAFT",
		2: "RELEASE",
	}
	StatusTranslation_value = map[string]int32{
		"HIDDEN":  0,
		"DRAFT":   1,
		"RELEASE": 2,
	}
)

func (x StatusTranslation) Enum() *StatusTranslation {
	p := new(StatusTranslation)
	*p = x
	return p
}

func (x StatusTranslation) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StatusTranslation) Descriptor() protoreflect.EnumDescriptor {
	return file_const_proto_enumTypes[3].Descriptor()
}

func (StatusTranslation) Type() protoreflect.EnumType {
	return &file_const_proto_enumTypes[3]
}

func (x StatusTranslation) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StatusTranslation.Descriptor instead.
func (StatusTranslation) EnumDescriptor() ([]byte, []int) {
	return file_const_proto_rawDescGZIP(), []int{3}
}

type StorageType int32

const (
	StorageType_STORAGE_TYPE_LOCAL StorageType = 0
	StorageType_STORAGE_TYPE_S3    StorageType = 1
)

// Enum value maps for StorageType.
var (
	StorageType_name = map[int32]string{
		0: "STORAGE_TYPE_LOCAL",
		1: "STORAGE_TYPE_S3",
	}
	StorageType_value = map[string]int32{
		"STORAGE_TYPE_LOCAL": 0,
		"STORAGE_TYPE_S3":    1,
	}
)

func (x StorageType) Enum() *StorageType {
	p := new(StorageType)
	*p = x
	return p
}

func (x StorageType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StorageType) Descriptor() protoreflect.EnumDescriptor {
	return file_const_proto_enumTypes[4].Descriptor()
}

func (StorageType) Type() protoreflect.EnumType {
	return &file_const_proto_enumTypes[4]
}

func (x StorageType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StorageType.Descriptor instead.
func (StorageType) EnumDescriptor() ([]byte, []int) {
	return file_const_proto_rawDescGZIP(), []int{4}
}

type PlatformType int32

const (
	PlatformType_PLATFORM_TYPE_WEB     PlatformType = 0
	PlatformType_PLATFORM_TYPE_IOS     PlatformType = 1
	PlatformType_PLATFORM_TYPE_ANDROID PlatformType = 2
	PlatformType_PLATFORM_TYPE_SKIP    PlatformType = 3 // skipped, using for checking new update translation with grpc
)

// Enum value maps for PlatformType.
var (
	PlatformType_name = map[int32]string{
		0: "PLATFORM_TYPE_WEB",
		1: "PLATFORM_TYPE_IOS",
		2: "PLATFORM_TYPE_ANDROID",
		3: "PLATFORM_TYPE_SKIP",
	}
	PlatformType_value = map[string]int32{
		"PLATFORM_TYPE_WEB":     0,
		"PLATFORM_TYPE_IOS":     1,
		"PLATFORM_TYPE_ANDROID": 2,
		"PLATFORM_TYPE_SKIP":    3,
	}
)

func (x PlatformType) Enum() *PlatformType {
	p := new(PlatformType)
	*p = x
	return p
}

func (x PlatformType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PlatformType) Descriptor() protoreflect.EnumDescriptor {
	return file_const_proto_enumTypes[5].Descriptor()
}

func (PlatformType) Type() protoreflect.EnumType {
	return &file_const_proto_enumTypes[5]
}

func (x PlatformType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PlatformType.Descriptor instead.
func (PlatformType) EnumDescriptor() ([]byte, []int) {
	return file_const_proto_rawDescGZIP(), []int{5}
}

type TranslationSource int32

const (
	TranslationSource_TRANSLATION_NONE   TranslationSource = 0
	TranslationSource_TRANSLATION_LIBRA  TranslationSource = 1
	TranslationSource_TRANSLATION_GOOGLE TranslationSource = 2
)

// Enum value maps for TranslationSource.
var (
	TranslationSource_name = map[int32]string{
		0: "TRANSLATION_NONE",
		1: "TRANSLATION_LIBRA",
		2: "TRANSLATION_GOOGLE",
	}
	TranslationSource_value = map[string]int32{
		"TRANSLATION_NONE":   0,
		"TRANSLATION_LIBRA":  1,
		"TRANSLATION_GOOGLE": 2,
	}
)

func (x TranslationSource) Enum() *TranslationSource {
	p := new(TranslationSource)
	*p = x
	return p
}

func (x TranslationSource) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TranslationSource) Descriptor() protoreflect.EnumDescriptor {
	return file_const_proto_enumTypes[6].Descriptor()
}

func (TranslationSource) Type() protoreflect.EnumType {
	return &file_const_proto_enumTypes[6]
}

func (x TranslationSource) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TranslationSource.Descriptor instead.
func (TranslationSource) EnumDescriptor() ([]byte, []int) {
	return file_const_proto_rawDescGZIP(), []int{6}
}

type TranslationConfiguration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Auto bool              `protobuf:"varint,1,opt,name=Auto,proto3" json:"Auto,omitempty"`
	Use  TranslationSource `protobuf:"varint,2,opt,name=Use,proto3,enum=freon.TranslationSource" json:"Use,omitempty"`
}

func (x *TranslationConfiguration) Reset() {
	*x = TranslationConfiguration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_const_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TranslationConfiguration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TranslationConfiguration) ProtoMessage() {}

func (x *TranslationConfiguration) ProtoReflect() protoreflect.Message {
	mi := &file_const_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TranslationConfiguration.ProtoReflect.Descriptor instead.
func (*TranslationConfiguration) Descriptor() ([]byte, []int) {
	return file_const_proto_rawDescGZIP(), []int{0}
}

func (x *TranslationConfiguration) GetAuto() bool {
	if x != nil {
		return x.Auto
	}
	return false
}

func (x *TranslationConfiguration) GetUse() TranslationSource {
	if x != nil {
		return x.Use
	}
	return TranslationSource_TRANSLATION_NONE
}

type StorageConfiguration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Use StorageType `protobuf:"varint,1,opt,name=Use,proto3,enum=freon.StorageType" json:"Use,omitempty"`
}

func (x *StorageConfiguration) Reset() {
	*x = StorageConfiguration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_const_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StorageConfiguration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StorageConfiguration) ProtoMessage() {}

func (x *StorageConfiguration) ProtoReflect() protoreflect.Message {
	mi := &file_const_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StorageConfiguration.ProtoReflect.Descriptor instead.
func (*StorageConfiguration) Descriptor() ([]byte, []int) {
	return file_const_proto_rawDescGZIP(), []int{1}
}

func (x *StorageConfiguration) GetUse() StorageType {
	if x != nil {
		return x.Use
	}
	return StorageType_STORAGE_TYPE_LOCAL
}

type SettingConfiguration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Storage     *StorageConfiguration     `protobuf:"bytes,1,opt,name=Storage,proto3" json:"Storage,omitempty"`
	Translation *TranslationConfiguration `protobuf:"bytes,2,opt,name=Translation,proto3" json:"Translation,omitempty"`
}

func (x *SettingConfiguration) Reset() {
	*x = SettingConfiguration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_const_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SettingConfiguration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SettingConfiguration) ProtoMessage() {}

func (x *SettingConfiguration) ProtoReflect() protoreflect.Message {
	mi := &file_const_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SettingConfiguration.ProtoReflect.Descriptor instead.
func (*SettingConfiguration) Descriptor() ([]byte, []int) {
	return file_const_proto_rawDescGZIP(), []int{2}
}

func (x *SettingConfiguration) GetStorage() *StorageConfiguration {
	if x != nil {
		return x.Storage
	}
	return nil
}

func (x *SettingConfiguration) GetTranslation() *TranslationConfiguration {
	if x != nil {
		return x.Translation
	}
	return nil
}

var File_const_proto protoreflect.FileDescriptor

var file_const_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x66,
	0x72, 0x65, 0x6f, 0x6e, 0x22, 0x5a, 0x0a, 0x18, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x12, 0x0a, 0x04, 0x41, 0x75, 0x74, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04,
	0x41, 0x75, 0x74, 0x6f, 0x12, 0x2a, 0x0a, 0x03, 0x55, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x18, 0x2e, 0x66, 0x72, 0x65, 0x6f, 0x6e, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x03, 0x55, 0x73, 0x65,
	0x22, 0x3c, 0x0a, 0x14, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x03, 0x55, 0x73, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x66, 0x72, 0x65, 0x6f, 0x6e, 0x2e, 0x53, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x03, 0x55, 0x73, 0x65, 0x22, 0x90,
	0x01, 0x0a, 0x14, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x35, 0x0a, 0x07, 0x53, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x66, 0x72, 0x65, 0x6f, 0x6e,
	0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x12, 0x41,
	0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x66, 0x72, 0x65, 0x6f, 0x6e, 0x2e, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2a, 0x46, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x0f, 0x0a, 0x0b, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x00,
	0x12, 0x13, 0x0a, 0x0f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x41, 0x43, 0x54,
	0x49, 0x56, 0x45, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x49, 0x53,
	0x5f, 0x42, 0x41, 0x4e, 0x4e, 0x45, 0x44, 0x10, 0x02, 0x2a, 0x66, 0x0a, 0x08, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x13, 0x0a, 0x0f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x52, 0x4f,
	0x4c, 0x45, 0x5f, 0x41, 0x44, 0x4d, 0x49, 0x4e, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x55, 0x53,
	0x45, 0x52, 0x5f, 0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x4c, 0x41, 0x54,
	0x4f, 0x52, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x52, 0x4f, 0x4c,
	0x45, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x52, 0x41, 0x54, 0x4f, 0x52, 0x10, 0x02, 0x12, 0x12, 0x0a,
	0x0e, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x10,
	0x03, 0x2a, 0x24, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0a, 0x0a, 0x06, 0x41,
	0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x4e, 0x4f, 0x54, 0x5f, 0x41,
	0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x01, 0x2a, 0x37, 0x0a, 0x11, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0a, 0x0a, 0x06,
	0x48, 0x49, 0x44, 0x44, 0x45, 0x4e, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x44, 0x52, 0x41, 0x46,
	0x54, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x45, 0x4c, 0x45, 0x41, 0x53, 0x45, 0x10, 0x02,
	0x2a, 0x3a, 0x0a, 0x0b, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x16, 0x0a, 0x12, 0x53, 0x54, 0x4f, 0x52, 0x41, 0x47, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x4c, 0x4f, 0x43, 0x41, 0x4c, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x54, 0x4f, 0x52, 0x41,
	0x47, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x33, 0x10, 0x01, 0x2a, 0x6f, 0x0a, 0x0c,
	0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x15, 0x0a, 0x11,
	0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x57, 0x45,
	0x42, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x4f, 0x53, 0x10, 0x01, 0x12, 0x19, 0x0a, 0x15, 0x50, 0x4c,
	0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x4e, 0x44, 0x52,
	0x4f, 0x49, 0x44, 0x10, 0x02, 0x12, 0x16, 0x0a, 0x12, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52,
	0x4d, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x4b, 0x49, 0x50, 0x10, 0x03, 0x2a, 0x58, 0x0a,
	0x11, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x4c, 0x41, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x54, 0x52, 0x41, 0x4e,
	0x53, 0x4c, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4c, 0x49, 0x42, 0x52, 0x41, 0x10, 0x01, 0x12,
	0x16, 0x0a, 0x12, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x4c, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x47,
	0x4f, 0x4f, 0x47, 0x4c, 0x45, 0x10, 0x02, 0x42, 0x0e, 0x5a, 0x0c, 0x70, 0x6b, 0x67, 0x2f, 0x66,
	0x72, 0x65, 0x6f, 0x6e, 0x41, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_const_proto_rawDescOnce sync.Once
	file_const_proto_rawDescData = file_const_proto_rawDesc
)

func file_const_proto_rawDescGZIP() []byte {
	file_const_proto_rawDescOnce.Do(func() {
		file_const_proto_rawDescData = protoimpl.X.CompressGZIP(file_const_proto_rawDescData)
	})
	return file_const_proto_rawDescData
}

var file_const_proto_enumTypes = make([]protoimpl.EnumInfo, 7)
var file_const_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_const_proto_goTypes = []interface{}{
	(UserStatus)(0),                  // 0: freon.UserStatus
	(UserRole)(0),                    // 1: freon.UserRole
	(Status)(0),                      // 2: freon.Status
	(StatusTranslation)(0),           // 3: freon.StatusTranslation
	(StorageType)(0),                 // 4: freon.StorageType
	(PlatformType)(0),                // 5: freon.PlatformType
	(TranslationSource)(0),           // 6: freon.TranslationSource
	(*TranslationConfiguration)(nil), // 7: freon.TranslationConfiguration
	(*StorageConfiguration)(nil),     // 8: freon.StorageConfiguration
	(*SettingConfiguration)(nil),     // 9: freon.SettingConfiguration
}
var file_const_proto_depIdxs = []int32{
	6, // 0: freon.TranslationConfiguration.Use:type_name -> freon.TranslationSource
	4, // 1: freon.StorageConfiguration.Use:type_name -> freon.StorageType
	8, // 2: freon.SettingConfiguration.Storage:type_name -> freon.StorageConfiguration
	7, // 3: freon.SettingConfiguration.Translation:type_name -> freon.TranslationConfiguration
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_const_proto_init() }
func file_const_proto_init() {
	if File_const_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_const_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TranslationConfiguration); i {
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
		file_const_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StorageConfiguration); i {
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
		file_const_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SettingConfiguration); i {
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
			RawDescriptor: file_const_proto_rawDesc,
			NumEnums:      7,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_const_proto_goTypes,
		DependencyIndexes: file_const_proto_depIdxs,
		EnumInfos:         file_const_proto_enumTypes,
		MessageInfos:      file_const_proto_msgTypes,
	}.Build()
	File_const_proto = out.File
	file_const_proto_rawDesc = nil
	file_const_proto_goTypes = nil
	file_const_proto_depIdxs = nil
}
