package main

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	handlers "github.com/thanapatfd/todolist/todo/handler"
	"github.com/thanapatfd/todolist/todo/repository"
	"github.com/thanapatfd/todolist/todo/server"
	"github.com/thanapatfd/todolist/todo/usecases"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

func main() {

	app := fiber.New()
	db := server.NewPosgrestDB()
	todoRepo := repository.NewTodoRepository(db.Db)
	todoUsecase := usecases.NewTodoUseCase(todoRepo)

	InitTracerWithOutput()

	todoHandler := handlers.NewTodoHandler(todoUsecase)

	// app.Use(middleware.Logging)

	app.Get("/lists", todoHandler.GetLists)
	app.Get("/lists/:id", todoHandler.GetListByID)
	app.Get("/lists/sort/id", todoHandler.SortListsByID)
	app.Post("/lists", todoHandler.CreateList)
	app.Put("/lists/:id", todoHandler.UpdateList)
	app.Patch("/lists/:id", todoHandler.PatchList)
	app.Delete("/lists/:id", todoHandler.DeleteList)

	app.Listen(":3000")
}

func newExporter(ctx context.Context) (trace.SpanExporter, error) {
	return otlptracehttp.New(ctx)
}
func InitTracerWithOutput() {
	traceExporter, err := stdouttrace.New(
		stdouttrace.WithPrettyPrint())
	if err != nil {
		log.Fatal(err)
	}
	tp := trace.NewTracerProvider(
		trace.WithBatcher(traceExporter),
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			attribute.String("environment", "Development"),
		)),
	)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
}
