// Code generated by Kitex v0.12.2. DO NOT EDIT.

package fabricebl

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	fabric_ebl "github.com/wxl-server/idl_gen/kitex_gen/fabric_ebl"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	CreateCompany(ctx context.Context, req *fabric_ebl.CreateCompanyReq, callOptions ...callopt.Option) (r *fabric_ebl.CreateCompanyResp, err error)
	Login(ctx context.Context, req *fabric_ebl.LoginReq, callOptions ...callopt.Option) (r *fabric_ebl.LoginResp, err error)
	GetUserInfo(ctx context.Context, req *fabric_ebl.GetUserInfoReq, callOptions ...callopt.Option) (r *fabric_ebl.GetUserInfoResp, err error)
	GetCompanyAllList(ctx context.Context, req *fabric_ebl.GetCompanyAllListReq, callOptions ...callopt.Option) (r *fabric_ebl.GetCompanyAllListResp, err error)
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
	return &kFabricEblClient{
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

type kFabricEblClient struct {
	*kClient
}

func (p *kFabricEblClient) CreateCompany(ctx context.Context, req *fabric_ebl.CreateCompanyReq, callOptions ...callopt.Option) (r *fabric_ebl.CreateCompanyResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateCompany(ctx, req)
}

func (p *kFabricEblClient) Login(ctx context.Context, req *fabric_ebl.LoginReq, callOptions ...callopt.Option) (r *fabric_ebl.LoginResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Login(ctx, req)
}

func (p *kFabricEblClient) GetUserInfo(ctx context.Context, req *fabric_ebl.GetUserInfoReq, callOptions ...callopt.Option) (r *fabric_ebl.GetUserInfoResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetUserInfo(ctx, req)
}

func (p *kFabricEblClient) GetCompanyAllList(ctx context.Context, req *fabric_ebl.GetCompanyAllListReq, callOptions ...callopt.Option) (r *fabric_ebl.GetCompanyAllListResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetCompanyAllList(ctx, req)
}
