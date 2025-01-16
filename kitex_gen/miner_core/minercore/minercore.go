// Code generated by Kitex v0.12.1. DO NOT EDIT.

package minercore

import (
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	miner_core "github.com/qcq1/idl_gen/kitex_gen/miner_core"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"QueryJobList": kitex.NewMethodInfo(
		queryJobListHandler,
		newMinerCoreQueryJobListArgs,
		newMinerCoreQueryJobListResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	minerCoreServiceInfo                = NewServiceInfo()
	minerCoreServiceInfoForClient       = NewServiceInfoForClient()
	minerCoreServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return minerCoreServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return minerCoreServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return minerCoreServiceInfoForClient
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
	serviceName := "MinerCore"
	handlerType := (*miner_core.MinerCore)(nil)
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
		"PackageName": "miner_core",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.12.1",
		Extra:           extra,
	}
	return svcInfo
}

func queryJobListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*miner_core.MinerCoreQueryJobListArgs)
	realResult := result.(*miner_core.MinerCoreQueryJobListResult)
	success, err := handler.(miner_core.MinerCore).QueryJobList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newMinerCoreQueryJobListArgs() interface{} {
	return miner_core.NewMinerCoreQueryJobListArgs()
}

func newMinerCoreQueryJobListResult() interface{} {
	return miner_core.NewMinerCoreQueryJobListResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) QueryJobList(ctx context.Context, req *miner_core.QueryJobListReq) (r *miner_core.QueryJobListResp, err error) {
	var _args miner_core.MinerCoreQueryJobListArgs
	_args.Req = req
	var _result miner_core.MinerCoreQueryJobListResult
	if err = p.c.Call(ctx, "QueryJobList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
