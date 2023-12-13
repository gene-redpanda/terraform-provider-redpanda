package mocks

import (
	"context"

	controlplanev1beta1 "github.com/redpanda-data/terraform-provider-redpanda/proto/gen/go/redpanda/api/controlplane/v1beta1"
	"google.golang.org/grpc"
)

var _ controlplanev1beta1.NamespaceServiceClient = MockNamespaceServiceClient{}

type MockNamespaceServiceClient struct {
	// You can add fields to store mock responses or any other information needed for your tests
}

func (MockNamespaceServiceClient) CreateNamespace(_ context.Context, _ *controlplanev1beta1.CreateNamespaceRequest, _ ...grpc.CallOption) (*controlplanev1beta1.Namespace, error) {
	// Implement mock logic here
	// Return a mock Namespace object and nil error for successful scenario
	// Return nil and a mock error for failure scenario
	return &controlplanev1beta1.Namespace{}, nil
}

func (MockNamespaceServiceClient) UpdateNamespace(_ context.Context, _ *controlplanev1beta1.UpdateNamespaceRequest, _ ...grpc.CallOption) (*controlplanev1beta1.Namespace, error) {
	// Implement mock logic here
	return &controlplanev1beta1.Namespace{}, nil
}

func (MockNamespaceServiceClient) GetNamespace(_ context.Context, _ *controlplanev1beta1.GetNamespaceRequest, _ ...grpc.CallOption) (*controlplanev1beta1.Namespace, error) {
	// Implement mock logic here
	return &controlplanev1beta1.Namespace{}, nil
}

func (MockNamespaceServiceClient) ListNamespaces(_ context.Context, _ *controlplanev1beta1.ListNamespacesRequest, _ ...grpc.CallOption) (*controlplanev1beta1.ListNamespacesResponse, error) {
	// Implement mock logic here
	return &controlplanev1beta1.ListNamespacesResponse{}, nil
}

func (MockNamespaceServiceClient) DeleteNamespace(_ context.Context, _ *controlplanev1beta1.DeleteNamespaceRequest, _ ...grpc.CallOption) (*controlplanev1beta1.DeleteNamespaceResponse, error) {
	// Implement mock logic here
	return &controlplanev1beta1.DeleteNamespaceResponse{}, nil
}
