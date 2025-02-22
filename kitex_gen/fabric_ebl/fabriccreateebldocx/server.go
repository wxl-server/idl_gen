// Code generated by Kitex v0.12.2. DO NOT EDIT.
package fabriccreateebldocx

import (
	server "github.com/cloudwego/kitex/server"
	fabric_ebl "github.com/wxl-server/idl_gen/kitex_gen/fabric_ebl"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler fabric_ebl.FabricCreateEblDocx, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)
	options = append(options, server.WithCompatibleMiddlewareForUnary())

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}

func RegisterService(svr server.Server, handler fabric_ebl.FabricCreateEblDocx, opts ...server.RegisterOption) error {
	return svr.RegisterService(serviceInfo(), handler, opts...)
}
