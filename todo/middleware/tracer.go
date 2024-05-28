package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
	"go.opentelemetry.io/otel/trace"
)

// Tracer เป็นตัวแทรก OpenTelemetry tracer
var Tracer = otel.GetTracerProvider().Tracer("fiber-server")

func TracerMiddleware(c *fiber.Ctx) error {
	// สร้าง propagator แบบ composite เพื่อจัดการการข้อมูล context และ baggage
	propagator := propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{})
	carrier := propagation.HeaderCarrier{}

	// แยกข้อมูล header เข้ามาและฝังลงใน context ปัจจุบัน
	c.Request().Header.VisitAll(func(key, value []byte) {
		carrier.Set(string(key), string(value))
	})

	propagator.Inject(c.Context(), carrier)

	// กำหนด options สำหรับ span
	spanOptions := []trace.SpanStartOption{
		trace.WithAttributes(semconv.HTTPMethodKey.String(c.Method())),
		trace.WithAttributes(semconv.HTTPTargetKey.String(string(c.Request().RequestURI()))),
		trace.WithAttributes(semconv.HTTPRouteKey.String(c.Route().Path)),
		trace.WithAttributes(semconv.HTTPURLKey.String(c.OriginalURL())),
		trace.WithAttributes(semconv.UserAgentOriginal(string(c.Request().Header.UserAgent()))),
		trace.WithAttributes(semconv.HTTPRequestContentLengthKey.Int(c.Request().Header.ContentLength())),
		trace.WithAttributes(semconv.HTTPSchemeKey.String(c.Protocol())),
		trace.WithAttributes(semconv.NetTransportTCP),
		trace.WithSpanKind(trace.SpanKindServer),
	}

	// เริ่ม span ใหม่ด้วย options และ context ที่กำหนด
	ctx, span := Tracer.Start(c.Context(), fmt.Sprintf("%s %s", c.Method(), c.Path()), spanOptions...)
	defer span.End()

	{
		// ฝังข้อมูลของ span context ลงใน header ของ response ที่จะส่งออก
		propagator := propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{})
		carrier := propagation.HeaderCarrier{}
		propagator.Inject(ctx, carrier)

		for _, k := range carrier.Keys() {
			c.Response().Header.Set(k, carrier.Get(k))
		}
	}

	// ดำเนินการคำขอ
	err := c.Next()

	// ตั้งค่า attributes เกี่ยวกับ response บน span
	span.SetAttributes(semconv.HTTPStatusCodeKey.Int(c.Response().StatusCode()))

	return err
}
