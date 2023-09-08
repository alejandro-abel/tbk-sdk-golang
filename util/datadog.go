package util

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
	"os"
)

func GetTraceContextFields(ctx context.Context) map[string]interface{} {
	span, _ := tracer.SpanFromContext(ctx)
	return log.Fields{
		"dd.trace_id": fmt.Sprintf("%d", span.Context().TraceID()),
		"dd.span_id":  fmt.Sprintf("%d", span.Context().SpanID()),
		"dd.service":  os.Getenv("DD_SERVICE"),
		"dd.env":      os.Getenv("DD_ENV"),
		"dd.version":  os.Getenv("DD_VERSION"),
	}
}
