// Code generated by thriftgo (0.4.1). DO NOT EDIT.

package miner_integrate

import (
	"context"
	"fmt"
)

type UpdateMockProducerQpsReq struct {
	ProductBaseInfoQps       *int64 `thrift:"product_base_info_qps,1,optional" frugal:"1,optional,i64" json:"product_base_info_qps,omitempty"`
	ProductShopInfoQps       *int64 `thrift:"product_shop_info_qps,2,optional" frugal:"2,optional,i64" json:"product_shop_info_qps,omitempty"`
	ProductLogoModelInfoQps  *int64 `thrift:"product_logo_model_info_qps,3,optional" frugal:"3,optional,i64" json:"product_logo_model_info_qps,omitempty"`
	ProductImageModelInfoQps *int64 `thrift:"product_image_model_info_qps,4,optional" frugal:"4,optional,i64" json:"product_image_model_info_qps,omitempty"`
	Test                     *int64 `thrift:"test,5,optional" frugal:"5,optional,i64" json:"test,omitempty"`
}

func NewUpdateMockProducerQpsReq() *UpdateMockProducerQpsReq {
	return &UpdateMockProducerQpsReq{}
}

func (p *UpdateMockProducerQpsReq) InitDefault() {
}

var UpdateMockProducerQpsReq_ProductBaseInfoQps_DEFAULT int64

func (p *UpdateMockProducerQpsReq) GetProductBaseInfoQps() (v int64) {
	if !p.IsSetProductBaseInfoQps() {
		return UpdateMockProducerQpsReq_ProductBaseInfoQps_DEFAULT
	}
	return *p.ProductBaseInfoQps
}

var UpdateMockProducerQpsReq_ProductShopInfoQps_DEFAULT int64

func (p *UpdateMockProducerQpsReq) GetProductShopInfoQps() (v int64) {
	if !p.IsSetProductShopInfoQps() {
		return UpdateMockProducerQpsReq_ProductShopInfoQps_DEFAULT
	}
	return *p.ProductShopInfoQps
}

var UpdateMockProducerQpsReq_ProductLogoModelInfoQps_DEFAULT int64

func (p *UpdateMockProducerQpsReq) GetProductLogoModelInfoQps() (v int64) {
	if !p.IsSetProductLogoModelInfoQps() {
		return UpdateMockProducerQpsReq_ProductLogoModelInfoQps_DEFAULT
	}
	return *p.ProductLogoModelInfoQps
}

var UpdateMockProducerQpsReq_ProductImageModelInfoQps_DEFAULT int64

func (p *UpdateMockProducerQpsReq) GetProductImageModelInfoQps() (v int64) {
	if !p.IsSetProductImageModelInfoQps() {
		return UpdateMockProducerQpsReq_ProductImageModelInfoQps_DEFAULT
	}
	return *p.ProductImageModelInfoQps
}

var UpdateMockProducerQpsReq_Test_DEFAULT int64

func (p *UpdateMockProducerQpsReq) GetTest() (v int64) {
	if !p.IsSetTest() {
		return UpdateMockProducerQpsReq_Test_DEFAULT
	}
	return *p.Test
}
func (p *UpdateMockProducerQpsReq) SetProductBaseInfoQps(val *int64) {
	p.ProductBaseInfoQps = val
}
func (p *UpdateMockProducerQpsReq) SetProductShopInfoQps(val *int64) {
	p.ProductShopInfoQps = val
}
func (p *UpdateMockProducerQpsReq) SetProductLogoModelInfoQps(val *int64) {
	p.ProductLogoModelInfoQps = val
}
func (p *UpdateMockProducerQpsReq) SetProductImageModelInfoQps(val *int64) {
	p.ProductImageModelInfoQps = val
}
func (p *UpdateMockProducerQpsReq) SetTest(val *int64) {
	p.Test = val
}

func (p *UpdateMockProducerQpsReq) IsSetProductBaseInfoQps() bool {
	return p.ProductBaseInfoQps != nil
}

func (p *UpdateMockProducerQpsReq) IsSetProductShopInfoQps() bool {
	return p.ProductShopInfoQps != nil
}

func (p *UpdateMockProducerQpsReq) IsSetProductLogoModelInfoQps() bool {
	return p.ProductLogoModelInfoQps != nil
}

func (p *UpdateMockProducerQpsReq) IsSetProductImageModelInfoQps() bool {
	return p.ProductImageModelInfoQps != nil
}

func (p *UpdateMockProducerQpsReq) IsSetTest() bool {
	return p.Test != nil
}

func (p *UpdateMockProducerQpsReq) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("UpdateMockProducerQpsReq(%+v)", *p)
}

var fieldIDToName_UpdateMockProducerQpsReq = map[int16]string{
	1: "product_base_info_qps",
	2: "product_shop_info_qps",
	3: "product_logo_model_info_qps",
	4: "product_image_model_info_qps",
	5: "test",
}

type UpdateMockProducerQpsResp struct {
}

func NewUpdateMockProducerQpsResp() *UpdateMockProducerQpsResp {
	return &UpdateMockProducerQpsResp{}
}

func (p *UpdateMockProducerQpsResp) InitDefault() {
}

func (p *UpdateMockProducerQpsResp) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("UpdateMockProducerQpsResp(%+v)", *p)
}

var fieldIDToName_UpdateMockProducerQpsResp = map[int16]string{}

type MinerIntegrate interface {
	UpdateMockProducerQps(ctx context.Context, req *UpdateMockProducerQpsReq) (r *UpdateMockProducerQpsResp, err error)
}

type MinerIntegrateUpdateMockProducerQpsArgs struct {
	Req *UpdateMockProducerQpsReq `thrift:"req,1" frugal:"1,default,UpdateMockProducerQpsReq" json:"req"`
}

func NewMinerIntegrateUpdateMockProducerQpsArgs() *MinerIntegrateUpdateMockProducerQpsArgs {
	return &MinerIntegrateUpdateMockProducerQpsArgs{}
}

func (p *MinerIntegrateUpdateMockProducerQpsArgs) InitDefault() {
}

var MinerIntegrateUpdateMockProducerQpsArgs_Req_DEFAULT *UpdateMockProducerQpsReq

func (p *MinerIntegrateUpdateMockProducerQpsArgs) GetReq() (v *UpdateMockProducerQpsReq) {
	if !p.IsSetReq() {
		return MinerIntegrateUpdateMockProducerQpsArgs_Req_DEFAULT
	}
	return p.Req
}
func (p *MinerIntegrateUpdateMockProducerQpsArgs) SetReq(val *UpdateMockProducerQpsReq) {
	p.Req = val
}

func (p *MinerIntegrateUpdateMockProducerQpsArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *MinerIntegrateUpdateMockProducerQpsArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("MinerIntegrateUpdateMockProducerQpsArgs(%+v)", *p)
}

var fieldIDToName_MinerIntegrateUpdateMockProducerQpsArgs = map[int16]string{
	1: "req",
}

type MinerIntegrateUpdateMockProducerQpsResult struct {
	Success *UpdateMockProducerQpsResp `thrift:"success,0,optional" frugal:"0,optional,UpdateMockProducerQpsResp" json:"success,omitempty"`
}

func NewMinerIntegrateUpdateMockProducerQpsResult() *MinerIntegrateUpdateMockProducerQpsResult {
	return &MinerIntegrateUpdateMockProducerQpsResult{}
}

func (p *MinerIntegrateUpdateMockProducerQpsResult) InitDefault() {
}

var MinerIntegrateUpdateMockProducerQpsResult_Success_DEFAULT *UpdateMockProducerQpsResp

func (p *MinerIntegrateUpdateMockProducerQpsResult) GetSuccess() (v *UpdateMockProducerQpsResp) {
	if !p.IsSetSuccess() {
		return MinerIntegrateUpdateMockProducerQpsResult_Success_DEFAULT
	}
	return p.Success
}
func (p *MinerIntegrateUpdateMockProducerQpsResult) SetSuccess(x interface{}) {
	p.Success = x.(*UpdateMockProducerQpsResp)
}

func (p *MinerIntegrateUpdateMockProducerQpsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MinerIntegrateUpdateMockProducerQpsResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("MinerIntegrateUpdateMockProducerQpsResult(%+v)", *p)
}

var fieldIDToName_MinerIntegrateUpdateMockProducerQpsResult = map[int16]string{
	0: "success",
}
