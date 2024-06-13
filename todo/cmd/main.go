package main

import (
	"github.com/gofiber/fiber/v2"
	handlers "github.com/thanapatfd/todolist/todo/handler"
	"github.com/thanapatfd/todolist/todo/middleware"
	"github.com/thanapatfd/todolist/todo/repository"
	"github.com/thanapatfd/todolist/todo/server"
	"github.com/thanapatfd/todolist/todo/usecases"
)

type config struct {
	AppName     string `env:"APP_NAME" envDefault:"TodoList-By-Ford"`
	AppVersion  string `env:"APP_VERSION" envDefault:"v0.0.0"`
	Environment string `env:"ENVIRONMENT" envDefault:"development"`
	Port        uint   `env:"PORT" envDefault:"5050"`
	Debuglog    bool   `env:"DEBUG_LOG" envDefault:"true"`

	Services struct {
		OtelGrpcEndpoint      string `env:"OTEL_GRPC_ENDPOINT" envDefault:"localhost:4317"`
		AddressServiceBaseUrl string `env:"SERVICE_ADDRESS_BASE_URL" envDefault:"http://localhost:8080"`
	}
}

func main() {

	app := fiber.New()
	db := server.NewPosgrestDB()
	todoRepo := repository.NewTodoRepository(db.Db)
	todoUsecase := usecases.NewUsecase(todoRepo)
	todoHandler := handlers.NewTodoHandler(todoUsecase)

	cfg := initEnvironment()
	initTracer(cfg)
	initLogger(cfg)

	app.Use(middleware.Logging)

	app.Get("/lists", todoHandler.GetLists)
	app.Get("/lists/:id", todoHandler.GetListByID)
	app.Get("/lists/sort/id", todoHandler.SortListsByID)
	app.Post("/lists", todoHandler.CreateList)
	app.Put("/lists/:id", todoHandler.UpdateList)
	app.Patch("/lists/:id", todoHandler.PatchList)
	app.Delete("/lists/:id", todoHandler.DeleteList)
	app.Put("/lists/changestatus/:id", todoHandler.ChangeStatus)

	app.Listen(":5050")
}
