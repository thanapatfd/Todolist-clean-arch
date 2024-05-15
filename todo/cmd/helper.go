package main

import (
	"context"
	"log"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

func initEnvironment() config {
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))

	err := godotenv.Load()
	if err != nil {
		log.Printf("Failed loading .env file: %s", err)
	}

	var cfg config
	err = env.Parse(&cfg)
	if err != nil {
		log.Fatalf("Error parse env to struct: %s", err)
	}

	return cfg
}

func initTracer(cfg config) {
	client := otlptracegrpc.NewClient(
		otlptracegrpc.WithInsecure(),
		otlptracegrpc.WithEndpoint(cfg.Services.OtelGrpcEndpoint),
	)

	exporter, err := otlptrace.New(context.Background(), client)
	if err != nil {
		log.Fatalf("Error initializing OTLP exporter: %v", err)
	}

	tp := trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(cfg.AppName),
			semconv.ServiceVersionKey.String(cfg.AppVersion),
			attribute.String("environment", cfg.Environment),
		)),
	)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
}
