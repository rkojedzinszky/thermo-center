// Configuration servicer interface

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.17.3
// source: configurator/api.proto

package configurator

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

type RadioCfgRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cluster uint32 `protobuf:"varint,1,opt,name=cluster,proto3" json:"cluster,omitempty"`
}

func (x *RadioCfgRequest) Reset() {
	*x = RadioCfgRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_configurator_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RadioCfgRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RadioCfgRequest) ProtoMessage() {}

func (x *RadioCfgRequest) ProtoReflect() protoreflect.Message {
	mi := &file_configurator_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RadioCfgRequest.ProtoReflect.Descriptor instead.
func (*RadioCfgRequest) Descriptor() ([]byte, []int) {
	return file_configurator_api_proto_rawDescGZIP(), []int{0}
}

func (x *RadioCfgRequest) GetCluster() uint32 {
	if x != nil {
		return x.Cluster
	}
	return 0
}

type RadioCfgResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Network     uint32 `protobuf:"varint,1,opt,name=network,proto3" json:"network,omitempty"`
	RadioConfig []byte `protobuf:"bytes,2,opt,name=radio_config,json=radioConfig,proto3" json:"radio_config,omitempty"`
	AesKey      []byte `protobuf:"bytes,3,opt,name=aes_key,json=aesKey,proto3" json:"aes_key,omitempty"`
}

func (x *RadioCfgResponse) Reset() {
	*x = RadioCfgResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_configurator_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RadioCfgResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RadioCfgResponse) ProtoMessage() {}

func (x *RadioCfgResponse) ProtoReflect() protoreflect.Message {
	mi := &file_configurator_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RadioCfgResponse.ProtoReflect.Descriptor instead.
func (*RadioCfgResponse) Descriptor() ([]byte, []int) {
	return file_configurator_api_proto_rawDescGZIP(), []int{1}
}

func (x *RadioCfgResponse) GetNetwork() uint32 {
	if x != nil {
		return x.Network
	}
	return 0
}

func (x *RadioCfgResponse) GetRadioConfig() []byte {
	if x != nil {
		return x.RadioConfig
	}
	return nil
}

func (x *RadioCfgResponse) GetAesKey() []byte {
	if x != nil {
		return x.AesKey
	}
	return nil
}

type Task struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskId uint32 `protobuf:"varint,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
}

func (x *Task) Reset() {
	*x = Task{}
	if protoimpl.UnsafeEnabled {
		mi := &file_configurator_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task) ProtoMessage() {}

func (x *Task) ProtoReflect() protoreflect.Message {
	mi := &file_configurator_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Task.ProtoReflect.Descriptor instead.
func (*Task) Descriptor() ([]byte, []int) {
	return file_configurator_api_proto_rawDescGZIP(), []int{2}
}

func (x *Task) GetTaskId() uint32 {
	if x != nil {
		return x.TaskId
	}
	return 0
}

type TaskDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskId   uint32            `protobuf:"varint,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	SensorId uint32            `protobuf:"varint,2,opt,name=sensor_id,json=sensorId,proto3" json:"sensor_id,omitempty"`
	Config   *RadioCfgResponse `protobuf:"bytes,3,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *TaskDetails) Reset() {
	*x = TaskDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_configurator_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskDetails) ProtoMessage() {}

func (x *TaskDetails) ProtoReflect() protoreflect.Message {
	mi := &file_configurator_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskDetails.ProtoReflect.Descriptor instead.
func (*TaskDetails) Descriptor() ([]byte, []int) {
	return file_configurator_api_proto_rawDescGZIP(), []int{3}
}

func (x *TaskDetails) GetTaskId() uint32 {
	if x != nil {
		return x.TaskId
	}
	return 0
}

func (x *TaskDetails) GetSensorId() uint32 {
	if x != nil {
		return x.SensorId
	}
	return 0
}

func (x *TaskDetails) GetConfig() *RadioCfgResponse {
	if x != nil {
		return x.Config
	}
	return nil
}

type TaskUpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *TaskUpdateResponse) Reset() {
	*x = TaskUpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_configurator_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskUpdateResponse) ProtoMessage() {}

func (x *TaskUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_configurator_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskUpdateResponse.ProtoReflect.Descriptor instead.
func (*TaskUpdateResponse) Descriptor() ([]byte, []int) {
	return file_configurator_api_proto_rawDescGZIP(), []int{4}
}

func (x *TaskUpdateResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type TaskFinishedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskId uint32 `protobuf:"varint,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	Error  string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *TaskFinishedRequest) Reset() {
	*x = TaskFinishedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_configurator_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskFinishedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskFinishedRequest) ProtoMessage() {}

func (x *TaskFinishedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_configurator_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskFinishedRequest.ProtoReflect.Descriptor instead.
func (*TaskFinishedRequest) Descriptor() ([]byte, []int) {
	return file_configurator_api_proto_rawDescGZIP(), []int{5}
}

func (x *TaskFinishedRequest) GetTaskId() uint32 {
	if x != nil {
		return x.TaskId
	}
	return 0
}

func (x *TaskFinishedRequest) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_configurator_api_proto protoreflect.FileDescriptor

var file_configurator_api_proto_rawDesc = []byte{
	0x0a, 0x16, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x61,
	0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x75, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x22, 0x2b, 0x0a, 0x0f, 0x52, 0x61, 0x64, 0x69, 0x6f, 0x43,
	0x66, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x22, 0x68, 0x0a, 0x10, 0x52, 0x61, 0x64, 0x69, 0x6f, 0x43, 0x66, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x61, 0x64, 0x69, 0x6f, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x72, 0x61, 0x64, 0x69, 0x6f, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x17, 0x0a, 0x07, 0x61, 0x65, 0x73, 0x5f, 0x6b, 0x65, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x61, 0x65, 0x73, 0x4b, 0x65, 0x79, 0x22, 0x1f, 0x0a,
	0x04, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x22, 0x7b,
	0x0a, 0x0b, 0x54, 0x61, 0x73, 0x6b, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x17, 0x0a,
	0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06,
	0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x73, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x49, 0x64, 0x12, 0x36, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x2e, 0x52, 0x61, 0x64, 0x69, 0x6f, 0x43, 0x66, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x2e, 0x0a, 0x12, 0x54,
	0x61, 0x73, 0x6b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x44, 0x0a, 0x13, 0x54,
	0x61, 0x73, 0x6b, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x32, 0xc6, 0x02, 0x0a, 0x0c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x12, 0x4e, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x52, 0x61, 0x64, 0x69, 0x6f, 0x43, 0x66,
	0x67, 0x12, 0x1d, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x2e, 0x52, 0x61, 0x64, 0x69, 0x6f, 0x43, 0x66, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e,
	0x52, 0x61, 0x64, 0x69, 0x6f, 0x43, 0x66, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x3e, 0x0a, 0x0b, 0x54, 0x61, 0x73, 0x6b, 0x41, 0x63, 0x71, 0x75, 0x69, 0x72,
	0x65, 0x12, 0x12, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x2e, 0x54, 0x61, 0x73, 0x6b, 0x1a, 0x19, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x22, 0x00, 0x12, 0x4f, 0x0a, 0x15, 0x54, 0x61, 0x73, 0x6b, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x79, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x12, 0x12, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x1a,
	0x20, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x54,
	0x61, 0x73, 0x6b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x55, 0x0a, 0x0c, 0x54, 0x61, 0x73, 0x6b, 0x46, 0x69, 0x6e, 0x69, 0x73,
	0x68, 0x65, 0x64, 0x12, 0x21, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75,
	0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x6b, 0x6f, 0x6a, 0x65, 0x64, 0x7a,
	0x69, 0x6e, 0x73, 0x7a, 0x6b, 0x79, 0x2f, 0x74, 0x68, 0x65, 0x72, 0x6d, 0x6f, 0x2d, 0x63, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x35, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_configurator_api_proto_rawDescOnce sync.Once
	file_configurator_api_proto_rawDescData = file_configurator_api_proto_rawDesc
)

func file_configurator_api_proto_rawDescGZIP() []byte {
	file_configurator_api_proto_rawDescOnce.Do(func() {
		file_configurator_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_configurator_api_proto_rawDescData)
	})
	return file_configurator_api_proto_rawDescData
}

var file_configurator_api_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_configurator_api_proto_goTypes = []interface{}{
	(*RadioCfgRequest)(nil),     // 0: configurator.RadioCfgRequest
	(*RadioCfgResponse)(nil),    // 1: configurator.RadioCfgResponse
	(*Task)(nil),                // 2: configurator.Task
	(*TaskDetails)(nil),         // 3: configurator.TaskDetails
	(*TaskUpdateResponse)(nil),  // 4: configurator.TaskUpdateResponse
	(*TaskFinishedRequest)(nil), // 5: configurator.TaskFinishedRequest
}
var file_configurator_api_proto_depIdxs = []int32{
	1, // 0: configurator.TaskDetails.config:type_name -> configurator.RadioCfgResponse
	0, // 1: configurator.Configurator.GetRadioCfg:input_type -> configurator.RadioCfgRequest
	2, // 2: configurator.Configurator.TaskAcquire:input_type -> configurator.Task
	2, // 3: configurator.Configurator.TaskDiscoveryReceived:input_type -> configurator.Task
	5, // 4: configurator.Configurator.TaskFinished:input_type -> configurator.TaskFinishedRequest
	1, // 5: configurator.Configurator.GetRadioCfg:output_type -> configurator.RadioCfgResponse
	3, // 6: configurator.Configurator.TaskAcquire:output_type -> configurator.TaskDetails
	4, // 7: configurator.Configurator.TaskDiscoveryReceived:output_type -> configurator.TaskUpdateResponse
	4, // 8: configurator.Configurator.TaskFinished:output_type -> configurator.TaskUpdateResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_configurator_api_proto_init() }
func file_configurator_api_proto_init() {
	if File_configurator_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_configurator_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RadioCfgRequest); i {
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
		file_configurator_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RadioCfgResponse); i {
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
		file_configurator_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Task); i {
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
		file_configurator_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskDetails); i {
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
		file_configurator_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskUpdateResponse); i {
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
		file_configurator_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskFinishedRequest); i {
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
			RawDescriptor: file_configurator_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_configurator_api_proto_goTypes,
		DependencyIndexes: file_configurator_api_proto_depIdxs,
		MessageInfos:      file_configurator_api_proto_msgTypes,
	}.Build()
	File_configurator_api_proto = out.File
	file_configurator_api_proto_rawDesc = nil
	file_configurator_api_proto_goTypes = nil
	file_configurator_api_proto_depIdxs = nil
}
