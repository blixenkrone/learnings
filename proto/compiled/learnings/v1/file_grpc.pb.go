// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: learnings/v1/file.proto

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

// FileUploadServiceClient is the client API for FileUploadService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FileUploadServiceClient interface {
	UploadFile(ctx context.Context, opts ...grpc.CallOption) (FileUploadService_UploadFileClient, error)
}

type fileUploadServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFileUploadServiceClient(cc grpc.ClientConnInterface) FileUploadServiceClient {
	return &fileUploadServiceClient{cc}
}

func (c *fileUploadServiceClient) UploadFile(ctx context.Context, opts ...grpc.CallOption) (FileUploadService_UploadFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &FileUploadService_ServiceDesc.Streams[0], "/learnings.v1.FileUploadService/UploadFile", opts...)
	if err != nil {
		return nil, err
	}
	x := &fileUploadServiceUploadFileClient{stream}
	return x, nil
}

type FileUploadService_UploadFileClient interface {
	Send(*UploadFileRequest) error
	CloseAndRecv() (*UploadFileResponse, error)
	grpc.ClientStream
}

type fileUploadServiceUploadFileClient struct {
	grpc.ClientStream
}

func (x *fileUploadServiceUploadFileClient) Send(m *UploadFileRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *fileUploadServiceUploadFileClient) CloseAndRecv() (*UploadFileResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadFileResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FileUploadServiceServer is the server API for FileUploadService service.
// All implementations should embed UnimplementedFileUploadServiceServer
// for forward compatibility
type FileUploadServiceServer interface {
	UploadFile(FileUploadService_UploadFileServer) error
}

// UnimplementedFileUploadServiceServer should be embedded to have forward compatible implementations.
type UnimplementedFileUploadServiceServer struct {
}

func (UnimplementedFileUploadServiceServer) UploadFile(FileUploadService_UploadFileServer) error {
	return status.Errorf(codes.Unimplemented, "method UploadFile not implemented")
}

// UnsafeFileUploadServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FileUploadServiceServer will
// result in compilation errors.
type UnsafeFileUploadServiceServer interface {
	mustEmbedUnimplementedFileUploadServiceServer()
}

func RegisterFileUploadServiceServer(s grpc.ServiceRegistrar, srv FileUploadServiceServer) {
	s.RegisterService(&FileUploadService_ServiceDesc, srv)
}

func _FileUploadService_UploadFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FileUploadServiceServer).UploadFile(&fileUploadServiceUploadFileServer{stream})
}

type FileUploadService_UploadFileServer interface {
	SendAndClose(*UploadFileResponse) error
	Recv() (*UploadFileRequest, error)
	grpc.ServerStream
}

type fileUploadServiceUploadFileServer struct {
	grpc.ServerStream
}

func (x *fileUploadServiceUploadFileServer) SendAndClose(m *UploadFileResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *fileUploadServiceUploadFileServer) Recv() (*UploadFileRequest, error) {
	m := new(UploadFileRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FileUploadService_ServiceDesc is the grpc.ServiceDesc for FileUploadService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FileUploadService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "learnings.v1.FileUploadService",
	HandlerType: (*FileUploadServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UploadFile",
			Handler:       _FileUploadService_UploadFile_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "learnings/v1/file.proto",
}
