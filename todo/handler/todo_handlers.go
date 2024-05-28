package handlers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/thanapatfd/todolist/todo/entity"
	"github.com/thanapatfd/todolist/todo/usecases"
)

type ListPayload struct {
	ID      int    `json:"id"`
	Name    string `json:"name" validate:"required"`
	Status  string `json:"status" validate:"required"`
	Details string `json:"details" validate:"required"`
}

type ListResponse struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Status  string `json:"status"`
	Details string `json:"details"`
}

type TodoHandler interface {
	CreateList(c *fiber.Ctx) error
	GetListByID(c *fiber.Ctx) error
	GetLists(c *fiber.Ctx) error
	UpdateList(c *fiber.Ctx) error
	PatchList(c *fiber.Ctx) error
	DeleteList(c *fiber.Ctx) error
	SortListsByID(c *fiber.Ctx) error
	Validation(payload ListPayload) (ListPayload, error)
}

type todoHandler struct {
	usecase usecases.TodoUseCase
}

func NewTodoHandler(usecase usecases.TodoUseCase) TodoHandler {
	return todoHandler{usecase: usecase}
}

func (h todoHandler) GetLists(c *fiber.Ctx) error {
	ctx, sp := tracer.Start(c.Context(), "handlers.GetLists")
	defer sp.End()
	res := []ListResponse{}
	name := c.Query("name")
	status := c.Query("status")
	lists, err := h.usecase.GetLists(ctx, name, status)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	for _, rows := range lists {
		res = append(res, ListResponse{
			ID:      rows.ID,
			Name:    rows.Name,
			Status:  rows.Status,
			Details: rows.Details,
		})
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

func (h todoHandler) GetListByID(c *fiber.Ctx) error {
	ctx, sp := tracer.Start(c.Context(), "handlers.GetListByID")
	defer sp.End()
	id := c.Params("id")
	list, err := h.usecase.GetListByID(ctx, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	res := ListResponse{
		ID:      list.ID,
		Name:    list.Name,
		Status:  list.Status,
		Details: list.Details,
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

func (h todoHandler) CreateList(c *fiber.Ctx) error {
	ctx, sp := tracer.Start(c.Context(), "handlers.CreateList")
	defer sp.End()
	payload := new(ListPayload)
	if err := c.BodyParser(payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	checkValid, err := h.Validation(*payload)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	list := entity.List{
		ID:      checkValid.ID,
		Name:    checkValid.Name,
		Status:  checkValid.Status,
		Details: checkValid.Details,
	}

	result, err := h.usecase.CreateList(ctx, list)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	res := ListResponse{
		ID:      result.ID,
		Name:    result.Name,
		Status:  result.Status,
		Details: result.Details,
	}
	return c.Status(fiber.StatusOK).JSON(res)
}

func (h todoHandler) UpdateList(c *fiber.Ctx) error {
	ctx, sp := tracer.Start(c.Context(), "handlers.UpdateList")
	defer sp.End()
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "missing id"})
	}
	payload := new(ListPayload)
	if err := c.BodyParser(payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	checkValid, err := h.Validation(*payload)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	list := entity.List{
		ID:      checkValid.ID,
		Name:    checkValid.Name,
		Status:  checkValid.Status,
		Details: checkValid.Details,
	}

	updateList, err := h.usecase.UpdateList(ctx, list, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	res := ListResponse{
		ID:      updateList.ID,
		Name:    updateList.Name,
		Status:  updateList.Status,
		Details: updateList.Details,
	}
	return c.Status(fiber.StatusOK).JSON(res)

}

func (h todoHandler) PatchList(c *fiber.Ctx) error {
	ctx, sp := tracer.Start(c.Context(), "handlers.PatchList")
	defer sp.End()
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "missing id"})
	}

	var payload ListPayload
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}

	listToUpdate := entity.List{
		ID:      payload.ID,
		Name:    payload.Name,
		Status:  payload.Status,
		Details: payload.Details,
	}

	updatedList, err := h.usecase.PatchList(ctx, listToUpdate, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	response := ListResponse{
		ID:      updatedList.ID,
		Name:    updatedList.Name,
		Status:  updatedList.Status,
		Details: updatedList.Details,
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

func (h todoHandler) DeleteList(c *fiber.Ctx) error {
	ctx, sp := tracer.Start(c.Context(), "handlers.DeleteList")
	defer sp.End()
	id := c.Params("id")
	err := h.usecase.DeleteList(ctx, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "List deleted successfully"})
}

func (h todoHandler) SortListsByID(c *fiber.Ctx) error {
	ctx, sp := tracer.Start(c.Context(), "handlers.SortListByID")
	defer sp.End()
	res := []ListResponse{}
	lists, err := h.usecase.SortListsByID(ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	for _, rows := range lists {
		res = append(res, ListResponse{
			ID:      rows.ID,
			Name:    rows.Name,
			Status:  rows.Status,
			Details: rows.Details,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"lists": res})
}

func (h todoHandler) Validation(payload ListPayload) (ListPayload, error) {
	if payload.Name == "" || payload.Details == "" || payload.Status == "" {
		return payload, errors.New("missing required fields: all fields must be non-empty")
	}

	return payload, nil
}
