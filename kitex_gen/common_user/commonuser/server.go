// Code generated by Kitex v0.12.2. DO NOT EDIT.
package commonuser

import (
	server "github.com/cloudwego/kitex/server"
	common_user "github.com/wxl-server/idl_gen/kitex_gen/common_user"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler common_user.CommonUser, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)
	options = append(options, server.WithCompatibleMiddlewareForUnary())

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}

func RegisterService(svr server.Server, handler common_user.CommonUser, opts ...server.RegisterOption) error {
	return svr.RegisterService(serviceInfo(), handler, opts...)
}
