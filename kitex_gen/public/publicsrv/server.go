// Code generated by Kitex v0.7.3. DO NOT EDIT.
package publicsrv

import (
	server "github.com/cloudwego/kitex/server"
	public "github.com/xu756/imlogic/kitex_gen/public"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler public.PublicSrv, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}
