package usecases

import (
	"context"

	"github.com/thanapatfd/todolist/todo/entity"
)

type TodoUseCase struct {
	repo todoRepository
}

func NewTodoUseCase(repo todoRepository) TodoUseCase {
	return TodoUseCase{repo: repo}
}

func (uc TodoUseCase) GetLists(ctx context.Context, name string, status string) ([]entity.List, error) {
	ctx, sp := tracer.Start(ctx, "usecases.GetLists")
	defer sp.End()

	list, err := uc.repo.GetLists(name, status)
	if err != nil {
		sp.RecordError(err)
		return nil, err
	}

	sp.AddEvent("Get Lists Success")

	return list, nil
}

func (uc TodoUseCase) GetListByID(ctx context.Context, id string) (entity.List, error) {
	ctx, sp := tracer.Start(ctx, "usecases.GetListByID")
	defer sp.End()

	list, err := uc.repo.GetListByID(id)
	if err != nil {
		sp.RecordError(err)
		return list, err
	}

	sp.AddEvent("Get List By ID Success")

	return list, nil
}
func (uc TodoUseCase) CreateList(ctx context.Context, list entity.List) (entity.List, error) {
	ctx, sp := tracer.Start(ctx, "usecases.CreateList")
	defer sp.End()

	list, err := uc.repo.CreateList(list)
	if err != nil {
		return list, err
	}

	sp.AddEvent("Create Lists Success")

	return list, nil
}

func (uc TodoUseCase) UpdateList(ctx context.Context, list entity.List, id string) (entity.List, error) {
	ctx, sp := tracer.Start(ctx, "usecases.UpdateList")
	defer sp.End()

	err1 := list.ChangeStatus(list.Status)
	if err1 != nil {
		return list, err1
	}
	sp.AddEvent("Change Status Success")

	list, err := uc.repo.UpdateList(list, id)
	if err != nil {
		return list, err
	}

	sp.AddEvent("Update List Success")

	return list, nil
}

func (uc TodoUseCase) PatchList(ctx context.Context, list entity.List, id string) (entity.List, error) {
	ctx, sp := tracer.Start(ctx, "usecases.PatchList")
	defer sp.End()

	list, err := uc.repo.PatchList(list, id)
	if err != nil {
		return list, err
	}

	sp.AddEvent("Patch List Success")

	return list, nil

}

func (uc TodoUseCase) DeleteList(ctx context.Context, id string) error {
	ctx, sp := tracer.Start(ctx, "usecases.DeleteList")
	defer sp.End()

	err := uc.repo.DeleteList(id)
	if err != nil {
		return err
	}

	sp.AddEvent("Delete List Success")

	return nil
}

func (uc TodoUseCase) SortListsByID(ctx context.Context) ([]entity.List, error) {
	ctx, sp := tracer.Start(ctx, "usecases.SortListByID")
	defer sp.End()

	list, err := uc.repo.SortListsByID()
	if err != nil {
		return nil, err
	}
	sp.AddEvent("Sort Lists Success")

	return list, nil
}
