// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: agent.proto

package agentpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Agent_SubmitJob_FullMethodName                = "/agentpb.Agent/SubmitJob"
	Agent_StoreCredential_FullMethodName          = "/agentpb.Agent/StoreCredential"
	Agent_Sync_FullMethodName                     = "/agentpb.Agent/Sync"
	Agent_ClimonAppInstall_FullMethodName         = "/agentpb.Agent/ClimonAppInstall"
	Agent_ClimonAppDelete_FullMethodName          = "/agentpb.Agent/ClimonAppDelete"
	Agent_DeployerAppInstall_FullMethodName       = "/agentpb.Agent/DeployerAppInstall"
	Agent_DeployerAppDelete_FullMethodName        = "/agentpb.Agent/DeployerAppDelete"
	Agent_ClusterAdd_FullMethodName               = "/agentpb.Agent/ClusterAdd"
	Agent_ClusterDelete_FullMethodName            = "/agentpb.Agent/ClusterDelete"
	Agent_RepositoryAdd_FullMethodName            = "/agentpb.Agent/RepositoryAdd"
	Agent_RepositoryDelete_FullMethodName         = "/agentpb.Agent/RepositoryDelete"
	Agent_ProjectAdd_FullMethodName               = "/agentpb.Agent/ProjectAdd"
	Agent_ProjectDelete_FullMethodName            = "/agentpb.Agent/ProjectDelete"
	Agent_SyncApp_FullMethodName                  = "/agentpb.Agent/SyncApp"
	Agent_GetAllApplicationConfigs_FullMethodName = "/agentpb.Agent/GetAllApplicationConfigs"
	Agent_GetApplicationConfig_FullMethodName     = "/agentpb.Agent/GetApplicationConfig"
	Agent_GetApplicationValues_FullMethodName     = "/agentpb.Agent/GetApplicationValues"
)

// AgentClient is the client API for Agent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AgentClient interface {
	SubmitJob(ctx context.Context, in *JobRequest, opts ...grpc.CallOption) (*JobResponse, error)
	StoreCredential(ctx context.Context, in *StoreCredentialRequest, opts ...grpc.CallOption) (*StoreCredentialResponse, error)
	Sync(ctx context.Context, in *SyncRequest, opts ...grpc.CallOption) (*SyncResponse, error)
	ClimonAppInstall(ctx context.Context, in *ClimonInstallRequest, opts ...grpc.CallOption) (*JobResponse, error)
	ClimonAppDelete(ctx context.Context, in *ClimonDeleteRequest, opts ...grpc.CallOption) (*JobResponse, error)
	DeployerAppInstall(ctx context.Context, in *ApplicationInstallRequest, opts ...grpc.CallOption) (*JobResponse, error)
	DeployerAppDelete(ctx context.Context, in *ApplicationDeleteRequest, opts ...grpc.CallOption) (*JobResponse, error)
	ClusterAdd(ctx context.Context, in *ClusterRequest, opts ...grpc.CallOption) (*JobResponse, error)
	ClusterDelete(ctx context.Context, in *ClusterRequest, opts ...grpc.CallOption) (*JobResponse, error)
	RepositoryAdd(ctx context.Context, in *RepositoryAddRequest, opts ...grpc.CallOption) (*JobResponse, error)
	RepositoryDelete(ctx context.Context, in *RepositoryDeleteRequest, opts ...grpc.CallOption) (*JobResponse, error)
	ProjectAdd(ctx context.Context, in *ProjectAddRequest, opts ...grpc.CallOption) (*JobResponse, error)
	ProjectDelete(ctx context.Context, in *ProjectDeleteRequest, opts ...grpc.CallOption) (*JobResponse, error)
	SyncApp(ctx context.Context, in *SyncAppRequest, opts ...grpc.CallOption) (*SyncAppResponse, error)
	GetAllApplicationConfigs(ctx context.Context, in *GetAllApplicationConfigsRequest, opts ...grpc.CallOption) (*GetAllApplicationConfigsResponse, error)
	GetApplicationConfig(ctx context.Context, in *GetApplicationConfigRequest, opts ...grpc.CallOption) (*GetApplicationConfigResponse, error)
	GetApplicationValues(ctx context.Context, in *GetApplicationValuesRequest, opts ...grpc.CallOption) (*GetApplicationValuesResponse, error)
}

type agentClient struct {
	cc grpc.ClientConnInterface
}

func NewAgentClient(cc grpc.ClientConnInterface) AgentClient {
	return &agentClient{cc}
}

func (c *agentClient) SubmitJob(ctx context.Context, in *JobRequest, opts ...grpc.CallOption) (*JobResponse, error) {
	out := new(JobResponse)
	err := c.cc.Invoke(ctx, Agent_SubmitJob_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) StoreCredential(ctx context.Context, in *StoreCredentialRequest, opts ...grpc.CallOption) (*StoreCredentialResponse, error) {
	out := new(StoreCredentialResponse)
	err := c.cc.Invoke(ctx, Agent_StoreCredential_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) Sync(ctx context.Context, in *SyncRequest, opts ...grpc.CallOption) (*SyncResponse, error) {
	out := new(SyncResponse)
	err := c.cc.Invoke(ctx, Agent_Sync_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) ClimonAppInstall(ctx context.Context, in *ClimonInstallRequest, opts ...grpc.CallOption) (*JobResponse, error) {
	out := new(JobResponse)
	err := c.cc.Invoke(ctx, Agent_ClimonAppInstall_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) ClimonAppDelete(ctx context.Context, in *ClimonDeleteRequest, opts ...grpc.CallOption) (*JobResponse, error) {
	out := new(JobResponse)
	err := c.cc.Invoke(ctx, Agent_ClimonAppDelete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) DeployerAppInstall(ctx context.Context, in *ApplicationInstallRequest, opts ...grpc.CallOption) (*JobResponse, error) {
	out := new(JobResponse)
	err := c.cc.Invoke(ctx, Agent_DeployerAppInstall_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) DeployerAppDelete(ctx context.Context, in *ApplicationDeleteRequest, opts ...grpc.CallOption) (*JobResponse, error) {
	out := new(JobResponse)
	err := c.cc.Invoke(ctx, Agent_DeployerAppDelete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) ClusterAdd(ctx context.Context, in *ClusterRequest, opts ...grpc.CallOption) (*JobResponse, error) {
	out := new(JobResponse)
	err := c.cc.Invoke(ctx, Agent_ClusterAdd_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) ClusterDelete(ctx context.Context, in *ClusterRequest, opts ...grpc.CallOption) (*JobResponse, error) {
	out := new(JobResponse)
	err := c.cc.Invoke(ctx, Agent_ClusterDelete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) RepositoryAdd(ctx context.Context, in *RepositoryAddRequest, opts ...grpc.CallOption) (*JobResponse, error) {
	out := new(JobResponse)
	err := c.cc.Invoke(ctx, Agent_RepositoryAdd_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) RepositoryDelete(ctx context.Context, in *RepositoryDeleteRequest, opts ...grpc.CallOption) (*JobResponse, error) {
	out := new(JobResponse)
	err := c.cc.Invoke(ctx, Agent_RepositoryDelete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) ProjectAdd(ctx context.Context, in *ProjectAddRequest, opts ...grpc.CallOption) (*JobResponse, error) {
	out := new(JobResponse)
	err := c.cc.Invoke(ctx, Agent_ProjectAdd_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) ProjectDelete(ctx context.Context, in *ProjectDeleteRequest, opts ...grpc.CallOption) (*JobResponse, error) {
	out := new(JobResponse)
	err := c.cc.Invoke(ctx, Agent_ProjectDelete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) SyncApp(ctx context.Context, in *SyncAppRequest, opts ...grpc.CallOption) (*SyncAppResponse, error) {
	out := new(SyncAppResponse)
	err := c.cc.Invoke(ctx, Agent_SyncApp_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) GetAllApplicationConfigs(ctx context.Context, in *GetAllApplicationConfigsRequest, opts ...grpc.CallOption) (*GetAllApplicationConfigsResponse, error) {
	out := new(GetAllApplicationConfigsResponse)
	err := c.cc.Invoke(ctx, Agent_GetAllApplicationConfigs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) GetApplicationConfig(ctx context.Context, in *GetApplicationConfigRequest, opts ...grpc.CallOption) (*GetApplicationConfigResponse, error) {
	out := new(GetApplicationConfigResponse)
	err := c.cc.Invoke(ctx, Agent_GetApplicationConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) GetApplicationValues(ctx context.Context, in *GetApplicationValuesRequest, opts ...grpc.CallOption) (*GetApplicationValuesResponse, error) {
	out := new(GetApplicationValuesResponse)
	err := c.cc.Invoke(ctx, Agent_GetApplicationValues_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AgentServer is the server API for Agent service.
// All implementations must embed UnimplementedAgentServer
// for forward compatibility
type AgentServer interface {
	SubmitJob(context.Context, *JobRequest) (*JobResponse, error)
	StoreCredential(context.Context, *StoreCredentialRequest) (*StoreCredentialResponse, error)
	Sync(context.Context, *SyncRequest) (*SyncResponse, error)
	ClimonAppInstall(context.Context, *ClimonInstallRequest) (*JobResponse, error)
	ClimonAppDelete(context.Context, *ClimonDeleteRequest) (*JobResponse, error)
	DeployerAppInstall(context.Context, *ApplicationInstallRequest) (*JobResponse, error)
	DeployerAppDelete(context.Context, *ApplicationDeleteRequest) (*JobResponse, error)
	ClusterAdd(context.Context, *ClusterRequest) (*JobResponse, error)
	ClusterDelete(context.Context, *ClusterRequest) (*JobResponse, error)
	RepositoryAdd(context.Context, *RepositoryAddRequest) (*JobResponse, error)
	RepositoryDelete(context.Context, *RepositoryDeleteRequest) (*JobResponse, error)
	ProjectAdd(context.Context, *ProjectAddRequest) (*JobResponse, error)
	ProjectDelete(context.Context, *ProjectDeleteRequest) (*JobResponse, error)
	SyncApp(context.Context, *SyncAppRequest) (*SyncAppResponse, error)
	GetAllApplicationConfigs(context.Context, *GetAllApplicationConfigsRequest) (*GetAllApplicationConfigsResponse, error)
	GetApplicationConfig(context.Context, *GetApplicationConfigRequest) (*GetApplicationConfigResponse, error)
	GetApplicationValues(context.Context, *GetApplicationValuesRequest) (*GetApplicationValuesResponse, error)
	mustEmbedUnimplementedAgentServer()
}

// UnimplementedAgentServer must be embedded to have forward compatible implementations.
type UnimplementedAgentServer struct {
}

func (UnimplementedAgentServer) SubmitJob(context.Context, *JobRequest) (*JobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitJob not implemented")
}
func (UnimplementedAgentServer) StoreCredential(context.Context, *StoreCredentialRequest) (*StoreCredentialResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoreCredential not implemented")
}
func (UnimplementedAgentServer) Sync(context.Context, *SyncRequest) (*SyncResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sync not implemented")
}
func (UnimplementedAgentServer) ClimonAppInstall(context.Context, *ClimonInstallRequest) (*JobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClimonAppInstall not implemented")
}
func (UnimplementedAgentServer) ClimonAppDelete(context.Context, *ClimonDeleteRequest) (*JobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClimonAppDelete not implemented")
}
func (UnimplementedAgentServer) DeployerAppInstall(context.Context, *ApplicationInstallRequest) (*JobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeployerAppInstall not implemented")
}
func (UnimplementedAgentServer) DeployerAppDelete(context.Context, *ApplicationDeleteRequest) (*JobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeployerAppDelete not implemented")
}
func (UnimplementedAgentServer) ClusterAdd(context.Context, *ClusterRequest) (*JobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClusterAdd not implemented")
}
func (UnimplementedAgentServer) ClusterDelete(context.Context, *ClusterRequest) (*JobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClusterDelete not implemented")
}
func (UnimplementedAgentServer) RepositoryAdd(context.Context, *RepositoryAddRequest) (*JobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RepositoryAdd not implemented")
}
func (UnimplementedAgentServer) RepositoryDelete(context.Context, *RepositoryDeleteRequest) (*JobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RepositoryDelete not implemented")
}
func (UnimplementedAgentServer) ProjectAdd(context.Context, *ProjectAddRequest) (*JobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProjectAdd not implemented")
}
func (UnimplementedAgentServer) ProjectDelete(context.Context, *ProjectDeleteRequest) (*JobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProjectDelete not implemented")
}
func (UnimplementedAgentServer) SyncApp(context.Context, *SyncAppRequest) (*SyncAppResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SyncApp not implemented")
}
func (UnimplementedAgentServer) GetAllApplicationConfigs(context.Context, *GetAllApplicationConfigsRequest) (*GetAllApplicationConfigsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllApplicationConfigs not implemented")
}
func (UnimplementedAgentServer) GetApplicationConfig(context.Context, *GetApplicationConfigRequest) (*GetApplicationConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetApplicationConfig not implemented")
}
func (UnimplementedAgentServer) GetApplicationValues(context.Context, *GetApplicationValuesRequest) (*GetApplicationValuesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetApplicationValues not implemented")
}
func (UnimplementedAgentServer) mustEmbedUnimplementedAgentServer() {}

// UnsafeAgentServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AgentServer will
// result in compilation errors.
type UnsafeAgentServer interface {
	mustEmbedUnimplementedAgentServer()
}

func RegisterAgentServer(s grpc.ServiceRegistrar, srv AgentServer) {
	s.RegisterService(&Agent_ServiceDesc, srv)
}

func _Agent_SubmitJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).SubmitJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Agent_SubmitJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).SubmitJob(ctx, req.(*JobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_StoreCredential_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreCredentialRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).StoreCredential(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Agent_StoreCredential_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).StoreCredential(ctx, req.(*StoreCredentialRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_Sync_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SyncRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).Sync(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Agent_Sync_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).Sync(ctx, req.(*SyncRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_ClimonAppInstall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClimonInstallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).ClimonAppInstall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Agent_ClimonAppInstall_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).ClimonAppInstall(ctx, req.(*ClimonInstallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_ClimonAppDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClimonDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).ClimonAppDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Agent_ClimonAppDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).ClimonAppDelete(ctx, req.(*ClimonDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_DeployerAppInstall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplicationInstallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).DeployerAppInstall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Agent_DeployerAppInstall_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).DeployerAppInstall(ctx, req.(*ApplicationInstallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_DeployerAppDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplicationDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).DeployerAppDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Agent_DeployerAppDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).DeployerAppDelete(ctx, req.(*ApplicationDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_ClusterAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).ClusterAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Agent_ClusterAdd_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).ClusterAdd(ctx, req.(*ClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_ClusterDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).ClusterDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Agent_ClusterDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).ClusterDelete(ctx, req.(*ClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_RepositoryAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RepositoryAddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).RepositoryAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Agent_RepositoryAdd_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).RepositoryAdd(ctx, req.(*RepositoryAddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_RepositoryDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RepositoryDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).RepositoryDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Agent_RepositoryDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).RepositoryDelete(ctx, req.(*RepositoryDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_ProjectAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProjectAddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).ProjectAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Agent_ProjectAdd_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).ProjectAdd(ctx, req.(*ProjectAddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_ProjectDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProjectDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).ProjectDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Agent_ProjectDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).ProjectDelete(ctx, req.(*ProjectDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_SyncApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SyncAppRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).SyncApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Agent_SyncApp_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).SyncApp(ctx, req.(*SyncAppRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_GetAllApplicationConfigs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllApplicationConfigsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).GetAllApplicationConfigs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Agent_GetAllApplicationConfigs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).GetAllApplicationConfigs(ctx, req.(*GetAllApplicationConfigsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_GetApplicationConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetApplicationConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).GetApplicationConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Agent_GetApplicationConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).GetApplicationConfig(ctx, req.(*GetApplicationConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_GetApplicationValues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetApplicationValuesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).GetApplicationValues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Agent_GetApplicationValues_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).GetApplicationValues(ctx, req.(*GetApplicationValuesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Agent_ServiceDesc is the grpc.ServiceDesc for Agent service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Agent_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "agentpb.Agent",
	HandlerType: (*AgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubmitJob",
			Handler:    _Agent_SubmitJob_Handler,
		},
		{
			MethodName: "StoreCredential",
			Handler:    _Agent_StoreCredential_Handler,
		},
		{
			MethodName: "Sync",
			Handler:    _Agent_Sync_Handler,
		},
		{
			MethodName: "ClimonAppInstall",
			Handler:    _Agent_ClimonAppInstall_Handler,
		},
		{
			MethodName: "ClimonAppDelete",
			Handler:    _Agent_ClimonAppDelete_Handler,
		},
		{
			MethodName: "DeployerAppInstall",
			Handler:    _Agent_DeployerAppInstall_Handler,
		},
		{
			MethodName: "DeployerAppDelete",
			Handler:    _Agent_DeployerAppDelete_Handler,
		},
		{
			MethodName: "ClusterAdd",
			Handler:    _Agent_ClusterAdd_Handler,
		},
		{
			MethodName: "ClusterDelete",
			Handler:    _Agent_ClusterDelete_Handler,
		},
		{
			MethodName: "RepositoryAdd",
			Handler:    _Agent_RepositoryAdd_Handler,
		},
		{
			MethodName: "RepositoryDelete",
			Handler:    _Agent_RepositoryDelete_Handler,
		},
		{
			MethodName: "ProjectAdd",
			Handler:    _Agent_ProjectAdd_Handler,
		},
		{
			MethodName: "ProjectDelete",
			Handler:    _Agent_ProjectDelete_Handler,
		},
		{
			MethodName: "SyncApp",
			Handler:    _Agent_SyncApp_Handler,
		},
		{
			MethodName: "GetAllApplicationConfigs",
			Handler:    _Agent_GetAllApplicationConfigs_Handler,
		},
		{
			MethodName: "GetApplicationConfig",
			Handler:    _Agent_GetApplicationConfig_Handler,
		},
		{
			MethodName: "GetApplicationValues",
			Handler:    _Agent_GetApplicationValues_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "agent.proto",
}
