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
	Agent_Ping_FullMethodName                   = "/agentpb.Agent/Ping"
	Agent_StoreCredential_FullMethodName        = "/agentpb.Agent/StoreCredential"
	Agent_SyncApp_FullMethodName                = "/agentpb.Agent/SyncApp"
	Agent_GetClusterApps_FullMethodName         = "/agentpb.Agent/GetClusterApps"
	Agent_GetClusterAppLaunches_FullMethodName  = "/agentpb.Agent/GetClusterAppLaunches"
	Agent_ConfigureAppSSO_FullMethodName        = "/agentpb.Agent/ConfigureAppSSO"
	Agent_GetClusterAppConfig_FullMethodName    = "/agentpb.Agent/GetClusterAppConfig"
	Agent_GetClusterAppValues_FullMethodName    = "/agentpb.Agent/GetClusterAppValues"
	Agent_GetClusterGlobalValues_FullMethodName = "/agentpb.Agent/GetClusterGlobalValues"
	Agent_InstallApp_FullMethodName             = "/agentpb.Agent/InstallApp"
	Agent_UnInstallApp_FullMethodName           = "/agentpb.Agent/UnInstallApp"
	Agent_UpgradeApp_FullMethodName             = "/agentpb.Agent/UpgradeApp"
	Agent_UpdateAppValues_FullMethodName        = "/agentpb.Agent/UpdateAppValues"
)

// AgentClient is the client API for Agent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AgentClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
	StoreCredential(ctx context.Context, in *StoreCredentialRequest, opts ...grpc.CallOption) (*StoreCredentialResponse, error)
	SyncApp(ctx context.Context, in *SyncAppRequest, opts ...grpc.CallOption) (*SyncAppResponse, error)
	GetClusterApps(ctx context.Context, in *GetClusterAppsRequest, opts ...grpc.CallOption) (*GetClusterAppsResponse, error)
	GetClusterAppLaunches(ctx context.Context, in *GetClusterAppLaunchesRequest, opts ...grpc.CallOption) (*GetClusterAppLaunchesResponse, error)
	ConfigureAppSSO(ctx context.Context, in *ConfigureAppSSORequest, opts ...grpc.CallOption) (*ConfigureAppSSOResponse, error)
	GetClusterAppConfig(ctx context.Context, in *GetClusterAppConfigRequest, opts ...grpc.CallOption) (*GetClusterAppConfigResponse, error)
	GetClusterAppValues(ctx context.Context, in *GetClusterAppValuesRequest, opts ...grpc.CallOption) (*GetClusterAppValuesResponse, error)
	GetClusterGlobalValues(ctx context.Context, in *GetClusterGlobalValuesRequest, opts ...grpc.CallOption) (*GetClusterGlobalValuesResponse, error)
	InstallApp(ctx context.Context, in *InstallAppRequest, opts ...grpc.CallOption) (*InstallAppResponse, error)
	UnInstallApp(ctx context.Context, in *UnInstallAppRequest, opts ...grpc.CallOption) (*UnInstallAppResponse, error)
	UpgradeApp(ctx context.Context, in *UpgradeAppRequest, opts ...grpc.CallOption) (*UpgradeAppResponse, error)
	UpdateAppValues(ctx context.Context, in *UpdateAppValuesRequest, opts ...grpc.CallOption) (*UpdateAppValuesResponse, error)
}

type agentClient struct {
	cc grpc.ClientConnInterface
}

func NewAgentClient(cc grpc.ClientConnInterface) AgentClient {
	return &agentClient{cc}
}

func (c *agentClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, Agent_Ping_FullMethodName, in, out, opts...)
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

func (c *agentClient) SyncApp(ctx context.Context, in *SyncAppRequest, opts ...grpc.CallOption) (*SyncAppResponse, error) {
	out := new(SyncAppResponse)
	err := c.cc.Invoke(ctx, Agent_SyncApp_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) GetClusterApps(ctx context.Context, in *GetClusterAppsRequest, opts ...grpc.CallOption) (*GetClusterAppsResponse, error) {
	out := new(GetClusterAppsResponse)
	err := c.cc.Invoke(ctx, Agent_GetClusterApps_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) GetClusterAppLaunches(ctx context.Context, in *GetClusterAppLaunchesRequest, opts ...grpc.CallOption) (*GetClusterAppLaunchesResponse, error) {
	out := new(GetClusterAppLaunchesResponse)
	err := c.cc.Invoke(ctx, Agent_GetClusterAppLaunches_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) ConfigureAppSSO(ctx context.Context, in *ConfigureAppSSORequest, opts ...grpc.CallOption) (*ConfigureAppSSOResponse, error) {
	out := new(ConfigureAppSSOResponse)
	err := c.cc.Invoke(ctx, Agent_ConfigureAppSSO_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) GetClusterAppConfig(ctx context.Context, in *GetClusterAppConfigRequest, opts ...grpc.CallOption) (*GetClusterAppConfigResponse, error) {
	out := new(GetClusterAppConfigResponse)
	err := c.cc.Invoke(ctx, Agent_GetClusterAppConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) GetClusterAppValues(ctx context.Context, in *GetClusterAppValuesRequest, opts ...grpc.CallOption) (*GetClusterAppValuesResponse, error) {
	out := new(GetClusterAppValuesResponse)
	err := c.cc.Invoke(ctx, Agent_GetClusterAppValues_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) GetClusterGlobalValues(ctx context.Context, in *GetClusterGlobalValuesRequest, opts ...grpc.CallOption) (*GetClusterGlobalValuesResponse, error) {
	out := new(GetClusterGlobalValuesResponse)
	err := c.cc.Invoke(ctx, Agent_GetClusterGlobalValues_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) InstallApp(ctx context.Context, in *InstallAppRequest, opts ...grpc.CallOption) (*InstallAppResponse, error) {
	out := new(InstallAppResponse)
	err := c.cc.Invoke(ctx, Agent_InstallApp_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) UnInstallApp(ctx context.Context, in *UnInstallAppRequest, opts ...grpc.CallOption) (*UnInstallAppResponse, error) {
	out := new(UnInstallAppResponse)
	err := c.cc.Invoke(ctx, Agent_UnInstallApp_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) UpgradeApp(ctx context.Context, in *UpgradeAppRequest, opts ...grpc.CallOption) (*UpgradeAppResponse, error) {
	out := new(UpgradeAppResponse)
	err := c.cc.Invoke(ctx, Agent_UpgradeApp_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) UpdateAppValues(ctx context.Context, in *UpdateAppValuesRequest, opts ...grpc.CallOption) (*UpdateAppValuesResponse, error) {
	out := new(UpdateAppValuesResponse)
	err := c.cc.Invoke(ctx, Agent_UpdateAppValues_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AgentServer is the server API for Agent service.
// All implementations must embed UnimplementedAgentServer
// for forward compatibility
type AgentServer interface {
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	StoreCredential(context.Context, *StoreCredentialRequest) (*StoreCredentialResponse, error)
	SyncApp(context.Context, *SyncAppRequest) (*SyncAppResponse, error)
	GetClusterApps(context.Context, *GetClusterAppsRequest) (*GetClusterAppsResponse, error)
	GetClusterAppLaunches(context.Context, *GetClusterAppLaunchesRequest) (*GetClusterAppLaunchesResponse, error)
	ConfigureAppSSO(context.Context, *ConfigureAppSSORequest) (*ConfigureAppSSOResponse, error)
	GetClusterAppConfig(context.Context, *GetClusterAppConfigRequest) (*GetClusterAppConfigResponse, error)
	GetClusterAppValues(context.Context, *GetClusterAppValuesRequest) (*GetClusterAppValuesResponse, error)
	GetClusterGlobalValues(context.Context, *GetClusterGlobalValuesRequest) (*GetClusterGlobalValuesResponse, error)
	InstallApp(context.Context, *InstallAppRequest) (*InstallAppResponse, error)
	UnInstallApp(context.Context, *UnInstallAppRequest) (*UnInstallAppResponse, error)
	UpgradeApp(context.Context, *UpgradeAppRequest) (*UpgradeAppResponse, error)
	UpdateAppValues(context.Context, *UpdateAppValuesRequest) (*UpdateAppValuesResponse, error)
	mustEmbedUnimplementedAgentServer()
}

// UnimplementedAgentServer must be embedded to have forward compatible implementations.
type UnimplementedAgentServer struct {
}

func (UnimplementedAgentServer) Ping(context.Context, *PingRequest) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedAgentServer) StoreCredential(context.Context, *StoreCredentialRequest) (*StoreCredentialResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoreCredential not implemented")
}
func (UnimplementedAgentServer) SyncApp(context.Context, *SyncAppRequest) (*SyncAppResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SyncApp not implemented")
}
func (UnimplementedAgentServer) GetClusterApps(context.Context, *GetClusterAppsRequest) (*GetClusterAppsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClusterApps not implemented")
}
func (UnimplementedAgentServer) GetClusterAppLaunches(context.Context, *GetClusterAppLaunchesRequest) (*GetClusterAppLaunchesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClusterAppLaunches not implemented")
}
func (UnimplementedAgentServer) ConfigureAppSSO(context.Context, *ConfigureAppSSORequest) (*ConfigureAppSSOResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfigureAppSSO not implemented")
}
func (UnimplementedAgentServer) GetClusterAppConfig(context.Context, *GetClusterAppConfigRequest) (*GetClusterAppConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClusterAppConfig not implemented")
}
func (UnimplementedAgentServer) GetClusterAppValues(context.Context, *GetClusterAppValuesRequest) (*GetClusterAppValuesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClusterAppValues not implemented")
}
func (UnimplementedAgentServer) GetClusterGlobalValues(context.Context, *GetClusterGlobalValuesRequest) (*GetClusterGlobalValuesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClusterGlobalValues not implemented")
}
func (UnimplementedAgentServer) InstallApp(context.Context, *InstallAppRequest) (*InstallAppResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InstallApp not implemented")
}
func (UnimplementedAgentServer) UnInstallApp(context.Context, *UnInstallAppRequest) (*UnInstallAppResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnInstallApp not implemented")
}
func (UnimplementedAgentServer) UpgradeApp(context.Context, *UpgradeAppRequest) (*UpgradeAppResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpgradeApp not implemented")
}
func (UnimplementedAgentServer) UpdateAppValues(context.Context, *UpdateAppValuesRequest) (*UpdateAppValuesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAppValues not implemented")
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

func _Agent_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Agent_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).Ping(ctx, req.(*PingRequest))
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

func _Agent_GetClusterApps_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClusterAppsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).GetClusterApps(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Agent_GetClusterApps_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).GetClusterApps(ctx, req.(*GetClusterAppsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_GetClusterAppLaunches_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClusterAppLaunchesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).GetClusterAppLaunches(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Agent_GetClusterAppLaunches_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).GetClusterAppLaunches(ctx, req.(*GetClusterAppLaunchesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_ConfigureAppSSO_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigureAppSSORequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).ConfigureAppSSO(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Agent_ConfigureAppSSO_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).ConfigureAppSSO(ctx, req.(*ConfigureAppSSORequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_GetClusterAppConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClusterAppConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).GetClusterAppConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Agent_GetClusterAppConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).GetClusterAppConfig(ctx, req.(*GetClusterAppConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_GetClusterAppValues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClusterAppValuesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).GetClusterAppValues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Agent_GetClusterAppValues_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).GetClusterAppValues(ctx, req.(*GetClusterAppValuesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_GetClusterGlobalValues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClusterGlobalValuesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).GetClusterGlobalValues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Agent_GetClusterGlobalValues_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).GetClusterGlobalValues(ctx, req.(*GetClusterGlobalValuesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_InstallApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InstallAppRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).InstallApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Agent_InstallApp_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).InstallApp(ctx, req.(*InstallAppRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_UnInstallApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnInstallAppRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).UnInstallApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Agent_UnInstallApp_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).UnInstallApp(ctx, req.(*UnInstallAppRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_UpgradeApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpgradeAppRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).UpgradeApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Agent_UpgradeApp_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).UpgradeApp(ctx, req.(*UpgradeAppRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_UpdateAppValues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAppValuesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).UpdateAppValues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Agent_UpdateAppValues_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).UpdateAppValues(ctx, req.(*UpdateAppValuesRequest))
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
			MethodName: "Ping",
			Handler:    _Agent_Ping_Handler,
		},
		{
			MethodName: "StoreCredential",
			Handler:    _Agent_StoreCredential_Handler,
		},
		{
			MethodName: "SyncApp",
			Handler:    _Agent_SyncApp_Handler,
		},
		{
			MethodName: "GetClusterApps",
			Handler:    _Agent_GetClusterApps_Handler,
		},
		{
			MethodName: "GetClusterAppLaunches",
			Handler:    _Agent_GetClusterAppLaunches_Handler,
		},
		{
			MethodName: "ConfigureAppSSO",
			Handler:    _Agent_ConfigureAppSSO_Handler,
		},
		{
			MethodName: "GetClusterAppConfig",
			Handler:    _Agent_GetClusterAppConfig_Handler,
		},
		{
			MethodName: "GetClusterAppValues",
			Handler:    _Agent_GetClusterAppValues_Handler,
		},
		{
			MethodName: "GetClusterGlobalValues",
			Handler:    _Agent_GetClusterGlobalValues_Handler,
		},
		{
			MethodName: "InstallApp",
			Handler:    _Agent_InstallApp_Handler,
		},
		{
			MethodName: "UnInstallApp",
			Handler:    _Agent_UnInstallApp_Handler,
		},
		{
			MethodName: "UpgradeApp",
			Handler:    _Agent_UpgradeApp_Handler,
		},
		{
			MethodName: "UpdateAppValues",
			Handler:    _Agent_UpdateAppValues_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "agent.proto",
}