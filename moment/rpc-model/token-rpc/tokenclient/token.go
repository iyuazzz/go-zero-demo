// Code generated by goctl. DO NOT EDIT!
// Source: token.proto

package tokenclient

import (
	"context"

	"moment/rpc-model/token-rpc/types/token"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	TokenReply   = token.TokenReply
	TokenRequest = token.TokenRequest

	Token interface {
		GetToken(ctx context.Context, in *TokenRequest, opts ...grpc.CallOption) (*TokenReply, error)
	}

	defaultToken struct {
		cli zrpc.Client
	}
)

func NewToken(cli zrpc.Client) Token {
	return &defaultToken{
		cli: cli,
	}
}

func (m *defaultToken) GetToken(ctx context.Context, in *TokenRequest, opts ...grpc.CallOption) (*TokenReply, error) {
	client := token.NewTokenClient(m.cli.Conn())
	return client.GetToken(ctx, in, opts...)
}
