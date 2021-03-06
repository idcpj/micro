package common

import (
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"io"
	"time"
)

func NewTracer(serverName string,addr string)(opentracing.Tracer,io.Closer,error){
	conf := &config.Configuration{
		ServiceName:         serverName,
		Sampler: &config.SamplerConfig{
			Type:                     jaeger.SamplerTypeConst,
			Param:                    1,
		},
		Reporter:            &config.ReporterConfig{
			BufferFlushInterval:        1*time.Second,
			LogSpans:                   true,
			LocalAgentHostPort:         addr,
		},
	}
	tracer, closer, err := conf.NewTracer()
	return tracer,closer,err

}