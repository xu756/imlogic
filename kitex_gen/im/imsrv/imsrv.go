// Code generated by Kitex v0.10.1. DO NOT EDIT.

package imsrv

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
	"MetaMessage": kitex.NewMethodInfo(
		metaMessageHandler,
		newMetaMessageArgs,
		newMetaMessageResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"PushMessage": kitex.NewMethodInfo(
		pushMessageHandler,
		newPushMessageArgs,
		newPushMessageResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
}

var (
	imSrvServiceInfo                = NewServiceInfo()
	imSrvServiceInfoForClient       = NewServiceInfoForClient()
	imSrvServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return imSrvServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return imSrvServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return imSrvServiceInfoForClient
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
	serviceName := "ImSrv"
	handlerType := (*im.ImSrv)(nil)
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

func metaMessageHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(im.MetaMsg)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(im.ImSrv).MetaMessage(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *MetaMessageArgs:
		success, err := handler.(im.ImSrv).MetaMessage(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*MetaMessageResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newMetaMessageArgs() interface{} {
	return &MetaMessageArgs{}
}

func newMetaMessageResult() interface{} {
	return &MetaMessageResult{}
}

type MetaMessageArgs struct {
	Req *im.MetaMsg
}

func (p *MetaMessageArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(im.MetaMsg)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *MetaMessageArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *MetaMessageArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *MetaMessageArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *MetaMessageArgs) Unmarshal(in []byte) error {
	msg := new(im.MetaMsg)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var MetaMessageArgs_Req_DEFAULT *im.MetaMsg

func (p *MetaMessageArgs) GetReq() *im.MetaMsg {
	if !p.IsSetReq() {
		return MetaMessageArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MetaMessageArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *MetaMessageArgs) GetFirstArgument() interface{} {
	return p.Req
}

type MetaMessageResult struct {
	Success *im.MessageRes
}

var MetaMessageResult_Success_DEFAULT *im.MessageRes

func (p *MetaMessageResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(im.MessageRes)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *MetaMessageResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *MetaMessageResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *MetaMessageResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *MetaMessageResult) Unmarshal(in []byte) error {
	msg := new(im.MessageRes)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MetaMessageResult) GetSuccess() *im.MessageRes {
	if !p.IsSetSuccess() {
		return MetaMessageResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MetaMessageResult) SetSuccess(x interface{}) {
	p.Success = x.(*im.MessageRes)
}

func (p *MetaMessageResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MetaMessageResult) GetResult() interface{} {
	return p.Success
}

func pushMessageHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(im.Message)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(im.ImSrv).PushMessage(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *PushMessageArgs:
		success, err := handler.(im.ImSrv).PushMessage(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*PushMessageResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newPushMessageArgs() interface{} {
	return &PushMessageArgs{}
}

func newPushMessageResult() interface{} {
	return &PushMessageResult{}
}

type PushMessageArgs struct {
	Req *im.Message
}

func (p *PushMessageArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(im.Message)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *PushMessageArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *PushMessageArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *PushMessageArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *PushMessageArgs) Unmarshal(in []byte) error {
	msg := new(im.Message)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var PushMessageArgs_Req_DEFAULT *im.Message

func (p *PushMessageArgs) GetReq() *im.Message {
	if !p.IsSetReq() {
		return PushMessageArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *PushMessageArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *PushMessageArgs) GetFirstArgument() interface{} {
	return p.Req
}

type PushMessageResult struct {
	Success *im.MessageRes
}

var PushMessageResult_Success_DEFAULT *im.MessageRes

func (p *PushMessageResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(im.MessageRes)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *PushMessageResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *PushMessageResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *PushMessageResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *PushMessageResult) Unmarshal(in []byte) error {
	msg := new(im.MessageRes)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *PushMessageResult) GetSuccess() *im.MessageRes {
	if !p.IsSetSuccess() {
		return PushMessageResult_Success_DEFAULT
	}
	return p.Success
}

func (p *PushMessageResult) SetSuccess(x interface{}) {
	p.Success = x.(*im.MessageRes)
}

func (p *PushMessageResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *PushMessageResult) GetResult() interface{} {
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

func (p *kClient) MetaMessage(ctx context.Context, Req *im.MetaMsg) (r *im.MessageRes, err error) {
	var _args MetaMessageArgs
	_args.Req = Req
	var _result MetaMessageResult
	if err = p.c.Call(ctx, "MetaMessage", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) PushMessage(ctx context.Context, Req *im.Message) (r *im.MessageRes, err error) {
	var _args PushMessageArgs
	_args.Req = Req
	var _result PushMessageResult
	if err = p.c.Call(ctx, "PushMessage", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
