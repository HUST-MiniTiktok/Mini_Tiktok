package tracer

import (
	"os"

	"github.com/HUST-MiniTiktok/mini_tiktok/pkg/conf"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
)

func init() {
	os.Setenv("JAEGER_AGENT_HOST", conf.GetConf().GetString("tracer.jaeger.host"))
	os.Setenv("JAEGER_AGENT_PORT", conf.GetConf().GetString("tracer.jaeger.port"))
	os.Setenv("JAEGER_DISABLED", "false")
	os.Setenv("JAEGER_SAMPLER_TYPE", "const")
	os.Setenv("JAEGER_SAMPLER_PARAM", "1")
	os.Setenv("JAEGER_REPORTER_LOG_SPANS", "true")
}

func InitJaeger(service string) {
	cfg, _ := jaegercfg.FromEnv()
	cfg.ServiceName = service
	tracer, _, err := cfg.NewTracer(jaegercfg.Logger(jaeger.StdLogger))
	if err != nil {
		klog.Fatalf("Could not initialize jaeger tracer: %s", err.Error())
	}
	opentracing.SetGlobalTracer(tracer)
}
