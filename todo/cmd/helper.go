package main

import (
	"context"
	"log"
	"log/slog"
	"os"

	"github.com/caarlos0/env/v10"
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
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime)) // ลบ timestamp จาก log

	// โหลด environment variables จาก env
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Printf("Failed loading .env file: %s", err)
	// }

	var cfg config
	// parse environment variables ให้เป็น struct config
	err := env.Parse(&cfg)
	if err != nil {
		log.Fatalf("Error parse env to struct: %s", err)
	}

	return cfg
}

func initLogger(cfg config) {
	// กำหนดระดับการ log เริ่มต้นเป็น Info

	logLevel := slog.LevelInfo
	if cfg.Debuglog {
		// ถ้าเปิดโหมด debug จะตั้งระดับการ log เป็น Debug
		logLevel = slog.LevelDebug
	}

	// สร้างตัวจัดการ log แบบ JSON และตั้งค่า options
	logger := slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		Level: logLevel,
	}))

	// ตั้งค่า logger เริ่มต้น
	slog.SetDefault(logger)
}

func initTracer(cfg config) {
	// สร้าง OTLP gRPC client
	client := otlptracegrpc.NewClient(
		otlptracegrpc.WithInsecure(),
		otlptracegrpc.WithEndpoint(cfg.Services.OtelGrpcEndpoint),
	)

	// สร้าง OTLP trace exporter
	exporter, err := otlptrace.New(context.Background(), client)
	if err != nil {
		slog.Error("Error initializing OTLP exporter: %v", err)
	}

	// สร้าง TracerProvider และกำหนดการตั้งค่า เป็น Batch Expoter และกำหรด Resource ต่างๆ
	tp := trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(cfg.AppName),
			semconv.ServiceVersionKey.String(cfg.AppVersion),
			attribute.String("environment", cfg.Environment),
		)),
	)
	// กำหนด TracerProvider ที่จะใช้ด้วย OpenTelemetry
	otel.SetTracerProvider(tp)

	// กำหนด propagator สำหรับ context propagation
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
}
