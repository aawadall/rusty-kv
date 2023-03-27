package api

import (
	"context"

	"github.com/aawadall/simple-kv/proto_api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GrpcApi must be embedded to have forward compatible implementations.

func (GrpcApi) Get(context.Context, *proto_api.GetRequest) (*proto_api.KeyValueRecord, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (GrpcApi) Set(context.Context, *proto_api.SetRequest) (*proto_api.SetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Set not implemented")
}
func (GrpcApi) Delete(context.Context, *proto_api.DeleteRequest) (*proto_api.DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (GrpcApi) SetMetadata(context.Context, *proto_api.SetMetadataRequest) (*proto_api.SetMetadataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetMetadata not implemented")
}
func (GrpcApi) DeleteMetadata(context.Context, *proto_api.DeleteMetadataRequest) (*proto_api.DeleteMetadataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMetadata not implemented")
}
func (GrpcApi) GetAllMetadata(context.Context, *proto_api.GetAllMetadataRequest) (*proto_api.GetAllMetadataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllMetadata not implemented")
}
func (GrpcApi) Find(context.Context, *proto_api.FindRequest) (*proto_api.FindResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Find not implemented")
}
func (GrpcApi) FindByMetadata(context.Context, *proto_api.FindByMetadataRequest) (*proto_api.FindByMetadataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByMetadata not implemented")
}
func (GrpcApi) mustEmbedGrpcApi() {}
