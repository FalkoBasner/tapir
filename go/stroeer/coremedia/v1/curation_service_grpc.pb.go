// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package coremedia

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

// CurationServiceClient is the client API for CurationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CurationServiceClient interface {
	// Responds with curation data for the provided request parameters.
	GetCuratedStages(ctx context.Context, in *GetCuratedStagesRequest, opts ...grpc.CallOption) (*GetCuratedStagesResponse, error)
}

type curationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCurationServiceClient(cc grpc.ClientConnInterface) CurationServiceClient {
	return &curationServiceClient{cc}
}

func (c *curationServiceClient) GetCuratedStages(ctx context.Context, in *GetCuratedStagesRequest, opts ...grpc.CallOption) (*GetCuratedStagesResponse, error) {
	out := new(GetCuratedStagesResponse)
	err := c.cc.Invoke(ctx, "/stroeer.coremedia.v1.CurationService/GetCuratedStages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CurationServiceServer is the server API for CurationService service.
// All implementations must embed UnimplementedCurationServiceServer
// for forward compatibility
type CurationServiceServer interface {
	// Responds with curation data for the provided request parameters.
	GetCuratedStages(context.Context, *GetCuratedStagesRequest) (*GetCuratedStagesResponse, error)
	mustEmbedUnimplementedCurationServiceServer()
}

// UnimplementedCurationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCurationServiceServer struct {
}

func (UnimplementedCurationServiceServer) GetCuratedStages(context.Context, *GetCuratedStagesRequest) (*GetCuratedStagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCuratedStages not implemented")
}
func (UnimplementedCurationServiceServer) mustEmbedUnimplementedCurationServiceServer() {}

// UnsafeCurationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CurationServiceServer will
// result in compilation errors.
type UnsafeCurationServiceServer interface {
	mustEmbedUnimplementedCurationServiceServer()
}

func RegisterCurationServiceServer(s grpc.ServiceRegistrar, srv CurationServiceServer) {
	s.RegisterService(&CurationService_ServiceDesc, srv)
}

func _CurationService_GetCuratedStages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCuratedStagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CurationServiceServer).GetCuratedStages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stroeer.coremedia.v1.CurationService/GetCuratedStages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CurationServiceServer).GetCuratedStages(ctx, req.(*GetCuratedStagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CurationService_ServiceDesc is the grpc.ServiceDesc for CurationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CurationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "stroeer.coremedia.v1.CurationService",
	HandlerType: (*CurationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCuratedStages",
			Handler:    _CurationService_GetCuratedStages_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stroeer/coremedia/v1/curation_service.proto",
}
