package usecases

import "github.com/thanapatfd/todolist/todo/usecases/repository"

func newMock() (*repository.MockRepository, TodoUseCase) {
	mr := new(repository.MockRepository) // สร้าง mock repository

	uc := NewUsecase(mr) // สร้าง use case โดยใช้ mock repository

	return mr, uc // คืนค่า mock repository และ use case
}
