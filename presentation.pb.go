// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: presentation.proto

package presentationContracts

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ImageDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AltText string `protobuf:"bytes,1,opt,name=AltText,proto3" json:"AltText,omitempty"`
	Height  int32  `protobuf:"varint,2,opt,name=Height,proto3" json:"Height,omitempty"`
	Url     string `protobuf:"bytes,3,opt,name=Url,proto3" json:"Url,omitempty"`
	Width   int32  `protobuf:"varint,4,opt,name=Width,proto3" json:"Width,omitempty"`
}

func (x *ImageDetails) Reset() {
	*x = ImageDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_presentation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImageDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageDetails) ProtoMessage() {}

func (x *ImageDetails) ProtoReflect() protoreflect.Message {
	mi := &file_presentation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageDetails.ProtoReflect.Descriptor instead.
func (*ImageDetails) Descriptor() ([]byte, []int) {
	return file_presentation_proto_rawDescGZIP(), []int{0}
}

func (x *ImageDetails) GetAltText() string {
	if x != nil {
		return x.AltText
	}
	return ""
}

func (x *ImageDetails) GetHeight() int32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *ImageDetails) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *ImageDetails) GetWidth() int32 {
	if x != nil {
		return x.Width
	}
	return 0
}

type ArtistRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *ArtistRequest) Reset() {
	*x = ArtistRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_presentation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArtistRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArtistRequest) ProtoMessage() {}

func (x *ArtistRequest) ProtoReflect() protoreflect.Message {
	mi := &file_presentation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArtistRequest.ProtoReflect.Descriptor instead.
func (*ArtistRequest) Descriptor() ([]byte, []int) {
	return file_presentation_proto_rawDescGZIP(), []int{1}
}

func (x *ArtistRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Artist struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int32          `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Name     string         `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Releases []*SlimRelease `protobuf:"bytes,3,rep,name=Releases,proto3" json:"Releases,omitempty"`
}

func (x *Artist) Reset() {
	*x = Artist{}
	if protoimpl.UnsafeEnabled {
		mi := &file_presentation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Artist) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Artist) ProtoMessage() {}

func (x *Artist) ProtoReflect() protoreflect.Message {
	mi := &file_presentation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Artist.ProtoReflect.Descriptor instead.
func (*Artist) Descriptor() ([]byte, []int) {
	return file_presentation_proto_rawDescGZIP(), []int{2}
}

func (x *Artist) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Artist) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Artist) GetReleases() []*SlimRelease {
	if x != nil {
		return x.Releases
	}
	return nil
}

type ArtistStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CompilationNumber int32 `protobuf:"varint,1,opt,name=CompilationNumber,proto3" json:"CompilationNumber,omitempty"`
	SoleReleaseNumber int32 `protobuf:"varint,2,opt,name=SoleReleaseNumber,proto3" json:"SoleReleaseNumber,omitempty"`
}

func (x *ArtistStats) Reset() {
	*x = ArtistStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_presentation_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArtistStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArtistStats) ProtoMessage() {}

func (x *ArtistStats) ProtoReflect() protoreflect.Message {
	mi := &file_presentation_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArtistStats.ProtoReflect.Descriptor instead.
func (*ArtistStats) Descriptor() ([]byte, []int) {
	return file_presentation_proto_rawDescGZIP(), []int{3}
}

func (x *ArtistStats) GetCompilationNumber() int32 {
	if x != nil {
		return x.CompilationNumber
	}
	return 0
}

func (x *ArtistStats) GetSoleReleaseNumber() int32 {
	if x != nil {
		return x.SoleReleaseNumber
	}
	return 0
}

type HealthCheckResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HealthCheckResponse) Reset() {
	*x = HealthCheckResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_presentation_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthCheckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthCheckResponse) ProtoMessage() {}

func (x *HealthCheckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_presentation_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthCheckResponse.ProtoReflect.Descriptor instead.
func (*HealthCheckResponse) Descriptor() ([]byte, []int) {
	return file_presentation_proto_rawDescGZIP(), []int{4}
}

type HealthCheckRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HealthCheckRequest) Reset() {
	*x = HealthCheckRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_presentation_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthCheckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthCheckRequest) ProtoMessage() {}

func (x *HealthCheckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_presentation_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthCheckRequest.ProtoReflect.Descriptor instead.
func (*HealthCheckRequest) Descriptor() ([]byte, []int) {
	return file_presentation_proto_rawDescGZIP(), []int{5}
}

type PlatformUrl struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlatformId string `protobuf:"bytes,1,opt,name=PlatformId,proto3" json:"PlatformId,omitempty"`
	Url        string `protobuf:"bytes,2,opt,name=Url,proto3" json:"Url,omitempty"`
}

func (x *PlatformUrl) Reset() {
	*x = PlatformUrl{}
	if protoimpl.UnsafeEnabled {
		mi := &file_presentation_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlatformUrl) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlatformUrl) ProtoMessage() {}

func (x *PlatformUrl) ProtoReflect() protoreflect.Message {
	mi := &file_presentation_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlatformUrl.ProtoReflect.Descriptor instead.
func (*PlatformUrl) Descriptor() ([]byte, []int) {
	return file_presentation_proto_rawDescGZIP(), []int{6}
}

func (x *PlatformUrl) GetPlatformId() string {
	if x != nil {
		return x.PlatformId
	}
	return ""
}

func (x *PlatformUrl) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type Release struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             int32                  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Description    string                 `protobuf:"bytes,2,opt,name=Description,proto3" json:"Description,omitempty"`
	ImageDetails   *ImageDetails          `protobuf:"bytes,3,opt,name=ImageDetails,proto3" json:"ImageDetails,omitempty"`
	Label          string                 `protobuf:"bytes,4,opt,name=Label,proto3" json:"Label,omitempty"`
	Name           string                 `protobuf:"bytes,5,opt,name=Name,proto3" json:"Name,omitempty"`
	PlatformUrls   []*PlatformUrl         `protobuf:"bytes,6,rep,name=PlatformUrls,proto3" json:"PlatformUrls,omitempty"`
	ReleaseArtists []*SlimArtist          `protobuf:"bytes,7,rep,name=ReleaseArtists,proto3" json:"ReleaseArtists,omitempty"`
	ReleaseDate    *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=ReleaseDate,proto3" json:"ReleaseDate,omitempty"`
	Tags           []string               `protobuf:"bytes,9,rep,name=Tags,proto3" json:"Tags,omitempty"`
	Tracks         []*Track               `protobuf:"bytes,10,rep,name=Tracks,proto3" json:"Tracks,omitempty"`
	Type           string                 `protobuf:"bytes,11,opt,name=Type,proto3" json:"Type,omitempty"`
	ArtistStats    *ArtistStats           `protobuf:"bytes,12,opt,name=ArtistStats,proto3" json:"ArtistStats,omitempty"`
}

func (x *Release) Reset() {
	*x = Release{}
	if protoimpl.UnsafeEnabled {
		mi := &file_presentation_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Release) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Release) ProtoMessage() {}

func (x *Release) ProtoReflect() protoreflect.Message {
	mi := &file_presentation_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Release.ProtoReflect.Descriptor instead.
func (*Release) Descriptor() ([]byte, []int) {
	return file_presentation_proto_rawDescGZIP(), []int{7}
}

func (x *Release) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Release) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Release) GetImageDetails() *ImageDetails {
	if x != nil {
		return x.ImageDetails
	}
	return nil
}

func (x *Release) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *Release) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Release) GetPlatformUrls() []*PlatformUrl {
	if x != nil {
		return x.PlatformUrls
	}
	return nil
}

func (x *Release) GetReleaseArtists() []*SlimArtist {
	if x != nil {
		return x.ReleaseArtists
	}
	return nil
}

func (x *Release) GetReleaseDate() *timestamppb.Timestamp {
	if x != nil {
		return x.ReleaseDate
	}
	return nil
}

func (x *Release) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *Release) GetTracks() []*Track {
	if x != nil {
		return x.Tracks
	}
	return nil
}

func (x *Release) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Release) GetArtistStats() *ArtistStats {
	if x != nil {
		return x.ArtistStats
	}
	return nil
}

type ReleaseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *ReleaseRequest) Reset() {
	*x = ReleaseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_presentation_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReleaseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReleaseRequest) ProtoMessage() {}

func (x *ReleaseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_presentation_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReleaseRequest.ProtoReflect.Descriptor instead.
func (*ReleaseRequest) Descriptor() ([]byte, []int) {
	return file_presentation_proto_rawDescGZIP(), []int{8}
}

func (x *ReleaseRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type SlimArtist struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
}

func (x *SlimArtist) Reset() {
	*x = SlimArtist{}
	if protoimpl.UnsafeEnabled {
		mi := &file_presentation_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SlimArtist) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SlimArtist) ProtoMessage() {}

func (x *SlimArtist) ProtoReflect() protoreflect.Message {
	mi := &file_presentation_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SlimArtist.ProtoReflect.Descriptor instead.
func (*SlimArtist) Descriptor() ([]byte, []int) {
	return file_presentation_proto_rawDescGZIP(), []int{9}
}

func (x *SlimArtist) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SlimArtist) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type SlimRelease struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int32                  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	ImageDetails *ImageDetails          `protobuf:"bytes,2,opt,name=ImageDetails,proto3" json:"ImageDetails,omitempty"`
	Name         string                 `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
	ReleaseDate  *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=ReleaseDate,proto3" json:"ReleaseDate,omitempty"`
	Type         string                 `protobuf:"bytes,5,opt,name=Type,proto3" json:"Type,omitempty"`
}

func (x *SlimRelease) Reset() {
	*x = SlimRelease{}
	if protoimpl.UnsafeEnabled {
		mi := &file_presentation_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SlimRelease) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SlimRelease) ProtoMessage() {}

func (x *SlimRelease) ProtoReflect() protoreflect.Message {
	mi := &file_presentation_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SlimRelease.ProtoReflect.Descriptor instead.
func (*SlimRelease) Descriptor() ([]byte, []int) {
	return file_presentation_proto_rawDescGZIP(), []int{10}
}

func (x *SlimRelease) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SlimRelease) GetImageDetails() *ImageDetails {
	if x != nil {
		return x.ImageDetails
	}
	return nil
}

func (x *SlimRelease) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SlimRelease) GetReleaseDate() *timestamppb.Timestamp {
	if x != nil {
		return x.ReleaseDate
	}
	return nil
}

func (x *SlimRelease) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type Track struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Artists     []*SlimArtist `protobuf:"bytes,1,rep,name=Artists,proto3" json:"Artists,omitempty"`
	DiscNumber  int32         `protobuf:"varint,2,opt,name=DiscNumber,proto3" json:"DiscNumber,omitempty"`
	IsExplicit  bool          `protobuf:"varint,3,opt,name=IsExplicit,proto3" json:"IsExplicit,omitempty"`
	Name        string        `protobuf:"bytes,4,opt,name=Name,proto3" json:"Name,omitempty"`
	TrackNumber int32         `protobuf:"varint,5,opt,name=TrackNumber,proto3" json:"TrackNumber,omitempty"`
}

func (x *Track) Reset() {
	*x = Track{}
	if protoimpl.UnsafeEnabled {
		mi := &file_presentation_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Track) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Track) ProtoMessage() {}

func (x *Track) ProtoReflect() protoreflect.Message {
	mi := &file_presentation_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Track.ProtoReflect.Descriptor instead.
func (*Track) Descriptor() ([]byte, []int) {
	return file_presentation_proto_rawDescGZIP(), []int{11}
}

func (x *Track) GetArtists() []*SlimArtist {
	if x != nil {
		return x.Artists
	}
	return nil
}

func (x *Track) GetDiscNumber() int32 {
	if x != nil {
		return x.DiscNumber
	}
	return 0
}

func (x *Track) GetIsExplicit() bool {
	if x != nil {
		return x.IsExplicit
	}
	return false
}

func (x *Track) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Track) GetTrackNumber() int32 {
	if x != nil {
		return x.TrackNumber
	}
	return 0
}

var File_presentation_proto protoreflect.FileDescriptor

var file_presentation_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x68, 0x0a, 0x0c, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x6c, 0x74, 0x54, 0x65, 0x78, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x6c, 0x74, 0x54, 0x65, 0x78, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x72, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x55, 0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x57, 0x69, 0x64,
	0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x57, 0x69, 0x64, 0x74, 0x68, 0x22,
	0x1f, 0x0a, 0x0d, 0x41, 0x72, 0x74, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x64,
	0x22, 0x56, 0x0a, 0x06, 0x41, 0x72, 0x74, 0x69, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x28,
	0x0a, 0x08, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x53, 0x6c, 0x69, 0x6d, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x52, 0x08,
	0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x73, 0x22, 0x69, 0x0a, 0x0b, 0x41, 0x72, 0x74, 0x69,
	0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x2c, 0x0a, 0x11, 0x43, 0x6f, 0x6d, 0x70, 0x69,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x11, 0x43, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x2c, 0x0a, 0x11, 0x53, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x6c,
	0x65, 0x61, 0x73, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x11, 0x53, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x22, 0x15, 0x0a, 0x13, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14, 0x0a, 0x12, 0x48, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x3f, 0x0a, 0x0b, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x55, 0x72, 0x6c, 0x12,
	0x1e, 0x0a, 0x0a, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x49, 0x64, 0x12,
	0x10, 0x0a, 0x03, 0x55, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x55, 0x72,
	0x6c, 0x22, 0xb5, 0x03, 0x0a, 0x07, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x64, 0x12, 0x20, 0x0a,
	0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x31, 0x0a, 0x0c, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x52, 0x0c, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x0c,
	0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x55, 0x72, 0x6c, 0x73, 0x18, 0x06, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x55, 0x72, 0x6c,
	0x52, 0x0c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x55, 0x72, 0x6c, 0x73, 0x12, 0x33,
	0x0a, 0x0e, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x41, 0x72, 0x74, 0x69, 0x73, 0x74, 0x73,
	0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x53, 0x6c, 0x69, 0x6d, 0x41, 0x72, 0x74,
	0x69, 0x73, 0x74, 0x52, 0x0e, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x41, 0x72, 0x74, 0x69,
	0x73, 0x74, 0x73, 0x12, 0x3c, 0x0a, 0x0b, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x44, 0x61,
	0x74, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x44, 0x61, 0x74,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x61, 0x67, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x04, 0x54, 0x61, 0x67, 0x73, 0x12, 0x1e, 0x0a, 0x06, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x73, 0x18,
	0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x52, 0x06, 0x54,
	0x72, 0x61, 0x63, 0x6b, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2e, 0x0a, 0x0b, 0x41, 0x72, 0x74,
	0x69, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x41, 0x72, 0x74, 0x69, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x0b, 0x41, 0x72,
	0x74, 0x69, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x22, 0x20, 0x0a, 0x0e, 0x52, 0x65, 0x6c,
	0x65, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x64, 0x22, 0x30, 0x0a, 0x0a, 0x53,
	0x6c, 0x69, 0x6d, 0x41, 0x72, 0x74, 0x69, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xb6, 0x01,
	0x0a, 0x0b, 0x53, 0x6c, 0x69, 0x6d, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x64, 0x12, 0x31, 0x0a,
	0x0c, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x52, 0x0c, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x3c, 0x0a, 0x0b, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x44,
	0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x44, 0x61,
	0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x22, 0xa4, 0x01, 0x0a, 0x05, 0x54, 0x72, 0x61, 0x63, 0x6b,
	0x12, 0x25, 0x0a, 0x07, 0x41, 0x72, 0x74, 0x69, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0b, 0x2e, 0x53, 0x6c, 0x69, 0x6d, 0x41, 0x72, 0x74, 0x69, 0x73, 0x74, 0x52, 0x07,
	0x41, 0x72, 0x74, 0x69, 0x73, 0x74, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x44, 0x69, 0x73, 0x63, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x44, 0x69, 0x73,
	0x63, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x49, 0x73, 0x45, 0x78, 0x70,
	0x6c, 0x69, 0x63, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x49, 0x73, 0x45,
	0x78, 0x70, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x54,
	0x72, 0x61, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0b, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x32, 0x9d, 0x01,
	0x0a, 0x0c, 0x50, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3a,
	0x0a, 0x0b, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x12, 0x13, 0x2e,
	0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x14, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x26, 0x0a, 0x09, 0x47, 0x65,
	0x74, 0x41, 0x72, 0x74, 0x69, 0x73, 0x74, 0x12, 0x0e, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x07, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x73, 0x74,
	0x22, 0x00, 0x12, 0x29, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65,
	0x12, 0x0f, 0x2e, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x08, 0x2e, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x22, 0x00, 0x42, 0x19, 0x5a,
	0x17, 0x2e, 0x2f, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43,
	0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_presentation_proto_rawDescOnce sync.Once
	file_presentation_proto_rawDescData = file_presentation_proto_rawDesc
)

func file_presentation_proto_rawDescGZIP() []byte {
	file_presentation_proto_rawDescOnce.Do(func() {
		file_presentation_proto_rawDescData = protoimpl.X.CompressGZIP(file_presentation_proto_rawDescData)
	})
	return file_presentation_proto_rawDescData
}

var file_presentation_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_presentation_proto_goTypes = []interface{}{
	(*ImageDetails)(nil),          // 0: ImageDetails
	(*ArtistRequest)(nil),         // 1: ArtistRequest
	(*Artist)(nil),                // 2: Artist
	(*ArtistStats)(nil),           // 3: ArtistStats
	(*HealthCheckResponse)(nil),   // 4: HealthCheckResponse
	(*HealthCheckRequest)(nil),    // 5: HealthCheckRequest
	(*PlatformUrl)(nil),           // 6: PlatformUrl
	(*Release)(nil),               // 7: Release
	(*ReleaseRequest)(nil),        // 8: ReleaseRequest
	(*SlimArtist)(nil),            // 9: SlimArtist
	(*SlimRelease)(nil),           // 10: SlimRelease
	(*Track)(nil),                 // 11: Track
	(*timestamppb.Timestamp)(nil), // 12: google.protobuf.Timestamp
}
var file_presentation_proto_depIdxs = []int32{
	10, // 0: Artist.Releases:type_name -> SlimRelease
	0,  // 1: Release.ImageDetails:type_name -> ImageDetails
	6,  // 2: Release.PlatformUrls:type_name -> PlatformUrl
	9,  // 3: Release.ReleaseArtists:type_name -> SlimArtist
	12, // 4: Release.ReleaseDate:type_name -> google.protobuf.Timestamp
	11, // 5: Release.Tracks:type_name -> Track
	3,  // 6: Release.ArtistStats:type_name -> ArtistStats
	0,  // 7: SlimRelease.ImageDetails:type_name -> ImageDetails
	12, // 8: SlimRelease.ReleaseDate:type_name -> google.protobuf.Timestamp
	9,  // 9: Track.Artists:type_name -> SlimArtist
	5,  // 10: Presentation.CheckHealth:input_type -> HealthCheckRequest
	1,  // 11: Presentation.GetArtist:input_type -> ArtistRequest
	8,  // 12: Presentation.GetRelease:input_type -> ReleaseRequest
	4,  // 13: Presentation.CheckHealth:output_type -> HealthCheckResponse
	2,  // 14: Presentation.GetArtist:output_type -> Artist
	7,  // 15: Presentation.GetRelease:output_type -> Release
	13, // [13:16] is the sub-list for method output_type
	10, // [10:13] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_presentation_proto_init() }
func file_presentation_proto_init() {
	if File_presentation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_presentation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImageDetails); i {
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
		file_presentation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArtistRequest); i {
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
		file_presentation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Artist); i {
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
		file_presentation_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArtistStats); i {
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
		file_presentation_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthCheckResponse); i {
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
		file_presentation_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthCheckRequest); i {
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
		file_presentation_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlatformUrl); i {
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
		file_presentation_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Release); i {
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
		file_presentation_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReleaseRequest); i {
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
		file_presentation_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SlimArtist); i {
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
		file_presentation_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SlimRelease); i {
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
		file_presentation_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Track); i {
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
			RawDescriptor: file_presentation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_presentation_proto_goTypes,
		DependencyIndexes: file_presentation_proto_depIdxs,
		MessageInfos:      file_presentation_proto_msgTypes,
	}.Build()
	File_presentation_proto = out.File
	file_presentation_proto_rawDesc = nil
	file_presentation_proto_goTypes = nil
	file_presentation_proto_depIdxs = nil
}
