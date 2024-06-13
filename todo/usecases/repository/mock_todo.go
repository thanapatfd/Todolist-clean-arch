package repository

import (
	"context"

	"github.com/stretchr/testify/mock"
	"github.com/thanapatfd/todolist/todo/entity"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) GetLists(ctx context.Context, name string, status string) ([]entity.List, error) {
	args := m.Called(ctx, name, status)
	return args.Get(0).([]entity.List), args.Error(1)
}

func (m *MockRepository) GetListByID(ctx context.Context, id string) (entity.List, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(entity.List), args.Error(1)

}

func (m *MockRepository) CreateList(ctx context.Context, list entity.List) (entity.List, error) {
	args := m.Called(ctx, list)
	return args.Get(0).(entity.List), args.Error(1)
}

func (m *MockRepository) UpdateList(ctx context.Context, list entity.List, id string) (entity.List, error) {
	args := m.Called(ctx, list, id)
	return args.Get(0).(entity.List), args.Error(1)
}

func (m *MockRepository) PatchList(ctx context.Context, list entity.List, id string) (entity.List, error) {
	args := m.Called(ctx, list, id)
	return args.Get(0).(entity.List), args.Error(1)
}

func (m *MockRepository) DeleteList(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockRepository) SortListsByID(ctx context.Context) ([]entity.List, error) {
	args := m.Called(ctx)
	return args.Get(0).([]entity.List), args.Error(1)
}
func (m *MockRepository) ChangeStatus(ctx context.Context, list entity.List, id string) (entity.List, error) {
	args := m.Called(ctx, list, id)
	return args.Get(0).(entity.List), args.Error(1)
}
