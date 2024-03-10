package influxdb

import (
	"context"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/xu756/imlogic/kitex_gen/im"
	"log"
	"time"
)

type MsgApi interface {
	CreateMsg(ctx context.Context, msg *im.Message) (err error)
	QueryMsg(ctx context.Context, from, to string) (res *[]im.Message, err error)
}

func (m *customClient) CreateMsg(ctx context.Context, msg *im.Message) (err error) {
	p := influxdb2.NewPoint("msg",
		map[string]string{"from": msg.From, "to": msg.To, "link_id": msg.LinkId, "msg_type": msg.MsgType},
		map[string]interface{}{"msg_id": msg.MsgId, "device": msg.Device, "msg_content": msg.MsgContent, "create_time": time.Now().Unix()},
		time.Now())
	return m.writeAPI.WritePoint(ctx, p)
}

func (m *customClient) QueryMsg(ctx context.Context, from, to string) (res *[]im.Message, err error) {
	query := `from(bucket: "imlogic")
	  |> range(start: -1d)
	  |> filter(fn: (r) => r._measurement == "msg" and r.from == "` + from + `" and r.to == "` + to + `")`
	result, err := m.queryAPI.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer result.Close()
	//var msgs []im.Message
	for result.Next() {
		log.Print(result.Record().Value())

	}
	return nil, nil
}
