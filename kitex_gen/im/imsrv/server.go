// Code generated by Kitex v0.8.0. DO NOT EDIT.
package imsrv

import (
	server "github.com/cloudwego/kitex/server"
	im "github.com/xu756/imlogic/kitex_gen/im"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler im.ImSrv, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}
