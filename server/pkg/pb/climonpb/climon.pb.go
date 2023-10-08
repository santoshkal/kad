// Code generated by protoc-gen-go. DO NOT EDIT.
// source: climon.proto

package climonpb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
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

type DeployRequest struct {
	Version              string   `protobuf:"bytes,1,opt,name=Version,proto3" json:"Version,omitempty"`
	RepoUrl              string   `protobuf:"bytes,2,opt,name=RepoUrl,proto3" json:"RepoUrl,omitempty"`
	RepoName             string   `protobuf:"bytes,3,opt,name=RepoName,proto3" json:"RepoName,omitempty"`
	Namespace            string   `protobuf:"bytes,4,opt,name=Namespace,proto3" json:"Namespace,omitempty"`
	ChartName            string   `protobuf:"bytes,5,opt,name=ChartName,proto3" json:"ChartName,omitempty"`
	ReleaseName          string   `protobuf:"bytes,6,opt,name=ReleaseName,proto3" json:"ReleaseName,omitempty"`
	ReferenceID          string   `protobuf:"bytes,7,opt,name=ReferenceID,proto3" json:"ReferenceID,omitempty"`
	Plugin               string   `protobuf:"bytes,8,opt,name=Plugin,proto3" json:"Plugin,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeployRequest) Reset()         { *m = DeployRequest{} }
func (m *DeployRequest) String() string { return proto.CompactTextString(m) }
func (*DeployRequest) ProtoMessage()    {}
func (*DeployRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_89182b4db867c195, []int{0}
}

func (m *DeployRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeployRequest.Unmarshal(m, b)
}
func (m *DeployRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeployRequest.Marshal(b, m, deterministic)
}
func (m *DeployRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeployRequest.Merge(m, src)
}
func (m *DeployRequest) XXX_Size() int {
	return xxx_messageInfo_DeployRequest.Size(m)
}
func (m *DeployRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeployRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeployRequest proto.InternalMessageInfo

func (m *DeployRequest) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *DeployRequest) GetRepoUrl() string {
	if m != nil {
		return m.RepoUrl
	}
	return ""
}

func (m *DeployRequest) GetRepoName() string {
	if m != nil {
		return m.RepoName
	}
	return ""
}

func (m *DeployRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *DeployRequest) GetChartName() string {
	if m != nil {
		return m.ChartName
	}
	return ""
}

func (m *DeployRequest) GetReleaseName() string {
	if m != nil {
		return m.ReleaseName
	}
	return ""
}

func (m *DeployRequest) GetReferenceID() string {
	if m != nil {
		return m.ReferenceID
	}
	return ""
}

func (m *DeployRequest) GetPlugin() string {
	if m != nil {
		return m.Plugin
	}
	return ""
}

type DeployResponse struct {
	Status                 string            `protobuf:"bytes,1,opt,name=Status,proto3" json:"Status,omitempty"`
	WorkFlowResponseStatus *WorkFlowResponse `protobuf:"bytes,2,opt,name=WorkFlowResponseStatus,proto3" json:"WorkFlowResponseStatus,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}          `json:"-"`
	XXX_unrecognized       []byte            `json:"-"`
	XXX_sizecache          int32             `json:"-"`
}

func (m *DeployResponse) Reset()         { *m = DeployResponse{} }
func (m *DeployResponse) String() string { return proto.CompactTextString(m) }
func (*DeployResponse) ProtoMessage()    {}
func (*DeployResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_89182b4db867c195, []int{1}
}

func (m *DeployResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeployResponse.Unmarshal(m, b)
}
func (m *DeployResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeployResponse.Marshal(b, m, deterministic)
}
func (m *DeployResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeployResponse.Merge(m, src)
}
func (m *DeployResponse) XXX_Size() int {
	return xxx_messageInfo_DeployResponse.Size(m)
}
func (m *DeployResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeployResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeployResponse proto.InternalMessageInfo

func (m *DeployResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *DeployResponse) GetWorkFlowResponseStatus() *WorkFlowResponse {
	if m != nil {
		return m.WorkFlowResponseStatus
	}
	return nil
}

type WorkFlowResponse struct {
	WorkflowID           string   `protobuf:"bytes,1,opt,name=WorkflowID,proto3" json:"WorkflowID,omitempty"`
	WorkFlowRunID        string   `protobuf:"bytes,2,opt,name=WorkFlowRunID,proto3" json:"WorkFlowRunID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WorkFlowResponse) Reset()         { *m = WorkFlowResponse{} }
func (m *WorkFlowResponse) String() string { return proto.CompactTextString(m) }
func (*WorkFlowResponse) ProtoMessage()    {}
func (*WorkFlowResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_89182b4db867c195, []int{2}
}

func (m *WorkFlowResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkFlowResponse.Unmarshal(m, b)
}
func (m *WorkFlowResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkFlowResponse.Marshal(b, m, deterministic)
}
func (m *WorkFlowResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkFlowResponse.Merge(m, src)
}
func (m *WorkFlowResponse) XXX_Size() int {
	return xxx_messageInfo_WorkFlowResponse.Size(m)
}
func (m *WorkFlowResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkFlowResponse.DiscardUnknown(m)
}

var xxx_messageInfo_WorkFlowResponse proto.InternalMessageInfo

func (m *WorkFlowResponse) GetWorkflowID() string {
	if m != nil {
		return m.WorkflowID
	}
	return ""
}

func (m *WorkFlowResponse) GetWorkFlowRunID() string {
	if m != nil {
		return m.WorkFlowRunID
	}
	return ""
}

func init() {
	proto.RegisterType((*DeployRequest)(nil), "climonpb.DeployRequest")
	proto.RegisterType((*DeployResponse)(nil), "climonpb.DeployResponse")
	proto.RegisterType((*WorkFlowResponse)(nil), "climonpb.WorkFlowResponse")
}

func init() { proto.RegisterFile("climon.proto", fileDescriptor_89182b4db867c195) }

var fileDescriptor_89182b4db867c195 = []byte{
	// 285 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0x4b, 0x4b, 0xc4, 0x30,
	0x14, 0x85, 0x19, 0x1f, 0x9d, 0xf6, 0xd6, 0x11, 0xc9, 0x62, 0x08, 0x83, 0xc8, 0x50, 0x5c, 0xb8,
	0xaa, 0xa0, 0xff, 0x40, 0x8b, 0xd0, 0x8d, 0x48, 0xc4, 0x07, 0xee, 0x3a, 0xe5, 0x8e, 0x16, 0x33,
	0x49, 0x4c, 0x5a, 0x06, 0xc1, 0xff, 0xed, 0x56, 0x92, 0x26, 0x75, 0x1c, 0x70, 0x95, 0x9c, 0xf3,
	0x9d, 0x0b, 0xf7, 0x01, 0x07, 0x35, 0x6f, 0x56, 0x52, 0xe4, 0x4a, 0xcb, 0x56, 0x92, 0xb8, 0x57,
	0x6a, 0x91, 0x7d, 0x8f, 0x60, 0x52, 0xa0, 0xe2, 0xf2, 0x93, 0xe1, 0x47, 0x87, 0xa6, 0x25, 0x14,
	0xc6, 0x8f, 0xa8, 0x4d, 0x23, 0x05, 0x1d, 0xcd, 0x47, 0x67, 0x09, 0x0b, 0xd2, 0x12, 0x86, 0x4a,
	0x3e, 0x68, 0x4e, 0x77, 0x7a, 0xe2, 0x25, 0x99, 0x41, 0x6c, 0xbf, 0xb7, 0xd5, 0x0a, 0xe9, 0xae,
	0x43, 0x83, 0x26, 0xc7, 0x90, 0xd8, 0xd7, 0xa8, 0xaa, 0x46, 0xba, 0xe7, 0xe0, 0xaf, 0x61, 0xe9,
	0xf5, 0x5b, 0xa5, 0x5b, 0x57, 0xba, 0xdf, 0xd3, 0xc1, 0x20, 0x73, 0x48, 0x19, 0x72, 0xac, 0x0c,
	0x3a, 0x1e, 0x39, 0xbe, 0x69, 0xf5, 0x89, 0x25, 0x6a, 0x14, 0x35, 0x96, 0x05, 0x1d, 0x87, 0xc4,
	0x60, 0x91, 0x29, 0x44, 0x77, 0xbc, 0x7b, 0x6d, 0x04, 0x8d, 0x1d, 0xf4, 0x2a, 0xfb, 0x82, 0xc3,
	0x30, 0xb8, 0x51, 0x52, 0x18, 0xb4, 0xc9, 0xfb, 0xb6, 0x6a, 0x3b, 0xe3, 0x07, 0xf7, 0x8a, 0x30,
	0x98, 0x3e, 0x49, 0xfd, 0x7e, 0xc3, 0xe5, 0x3a, 0x64, 0x7d, 0xce, 0xae, 0x21, 0xbd, 0x98, 0xe5,
	0x61, 0x9d, 0xf9, 0x76, 0x8e, 0xfd, 0x53, 0x99, 0x3d, 0xc3, 0xd1, 0x36, 0x21, 0x27, 0x00, 0xd6,
	0x5b, 0x72, 0xb9, 0x2e, 0x0b, 0xdf, 0xc3, 0x86, 0x43, 0x4e, 0x61, 0x32, 0xd4, 0x74, 0xa2, 0x2c,
	0xfc, 0x15, 0xfe, 0x9a, 0x57, 0xe9, 0x4b, 0x72, 0x1e, 0xfa, 0x59, 0x44, 0xee, 0xde, 0x97, 0x3f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x2f, 0x58, 0xa4, 0xcd, 0xff, 0x01, 0x00, 0x00,
}
