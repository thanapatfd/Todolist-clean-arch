package usecases

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/thanapatfd/todolist/todo/entity"
)

func TestUseCase_TodoList(t *testing.T) {
	t.Run("TestCase 1", func(t *testing.T) {
		// Setup
		ctx := context.Background()
		list := entity.List{
			ID:      1,
			Name:    "Example List",
			Status:  "Todo",
			Details: "Example Details",
		}

		mr, uc := newMock()

		mr.On("CreateList", mock.Anything, list).Return(list, nil)

		result, err := uc.CreateList(ctx, list)

		assert.NoError(t, err)
		assert.Equal(t, result, list)
		mr.AssertExpectations(t)
	})
}
