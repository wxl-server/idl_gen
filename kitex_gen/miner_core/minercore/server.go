// Code generated by Kitex v1.18.1. DO NOT EDIT.
package minercore

import (
	byted "code.byted.org/kite/kitex/byted"
	server "code.byted.org/kite/kitex/server"
	miner_core "github.com/qcq1/rpc_miner_core/kitex_gen/miner_core"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler miner_core.MinerCore, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, byted.ServerSuite(serviceInfo()))

	options = append(options, opts...)
	options = append(options, server.WithCompatibleMiddlewareForUnary())

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}

// NewServerWithBytedConfig creates a server.Server with the given handler and options.
func NewServerWithBytedConfig(handler miner_core.MinerCore, config *byted.ServerConfig, opts ...server.Option) server.Server {
	var options []server.Option
	options = append(options, byted.ServerSuiteWithConfig(serviceInfo(), config))
	options = append(options, server.WithCompatibleMiddlewareForUnary())
	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}

func RegisterService(svr server.Server, handler miner_core.MinerCore, opts ...server.RegisterOption) error {
	return svr.RegisterService(serviceInfo(), handler, opts...)
}
