package main

import (
	"context"
	im "imlogic/kitex_gen/im"
)

// ImServerImpl implements the last service interface defined in the IDL.
type ImServerImpl struct{}

// SendMsg implements the ImServerImpl interface.
func (s *ImServerImpl) SendMsg(ctx context.Context, req *im.Message) (resp *im.MessageRes, err error) {
	// TODO: Your code here...
	return
}
