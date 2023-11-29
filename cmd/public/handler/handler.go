package handler

import (
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/xu756/imlogic/common/db"
	"github.com/xu756/imlogic/internal/xjwt"
)

type PublicSrvImpl struct {
	WriteAPI api.WriteAPIBlocking
	ReadAPI  api.QueryAPI
	Model    db.Model
	Jwt      *xjwt.Jwt
}

func NewPublicSrvImpl() *PublicSrvImpl {
	token := "eXXq1MXbFbZ0KvEEpV9cFRMAlhjtMoAYLONheghCgXIOvUtqUaGTOb2SKRP3MMTPvon9jCggr4B6YI4AuH0qjw=="
	url := "https://influxdb.imlogic.cn"
	org := "influxdb-org"
	bucket := "public"
	client := influxdb2.NewClient(url, token)
	return &PublicSrvImpl{
		WriteAPI: client.WriteAPIBlocking(org, bucket),
		ReadAPI:  client.QueryAPI(org),
		Model:    db.NewModel(),
		Jwt:      xjwt.NewJwt(),
	}

}
