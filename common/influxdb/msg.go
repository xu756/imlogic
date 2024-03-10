package influxdb

import (
	"context"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/xu756/imlogic/kitex_gen/im"
	"time"
)

type MsgApi interface {
	CreateMsg(ctx context.Context, msg *im.Message) (err error)
	//QueryMsg(ctx context.Context, from, to string) (res *im.Message, err error)
}

func (m *customClient) CreateMsg(ctx context.Context, msg *im.Message) (err error) {
	p := influxdb2.NewPoint("msg",
		map[string]string{"from": msg.From, "to": msg.To, "link_id": msg.From, "msg_type": msg.MsgType},
		map[string]interface{}{"msg_id": msg.MsgId, "device": msg.Device, "msg_content": msg.MsgContent, "create_time": time.Now().Unix()},
		time.Now())
	return m.writeAPI.WritePoint(ctx, p)
}

//func (m *customClient) QueryMsg(ctx context.Context, from, to string) (res *im.Message, err error) {
//	query := `from(bucket: "msg")
//  |> range(start: -1h)
//  |> filter(fn: (r) => r._measurement == "msg" and r._field == "from" and r._value == "from" and r._field == "to" and r._value == "to")`
//
//	return nil, nil
//}
