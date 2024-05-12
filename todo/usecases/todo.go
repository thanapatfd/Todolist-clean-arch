package usecases

import (
	"github.com/thanapatfd/todolist/todo/entity"
	"go.opentelemetry.io/otel"
)

var tracer = otel.Tracer("usecases")

type todoRepository interface {
	CreateList(entity.List) (entity.List, error)
	GetListByID(id string) (entity.List, error)
	GetLists(name string, status string) ([]entity.List, error)
	UpdateList(list entity.List, id string) (entity.List, error)
	PatchList(list entity.List, id string) (entity.List, error)
	DeleteList(id string) error
	SortListsByID() ([]entity.List, error)
}
