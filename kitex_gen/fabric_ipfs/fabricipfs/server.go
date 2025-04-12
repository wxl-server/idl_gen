// Code generated by Kitex v0.12.3. DO NOT EDIT.
package fabricipfs

import (
	server "github.com/cloudwego/kitex/server"
	fabric_ipfs "github.com/wxl-server/idl_gen/kitex_gen/fabric_ipfs"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler fabric_ipfs.FabricIpfs, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)
	options = append(options, server.WithCompatibleMiddlewareForUnary())

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}

func RegisterService(svr server.Server, handler fabric_ipfs.FabricIpfs, opts ...server.RegisterOption) error {
	return svr.RegisterService(serviceInfo(), handler, opts...)
}
