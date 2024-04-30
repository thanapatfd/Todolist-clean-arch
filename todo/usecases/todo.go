package usecases

import "github.com/thanapatfd/todolist/todo/entity"

type todoRepository interface {
	CreateList(entity.List) (entity.List, error)
	GetListByID(id string) (entity.List, error)
	GetLists() ([]entity.List, error)
	UpdateList(list entity.List, id string) (entity.List, error)
	DeleteList(id string) error
	SortListsByID() ([]entity.List, error)
}
