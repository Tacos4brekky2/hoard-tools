// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: hoard-data.proto

package hoard_data

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

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID   string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Body string `protobuf:"bytes,2,opt,name=Body,proto3" json:"Body,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	mi := &file_hoard_data_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_hoard_data_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_hoard_data_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Message) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

type FetchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID            string            `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	URL           string            `protobuf:"bytes,2,opt,name=URL,proto3" json:"URL,omitempty"`
	RequestMethod string            `protobuf:"bytes,3,opt,name=RequestMethod,proto3" json:"RequestMethod,omitempty"`
	Headers       map[string]string `protobuf:"bytes,4,rep,name=Headers,proto3" json:"Headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *FetchRequest) Reset() {
	*x = FetchRequest{}
	mi := &file_hoard_data_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FetchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchRequest) ProtoMessage() {}

func (x *FetchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hoard_data_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchRequest.ProtoReflect.Descriptor instead.
func (*FetchRequest) Descriptor() ([]byte, []int) {
	return file_hoard_data_proto_rawDescGZIP(), []int{1}
}

func (x *FetchRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *FetchRequest) GetURL() string {
	if x != nil {
		return x.URL
	}
	return ""
}

func (x *FetchRequest) GetRequestMethod() string {
	if x != nil {
		return x.RequestMethod
	}
	return ""
}

func (x *FetchRequest) GetHeaders() map[string]string {
	if x != nil {
		return x.Headers
	}
	return nil
}

type ReadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID         string            `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Database   string            `protobuf:"bytes,2,opt,name=Database,proto3" json:"Database,omitempty"`
	Collection string            `protobuf:"bytes,3,opt,name=Collection,proto3" json:"Collection,omitempty"`
	Filters    map[string]string `protobuf:"bytes,4,rep,name=Filters,proto3" json:"Filters,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ReadRequest) Reset() {
	*x = ReadRequest{}
	mi := &file_hoard_data_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadRequest) ProtoMessage() {}

func (x *ReadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hoard_data_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadRequest.ProtoReflect.Descriptor instead.
func (*ReadRequest) Descriptor() ([]byte, []int) {
	return file_hoard_data_proto_rawDescGZIP(), []int{2}
}

func (x *ReadRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *ReadRequest) GetDatabase() string {
	if x != nil {
		return x.Database
	}
	return ""
}

func (x *ReadRequest) GetCollection() string {
	if x != nil {
		return x.Collection
	}
	return ""
}

func (x *ReadRequest) GetFilters() map[string]string {
	if x != nil {
		return x.Filters
	}
	return nil
}

type Document struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID           string            `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	DocumentMeta map[string]string `protobuf:"bytes,2,rep,name=DocumentMeta,proto3" json:"DocumentMeta,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Data         map[string]string `protobuf:"bytes,3,rep,name=Data,proto3" json:"Data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Document) Reset() {
	*x = Document{}
	mi := &file_hoard_data_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Document) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Document) ProtoMessage() {}

func (x *Document) ProtoReflect() protoreflect.Message {
	mi := &file_hoard_data_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Document.ProtoReflect.Descriptor instead.
func (*Document) Descriptor() ([]byte, []int) {
	return file_hoard_data_proto_rawDescGZIP(), []int{3}
}

func (x *Document) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Document) GetDocumentMeta() map[string]string {
	if x != nil {
		return x.DocumentMeta
	}
	return nil
}

func (x *Document) GetData() map[string]string {
	if x != nil {
		return x.Data
	}
	return nil
}

type DocumentMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LastModified string `protobuf:"bytes,1,opt,name=LastModified,proto3" json:"LastModified,omitempty"`
	Source       string `protobuf:"bytes,2,opt,name=Source,proto3" json:"Source,omitempty"`
}

func (x *DocumentMeta) Reset() {
	*x = DocumentMeta{}
	mi := &file_hoard_data_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DocumentMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DocumentMeta) ProtoMessage() {}

func (x *DocumentMeta) ProtoReflect() protoreflect.Message {
	mi := &file_hoard_data_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DocumentMeta.ProtoReflect.Descriptor instead.
func (*DocumentMeta) Descriptor() ([]byte, []int) {
	return file_hoard_data_proto_rawDescGZIP(), []int{4}
}

func (x *DocumentMeta) GetLastModified() string {
	if x != nil {
		return x.LastModified
	}
	return ""
}

func (x *DocumentMeta) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

var File_hoard_data_proto protoreflect.FileDescriptor

var file_hoard_data_proto_rawDesc = []byte{
	0x0a, 0x10, 0x68, 0x6f, 0x61, 0x72, 0x64, 0x2d, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x2d, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a,
	0x04, 0x42, 0x6f, 0x64, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x42, 0x6f, 0x64,
	0x79, 0x22, 0xc8, 0x01, 0x0a, 0x0c, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x49, 0x44, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x52, 0x4c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x55, 0x52, 0x4c, 0x12, 0x24, 0x0a, 0x0d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x34, 0x0a, 0x07, 0x48, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x46, 0x65,
	0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73,
	0x1a, 0x3a, 0x0a, 0x0c, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xca, 0x01, 0x0a,
	0x0b, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08,
	0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x43, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x33, 0x0a, 0x07, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x52, 0x65, 0x61, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x1a, 0x3a, 0x0a,
	0x0c, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xfe, 0x01, 0x0a, 0x08, 0x44, 0x6f,
	0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x3f, 0x0a, 0x0c, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65,
	0x6e, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x44,
	0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74,
	0x4d, 0x65, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x44, 0x6f, 0x63, 0x75, 0x6d,
	0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x27, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61,
	0x1a, 0x3f, 0x0a, 0x11, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x61,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x1a, 0x37, 0x0a, 0x09, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x4a, 0x0a, 0x0c, 0x44, 0x6f,
	0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x22, 0x0a, 0x0c, 0x4c, 0x61,
	0x73, 0x74, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x4c, 0x61, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x32, 0x7b, 0x0a, 0x0b, 0x44, 0x61, 0x74, 0x61, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x25, 0x0a, 0x09, 0x46, 0x65, 0x74, 0x63, 0x68, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x0d, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x09, 0x2e, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x08,
	0x52, 0x65, 0x61, 0x64, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0c, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x09, 0x2e, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x20, 0x0a, 0x09, 0x57, 0x72, 0x69, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x09,
	0x2e, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x08, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x54, 0x61, 0x63, 0x6f, 0x73, 0x34, 0x62, 0x72, 0x65, 0x6b, 0x6b, 0x79, 0x32, 0x2f,
	0x68, 0x6f, 0x61, 0x72, 0x64, 0x2d, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2f, 0x67, 0x6f, 0x2d, 0x67,
	0x72, 0x70, 0x63, 0x2f, 0x68, 0x6f, 0x61, 0x72, 0x64, 0x2d, 0x64, 0x61, 0x74, 0x61, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_hoard_data_proto_rawDescOnce sync.Once
	file_hoard_data_proto_rawDescData = file_hoard_data_proto_rawDesc
)

func file_hoard_data_proto_rawDescGZIP() []byte {
	file_hoard_data_proto_rawDescOnce.Do(func() {
		file_hoard_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_hoard_data_proto_rawDescData)
	})
	return file_hoard_data_proto_rawDescData
}

var file_hoard_data_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_hoard_data_proto_goTypes = []any{
	(*Message)(nil),      // 0: Message
	(*FetchRequest)(nil), // 1: FetchRequest
	(*ReadRequest)(nil),  // 2: ReadRequest
	(*Document)(nil),     // 3: Document
	(*DocumentMeta)(nil), // 4: DocumentMeta
	nil,                  // 5: FetchRequest.HeadersEntry
	nil,                  // 6: ReadRequest.FiltersEntry
	nil,                  // 7: Document.DocumentMetaEntry
	nil,                  // 8: Document.DataEntry
}
var file_hoard_data_proto_depIdxs = []int32{
	5, // 0: FetchRequest.Headers:type_name -> FetchRequest.HeadersEntry
	6, // 1: ReadRequest.Filters:type_name -> ReadRequest.FiltersEntry
	7, // 2: Document.DocumentMeta:type_name -> Document.DocumentMetaEntry
	8, // 3: Document.Data:type_name -> Document.DataEntry
	1, // 4: DataService.FetchData:input_type -> FetchRequest
	2, // 5: DataService.ReadData:input_type -> ReadRequest
	3, // 6: DataService.WriteData:input_type -> Document
	3, // 7: DataService.FetchData:output_type -> Document
	3, // 8: DataService.ReadData:output_type -> Document
	0, // 9: DataService.WriteData:output_type -> Message
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_hoard_data_proto_init() }
func file_hoard_data_proto_init() {
	if File_hoard_data_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_hoard_data_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_hoard_data_proto_goTypes,
		DependencyIndexes: file_hoard_data_proto_depIdxs,
		MessageInfos:      file_hoard_data_proto_msgTypes,
	}.Build()
	File_hoard_data_proto = out.File
	file_hoard_data_proto_rawDesc = nil
	file_hoard_data_proto_goTypes = nil
	file_hoard_data_proto_depIdxs = nil
}
