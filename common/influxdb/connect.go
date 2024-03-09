package influxdb

import (
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/xu756/imlogic/common/config"
)

type customClient struct {
	writeAPI api.WriteAPIBlocking
	queryAPI api.QueryAPI
}

func NewClient() Client {
	var InfluxdbConfig = config.RunData.InfluxdbConfig
	client := influxdb2.NewClient(InfluxdbConfig.Addr, InfluxdbConfig.Token)
	writeAPI := client.WriteAPIBlocking(InfluxdbConfig.Org, InfluxdbConfig.Bucket)
	queryAPI := client.QueryAPI(InfluxdbConfig.Org)
	return &customClient{
		writeAPI: writeAPI,
		queryAPI: queryAPI,
	}
}
