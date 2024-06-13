package usecases

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/thanapatfd/todolist/todo/entity"
)

func TestTodoUseCase_GetLists(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		ctx := context.Background()
		expectedLists := []entity.List{
			{
				ID:      1,
				Name:    "List 1",
				Status:  "Todo",
				Details: "Example Details",
			},

			{
				ID:      2,
				Name:    "List 2",
				Status:  "Doing",
				Details: "Example Details",
			},
		}

		mr, uc := newMock()
		mr.On("GetLists", mock.Anything, "TestName", "TestStatus").Return(expectedLists, nil)

		lists, err := uc.GetLists(ctx, "TestName", "TestStatus")

		assert.NoError(t, err)
		assert.Equal(t, expectedLists, lists)
		mr.AssertExpectations(t)
	})

	t.Run("Error", func(t *testing.T) {
		ctx := context.Background()
		expectedError := errors.New("cannot get lists ")
		expectedLists := []entity.List{}

		mr, uc := newMock()
		mr.On("GetLists", mock.Anything, "TestName", "TestStatus").Return(expectedLists, expectedError)

		lists, err := uc.GetLists(ctx, "TestName", "TestStatus")

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		assert.Nil(t, lists)
		mr.AssertExpectations(t)
	})

}

func TestTodoUseCase_GetListByID(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		ctx := context.Background()
		expectedList := entity.List{
			ID:      1,
			Name:    "List 1",
			Status:  "Todo",
			Details: "Example Details",
		}
		mr, uc := newMock()

		mr.On("GetListByID", mock.Anything, "1").Return(expectedList, nil)

		list, err := uc.GetListByID(ctx, "1")

		assert.NoError(t, err)
		assert.Equal(t, expectedList, list)
		mr.AssertExpectations(t)
	})

	t.Run("Error", func(t *testing.T) {
		ctx := context.Background()
		expectedError := errors.New("cannot get list")
		expectedLists := entity.List{}
		mr, uc := newMock()

		mr.On("GetListByID", mock.Anything, "1").Return(expectedLists, expectedError)

		list, err := uc.GetListByID(ctx, "1")

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		assert.Equal(t, expectedLists, list)
		mr.AssertExpectations(t)
	})
}

func TestTodoUseCase_CreateList(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		ctx := context.Background()
		newList := entity.List{
			ID:      1,
			Name:    "List 1",
			Status:  "Todo",
			Details: "Example Details",
		}
		createdList := entity.List{
			ID:      1,
			Name:    "List 1",
			Status:  "Todo",
			Details: "Example Details",
		}

		mr, uc := newMock()
		mr.On("CreateList", mock.Anything, newList).Return(createdList, nil)

		list, err := uc.CreateList(ctx, newList)

		assert.NoError(t, err)
		assert.Equal(t, createdList, list)
		mr.AssertExpectations(t)
	})

	t.Run("Error", func(t *testing.T) {
		ctx := context.Background()
		newList := entity.List{
			ID:      1,
			Name:    "List 1",
			Status:  "Todo",
			Details: "Example Details",
		}

		expectedError := errors.New("cannot create list")
		mr, uc := newMock()

		mr.On("CreateList", mock.Anything, newList).Return(newList, expectedError)

		list, err := uc.CreateList(ctx, newList)

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		assert.Equal(t, newList, list)
		mr.AssertExpectations(t)
	})
}

func TestTodoUseCase_UpdateList(t *testing.T) {

	t.Run("UpdateList Success", func(t *testing.T) {
		ctx := context.Background()
		id := "1"
		expectedList := entity.List{
			ID:      1,
			Name:    "List 1",
			Status:  "Todo",
			Details: "Example Details",
		}

		mr, uc := newMock()

		mr.On("UpdateList", mock.Anything, expectedList, id).Return(expectedList, nil)

		list, err := uc.UpdateList(ctx, expectedList, id)

		assert.NoError(t, err)
		assert.Equal(t, expectedList, list)
		mr.AssertExpectations(t)
	})

	t.Run("UpdateList Error", func(t *testing.T) {
		ctx := context.Background()
		id := "1"
		expectedList := entity.List{
			ID:      1,
			Name:    "List 1",
			Status:  "Todo",
			Details: "Example Details",
		}
		expectedError := errors.New("cannot update list")

		mr, uc := newMock()

		mr.On("UpdateList", mock.Anything, expectedList, id).Return(expectedList, expectedError)

		list, err := uc.UpdateList(ctx, expectedList, id)

		assert.Error(t, err)
		assert.Equal(t, expectedList, list)
		assert.Equal(t, expectedError, err)
		mr.AssertExpectations(t)
	})

}

func TestTodoUseCase_SortListsByID(t *testing.T) {
	t.Run("SortList Success", func(t *testing.T) {

		ctx := context.Background()

		expectedList := []entity.List{
			{ID: 3, Name: "List 3", Status: "Todo", Details: "Details 3"},
			{ID: 1, Name: "List 1", Status: "Doing", Details: "Details 1"},
			{ID: 2, Name: "List 2", Status: "Done", Details: "Details 2"},
		}
		mr, uc := newMock()

		mr.On("SortListsByID", mock.Anything).Return(expectedList, nil)

		lists, err := uc.SortListsByID(ctx)

		assert.NoError(t, err)
		assert.Equal(t, expectedList, lists)
		mr.AssertExpectations(t)
	})

	t.Run("SortList Error", func(t *testing.T) {
		ctx := context.Background()

		expectedError := errors.New("cannot sort")
		expectedList := []entity.List{
			{ID: 3, Name: "List 3", Status: "Todo", Details: "Details 3"},
			{ID: 1, Name: "List 1", Status: "Doing", Details: "Details 1"},
			{ID: 2, Name: "List 2", Status: "Done", Details: "Details 2"},
		}
		mr, uc := newMock()

		mr.On("SortListsByID", mock.Anything).Return(expectedList, expectedError)

		lists, err := uc.SortListsByID(ctx)

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		assert.Equal(t, expectedList, lists)
		mr.AssertExpectations(t)
	})
}

func TestTodoUseCase_DeleteList(t *testing.T) {
	t.Run("DeleteList Success", func(t *testing.T) {

		ctx := context.Background()

		id := "1"

		mr, uc := newMock()

		mr.On("DeleteList", mock.Anything, id).Return(nil)

		err := uc.DeleteList(ctx, id)

		assert.NoError(t, err)
		mr.AssertExpectations(t)
	})

	t.Run("DeleteList Error", func(t *testing.T) {

		ctx := context.Background()
		id := "1"
		expectedError := errors.New("cannot delete list")

		mr, uc := newMock()

		mr.On("DeleteList", mock.Anything, id).Return(expectedError)

		err := uc.DeleteList(ctx, id)

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		mr.AssertExpectations(t)
	})
}
func TestTodoUseCase_PatchList(t *testing.T) {

	t.Run("Patch List Success", func(t *testing.T) {
		ctx := context.Background()
		id := "1"
		expectedList := entity.List{
			ID:      1,
			Name:    "List 1",
			Status:  "Todo",
			Details: "Example Details",
		}

		mr, uc := newMock()

		mr.On("PatchList", mock.Anything, expectedList, id).Return(expectedList, nil)

		list, err := uc.PatchList(ctx, expectedList, id)

		assert.NoError(t, err)
		assert.Equal(t, expectedList, list)
		mr.AssertExpectations(t)
	})

	t.Run("PatchList Error", func(t *testing.T) {
		ctx := context.Background()
		id := "1"
		expectedList := entity.List{
			ID:      1,
			Name:    "List 1",
			Status:  "Todo",
			Details: "Example Details",
		}
		expectedError := errors.New("cannot patch list")

		mr, uc := newMock()

		mr.On("PatchList", mock.Anything, expectedList, id).Return(expectedList, expectedError)

		list, err := uc.PatchList(ctx, expectedList, id)

		assert.Error(t, err)
		assert.Equal(t, expectedList, list)
		assert.Equal(t, expectedError, err)
		mr.AssertExpectations(t)
	})
}

func TestTodoUseCase_ChangeStatus(t *testing.T) {
	t.Run("ChangeStatus Success", func(t *testing.T) {
		ctx := context.Background()
		id := "1"

		expectedList := entity.List{
			ID:      1,
			Name:    "List 1",
			Status:  "Todo",
			Details: "Example Details",
		}

		mr, uc := newMock()

		mr.On("ChangeStatus", mock.Anything, expectedList, id).Return(expectedList, nil)

		list, err := uc.ChangeStatus(ctx, expectedList, id)

		assert.Equal(t, expectedList, list)
		assert.NoError(t, err)
		mr.AssertExpectations(t)
	})

	t.Run("ChangeStatus Error", func(t *testing.T) {
		ctx := context.Background()
		id := "1"

		expectedList := entity.List{
			ID:      1,
			Name:    "List 1",
			Status:  "Todo",
			Details: "Example Details",
		}
		expectedError := errors.New("invalid status")

		mr, uc := newMock()

		mr.On("ChangeStatus", mock.Anything, expectedList, id).Return(expectedList, expectedError)

		list, err := uc.ChangeStatus(ctx, expectedList, id)

		assert.Equal(t, expectedList, list)
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
		mr.AssertExpectations(t)
	})

}
