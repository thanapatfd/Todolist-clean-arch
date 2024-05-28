package usecases

import (
	"github.com/thanapatfd/todolist/todo/usecases/repository"
	"go.opentelemetry.io/otel"
)

var tracer = otel.Tracer("usecases")

// TodoUseCase struct คือการจัดกลุ่มฟังก์ชันและข้อมูลที่เกี่ยวข้องกับ use case ของ todo list
type TodoUseCase struct {
	todoRepo repository.TodoRepository // Dependency: todoRepo ชนิด TodoRepository ที่จะถูก inject เข้ามา
}

// NewUsecase คือ constructor function สำหรับสร้าง instance ของ TodoUseCase
func NewUsecase(
	todoRepo repository.TodoRepository,
) TodoUseCase {
	return TodoUseCase{
		todoRepo: todoRepo, // กำหนดค่าให้กับฟิลด์ todoRepo ใน struct TodoUseCase
	}
}
