// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: learnings/v1/learnings.proto

package learningsv1

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

// LearningsServiceClient is the client API for LearningsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LearningsServiceClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
	AddCourse(ctx context.Context, in *AddCourseRequest, opts ...grpc.CallOption) (*AddCourseResponse, error)
	ListCourses(ctx context.Context, in *ListCoursesRequest, opts ...grpc.CallOption) (LearningsService_ListCoursesClient, error)
}

type learningsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLearningsServiceClient(cc grpc.ClientConnInterface) LearningsServiceClient {
	return &learningsServiceClient{cc}
}

func (c *learningsServiceClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, "/learnings.v1.LearningsService/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *learningsServiceClient) AddCourse(ctx context.Context, in *AddCourseRequest, opts ...grpc.CallOption) (*AddCourseResponse, error) {
	out := new(AddCourseResponse)
	err := c.cc.Invoke(ctx, "/learnings.v1.LearningsService/AddCourse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *learningsServiceClient) ListCourses(ctx context.Context, in *ListCoursesRequest, opts ...grpc.CallOption) (LearningsService_ListCoursesClient, error) {
	stream, err := c.cc.NewStream(ctx, &LearningsService_ServiceDesc.Streams[0], "/learnings.v1.LearningsService/ListCourses", opts...)
	if err != nil {
		return nil, err
	}
	x := &learningsServiceListCoursesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type LearningsService_ListCoursesClient interface {
	Recv() (*ListCoursesResponse, error)
	grpc.ClientStream
}

type learningsServiceListCoursesClient struct {
	grpc.ClientStream
}

func (x *learningsServiceListCoursesClient) Recv() (*ListCoursesResponse, error) {
	m := new(ListCoursesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// LearningsServiceServer is the server API for LearningsService service.
// All implementations should embed UnimplementedLearningsServiceServer
// for forward compatibility
type LearningsServiceServer interface {
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	AddCourse(context.Context, *AddCourseRequest) (*AddCourseResponse, error)
	ListCourses(*ListCoursesRequest, LearningsService_ListCoursesServer) error
}

// UnimplementedLearningsServiceServer should be embedded to have forward compatible implementations.
type UnimplementedLearningsServiceServer struct {
}

func (UnimplementedLearningsServiceServer) Ping(context.Context, *PingRequest) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedLearningsServiceServer) AddCourse(context.Context, *AddCourseRequest) (*AddCourseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCourse not implemented")
}
func (UnimplementedLearningsServiceServer) ListCourses(*ListCoursesRequest, LearningsService_ListCoursesServer) error {
	return status.Errorf(codes.Unimplemented, "method ListCourses not implemented")
}

// UnsafeLearningsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LearningsServiceServer will
// result in compilation errors.
type UnsafeLearningsServiceServer interface {
	mustEmbedUnimplementedLearningsServiceServer()
}

func RegisterLearningsServiceServer(s grpc.ServiceRegistrar, srv LearningsServiceServer) {
	s.RegisterService(&LearningsService_ServiceDesc, srv)
}

func _LearningsService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LearningsServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/learnings.v1.LearningsService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LearningsServiceServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LearningsService_AddCourse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCourseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LearningsServiceServer).AddCourse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/learnings.v1.LearningsService/AddCourse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LearningsServiceServer).AddCourse(ctx, req.(*AddCourseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LearningsService_ListCourses_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListCoursesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LearningsServiceServer).ListCourses(m, &learningsServiceListCoursesServer{stream})
}

type LearningsService_ListCoursesServer interface {
	Send(*ListCoursesResponse) error
	grpc.ServerStream
}

type learningsServiceListCoursesServer struct {
	grpc.ServerStream
}

func (x *learningsServiceListCoursesServer) Send(m *ListCoursesResponse) error {
	return x.ServerStream.SendMsg(m)
}

// LearningsService_ServiceDesc is the grpc.ServiceDesc for LearningsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LearningsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "learnings.v1.LearningsService",
	HandlerType: (*LearningsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _LearningsService_Ping_Handler,
		},
		{
			MethodName: "AddCourse",
			Handler:    _LearningsService_AddCourse_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListCourses",
			Handler:       _LearningsService_ListCourses_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "learnings/v1/learnings.proto",
}
