package influxdb

import (
	"context"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/xu756/imlogic/kitex_gen/im"
	"time"
)

type MsgApi interface {
	CreateMsg(ctx context.Context, msg *im.Message) (err error)
}

func (m *customClient) CreateMsg(ctx context.Context, msg *im.Message) (err error) {
	p := influxdb2.NewPoint("msg",
		map[string]string{"user_id": msg.UserId, "link_id": msg.From, "device": msg.Device},
		map[string]interface{}{"msg_id": msg.MsgId, "msg_type": msg.MsgType, "msg_content": msg.MsgContent, "create_time": time.Now().Unix()},
		time.Now())

	return m.writeAPI.WritePoint(ctx, p)
}
