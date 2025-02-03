package tracing

import (
	"context"
	"log"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
	"go.opentelemetry.io/otel/trace"
)

var Tracer trace.Tracer

func InitTracer() {

	// Create a background context
	ctx := context.Background()

	exp, err := otlptracehttp.New(
		ctx,
		otlptracehttp.WithEndpoint("10.103.166.182:4318"), // Jaeger collector endpoint
		otlptracehttp.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("Failed to initialize OTLP exporter: %v", err)
	}
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exp),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("SessionManagerService"),
		)),
	)
	otel.SetTracerProvider(tp)
	Tracer = tp.Tracer("SessionManager")
	log.Println("Tracer initialized")
}
// type TracerManager struct{}
// // InitTracer initializes the OpenTelemetry tracer
// func (t *TracerManager) Start() {
// 	log.Println("Starting Tracer initialization...")
// 	ctx := context.Background()

// 	exp, err := otlptracehttp.New(
// 		ctx,
// 		otlptracehttp.WithEndpoint("10.105.78.203:14250"), // Jaeger collector endpoint
// 		otlptracehttp.WithInsecure(),
// 	)
// 	if err != nil {
// 		log.Fatalf("Failed to initialize OTLP exporter: %v", err)
// 	}
// 	tp := sdktrace.NewTracerProvider(
// 		sdktrace.WithBatcher(exp),
// 		sdktrace.WithResource(resource.NewWithAttributes(
// 			semconv.SchemaURL,
// 			semconv.ServiceNameKey.String("SessionManagerService"),
// 		)),
// 	)
// 	otel.SetTracerProvider(tp)
// 	Tracer = tp.Tracer("SessionManager")
// 	log.Println("Tracer initialized successfully")
// }
