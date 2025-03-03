// Code generated by Kitex v0.12.3. DO NOT EDIT.

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
	CreateEbl(ctx context.Context, req *fabric_ebl.CreateEblReq, callOptions ...callopt.Option) (r *fabric_ebl.CreateEblResp, err error)
	QueryAllEblList(ctx context.Context, req *fabric_ebl.QueryAllEblListReq, callOptions ...callopt.Option) (r *fabric_ebl.QueryAllEblListResp, err error)
	QueryEblList(ctx context.Context, req *fabric_ebl.QueryEblListReq, callOptions ...callopt.Option) (r *fabric_ebl.QueryEblListResp, err error)
	OperateEbl(ctx context.Context, req *fabric_ebl.OperateEblReq, callOptions ...callopt.Option) (r *fabric_ebl.OperateEblResp, err error)
	UploadSeal(ctx context.Context, req *fabric_ebl.UploadSealReq, callOptions ...callopt.Option) (r *fabric_ebl.UploadSealResp, err error)
	CheckToken(ctx context.Context, req *fabric_ebl.CheckTokenReq, callOptions ...callopt.Option) (r *fabric_ebl.CheckTokenResp, err error)
	CreateInvoice(ctx context.Context, req *fabric_ebl.CreateInvoiceReq, callOptions ...callopt.Option) (r *fabric_ebl.CreateInvoiceResp, err error)
	CreateContract(ctx context.Context, req *fabric_ebl.CreateContractReq, callOptions ...callopt.Option) (r *fabric_ebl.CreateContractResp, err error)
	CreateDocument(ctx context.Context, req *fabric_ebl.CreateDocumentReq, callOptions ...callopt.Option) (r *fabric_ebl.CreateDocumentResp, err error)
	QueryInvoiceList(ctx context.Context, req *fabric_ebl.QueryInvoiceListReq, callOptions ...callopt.Option) (r *fabric_ebl.QueryInvoiceListResp, err error)
	QueryContractList(ctx context.Context, req *fabric_ebl.QueryContractListReq, callOptions ...callopt.Option) (r *fabric_ebl.QueryContractListResp, err error)
	QueryDocumentList(ctx context.Context, req *fabric_ebl.QueryDocumentListReq, callOptions ...callopt.Option) (r *fabric_ebl.QueryDocumentListResp, err error)
	GetInvoice(ctx context.Context, req *fabric_ebl.GetInvoiceReq, callOptions ...callopt.Option) (r *fabric_ebl.GetInvoiceResp, err error)
	GetContract(ctx context.Context, req *fabric_ebl.GetContractReq, callOptions ...callopt.Option) (r *fabric_ebl.GetContractResp, err error)
	GetDocument(ctx context.Context, req *fabric_ebl.GetDocumentReq, callOptions ...callopt.Option) (r *fabric_ebl.GetDocumentResp, err error)
	QuerySeal(ctx context.Context, req *fabric_ebl.QuerySealReq, callOptions ...callopt.Option) (r *fabric_ebl.QuerySealResp, err error)
	DeleteSeal(ctx context.Context, req *fabric_ebl.DeleteSealReq, callOptions ...callopt.Option) (r *fabric_ebl.DeleteSealResp, err error)
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

func (p *kFabricEblClient) CreateEbl(ctx context.Context, req *fabric_ebl.CreateEblReq, callOptions ...callopt.Option) (r *fabric_ebl.CreateEblResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateEbl(ctx, req)
}

func (p *kFabricEblClient) QueryAllEblList(ctx context.Context, req *fabric_ebl.QueryAllEblListReq, callOptions ...callopt.Option) (r *fabric_ebl.QueryAllEblListResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.QueryAllEblList(ctx, req)
}

func (p *kFabricEblClient) QueryEblList(ctx context.Context, req *fabric_ebl.QueryEblListReq, callOptions ...callopt.Option) (r *fabric_ebl.QueryEblListResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.QueryEblList(ctx, req)
}

func (p *kFabricEblClient) OperateEbl(ctx context.Context, req *fabric_ebl.OperateEblReq, callOptions ...callopt.Option) (r *fabric_ebl.OperateEblResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.OperateEbl(ctx, req)
}

func (p *kFabricEblClient) UploadSeal(ctx context.Context, req *fabric_ebl.UploadSealReq, callOptions ...callopt.Option) (r *fabric_ebl.UploadSealResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UploadSeal(ctx, req)
}

func (p *kFabricEblClient) CheckToken(ctx context.Context, req *fabric_ebl.CheckTokenReq, callOptions ...callopt.Option) (r *fabric_ebl.CheckTokenResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CheckToken(ctx, req)
}

func (p *kFabricEblClient) CreateInvoice(ctx context.Context, req *fabric_ebl.CreateInvoiceReq, callOptions ...callopt.Option) (r *fabric_ebl.CreateInvoiceResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateInvoice(ctx, req)
}

func (p *kFabricEblClient) CreateContract(ctx context.Context, req *fabric_ebl.CreateContractReq, callOptions ...callopt.Option) (r *fabric_ebl.CreateContractResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateContract(ctx, req)
}

func (p *kFabricEblClient) CreateDocument(ctx context.Context, req *fabric_ebl.CreateDocumentReq, callOptions ...callopt.Option) (r *fabric_ebl.CreateDocumentResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateDocument(ctx, req)
}

func (p *kFabricEblClient) QueryInvoiceList(ctx context.Context, req *fabric_ebl.QueryInvoiceListReq, callOptions ...callopt.Option) (r *fabric_ebl.QueryInvoiceListResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.QueryInvoiceList(ctx, req)
}

func (p *kFabricEblClient) QueryContractList(ctx context.Context, req *fabric_ebl.QueryContractListReq, callOptions ...callopt.Option) (r *fabric_ebl.QueryContractListResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.QueryContractList(ctx, req)
}

func (p *kFabricEblClient) QueryDocumentList(ctx context.Context, req *fabric_ebl.QueryDocumentListReq, callOptions ...callopt.Option) (r *fabric_ebl.QueryDocumentListResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.QueryDocumentList(ctx, req)
}

func (p *kFabricEblClient) GetInvoice(ctx context.Context, req *fabric_ebl.GetInvoiceReq, callOptions ...callopt.Option) (r *fabric_ebl.GetInvoiceResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetInvoice(ctx, req)
}

func (p *kFabricEblClient) GetContract(ctx context.Context, req *fabric_ebl.GetContractReq, callOptions ...callopt.Option) (r *fabric_ebl.GetContractResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetContract(ctx, req)
}

func (p *kFabricEblClient) GetDocument(ctx context.Context, req *fabric_ebl.GetDocumentReq, callOptions ...callopt.Option) (r *fabric_ebl.GetDocumentResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetDocument(ctx, req)
}

func (p *kFabricEblClient) QuerySeal(ctx context.Context, req *fabric_ebl.QuerySealReq, callOptions ...callopt.Option) (r *fabric_ebl.QuerySealResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.QuerySeal(ctx, req)
}

func (p *kFabricEblClient) DeleteSeal(ctx context.Context, req *fabric_ebl.DeleteSealReq, callOptions ...callopt.Option) (r *fabric_ebl.DeleteSealResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteSeal(ctx, req)
}
