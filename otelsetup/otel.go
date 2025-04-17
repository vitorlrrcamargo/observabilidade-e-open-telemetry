package otelsetup

import (
	"context"
	"log"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/zipkin"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

func InitTracer(zipkinURL string) func(context.Context) error {
	exporter, err := zipkin.New(zipkinURL)
	if err != nil {
		log.Fatal(err)
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
	)

	otel.SetTracerProvider(tp)

	return tp.Shutdown
}
