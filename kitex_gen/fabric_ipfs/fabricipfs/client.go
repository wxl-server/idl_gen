// Code generated by Kitex v0.13.1. DO NOT EDIT.

package fabricipfs

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	fabric_ipfs "github.com/wxl-server/idl_gen/kitex_gen/fabric_ipfs"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	CreateEblDocx(ctx context.Context, req *fabric_ipfs.CreateEblDocxReq, callOptions ...callopt.Option) (r *fabric_ipfs.CreateEblDocxResp, err error)
	SealEbl(ctx context.Context, req *fabric_ipfs.SealEblReq, callOptions ...callopt.Option) (r *fabric_ipfs.SealEblResp, err error)
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
	return &kFabricIpfsClient{
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

type kFabricIpfsClient struct {
	*kClient
}

func (p *kFabricIpfsClient) CreateEblDocx(ctx context.Context, req *fabric_ipfs.CreateEblDocxReq, callOptions ...callopt.Option) (r *fabric_ipfs.CreateEblDocxResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateEblDocx(ctx, req)
}

func (p *kFabricIpfsClient) SealEbl(ctx context.Context, req *fabric_ipfs.SealEblReq, callOptions ...callopt.Option) (r *fabric_ipfs.SealEblResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SealEbl(ctx, req)
}
