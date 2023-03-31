// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package userclient

import (
	"context"

	"boat/rpc/user/pb/user"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	IdRequest    = user.IdRequest
	TestRequest  = user.TestRequest
	UserResponse = user.UserResponse

	User interface {
		GetUser(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*UserResponse, error)
		TestErr(ctx context.Context, in *TestRequest, opts ...grpc.CallOption) (*UserResponse, error)
	}

	defaultUser struct {
		cli zrpc.Client
	}
)

func NewUser(cli zrpc.Client) User {
	return &defaultUser{
		cli: cli,
	}
}

func (m *defaultUser) GetUser(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.GetUser(ctx, in, opts...)
}

func (m *defaultUser) TestErr(ctx context.Context, in *TestRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.TestErr(ctx, in, opts...)
}
