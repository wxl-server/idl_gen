// Code generated by Kitex v0.12.3. DO NOT EDIT.

package fabricipfs

import (
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	fabric_ipfs "github.com/wxl-server/idl_gen/kitex_gen/fabric_ipfs"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"CreateEblDocx": kitex.NewMethodInfo(
		createEblDocxHandler,
		newFabricIpfsCreateEblDocxArgs,
		newFabricIpfsCreateEblDocxResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"SealEbl": kitex.NewMethodInfo(
		sealEblHandler,
		newFabricIpfsSealEblArgs,
		newFabricIpfsSealEblResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	fabricIpfsServiceInfo                = NewServiceInfo()
	fabricIpfsServiceInfoForClient       = NewServiceInfoForClient()
	fabricIpfsServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return fabricIpfsServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return fabricIpfsServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return fabricIpfsServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "FabricIpfs"
	handlerType := (*fabric_ipfs.FabricIpfs)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "fabric_ipfs",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.12.3",
		Extra:           extra,
	}
	return svcInfo
}

func createEblDocxHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*fabric_ipfs.FabricIpfsCreateEblDocxArgs)
	realResult := result.(*fabric_ipfs.FabricIpfsCreateEblDocxResult)
	success, err := handler.(fabric_ipfs.FabricIpfs).CreateEblDocx(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newFabricIpfsCreateEblDocxArgs() interface{} {
	return fabric_ipfs.NewFabricIpfsCreateEblDocxArgs()
}

func newFabricIpfsCreateEblDocxResult() interface{} {
	return fabric_ipfs.NewFabricIpfsCreateEblDocxResult()
}

func sealEblHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*fabric_ipfs.FabricIpfsSealEblArgs)
	realResult := result.(*fabric_ipfs.FabricIpfsSealEblResult)
	success, err := handler.(fabric_ipfs.FabricIpfs).SealEbl(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newFabricIpfsSealEblArgs() interface{} {
	return fabric_ipfs.NewFabricIpfsSealEblArgs()
}

func newFabricIpfsSealEblResult() interface{} {
	return fabric_ipfs.NewFabricIpfsSealEblResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) CreateEblDocx(ctx context.Context, req *fabric_ipfs.CreateEblDocxReq) (r *fabric_ipfs.CreateEblDocxResp, err error) {
	var _args fabric_ipfs.FabricIpfsCreateEblDocxArgs
	_args.Req = req
	var _result fabric_ipfs.FabricIpfsCreateEblDocxResult
	if err = p.c.Call(ctx, "CreateEblDocx", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) SealEbl(ctx context.Context, req *fabric_ipfs.SealEblReq) (r *fabric_ipfs.SealEblResp, err error) {
	var _args fabric_ipfs.FabricIpfsSealEblArgs
	_args.Req = req
	var _result fabric_ipfs.FabricIpfsSealEblResult
	if err = p.c.Call(ctx, "SealEbl", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
