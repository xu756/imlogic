package trace

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	hertzlogrus "github.com/hertz-contrib/obs-opentelemetry/logging/logrus"
	"github.com/hertz-contrib/obs-opentelemetry/provider"
	"github.com/sirupsen/logrus"
	"imlogic/common/config"
	"imlogic/common/xlog"
)

type HertzTrace struct {
	provider provider.OtelProvider
}

func HertzTraceSetUp(serviceName string) *HertzTrace {
	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceNamespace("default"),
		provider.WithServiceName(serviceName),
		provider.WithExportEndpoint(config.RunData().OpenTelemetry.Endpoint),
		provider.WithEnableTracing(true),
		provider.WithEnableMetrics(true),
		provider.WithHeaders(map[string]string{
			"Authorization": "Basic " + base64.StdEncoding.EncodeToString([]byte(
				fmt.Sprintf("%s:%s", config.RunData().OpenTelemetry.UserName, config.RunData().OpenTelemetry.PassWord),
			)),
			"stream-name":  config.RunData().OpenTelemetry.Straem,
			"organization": "default",
		}),
		provider.WithInsecure(),
	)
	logger := hertzlogrus.NewLogger(
		// 打印错误日志
		hertzlogrus.WithTraceHookErrorSpanLevel(logrus.ErrorLevel),
		// 打印所有日志
		hertzlogrus.WithTraceHookLevels(logrus.AllLevels),
		// 打印日志的字段
		hertzlogrus.WithRecordStackTraceInSpan(true),
	)
	hlog.SetLogger(logger)
	hlog.SetLevel(hlog.LevelTrace)
	hlog.SetOutput(xlog.LogOutFile(serviceName))
	return &HertzTrace{provider: p}
}

//defer p.Shutdown(context.Background())

func (t *HertzTrace) Shutdown() {
	t.provider.Shutdown(context.Background())
}
