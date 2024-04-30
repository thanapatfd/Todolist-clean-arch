package usecases

import "github.com/thanapatfd/todolist/todo/entity"

type TodoUseCase struct {
	repo todoRepository
}

func NewTodoUseCase(repo todoRepository) TodoUseCase {
	return TodoUseCase{repo: repo}
}

func (usecase TodoUseCase) GetLists() ([]entity.List, error) {
	list, err := usecase.repo.GetLists()
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (usecase TodoUseCase) GetListByID(id string) (entity.List, error) {
	list, err := usecase.repo.GetListByID(id)
	if err != nil {
		return list, err
	}
	return list, nil
}
func (usecase TodoUseCase) CreateList(list entity.List) (entity.List, error) {
	list, err := usecase.repo.CreateList(list)
	if err != nil {
		return list, err
	}
	return list, nil
}

func (usecase TodoUseCase) UpdateList(list entity.List, id string) (entity.List, error) {
	err1 := list.ChangeStatus(list.Status)
	if err1 != nil {
		return list, err1
	}

	list, err := usecase.repo.UpdateList(list, id)
	if err != nil {
		return list, err
	}
	return list, nil
}

func (usecase TodoUseCase) DeleteList(id string) error {
	err := usecase.repo.DeleteList(id)
	if err != nil {
		return err
	}
	return nil
}

func (usecase TodoUseCase) SortListsByID() ([]entity.List, error) {
	list, err := usecase.repo.SortListsByID()
	if err != nil {
		return nil, err
	}
	return list, nil
}
