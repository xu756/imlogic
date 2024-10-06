package trace

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/sirupsen/logrus"
	"imlogic/common/config"
	"imlogic/common/xlog"
)

type KitexTrace struct {
	provider provider.OtelProvider
}

func KitexTraceSetUp(serviceName string) *KitexTrace {
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
	logger := kitexlogrus.NewLogger(
		// 打印错误日志
		kitexlogrus.WithTraceHookErrorSpanLevel(logrus.ErrorLevel),
		// 打印所有日志
		kitexlogrus.WithTraceHookLevels(logrus.AllLevels),
		// 打印日志的字段
		kitexlogrus.WithRecordStackTraceInSpan(true),
	)
	klog.SetLogger(logger)
	klog.SetLevel(klog.LevelTrace)
	klog.SetOutput(xlog.LogOutFile(serviceName))
	return &KitexTrace{provider: p}
}

// defer p.Shutdown(context.Background())
func (t *KitexTrace) Shutdown() {
	t.provider.Shutdown(context.Background())
}
