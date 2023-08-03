//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2023 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

package grpc

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ServerLiveRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServerLiveRequest) Reset()         { *m = ServerLiveRequest{} }
func (m *ServerLiveRequest) String() string { return proto.CompactTextString(m) }
func (*ServerLiveRequest) ProtoMessage()    {}
func (*ServerLiveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_69a05cebb4f54974, []int{0}
}

func (m *ServerLiveRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerLiveRequest.Unmarshal(m, b)
}
func (m *ServerLiveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerLiveRequest.Marshal(b, m, deterministic)
}
func (m *ServerLiveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerLiveRequest.Merge(m, src)
}
func (m *ServerLiveRequest) XXX_Size() int {
	return xxx_messageInfo_ServerLiveRequest.Size(m)
}
func (m *ServerLiveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerLiveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ServerLiveRequest proto.InternalMessageInfo

type ServerLiveResponse struct {
	// True if the inference server is live, false if not live.
	Live                 bool     `protobuf:"varint,1,opt,name=live,proto3" json:"live,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServerLiveResponse) Reset()         { *m = ServerLiveResponse{} }
func (m *ServerLiveResponse) String() string { return proto.CompactTextString(m) }
func (*ServerLiveResponse) ProtoMessage()    {}
func (*ServerLiveResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_69a05cebb4f54974, []int{1}
}

func (m *ServerLiveResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerLiveResponse.Unmarshal(m, b)
}
func (m *ServerLiveResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerLiveResponse.Marshal(b, m, deterministic)
}
func (m *ServerLiveResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerLiveResponse.Merge(m, src)
}
func (m *ServerLiveResponse) XXX_Size() int {
	return xxx_messageInfo_ServerLiveResponse.Size(m)
}
func (m *ServerLiveResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerLiveResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ServerLiveResponse proto.InternalMessageInfo

func (m *ServerLiveResponse) GetLive() bool {
	if m != nil {
		return m.Live
	}
	return false
}

type ServerReadyRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServerReadyRequest) Reset()         { *m = ServerReadyRequest{} }
func (m *ServerReadyRequest) String() string { return proto.CompactTextString(m) }
func (*ServerReadyRequest) ProtoMessage()    {}
func (*ServerReadyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_69a05cebb4f54974, []int{2}
}

func (m *ServerReadyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerReadyRequest.Unmarshal(m, b)
}
func (m *ServerReadyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerReadyRequest.Marshal(b, m, deterministic)
}
func (m *ServerReadyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerReadyRequest.Merge(m, src)
}
func (m *ServerReadyRequest) XXX_Size() int {
	return xxx_messageInfo_ServerReadyRequest.Size(m)
}
func (m *ServerReadyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerReadyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ServerReadyRequest proto.InternalMessageInfo

type ServerReadyResponse struct {
	// True if the inference server is ready, false if not ready.
	Ready                bool     `protobuf:"varint,1,opt,name=ready,proto3" json:"ready,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServerReadyResponse) Reset()         { *m = ServerReadyResponse{} }
func (m *ServerReadyResponse) String() string { return proto.CompactTextString(m) }
func (*ServerReadyResponse) ProtoMessage()    {}
func (*ServerReadyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_69a05cebb4f54974, []int{3}
}

func (m *ServerReadyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerReadyResponse.Unmarshal(m, b)
}
func (m *ServerReadyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerReadyResponse.Marshal(b, m, deterministic)
}
func (m *ServerReadyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerReadyResponse.Merge(m, src)
}
func (m *ServerReadyResponse) XXX_Size() int {
	return xxx_messageInfo_ServerReadyResponse.Size(m)
}
func (m *ServerReadyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerReadyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ServerReadyResponse proto.InternalMessageInfo

func (m *ServerReadyResponse) GetReady() bool {
	if m != nil {
		return m.Ready
	}
	return false
}

type ModelReadyRequest struct {
	// The name of the model to check for readiness.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The version of the model to check for readiness. If not given the
	// server will choose a version based on the model and internal policy.
	Version              string   `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ModelReadyRequest) Reset()         { *m = ModelReadyRequest{} }
func (m *ModelReadyRequest) String() string { return proto.CompactTextString(m) }
func (*ModelReadyRequest) ProtoMessage()    {}
func (*ModelReadyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_69a05cebb4f54974, []int{4}
}

func (m *ModelReadyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModelReadyRequest.Unmarshal(m, b)
}
func (m *ModelReadyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModelReadyRequest.Marshal(b, m, deterministic)
}
func (m *ModelReadyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModelReadyRequest.Merge(m, src)
}
func (m *ModelReadyRequest) XXX_Size() int {
	return xxx_messageInfo_ModelReadyRequest.Size(m)
}
func (m *ModelReadyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ModelReadyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ModelReadyRequest proto.InternalMessageInfo

func (m *ModelReadyRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ModelReadyRequest) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

type ModelReadyResponse struct {
	// True if the model is ready, false if not ready.
	Ready                bool     `protobuf:"varint,1,opt,name=ready,proto3" json:"ready,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ModelReadyResponse) Reset()         { *m = ModelReadyResponse{} }
func (m *ModelReadyResponse) String() string { return proto.CompactTextString(m) }
func (*ModelReadyResponse) ProtoMessage()    {}
func (*ModelReadyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_69a05cebb4f54974, []int{5}
}

func (m *ModelReadyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModelReadyResponse.Unmarshal(m, b)
}
func (m *ModelReadyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModelReadyResponse.Marshal(b, m, deterministic)
}
func (m *ModelReadyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModelReadyResponse.Merge(m, src)
}
func (m *ModelReadyResponse) XXX_Size() int {
	return xxx_messageInfo_ModelReadyResponse.Size(m)
}
func (m *ModelReadyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ModelReadyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ModelReadyResponse proto.InternalMessageInfo

func (m *ModelReadyResponse) GetReady() bool {
	if m != nil {
		return m.Ready
	}
	return false
}

type ServerMetadataRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServerMetadataRequest) Reset()         { *m = ServerMetadataRequest{} }
func (m *ServerMetadataRequest) String() string { return proto.CompactTextString(m) }
func (*ServerMetadataRequest) ProtoMessage()    {}
func (*ServerMetadataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_69a05cebb4f54974, []int{6}
}

func (m *ServerMetadataRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerMetadataRequest.Unmarshal(m, b)
}
func (m *ServerMetadataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerMetadataRequest.Marshal(b, m, deterministic)
}
func (m *ServerMetadataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerMetadataRequest.Merge(m, src)
}
func (m *ServerMetadataRequest) XXX_Size() int {
	return xxx_messageInfo_ServerMetadataRequest.Size(m)
}
func (m *ServerMetadataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerMetadataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ServerMetadataRequest proto.InternalMessageInfo

type ServerMetadataResponse struct {
	// The server name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The server version.
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	// The extensions supported by the server.
	Extensions           []string `protobuf:"bytes,3,rep,name=extensions,proto3" json:"extensions,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServerMetadataResponse) Reset()         { *m = ServerMetadataResponse{} }
func (m *ServerMetadataResponse) String() string { return proto.CompactTextString(m) }
func (*ServerMetadataResponse) ProtoMessage()    {}
func (*ServerMetadataResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_69a05cebb4f54974, []int{7}
}

func (m *ServerMetadataResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerMetadataResponse.Unmarshal(m, b)
}
func (m *ServerMetadataResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerMetadataResponse.Marshal(b, m, deterministic)
}
func (m *ServerMetadataResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerMetadataResponse.Merge(m, src)
}
func (m *ServerMetadataResponse) XXX_Size() int {
	return xxx_messageInfo_ServerMetadataResponse.Size(m)
}
func (m *ServerMetadataResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerMetadataResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ServerMetadataResponse proto.InternalMessageInfo

func (m *ServerMetadataResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ServerMetadataResponse) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *ServerMetadataResponse) GetExtensions() []string {
	if m != nil {
		return m.Extensions
	}
	return nil
}

type ModelMetadataRequest struct {
	// The name of the model.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The version of the model to check for readiness. If not given the
	// server will choose a version based on the model and internal policy.
	Version              string   `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ModelMetadataRequest) Reset()         { *m = ModelMetadataRequest{} }
func (m *ModelMetadataRequest) String() string { return proto.CompactTextString(m) }
func (*ModelMetadataRequest) ProtoMessage()    {}
func (*ModelMetadataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_69a05cebb4f54974, []int{8}
}

func (m *ModelMetadataRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModelMetadataRequest.Unmarshal(m, b)
}
func (m *ModelMetadataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModelMetadataRequest.Marshal(b, m, deterministic)
}
func (m *ModelMetadataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModelMetadataRequest.Merge(m, src)
}
func (m *ModelMetadataRequest) XXX_Size() int {
	return xxx_messageInfo_ModelMetadataRequest.Size(m)
}
func (m *ModelMetadataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ModelMetadataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ModelMetadataRequest proto.InternalMessageInfo

func (m *ModelMetadataRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ModelMetadataRequest) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

type ModelMetadataResponse struct {
	// The model name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The versions of the model available on the server.
	Versions []string `protobuf:"bytes,2,rep,name=versions,proto3" json:"versions,omitempty"`
	// The model's platform. See Platforms.
	Platform string `protobuf:"bytes,3,opt,name=platform,proto3" json:"platform,omitempty"`
	// The model's inputs.
	Inputs []*ModelMetadataResponse_TensorMetadata `protobuf:"bytes,4,rep,name=inputs,proto3" json:"inputs,omitempty"`
	// The model's outputs.
	Outputs              []*ModelMetadataResponse_TensorMetadata `protobuf:"bytes,5,rep,name=outputs,proto3" json:"outputs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                `json:"-"`
	XXX_unrecognized     []byte                                  `json:"-"`
	XXX_sizecache        int32                                   `json:"-"`
}

func (m *ModelMetadataResponse) Reset()         { *m = ModelMetadataResponse{} }
func (m *ModelMetadataResponse) String() string { return proto.CompactTextString(m) }
func (*ModelMetadataResponse) ProtoMessage()    {}
func (*ModelMetadataResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_69a05cebb4f54974, []int{9}
}

func (m *ModelMetadataResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModelMetadataResponse.Unmarshal(m, b)
}
func (m *ModelMetadataResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModelMetadataResponse.Marshal(b, m, deterministic)
}
func (m *ModelMetadataResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModelMetadataResponse.Merge(m, src)
}
func (m *ModelMetadataResponse) XXX_Size() int {
	return xxx_messageInfo_ModelMetadataResponse.Size(m)
}
func (m *ModelMetadataResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ModelMetadataResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ModelMetadataResponse proto.InternalMessageInfo

func (m *ModelMetadataResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ModelMetadataResponse) GetVersions() []string {
	if m != nil {
		return m.Versions
	}
	return nil
}

func (m *ModelMetadataResponse) GetPlatform() string {
	if m != nil {
		return m.Platform
	}
	return ""
}

func (m *ModelMetadataResponse) GetInputs() []*ModelMetadataResponse_TensorMetadata {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *ModelMetadataResponse) GetOutputs() []*ModelMetadataResponse_TensorMetadata {
	if m != nil {
		return m.Outputs
	}
	return nil
}

// Metadata for a tensor.
type ModelMetadataResponse_TensorMetadata struct {
	// The tensor name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The tensor data type.
	Datatype string `protobuf:"bytes,2,opt,name=datatype,proto3" json:"datatype,omitempty"`
	// The tensor shape. A variable-size dimension is represented
	// by a -1 value.
	Shape                []int64  `protobuf:"varint,3,rep,packed,name=shape,proto3" json:"shape,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ModelMetadataResponse_TensorMetadata) Reset()         { *m = ModelMetadataResponse_TensorMetadata{} }
func (m *ModelMetadataResponse_TensorMetadata) String() string { return proto.CompactTextString(m) }
func (*ModelMetadataResponse_TensorMetadata) ProtoMessage()    {}
func (*ModelMetadataResponse_TensorMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_69a05cebb4f54974, []int{9, 0}
}

func (m *ModelMetadataResponse_TensorMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModelMetadataResponse_TensorMetadata.Unmarshal(m, b)
}
func (m *ModelMetadataResponse_TensorMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModelMetadataResponse_TensorMetadata.Marshal(b, m, deterministic)
}
func (m *ModelMetadataResponse_TensorMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModelMetadataResponse_TensorMetadata.Merge(m, src)
}
func (m *ModelMetadataResponse_TensorMetadata) XXX_Size() int {
	return xxx_messageInfo_ModelMetadataResponse_TensorMetadata.Size(m)
}
func (m *ModelMetadataResponse_TensorMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_ModelMetadataResponse_TensorMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_ModelMetadataResponse_TensorMetadata proto.InternalMessageInfo

func (m *ModelMetadataResponse_TensorMetadata) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ModelMetadataResponse_TensorMetadata) GetDatatype() string {
	if m != nil {
		return m.Datatype
	}
	return ""
}

func (m *ModelMetadataResponse_TensorMetadata) GetShape() []int64 {
	if m != nil {
		return m.Shape
	}
	return nil
}

type ModelInferRequest struct {
	// The name of the model to use for inferencing.
	ModelName string `protobuf:"bytes,1,opt,name=model_name,json=modelName,proto3" json:"model_name,omitempty"`
	// The version of the model to use for inference. If not given the
	// server will choose a version based on the model and internal policy.
	ModelVersion string `protobuf:"bytes,2,opt,name=model_version,json=modelVersion,proto3" json:"model_version,omitempty"`
	// Optional identifier for the request. If specified will be
	// returned in the response.
	Id string `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	// Optional inference parameters.
	Parameters map[string]*InferParameter `protobuf:"bytes,4,rep,name=parameters,proto3" json:"parameters,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// The input tensors for the inference.
	Inputs []*ModelInferRequest_InferInputTensor `protobuf:"bytes,5,rep,name=inputs,proto3" json:"inputs,omitempty"`
	// The requested output tensors for the inference. Optional, if not
	// specified all outputs produced by the model will be returned.
	Outputs []*ModelInferRequest_InferRequestedOutputTensor `protobuf:"bytes,6,rep,name=outputs,proto3" json:"outputs,omitempty"`
	// The data contained in an input tensor can be represented in "raw"
	// bytes form or in the repeated type that matches the tensor's data
	// type. To use the raw representation 'raw_input_contents' must be
	// initialized with data for each tensor in the same order as
	// 'inputs'. For each tensor, the size of this content must match
	// what is expected by the tensor's shape and data type. The raw
	// data must be the flattened, one-dimensional, row-major order of
	// the tensor elements without any stride or padding between the
	// elements. Note that the FP16 and BF16 data types must be represented as
	// raw content as there is no specific data type for a 16-bit float type.
	//
	// If this field is specified then InferInputTensor::contents must
	// not be specified for any input tensor.
	RawInputContents     [][]byte `protobuf:"bytes,7,rep,name=raw_input_contents,json=rawInputContents,proto3" json:"raw_input_contents,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ModelInferRequest) Reset()         { *m = ModelInferRequest{} }
func (m *ModelInferRequest) String() string { return proto.CompactTextString(m) }
func (*ModelInferRequest) ProtoMessage()    {}
func (*ModelInferRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_69a05cebb4f54974, []int{10}
}

func (m *ModelInferRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModelInferRequest.Unmarshal(m, b)
}
func (m *ModelInferRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModelInferRequest.Marshal(b, m, deterministic)
}
func (m *ModelInferRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModelInferRequest.Merge(m, src)
}
func (m *ModelInferRequest) XXX_Size() int {
	return xxx_messageInfo_ModelInferRequest.Size(m)
}
func (m *ModelInferRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ModelInferRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ModelInferRequest proto.InternalMessageInfo

func (m *ModelInferRequest) GetModelName() string {
	if m != nil {
		return m.ModelName
	}
	return ""
}

func (m *ModelInferRequest) GetModelVersion() string {
	if m != nil {
		return m.ModelVersion
	}
	return ""
}

func (m *ModelInferRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ModelInferRequest) GetParameters() map[string]*InferParameter {
	if m != nil {
		return m.Parameters
	}
	return nil
}

func (m *ModelInferRequest) GetInputs() []*ModelInferRequest_InferInputTensor {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *ModelInferRequest) GetOutputs() []*ModelInferRequest_InferRequestedOutputTensor {
	if m != nil {
		return m.Outputs
	}
	return nil
}

func (m *ModelInferRequest) GetRawInputContents() [][]byte {
	if m != nil {
		return m.RawInputContents
	}
	return nil
}

// An input tensor for an inference request.
type ModelInferRequest_InferInputTensor struct {
	// The tensor name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The tensor data type.
	Datatype string `protobuf:"bytes,2,opt,name=datatype,proto3" json:"datatype,omitempty"`
	// The tensor shape.
	Shape []int64 `protobuf:"varint,3,rep,packed,name=shape,proto3" json:"shape,omitempty"`
	// Optional inference input tensor parameters.
	Parameters map[string]*InferParameter `protobuf:"bytes,4,rep,name=parameters,proto3" json:"parameters,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// The tensor contents using a data-type format. This field must
	// not be specified if "raw" tensor contents are being used for
	// the inference request.
	Contents             *InferTensorContents `protobuf:"bytes,5,opt,name=contents,proto3" json:"contents,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ModelInferRequest_InferInputTensor) Reset()         { *m = ModelInferRequest_InferInputTensor{} }
func (m *ModelInferRequest_InferInputTensor) String() string { return proto.CompactTextString(m) }
func (*ModelInferRequest_InferInputTensor) ProtoMessage()    {}
func (*ModelInferRequest_InferInputTensor) Descriptor() ([]byte, []int) {
	return fileDescriptor_69a05cebb4f54974, []int{10, 0}
}

func (m *ModelInferRequest_InferInputTensor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModelInferRequest_InferInputTensor.Unmarshal(m, b)
}
func (m *ModelInferRequest_InferInputTensor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModelInferRequest_InferInputTensor.Marshal(b, m, deterministic)
}
func (m *ModelInferRequest_InferInputTensor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModelInferRequest_InferInputTensor.Merge(m, src)
}
func (m *ModelInferRequest_InferInputTensor) XXX_Size() int {
	return xxx_messageInfo_ModelInferRequest_InferInputTensor.Size(m)
}
func (m *ModelInferRequest_InferInputTensor) XXX_DiscardUnknown() {
	xxx_messageInfo_ModelInferRequest_InferInputTensor.DiscardUnknown(m)
}

var xxx_messageInfo_ModelInferRequest_InferInputTensor proto.InternalMessageInfo

func (m *ModelInferRequest_InferInputTensor) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ModelInferRequest_InferInputTensor) GetDatatype() string {
	if m != nil {
		return m.Datatype
	}
	return ""
}

func (m *ModelInferRequest_InferInputTensor) GetShape() []int64 {
	if m != nil {
		return m.Shape
	}
	return nil
}

func (m *ModelInferRequest_InferInputTensor) GetParameters() map[string]*InferParameter {
	if m != nil {
		return m.Parameters
	}
	return nil
}

func (m *ModelInferRequest_InferInputTensor) GetContents() *InferTensorContents {
	if m != nil {
		return m.Contents
	}
	return nil
}

// An output tensor requested for an inference request.
type ModelInferRequest_InferRequestedOutputTensor struct {
	// The tensor name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Optional requested output tensor parameters.
	Parameters           map[string]*InferParameter `protobuf:"bytes,2,rep,name=parameters,proto3" json:"parameters,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *ModelInferRequest_InferRequestedOutputTensor) Reset() {
	*m = ModelInferRequest_InferRequestedOutputTensor{}
}
func (m *ModelInferRequest_InferRequestedOutputTensor) String() string {
	return proto.CompactTextString(m)
}
func (*ModelInferRequest_InferRequestedOutputTensor) ProtoMessage() {}
func (*ModelInferRequest_InferRequestedOutputTensor) Descriptor() ([]byte, []int) {
	return fileDescriptor_69a05cebb4f54974, []int{10, 1}
}

func (m *ModelInferRequest_InferRequestedOutputTensor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModelInferRequest_InferRequestedOutputTensor.Unmarshal(m, b)
}
func (m *ModelInferRequest_InferRequestedOutputTensor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModelInferRequest_InferRequestedOutputTensor.Marshal(b, m, deterministic)
}
func (m *ModelInferRequest_InferRequestedOutputTensor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModelInferRequest_InferRequestedOutputTensor.Merge(m, src)
}
func (m *ModelInferRequest_InferRequestedOutputTensor) XXX_Size() int {
	return xxx_messageInfo_ModelInferRequest_InferRequestedOutputTensor.Size(m)
}
func (m *ModelInferRequest_InferRequestedOutputTensor) XXX_DiscardUnknown() {
	xxx_messageInfo_ModelInferRequest_InferRequestedOutputTensor.DiscardUnknown(m)
}

var xxx_messageInfo_ModelInferRequest_InferRequestedOutputTensor proto.InternalMessageInfo

func (m *ModelInferRequest_InferRequestedOutputTensor) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ModelInferRequest_InferRequestedOutputTensor) GetParameters() map[string]*InferParameter {
	if m != nil {
		return m.Parameters
	}
	return nil
}

type ModelInferResponse struct {
	// The name of the model used for inference.
	ModelName string `protobuf:"bytes,1,opt,name=model_name,json=modelName,proto3" json:"model_name,omitempty"`
	// The version of the model used for inference.
	ModelVersion string `protobuf:"bytes,2,opt,name=model_version,json=modelVersion,proto3" json:"model_version,omitempty"`
	// The id of the inference request if one was specified.
	Id string `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	// Optional inference response parameters.
	Parameters map[string]*InferParameter `protobuf:"bytes,4,rep,name=parameters,proto3" json:"parameters,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// The output tensors holding inference results.
	Outputs []*ModelInferResponse_InferOutputTensor `protobuf:"bytes,5,rep,name=outputs,proto3" json:"outputs,omitempty"`
	// The data contained in an output tensor can be represented in
	// "raw" bytes form or in the repeated type that matches the
	// tensor's data type. To use the raw representation 'raw_output_contents'
	// must be initialized with data for each tensor in the same order as
	// 'outputs'. For each tensor, the size of this content must match
	// what is expected by the tensor's shape and data type. The raw
	// data must be the flattened, one-dimensional, row-major order of
	// the tensor elements without any stride or padding between the
	// elements. Note that the FP16 and BF16 data types must be represented as
	// raw content as there is no specific data type for a 16-bit float type.
	//
	// If this field is specified then InferOutputTensor::contents must
	// not be specified for any output tensor.
	RawOutputContents    [][]byte `protobuf:"bytes,6,rep,name=raw_output_contents,json=rawOutputContents,proto3" json:"raw_output_contents,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ModelInferResponse) Reset()         { *m = ModelInferResponse{} }
func (m *ModelInferResponse) String() string { return proto.CompactTextString(m) }
func (*ModelInferResponse) ProtoMessage()    {}
func (*ModelInferResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_69a05cebb4f54974, []int{11}
}

func (m *ModelInferResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModelInferResponse.Unmarshal(m, b)
}
func (m *ModelInferResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModelInferResponse.Marshal(b, m, deterministic)
}
func (m *ModelInferResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModelInferResponse.Merge(m, src)
}
func (m *ModelInferResponse) XXX_Size() int {
	return xxx_messageInfo_ModelInferResponse.Size(m)
}
func (m *ModelInferResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ModelInferResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ModelInferResponse proto.InternalMessageInfo

func (m *ModelInferResponse) GetModelName() string {
	if m != nil {
		return m.ModelName
	}
	return ""
}

func (m *ModelInferResponse) GetModelVersion() string {
	if m != nil {
		return m.ModelVersion
	}
	return ""
}

func (m *ModelInferResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ModelInferResponse) GetParameters() map[string]*InferParameter {
	if m != nil {
		return m.Parameters
	}
	return nil
}

func (m *ModelInferResponse) GetOutputs() []*ModelInferResponse_InferOutputTensor {
	if m != nil {
		return m.Outputs
	}
	return nil
}

func (m *ModelInferResponse) GetRawOutputContents() [][]byte {
	if m != nil {
		return m.RawOutputContents
	}
	return nil
}

// An output tensor returned for an inference request.
type ModelInferResponse_InferOutputTensor struct {
	// The tensor name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The tensor data type.
	Datatype string `protobuf:"bytes,2,opt,name=datatype,proto3" json:"datatype,omitempty"`
	// The tensor shape.
	Shape []int64 `protobuf:"varint,3,rep,packed,name=shape,proto3" json:"shape,omitempty"`
	// Optional output tensor parameters.
	Parameters map[string]*InferParameter `protobuf:"bytes,4,rep,name=parameters,proto3" json:"parameters,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// The tensor contents using a data-type format. This field must
	// not be specified if "raw" tensor contents are being used for
	// the inference response.
	Contents             *InferTensorContents `protobuf:"bytes,5,opt,name=contents,proto3" json:"contents,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ModelInferResponse_InferOutputTensor) Reset()         { *m = ModelInferResponse_InferOutputTensor{} }
func (m *ModelInferResponse_InferOutputTensor) String() string { return proto.CompactTextString(m) }
func (*ModelInferResponse_InferOutputTensor) ProtoMessage()    {}
func (*ModelInferResponse_InferOutputTensor) Descriptor() ([]byte, []int) {
	return fileDescriptor_69a05cebb4f54974, []int{11, 0}
}

func (m *ModelInferResponse_InferOutputTensor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModelInferResponse_InferOutputTensor.Unmarshal(m, b)
}
func (m *ModelInferResponse_InferOutputTensor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModelInferResponse_InferOutputTensor.Marshal(b, m, deterministic)
}
func (m *ModelInferResponse_InferOutputTensor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModelInferResponse_InferOutputTensor.Merge(m, src)
}
func (m *ModelInferResponse_InferOutputTensor) XXX_Size() int {
	return xxx_messageInfo_ModelInferResponse_InferOutputTensor.Size(m)
}
func (m *ModelInferResponse_InferOutputTensor) XXX_DiscardUnknown() {
	xxx_messageInfo_ModelInferResponse_InferOutputTensor.DiscardUnknown(m)
}

var xxx_messageInfo_ModelInferResponse_InferOutputTensor proto.InternalMessageInfo

func (m *ModelInferResponse_InferOutputTensor) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ModelInferResponse_InferOutputTensor) GetDatatype() string {
	if m != nil {
		return m.Datatype
	}
	return ""
}

func (m *ModelInferResponse_InferOutputTensor) GetShape() []int64 {
	if m != nil {
		return m.Shape
	}
	return nil
}

func (m *ModelInferResponse_InferOutputTensor) GetParameters() map[string]*InferParameter {
	if m != nil {
		return m.Parameters
	}
	return nil
}

func (m *ModelInferResponse_InferOutputTensor) GetContents() *InferTensorContents {
	if m != nil {
		return m.Contents
	}
	return nil
}

// An inference parameter value. The Parameters message describes a
// “name”/”value” pair, where the “name” is the name of the parameter
// and the “value” is a boolean, integer, or string corresponding to
// the parameter.
type InferParameter struct {
	// The parameter value can be a string, an int64, a boolean
	// or a message specific to a predefined parameter.
	//
	// Types that are valid to be assigned to ParameterChoice:
	//	*InferParameter_BoolParam
	//	*InferParameter_Int64Param
	//	*InferParameter_StringParam
	ParameterChoice      isInferParameter_ParameterChoice `protobuf_oneof:"parameter_choice"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *InferParameter) Reset()         { *m = InferParameter{} }
func (m *InferParameter) String() string { return proto.CompactTextString(m) }
func (*InferParameter) ProtoMessage()    {}
func (*InferParameter) Descriptor() ([]byte, []int) {
	return fileDescriptor_69a05cebb4f54974, []int{12}
}

func (m *InferParameter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InferParameter.Unmarshal(m, b)
}
func (m *InferParameter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InferParameter.Marshal(b, m, deterministic)
}
func (m *InferParameter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InferParameter.Merge(m, src)
}
func (m *InferParameter) XXX_Size() int {
	return xxx_messageInfo_InferParameter.Size(m)
}
func (m *InferParameter) XXX_DiscardUnknown() {
	xxx_messageInfo_InferParameter.DiscardUnknown(m)
}

var xxx_messageInfo_InferParameter proto.InternalMessageInfo

type isInferParameter_ParameterChoice interface {
	isInferParameter_ParameterChoice()
}

type InferParameter_BoolParam struct {
	BoolParam bool `protobuf:"varint,1,opt,name=bool_param,json=boolParam,proto3,oneof"`
}

type InferParameter_Int64Param struct {
	Int64Param int64 `protobuf:"varint,2,opt,name=int64_param,json=int64Param,proto3,oneof"`
}

type InferParameter_StringParam struct {
	StringParam string `protobuf:"bytes,3,opt,name=string_param,json=stringParam,proto3,oneof"`
}

func (*InferParameter_BoolParam) isInferParameter_ParameterChoice() {}

func (*InferParameter_Int64Param) isInferParameter_ParameterChoice() {}

func (*InferParameter_StringParam) isInferParameter_ParameterChoice() {}

func (m *InferParameter) GetParameterChoice() isInferParameter_ParameterChoice {
	if m != nil {
		return m.ParameterChoice
	}
	return nil
}

func (m *InferParameter) GetBoolParam() bool {
	if x, ok := m.GetParameterChoice().(*InferParameter_BoolParam); ok {
		return x.BoolParam
	}
	return false
}

func (m *InferParameter) GetInt64Param() int64 {
	if x, ok := m.GetParameterChoice().(*InferParameter_Int64Param); ok {
		return x.Int64Param
	}
	return 0
}

func (m *InferParameter) GetStringParam() string {
	if x, ok := m.GetParameterChoice().(*InferParameter_StringParam); ok {
		return x.StringParam
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*InferParameter) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*InferParameter_BoolParam)(nil),
		(*InferParameter_Int64Param)(nil),
		(*InferParameter_StringParam)(nil),
	}
}

// The data contained in a tensor represented by the repeated type
// that matches the tensor's data type. Protobuf oneof is not used
// because oneofs cannot contain repeated fields.
type InferTensorContents struct {
	// Representation for BOOL data type. The size must match what is
	// expected by the tensor's shape. The contents must be the flattened,
	// one-dimensional, row-major order of the tensor elements.
	BoolContents []bool `protobuf:"varint,1,rep,packed,name=bool_contents,json=boolContents,proto3" json:"bool_contents,omitempty"`
	// Representation for INT8, INT16, and INT32 data types. The size
	// must match what is expected by the tensor's shape. The contents
	// must be the flattened, one-dimensional, row-major order of the
	// tensor elements.
	IntContents []int32 `protobuf:"varint,2,rep,packed,name=int_contents,json=intContents,proto3" json:"int_contents,omitempty"`
	// Representation for INT64 data types. The size must match what
	// is expected by the tensor's shape. The contents must be the
	// flattened, one-dimensional, row-major order of the tensor elements.
	Int64Contents []int64 `protobuf:"varint,3,rep,packed,name=int64_contents,json=int64Contents,proto3" json:"int64_contents,omitempty"`
	// Representation for UINT8, UINT16, and UINT32 data types. The size
	// must match what is expected by the tensor's shape. The contents
	// must be the flattened, one-dimensional, row-major order of the
	// tensor elements.
	UintContents []uint32 `protobuf:"varint,4,rep,packed,name=uint_contents,json=uintContents,proto3" json:"uint_contents,omitempty"`
	// Representation for UINT64 data types. The size must match what
	// is expected by the tensor's shape. The contents must be the
	// flattened, one-dimensional, row-major order of the tensor elements.
	Uint64Contents []uint64 `protobuf:"varint,5,rep,packed,name=uint64_contents,json=uint64Contents,proto3" json:"uint64_contents,omitempty"`
	// Representation for FP32 data type. The size must match what is
	// expected by the tensor's shape. The contents must be the flattened,
	// one-dimensional, row-major order of the tensor elements.
	Fp32Contents []float32 `protobuf:"fixed32,6,rep,packed,name=fp32_contents,json=fp32Contents,proto3" json:"fp32_contents,omitempty"`
	// Representation for FP64 data type. The size must match what is
	// expected by the tensor's shape. The contents must be the flattened,
	// one-dimensional, row-major order of the tensor elements.
	Fp64Contents []float64 `protobuf:"fixed64,7,rep,packed,name=fp64_contents,json=fp64Contents,proto3" json:"fp64_contents,omitempty"`
	// Representation for BYTES data type. The size must match what is
	// expected by the tensor's shape. The contents must be the flattened,
	// one-dimensional, row-major order of the tensor elements.
	BytesContents        [][]byte `protobuf:"bytes,8,rep,name=bytes_contents,json=bytesContents,proto3" json:"bytes_contents,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InferTensorContents) Reset()         { *m = InferTensorContents{} }
func (m *InferTensorContents) String() string { return proto.CompactTextString(m) }
func (*InferTensorContents) ProtoMessage()    {}
func (*InferTensorContents) Descriptor() ([]byte, []int) {
	return fileDescriptor_69a05cebb4f54974, []int{13}
}

func (m *InferTensorContents) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InferTensorContents.Unmarshal(m, b)
}
func (m *InferTensorContents) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InferTensorContents.Marshal(b, m, deterministic)
}
func (m *InferTensorContents) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InferTensorContents.Merge(m, src)
}
func (m *InferTensorContents) XXX_Size() int {
	return xxx_messageInfo_InferTensorContents.Size(m)
}
func (m *InferTensorContents) XXX_DiscardUnknown() {
	xxx_messageInfo_InferTensorContents.DiscardUnknown(m)
}

var xxx_messageInfo_InferTensorContents proto.InternalMessageInfo

func (m *InferTensorContents) GetBoolContents() []bool {
	if m != nil {
		return m.BoolContents
	}
	return nil
}

func (m *InferTensorContents) GetIntContents() []int32 {
	if m != nil {
		return m.IntContents
	}
	return nil
}

func (m *InferTensorContents) GetInt64Contents() []int64 {
	if m != nil {
		return m.Int64Contents
	}
	return nil
}

func (m *InferTensorContents) GetUintContents() []uint32 {
	if m != nil {
		return m.UintContents
	}
	return nil
}

func (m *InferTensorContents) GetUint64Contents() []uint64 {
	if m != nil {
		return m.Uint64Contents
	}
	return nil
}

func (m *InferTensorContents) GetFp32Contents() []float32 {
	if m != nil {
		return m.Fp32Contents
	}
	return nil
}

func (m *InferTensorContents) GetFp64Contents() []float64 {
	if m != nil {
		return m.Fp64Contents
	}
	return nil
}

func (m *InferTensorContents) GetBytesContents() [][]byte {
	if m != nil {
		return m.BytesContents
	}
	return nil
}

func init() {
	proto.RegisterType((*ServerLiveRequest)(nil), "inference.ServerLiveRequest")
	proto.RegisterType((*ServerLiveResponse)(nil), "inference.ServerLiveResponse")
	proto.RegisterType((*ServerReadyRequest)(nil), "inference.ServerReadyRequest")
	proto.RegisterType((*ServerReadyResponse)(nil), "inference.ServerReadyResponse")
	proto.RegisterType((*ModelReadyRequest)(nil), "inference.ModelReadyRequest")
	proto.RegisterType((*ModelReadyResponse)(nil), "inference.ModelReadyResponse")
	proto.RegisterType((*ServerMetadataRequest)(nil), "inference.ServerMetadataRequest")
	proto.RegisterType((*ServerMetadataResponse)(nil), "inference.ServerMetadataResponse")
	proto.RegisterType((*ModelMetadataRequest)(nil), "inference.ModelMetadataRequest")
	proto.RegisterType((*ModelMetadataResponse)(nil), "inference.ModelMetadataResponse")
	proto.RegisterType((*ModelMetadataResponse_TensorMetadata)(nil), "inference.ModelMetadataResponse.TensorMetadata")
	proto.RegisterType((*ModelInferRequest)(nil), "inference.ModelInferRequest")
	proto.RegisterMapType((map[string]*InferParameter)(nil), "inference.ModelInferRequest.ParametersEntry")
	proto.RegisterType((*ModelInferRequest_InferInputTensor)(nil), "inference.ModelInferRequest.InferInputTensor")
	proto.RegisterMapType((map[string]*InferParameter)(nil), "inference.ModelInferRequest.InferInputTensor.ParametersEntry")
	proto.RegisterType((*ModelInferRequest_InferRequestedOutputTensor)(nil), "inference.ModelInferRequest.InferRequestedOutputTensor")
	proto.RegisterMapType((map[string]*InferParameter)(nil), "inference.ModelInferRequest.InferRequestedOutputTensor.ParametersEntry")
	proto.RegisterType((*ModelInferResponse)(nil), "inference.ModelInferResponse")
	proto.RegisterMapType((map[string]*InferParameter)(nil), "inference.ModelInferResponse.ParametersEntry")
	proto.RegisterType((*ModelInferResponse_InferOutputTensor)(nil), "inference.ModelInferResponse.InferOutputTensor")
	proto.RegisterMapType((map[string]*InferParameter)(nil), "inference.ModelInferResponse.InferOutputTensor.ParametersEntry")
	proto.RegisterType((*InferParameter)(nil), "inference.InferParameter")
	proto.RegisterType((*InferTensorContents)(nil), "inference.InferTensorContents")
}

func init() {
	proto.RegisterFile("modules/text2vec-kserve/grpc/service.proto", fileDescriptor_69a05cebb4f54974)
}

var fileDescriptor_69a05cebb4f54974 = []byte{
	// 1008 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x57, 0xdd, 0x6e, 0x1b, 0x45,
	0x14, 0xb6, 0x77, 0xe3, 0xc4, 0x3e, 0xfe, 0x69, 0x32, 0x49, 0xc1, 0xac, 0xda, 0xd4, 0xd9, 0xaa,
	0xc2, 0x2a, 0xad, 0x2d, 0xb9, 0x08, 0x50, 0x25, 0x84, 0x68, 0xa9, 0x92, 0xa8, 0x4d, 0x29, 0xd3,
	0xaa, 0x20, 0x24, 0x64, 0x6d, 0xec, 0x49, 0xba, 0xaa, 0xbd, 0xbb, 0xec, 0xce, 0x3a, 0xf5, 0x43,
	0xf0, 0x08, 0x3c, 0x0d, 0xd7, 0x08, 0xde, 0x80, 0x87, 0xe0, 0x96, 0x0b, 0x34, 0x3f, 0x3b, 0x9e,
	0xdd, 0xb5, 0x37, 0xa4, 0x4a, 0x2e, 0xb8, 0xdb, 0x73, 0xe6, 0x9b, 0x6f, 0xce, 0xef, 0xec, 0x19,
	0xb8, 0x3b, 0xf5, 0xc7, 0xf1, 0x84, 0x44, 0x7d, 0x4a, 0xde, 0xd1, 0xc1, 0x8c, 0x8c, 0xee, 0xbf,
	0x8d, 0x48, 0x38, 0x23, 0xfd, 0xd3, 0x30, 0x18, 0xf5, 0xd9, 0xa7, 0x3b, 0x22, 0xbd, 0x20, 0xf4,
	0xa9, 0x8f, 0x6a, 0xae, 0x77, 0x42, 0x42, 0xe2, 0x8d, 0x88, 0xbd, 0x0d, 0x5b, 0x2f, 0x19, 0x2c,
	0x7c, 0xe6, 0xce, 0x08, 0x26, 0x3f, 0xc7, 0x24, 0xa2, 0x76, 0x17, 0x90, 0xae, 0x8c, 0x02, 0xdf,
	0x8b, 0x08, 0x42, 0xb0, 0x36, 0x71, 0x67, 0xa4, 0x5d, 0xee, 0x94, 0xbb, 0x55, 0xcc, 0xbf, 0xed,
	0x9d, 0x04, 0x89, 0x89, 0x33, 0x9e, 0x27, 0xfb, 0x3f, 0x81, 0xed, 0x94, 0x56, 0x12, 0xec, 0x40,
	0x25, 0x64, 0x0a, 0xc9, 0x20, 0x04, 0xfb, 0x6b, 0xd8, 0x3a, 0xf2, 0xc7, 0x64, 0xa2, 0x33, 0xb0,
	0xb3, 0x3c, 0x67, 0x2a, 0xce, 0xaa, 0x61, 0xfe, 0x8d, 0xda, 0xb0, 0x31, 0x23, 0x61, 0xe4, 0xfa,
	0x5e, 0xdb, 0xe0, 0xea, 0x44, 0xb4, 0xef, 0x02, 0xd2, 0x29, 0x0a, 0x8f, 0xfb, 0x10, 0xae, 0x0b,
	0xdb, 0x8e, 0x08, 0x75, 0xc6, 0x0e, 0x75, 0x12, 0xa3, 0x4f, 0xe0, 0x83, 0xec, 0xc2, 0xc2, 0xf1,
	0xff, 0x6e, 0x0c, 0xda, 0x05, 0x20, 0xef, 0x28, 0xf1, 0x98, 0x10, 0xb5, 0xcd, 0x8e, 0xd9, 0xad,
	0x61, 0x4d, 0x63, 0x7f, 0x03, 0x3b, 0xdc, 0xd8, 0xcc, 0xf9, 0x17, 0x74, 0xf9, 0x0f, 0x03, 0xae,
	0x67, 0x68, 0x0a, 0xac, 0xb5, 0xa0, 0x2a, 0x37, 0x46, 0x6d, 0x83, 0x5b, 0xa4, 0x64, 0xb6, 0x16,
	0x4c, 0x1c, 0x7a, 0xe2, 0x87, 0xd3, 0xb6, 0xc9, 0xf7, 0x28, 0x19, 0xed, 0xc3, 0xba, 0xeb, 0x05,
	0x31, 0x8d, 0xda, 0x6b, 0x1d, 0xb3, 0x5b, 0x1f, 0xf4, 0x7b, 0xaa, 0x72, 0x7a, 0x4b, 0x4f, 0xef,
	0xbd, 0x22, 0x5e, 0xe4, 0x2f, 0x42, 0x28, 0xb7, 0xa3, 0x43, 0xd8, 0xf0, 0x63, 0xca, 0x99, 0x2a,
	0xef, 0xc7, 0x94, 0xec, 0xb7, 0x5e, 0x43, 0x2b, 0xbd, 0xb4, 0xca, 0x63, 0xb6, 0x46, 0xe7, 0x01,
	0x91, 0xa1, 0x53, 0x32, 0x2b, 0x8c, 0xe8, 0x8d, 0x13, 0x10, 0x9e, 0x1c, 0x13, 0x0b, 0xc1, 0xfe,
	0xb5, 0x2a, 0x0b, 0xf1, 0x90, 0x19, 0x96, 0x64, 0xe5, 0x26, 0xc0, 0x94, 0x29, 0x87, 0xda, 0x09,
	0x35, 0xae, 0x79, 0xce, 0x8e, 0xb9, 0x0d, 0x4d, 0xb1, 0x9c, 0x4e, 0x53, 0x83, 0x2b, 0x5f, 0xcb,
	0x8a, 0x68, 0x81, 0xe1, 0x8e, 0x65, 0x6c, 0x0d, 0x77, 0x8c, 0x9e, 0x01, 0x04, 0x4e, 0xe8, 0x4c,
	0x09, 0x25, 0x61, 0x12, 0xd9, 0x7b, 0xd9, 0x78, 0xe8, 0x56, 0xf4, 0x5e, 0x28, 0xf8, 0x13, 0x8f,
	0x86, 0x73, 0xac, 0xed, 0x47, 0x4f, 0x54, 0x8e, 0x44, 0x64, 0xef, 0x17, 0x32, 0x71, 0xe1, 0x90,
	0xe1, 0x45, 0x10, 0x55, 0x86, 0xbe, 0x5b, 0x64, 0x68, 0x9d, 0xf3, 0x7c, 0x7e, 0x3e, 0x8f, 0x14,
	0xc8, 0xf8, 0x5b, 0xbe, 0x53, 0x32, 0x26, 0x3c, 0xe8, 0x1e, 0xa0, 0xd0, 0x39, 0x1b, 0xf2, 0x03,
	0x86, 0x23, 0xdf, 0xa3, 0xc4, 0xa3, 0x51, 0x7b, 0xa3, 0x63, 0x76, 0x1b, 0x78, 0x33, 0x74, 0xce,
	0xb8, 0x19, 0x8f, 0xa5, 0xde, 0xfa, 0xdd, 0x80, 0xcd, 0xac, 0x75, 0x97, 0x93, 0x5a, 0xf4, 0xd3,
	0x92, 0x80, 0x7f, 0x79, 0xa1, 0x30, 0x15, 0x66, 0xe0, 0x21, 0x54, 0x95, 0x77, 0x95, 0x4e, 0xb9,
	0x5b, 0x1f, 0xec, 0x6a, 0xe4, 0x9c, 0x4a, 0xb0, 0x24, 0xbe, 0x62, 0x85, 0xb7, 0x7e, 0x80, 0x6b,
	0x19, 0x6a, 0xb4, 0x09, 0xe6, 0x5b, 0x32, 0x97, 0x2e, 0xb3, 0x4f, 0xd4, 0x87, 0xca, 0xcc, 0x99,
	0xc4, 0xc2, 0xdd, 0xfa, 0xe0, 0xa3, 0x2c, 0xbb, 0x62, 0xc0, 0x02, 0xf7, 0xd0, 0xf8, 0xa2, 0x6c,
	0xfd, 0x5d, 0x06, 0x6b, 0x75, 0x96, 0x96, 0x46, 0xf6, 0x34, 0x15, 0x27, 0x83, 0xc7, 0x69, 0xff,
	0x3d, 0xcb, 0xa0, 0x28, 0x62, 0x57, 0xe8, 0xf5, 0x95, 0x31, 0xdb, 0x7f, 0x55, 0xe4, 0x5f, 0x46,
	0xfa, 0x2c, 0xaf, 0xdb, 0xab, 0xb8, 0x20, 0x8e, 0x96, 0xd4, 0xeb, 0xaa, 0xb6, 0x96, 0xb7, 0x65,
	0x51, 0x7d, 0x9e, 0x7f, 0xf9, 0xa6, 0xb9, 0xb8, 0xb4, 0xbc, 0xa5, 0x7b, 0xb0, 0xcd, 0x5a, 0x5a,
	0x88, 0x8b, 0x9e, 0x5e, 0xe7, 0x3d, 0xbd, 0x15, 0x3a, 0x67, 0x62, 0x9b, 0x6a, 0xea, 0x3f, 0x0d,
	0xd8, 0xca, 0xd1, 0x5d, 0x52, 0x57, 0x0f, 0x97, 0x44, 0xe9, 0xab, 0x0b, 0x7a, 0xf6, 0x3f, 0xec,
	0xeb, 0xab, 0xab, 0xf0, 0x5f, 0xca, 0xd0, 0x4a, 0xaf, 0xa2, 0x5b, 0x00, 0xc7, 0xbe, 0x3f, 0x19,
	0xf2, 0xa8, 0x88, 0x41, 0xea, 0xa0, 0x84, 0x6b, 0x4c, 0xc7, 0x41, 0x68, 0x0f, 0xea, 0xae, 0x47,
	0x3f, 0xfb, 0x54, 0x22, 0xd8, 0x71, 0xe6, 0x41, 0x09, 0x03, 0x57, 0x0a, 0xc8, 0x6d, 0x68, 0x44,
	0x34, 0x74, 0xbd, 0x53, 0x89, 0xe1, 0x75, 0x7e, 0x50, 0xc2, 0x75, 0xa1, 0xe5, 0xa0, 0x47, 0x08,
	0x36, 0x55, 0xe4, 0x87, 0xa3, 0x37, 0xbe, 0x3b, 0x22, 0xf6, 0x6f, 0x06, 0x6c, 0x2f, 0x89, 0x32,
	0xeb, 0x29, 0x6e, 0x94, 0x4a, 0x4e, 0xb9, 0x63, 0x76, 0xab, 0xb8, 0xc1, 0x94, 0x0a, 0xb4, 0x07,
	0x0d, 0xd7, 0xd3, 0x4a, 0x94, 0xdd, 0x66, 0x15, 0xcc, 0x8c, 0x55, 0x90, 0x3b, 0xd0, 0x12, 0xb6,
	0x2b, 0x90, 0xa8, 0xaf, 0x26, 0xd7, 0xea, 0xc7, 0xc5, 0x29, 0x2a, 0x56, 0x6a, 0x4d, 0xdc, 0x88,
	0x75, 0xae, 0x8f, 0xe1, 0x5a, 0x9c, 0x21, 0x63, 0xbd, 0xb6, 0x86, 0x5b, 0x71, 0x8e, 0xed, 0x24,
	0x78, 0x30, 0x48, 0xf7, 0x8e, 0x81, 0x1b, 0x4c, 0x99, 0x06, 0xe9, 0x5c, 0xec, 0xa7, 0x59, 0x66,
	0x20, 0x8d, 0xe9, 0x0e, 0xb4, 0x8e, 0xe7, 0x94, 0x44, 0x0b, 0x54, 0x95, 0xb7, 0x61, 0x93, 0x6b,
	0x13, 0xd8, 0xe0, 0x1f, 0x13, 0x76, 0xf6, 0xf1, 0x8b, 0xc7, 0x87, 0x49, 0x05, 0xbc, 0x14, 0x6f,
	0x01, 0xf4, 0x14, 0x60, 0x31, 0xe5, 0xa3, 0x1b, 0x5a, 0x89, 0xe4, 0x5e, 0x04, 0xd6, 0xcd, 0x15,
	0xab, 0xa2, 0xaf, 0xec, 0x12, 0x7a, 0x0e, 0x75, 0x6d, 0xe4, 0x47, 0x79, 0xbc, 0x3e, 0xde, 0x5b,
	0xbb, 0xab, 0x96, 0x15, 0xdf, 0x53, 0x80, 0xc5, 0x48, 0x9f, 0x32, 0x2e, 0xf7, 0x58, 0x48, 0x19,
	0x97, 0x7f, 0x07, 0xd8, 0x25, 0xf4, 0x3d, 0xb4, 0xd2, 0xa3, 0x3d, 0xea, 0xe4, 0x0c, 0xc8, 0x8c,
	0xe3, 0xd6, 0x5e, 0x01, 0x42, 0x11, 0xbf, 0x82, 0x66, 0x6a, 0x78, 0x45, 0xb7, 0x56, 0x8f, 0xb5,
	0x82, 0xb6, 0x73, 0xde, 0xdc, 0xab, 0xf9, 0xce, 0x33, 0x96, 0xf7, 0x5d, 0xff, 0xe7, 0xe6, 0x7d,
	0x4f, 0x5d, 0x78, 0x76, 0xe9, 0xd1, 0xee, 0x8f, 0x37, 0x8a, 0x5e, 0x86, 0xc7, 0xeb, 0xfc, 0x49,
	0xf8, 0xe0, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb9, 0x82, 0xd2, 0x56, 0x40, 0x0e, 0x00, 0x00,
}
