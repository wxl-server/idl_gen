// Code generated by Kitex v0.12.2. DO NOT EDIT.

package commonuser

import (
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	common_user "github.com/wxl-server/idl_gen/kitex_gen/common_user"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"SignUp": kitex.NewMethodInfo(
		signUpHandler,
		newCommonUserSignUpArgs,
		newCommonUserSignUpResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"UpdatePassword": kitex.NewMethodInfo(
		updatePasswordHandler,
		newCommonUserUpdatePasswordArgs,
		newCommonUserUpdatePasswordResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"Login": kitex.NewMethodInfo(
		loginHandler,
		newCommonUserLoginArgs,
		newCommonUserLoginResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"ValidateToken": kitex.NewMethodInfo(
		validateTokenHandler,
		newCommonUserValidateTokenArgs,
		newCommonUserValidateTokenResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	commonUserServiceInfo                = NewServiceInfo()
	commonUserServiceInfoForClient       = NewServiceInfoForClient()
	commonUserServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return commonUserServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return commonUserServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return commonUserServiceInfoForClient
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
	serviceName := "CommonUser"
	handlerType := (*common_user.CommonUser)(nil)
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
		"PackageName": "common_user",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.12.2",
		Extra:           extra,
	}
	return svcInfo
}

func signUpHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*common_user.CommonUserSignUpArgs)
	realResult := result.(*common_user.CommonUserSignUpResult)
	success, err := handler.(common_user.CommonUser).SignUp(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newCommonUserSignUpArgs() interface{} {
	return common_user.NewCommonUserSignUpArgs()
}

func newCommonUserSignUpResult() interface{} {
	return common_user.NewCommonUserSignUpResult()
}

func updatePasswordHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*common_user.CommonUserUpdatePasswordArgs)
	realResult := result.(*common_user.CommonUserUpdatePasswordResult)
	success, err := handler.(common_user.CommonUser).UpdatePassword(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newCommonUserUpdatePasswordArgs() interface{} {
	return common_user.NewCommonUserUpdatePasswordArgs()
}

func newCommonUserUpdatePasswordResult() interface{} {
	return common_user.NewCommonUserUpdatePasswordResult()
}

func loginHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*common_user.CommonUserLoginArgs)
	realResult := result.(*common_user.CommonUserLoginResult)
	success, err := handler.(common_user.CommonUser).Login(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newCommonUserLoginArgs() interface{} {
	return common_user.NewCommonUserLoginArgs()
}

func newCommonUserLoginResult() interface{} {
	return common_user.NewCommonUserLoginResult()
}

func validateTokenHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*common_user.CommonUserValidateTokenArgs)
	realResult := result.(*common_user.CommonUserValidateTokenResult)
	success, err := handler.(common_user.CommonUser).ValidateToken(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newCommonUserValidateTokenArgs() interface{} {
	return common_user.NewCommonUserValidateTokenArgs()
}

func newCommonUserValidateTokenResult() interface{} {
	return common_user.NewCommonUserValidateTokenResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) SignUp(ctx context.Context, req *common_user.SignUpReq) (r *common_user.SignUpResp, err error) {
	var _args common_user.CommonUserSignUpArgs
	_args.Req = req
	var _result common_user.CommonUserSignUpResult
	if err = p.c.Call(ctx, "SignUp", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdatePassword(ctx context.Context, req *common_user.UpdatePasswordReq) (r *common_user.UpdatePasswordResp, err error) {
	var _args common_user.CommonUserUpdatePasswordArgs
	_args.Req = req
	var _result common_user.CommonUserUpdatePasswordResult
	if err = p.c.Call(ctx, "UpdatePassword", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Login(ctx context.Context, req *common_user.LoginReq) (r *common_user.LoginResp, err error) {
	var _args common_user.CommonUserLoginArgs
	_args.Req = req
	var _result common_user.CommonUserLoginResult
	if err = p.c.Call(ctx, "Login", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ValidateToken(ctx context.Context, req *common_user.ValidateTokenReq) (r *common_user.ValidateTokenResp, err error) {
	var _args common_user.CommonUserValidateTokenArgs
	_args.Req = req
	var _result common_user.CommonUserValidateTokenResult
	if err = p.c.Call(ctx, "ValidateToken", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
