package entity

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListChangeStatus(t *testing.T) {
	tests := []struct {
		name           string
		initialStatus  string
		newStatus      string
		expectedStatus string
		expectedError  error
	}{
		{
			name:           "Change from Todo to Doing",
			initialStatus:  "Todo",
			newStatus:      "Doing",
			expectedStatus: "Doing",
			expectedError:  nil,
		},
		{
			name:           "Change from Doing to Done",
			initialStatus:  "Doing",
			newStatus:      "Done",
			expectedStatus: "Done",
			expectedError:  nil,
		},
		{
			name:           "Invalid change from Done to Todo",
			initialStatus:  "Done",
			newStatus:      "Todo",
			expectedStatus: "Done",
			expectedError:  errors.New("cannot change status from Done"),
		},
		{
			name:           "Invalid change from Done to Doing",
			initialStatus:  "Done",
			newStatus:      "Doing",
			expectedStatus: "Done",
			expectedError:  errors.New("cannot change status from Done"),
		},
		{
			name:           "No change from Done to Done",
			initialStatus:  "Done",
			newStatus:      "Done",
			expectedStatus: "Done",
			expectedError:  nil,
		},
		{
			name:           "Invalid empty status",
			initialStatus:  "Todo",
			newStatus:      "",
			expectedStatus: "Todo",
			expectedError:  errors.New("invalid status"),
		},
		{
			name:           "Invalid status transition",
			initialStatus:  "Todo",
			newStatus:      "Done",
			expectedStatus: "Todo",
			expectedError:  errors.New("invalid status"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := &List{
				ID:      1,
				Name:    "Example List",
				Status:  tt.initialStatus,
				Details: "Example Details",
			}

			err := list.ChangeStatus(tt.newStatus)

			// Assert
			assert.Equal(t, tt.expectedStatus, list.Status)
			if tt.expectedError != nil {
				assert.EqualError(t, err, tt.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
