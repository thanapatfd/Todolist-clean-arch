package usecases

import (
	"context"
	"log/slog"

	"github.com/thanapatfd/todolist/todo/entity"
)

func (uc TodoUseCase) GetLists(ctx context.Context, name string, status string) ([]entity.List, error) {
	ctx, sp := tracer.Start(ctx, "usecases.GetLists")
	defer sp.End()

	list, err := uc.todoRepo.GetLists(ctx, name, status)
	if err != nil {
		sp.RecordError(err)
		slog.Error("query error")
		return nil, err
	}

	sp.AddEvent("Get Lists Success")

	return list, nil
}

func (uc TodoUseCase) GetListByID(ctx context.Context, id string) (entity.List, error) {
	ctx, sp := tracer.Start(ctx, "usecases.GetListByID")
	defer sp.End()

	list, err := uc.todoRepo.GetListByID(ctx, id)
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

	list, err := uc.todoRepo.CreateList(ctx, list)
	if err != nil {
		sp.RecordError(err)
		return list, err
	}

	sp.AddEvent("Create List Success")

	return list, nil
}

func (uc TodoUseCase) UpdateList(ctx context.Context, list entity.List, id string) (entity.List, error) {
	ctx, sp := tracer.Start(ctx, "usecases.UpdateList")
	defer sp.End()

	err1 := list.ChangeStatus(list.Status)
	if err1 != nil {
		sp.RecordError(err1)
		return list, err1
	}

	list, err := uc.todoRepo.UpdateList(ctx, list, id)
	if err != nil {
		sp.RecordError(err)
		return list, err
	}

	sp.AddEvent("Update List Success")

	return list, nil
}

func (uc TodoUseCase) PatchList(ctx context.Context, list entity.List, id string) (entity.List, error) {
	ctx, sp := tracer.Start(ctx, "usecases.PatchList")
	defer sp.End()

	list, err := uc.todoRepo.PatchList(ctx, list, id)
	if err != nil {
		sp.RecordError(err)
		return list, err
	}

	sp.AddEvent("Patch List Success")

	return list, nil

}

func (uc TodoUseCase) DeleteList(ctx context.Context, id string) error {
	ctx, sp := tracer.Start(ctx, "usecases.DeleteList")
	defer sp.End()

	err := uc.todoRepo.DeleteList(ctx, id)
	if err != nil {
		sp.RecordError(err)
		return err
	}

	sp.AddEvent("Delete List Success")

	return nil
}

func (uc TodoUseCase) SortListsByID(ctx context.Context) ([]entity.List, error) {
	ctx, sp := tracer.Start(ctx, "usecases.SortListByID")
	defer sp.End()

	list, err := uc.todoRepo.SortListsByID(ctx)
	if err != nil {
		sp.RecordError(err)
		return nil, err
	}
	sp.AddEvent("Sort Lists Success")

	return list, nil
}
