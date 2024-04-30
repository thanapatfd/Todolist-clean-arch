package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thanapatfd/todolist/todo/database"
	handlers "github.com/thanapatfd/todolist/todo/handler"
	"github.com/thanapatfd/todolist/todo/repository"
	"github.com/thanapatfd/todolist/todo/usecases"
)

func main() {

	app := fiber.New()
	db := database.NewPosgrestDB()
	todoRepo := repository.NewTodoRepository(db.Db)
	todoUsecase := usecases.NewTodoUseCase(todoRepo)
	todoHandler := handlers.NewTodoHandler(todoUsecase)

	app.Get("/lists", todoHandler.GetLists)
	app.Get("/lists/:id", todoHandler.GetListByID)
	app.Get("/lists/sort/id", todoHandler.SortListsByID)
	app.Post("/lists", todoHandler.CreateList)
	app.Put("/lists/:id", todoHandler.UpdateList)
	app.Delete("/lists/:id", todoHandler.DeleteList)

	app.Listen(":3000")
}
