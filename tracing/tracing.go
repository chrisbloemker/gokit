package tracing

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/sdk/trace"
)

func Init(serviceName string) func(context.Context) error {
	tp := trace.NewTracerProvider()
	otel.SetTracerProvider(tp)
	return tp.Shutdown
}
