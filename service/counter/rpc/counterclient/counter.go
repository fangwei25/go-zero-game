// Code generated by goctl. DO NOT EDIT!
// Source: counter.proto

package counterclient

import (
	"context"

	"web_game/service/counter/rpc/counter"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	ReqQuery  = counter.ReqQuery
	ReqUpdate = counter.ReqUpdate
	ResQuery  = counter.ResQuery
	ResUpdate = counter.ResUpdate

	Counter interface {
		Query(ctx context.Context, in *ReqQuery, opts ...grpc.CallOption) (*ResQuery, error)
		Update(ctx context.Context, in *ReqUpdate, opts ...grpc.CallOption) (*ResUpdate, error)
	}

	defaultCounter struct {
		cli zrpc.Client
	}
)

func NewCounter(cli zrpc.Client) Counter {
	return &defaultCounter{
		cli: cli,
	}
}

func (m *defaultCounter) Query(ctx context.Context, in *ReqQuery, opts ...grpc.CallOption) (*ResQuery, error) {
	client := counter.NewCounterClient(m.cli.Conn())
	return client.Query(ctx, in, opts...)
}

func (m *defaultCounter) Update(ctx context.Context, in *ReqUpdate, opts ...grpc.CallOption) (*ResUpdate, error) {
	client := counter.NewCounterClient(m.cli.Conn())
	return client.Update(ctx, in, opts...)
}