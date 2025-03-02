// Code generated by Kitex v0.12.3. DO NOT EDIT.

package minercore

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	miner_core "github.com/wxl-server/idl_gen/kitex_gen/miner_core"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	SignUp(ctx context.Context, req *miner_core.SignUpReq, callOptions ...callopt.Option) (r *miner_core.SignUpResp, err error)
	Login(ctx context.Context, req *miner_core.LoginReq, callOptions ...callopt.Option) (r *miner_core.LoginResp, err error)
	QueryJobList(ctx context.Context, req *miner_core.QueryJobListReq, callOptions ...callopt.Option) (r *miner_core.QueryJobListResp, err error)
	CreateJob(ctx context.Context, req *miner_core.CreateJobReq, callOptions ...callopt.Option) (r *miner_core.CreateJobResp, err error)
	DeleteJob(ctx context.Context, req *miner_core.DeleteJobReq, callOptions ...callopt.Option) (r *miner_core.DeleteJobResp, err error)
	QueryIndicatorList(ctx context.Context, req *miner_core.QueryIndicatorListReq, callOptions ...callopt.Option) (r *miner_core.QueryIndicatorListResp, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfoForClient(), options...)
	if err != nil {
		return nil, err
	}
	return &kMinerCoreClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kMinerCoreClient struct {
	*kClient
}

func (p *kMinerCoreClient) SignUp(ctx context.Context, req *miner_core.SignUpReq, callOptions ...callopt.Option) (r *miner_core.SignUpResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SignUp(ctx, req)
}

func (p *kMinerCoreClient) Login(ctx context.Context, req *miner_core.LoginReq, callOptions ...callopt.Option) (r *miner_core.LoginResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Login(ctx, req)
}

func (p *kMinerCoreClient) QueryJobList(ctx context.Context, req *miner_core.QueryJobListReq, callOptions ...callopt.Option) (r *miner_core.QueryJobListResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.QueryJobList(ctx, req)
}

func (p *kMinerCoreClient) CreateJob(ctx context.Context, req *miner_core.CreateJobReq, callOptions ...callopt.Option) (r *miner_core.CreateJobResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateJob(ctx, req)
}

func (p *kMinerCoreClient) DeleteJob(ctx context.Context, req *miner_core.DeleteJobReq, callOptions ...callopt.Option) (r *miner_core.DeleteJobResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteJob(ctx, req)
}

func (p *kMinerCoreClient) QueryIndicatorList(ctx context.Context, req *miner_core.QueryIndicatorListReq, callOptions ...callopt.Option) (r *miner_core.QueryIndicatorListResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.QueryIndicatorList(ctx, req)
}
