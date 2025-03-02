// Code generated by Kitex v0.12.3. DO NOT EDIT.

package minercore

import (
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	miner_core "github.com/wxl-server/idl_gen/kitex_gen/miner_core"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"SignUp": kitex.NewMethodInfo(
		signUpHandler,
		newMinerCoreSignUpArgs,
		newMinerCoreSignUpResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"Login": kitex.NewMethodInfo(
		loginHandler,
		newMinerCoreLoginArgs,
		newMinerCoreLoginResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"QueryUserList": kitex.NewMethodInfo(
		queryUserListHandler,
		newMinerCoreQueryUserListArgs,
		newMinerCoreQueryUserListResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"QueryJobList": kitex.NewMethodInfo(
		queryJobListHandler,
		newMinerCoreQueryJobListArgs,
		newMinerCoreQueryJobListResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"CreateJob": kitex.NewMethodInfo(
		createJobHandler,
		newMinerCoreCreateJobArgs,
		newMinerCoreCreateJobResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"DeleteJob": kitex.NewMethodInfo(
		deleteJobHandler,
		newMinerCoreDeleteJobArgs,
		newMinerCoreDeleteJobResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"QueryIndicatorList": kitex.NewMethodInfo(
		queryIndicatorListHandler,
		newMinerCoreQueryIndicatorListArgs,
		newMinerCoreQueryIndicatorListResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"RunTask": kitex.NewMethodInfo(
		runTaskHandler,
		newMinerCoreRunTaskArgs,
		newMinerCoreRunTaskResult,
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
		KiteXGenVersion: "v0.12.3",
		Extra:           extra,
	}
	return svcInfo
}

func signUpHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*miner_core.MinerCoreSignUpArgs)
	realResult := result.(*miner_core.MinerCoreSignUpResult)
	success, err := handler.(miner_core.MinerCore).SignUp(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newMinerCoreSignUpArgs() interface{} {
	return miner_core.NewMinerCoreSignUpArgs()
}

func newMinerCoreSignUpResult() interface{} {
	return miner_core.NewMinerCoreSignUpResult()
}

func loginHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*miner_core.MinerCoreLoginArgs)
	realResult := result.(*miner_core.MinerCoreLoginResult)
	success, err := handler.(miner_core.MinerCore).Login(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newMinerCoreLoginArgs() interface{} {
	return miner_core.NewMinerCoreLoginArgs()
}

func newMinerCoreLoginResult() interface{} {
	return miner_core.NewMinerCoreLoginResult()
}

func queryUserListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*miner_core.MinerCoreQueryUserListArgs)
	realResult := result.(*miner_core.MinerCoreQueryUserListResult)
	success, err := handler.(miner_core.MinerCore).QueryUserList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newMinerCoreQueryUserListArgs() interface{} {
	return miner_core.NewMinerCoreQueryUserListArgs()
}

func newMinerCoreQueryUserListResult() interface{} {
	return miner_core.NewMinerCoreQueryUserListResult()
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

func createJobHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*miner_core.MinerCoreCreateJobArgs)
	realResult := result.(*miner_core.MinerCoreCreateJobResult)
	success, err := handler.(miner_core.MinerCore).CreateJob(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newMinerCoreCreateJobArgs() interface{} {
	return miner_core.NewMinerCoreCreateJobArgs()
}

func newMinerCoreCreateJobResult() interface{} {
	return miner_core.NewMinerCoreCreateJobResult()
}

func deleteJobHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*miner_core.MinerCoreDeleteJobArgs)
	realResult := result.(*miner_core.MinerCoreDeleteJobResult)
	success, err := handler.(miner_core.MinerCore).DeleteJob(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newMinerCoreDeleteJobArgs() interface{} {
	return miner_core.NewMinerCoreDeleteJobArgs()
}

func newMinerCoreDeleteJobResult() interface{} {
	return miner_core.NewMinerCoreDeleteJobResult()
}

func queryIndicatorListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*miner_core.MinerCoreQueryIndicatorListArgs)
	realResult := result.(*miner_core.MinerCoreQueryIndicatorListResult)
	success, err := handler.(miner_core.MinerCore).QueryIndicatorList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newMinerCoreQueryIndicatorListArgs() interface{} {
	return miner_core.NewMinerCoreQueryIndicatorListArgs()
}

func newMinerCoreQueryIndicatorListResult() interface{} {
	return miner_core.NewMinerCoreQueryIndicatorListResult()
}

func runTaskHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*miner_core.MinerCoreRunTaskArgs)
	realResult := result.(*miner_core.MinerCoreRunTaskResult)
	success, err := handler.(miner_core.MinerCore).RunTask(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newMinerCoreRunTaskArgs() interface{} {
	return miner_core.NewMinerCoreRunTaskArgs()
}

func newMinerCoreRunTaskResult() interface{} {
	return miner_core.NewMinerCoreRunTaskResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) SignUp(ctx context.Context, req *miner_core.SignUpReq) (r *miner_core.SignUpResp, err error) {
	var _args miner_core.MinerCoreSignUpArgs
	_args.Req = req
	var _result miner_core.MinerCoreSignUpResult
	if err = p.c.Call(ctx, "SignUp", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Login(ctx context.Context, req *miner_core.LoginReq) (r *miner_core.LoginResp, err error) {
	var _args miner_core.MinerCoreLoginArgs
	_args.Req = req
	var _result miner_core.MinerCoreLoginResult
	if err = p.c.Call(ctx, "Login", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) QueryUserList(ctx context.Context, req *miner_core.QueryUserListReq) (r *miner_core.QueryUserListResp, err error) {
	var _args miner_core.MinerCoreQueryUserListArgs
	_args.Req = req
	var _result miner_core.MinerCoreQueryUserListResult
	if err = p.c.Call(ctx, "QueryUserList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
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

func (p *kClient) CreateJob(ctx context.Context, req *miner_core.CreateJobReq) (r *miner_core.CreateJobResp, err error) {
	var _args miner_core.MinerCoreCreateJobArgs
	_args.Req = req
	var _result miner_core.MinerCoreCreateJobResult
	if err = p.c.Call(ctx, "CreateJob", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DeleteJob(ctx context.Context, req *miner_core.DeleteJobReq) (r *miner_core.DeleteJobResp, err error) {
	var _args miner_core.MinerCoreDeleteJobArgs
	_args.Req = req
	var _result miner_core.MinerCoreDeleteJobResult
	if err = p.c.Call(ctx, "DeleteJob", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) QueryIndicatorList(ctx context.Context, req *miner_core.QueryIndicatorListReq) (r *miner_core.QueryIndicatorListResp, err error) {
	var _args miner_core.MinerCoreQueryIndicatorListArgs
	_args.Req = req
	var _result miner_core.MinerCoreQueryIndicatorListResult
	if err = p.c.Call(ctx, "QueryIndicatorList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) RunTask(ctx context.Context, req *miner_core.RunTaskReq) (r *miner_core.RunTaskResp, err error) {
	var _args miner_core.MinerCoreRunTaskArgs
	_args.Req = req
	var _result miner_core.MinerCoreRunTaskResult
	if err = p.c.Call(ctx, "RunTask", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
