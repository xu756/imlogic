// Code generated by Kitex v0.10.1. DO NOT EDIT.

package imserver

import (
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
	im "imlogic/kitex_gen/im"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"SendMsg": kitex.NewMethodInfo(
		sendMsgHandler,
		newSendMsgArgs,
		newSendMsgResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
}

var (
	imServerServiceInfo                = NewServiceInfo()
	imServerServiceInfoForClient       = NewServiceInfoForClient()
	imServerServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return imServerServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return imServerServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return imServerServiceInfoForClient
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
	serviceName := "ImServer"
	handlerType := (*im.ImServer)(nil)
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
		"PackageName": "im",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.10.1",
		Extra:           extra,
	}
	return svcInfo
}

func sendMsgHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(im.Message)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(im.ImServer).SendMsg(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *SendMsgArgs:
		success, err := handler.(im.ImServer).SendMsg(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*SendMsgResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newSendMsgArgs() interface{} {
	return &SendMsgArgs{}
}

func newSendMsgResult() interface{} {
	return &SendMsgResult{}
}

type SendMsgArgs struct {
	Req *im.Message
}

func (p *SendMsgArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(im.Message)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *SendMsgArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *SendMsgArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *SendMsgArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *SendMsgArgs) Unmarshal(in []byte) error {
	msg := new(im.Message)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var SendMsgArgs_Req_DEFAULT *im.Message

func (p *SendMsgArgs) GetReq() *im.Message {
	if !p.IsSetReq() {
		return SendMsgArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *SendMsgArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *SendMsgArgs) GetFirstArgument() interface{} {
	return p.Req
}

type SendMsgResult struct {
	Success *im.MessageRes
}

var SendMsgResult_Success_DEFAULT *im.MessageRes

func (p *SendMsgResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(im.MessageRes)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *SendMsgResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *SendMsgResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *SendMsgResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *SendMsgResult) Unmarshal(in []byte) error {
	msg := new(im.MessageRes)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SendMsgResult) GetSuccess() *im.MessageRes {
	if !p.IsSetSuccess() {
		return SendMsgResult_Success_DEFAULT
	}
	return p.Success
}

func (p *SendMsgResult) SetSuccess(x interface{}) {
	p.Success = x.(*im.MessageRes)
}

func (p *SendMsgResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *SendMsgResult) GetResult() interface{} {
	return p.Success
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) SendMsg(ctx context.Context, Req *im.Message) (r *im.MessageRes, err error) {
	var _args SendMsgArgs
	_args.Req = Req
	var _result SendMsgResult
	if err = p.c.Call(ctx, "SendMsg", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
