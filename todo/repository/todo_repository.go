package repository

import (
	"context"
	"log/slog"

	"github.com/thanapatfd/todolist/todo/entity"
	"gorm.io/gorm"
)

type TodoModel struct {
	ID      int
	Name    string
	Status  string
	Details string
}

// todoRepositoryDB struct ใช้เก็บการเชื่อมต่อกับฐานข้อมูล
type todoRepositoryDB struct {
	db *gorm.DB
}

// NewTodoRepository เป็น constructor function สำหรับสร้าง instance ใหม่ของ todoRepositoryDB
func NewTodoRepository(db *gorm.DB) todoRepositoryDB {
	return todoRepositoryDB{db: db}
}

func (r todoRepositoryDB) GetLists(ctx context.Context, name string, status string) ([]entity.List, error) {
	ctx, sp := tracer.Start(ctx, "repositories.GetLists") 
	defer sp.End()                                      

	listRepo := []TodoModel{}
	result := r.db.WithContext(ctx) 
	if name != "" {
		result = result.Where("name LIKE ?", "%"+name+"%") // ตามชื่อ
	}

	if status != "" {
		result = result.Where("status = ?", status) // ตามสถานะ
	}

	result = result.Find(&listRepo) // ดึงข้อมูลจากฐานข้อมูล
	if result.Error != nil {
		slog.Warn("query error")
		return nil, result.Error
	}

	var rows []entity.List
	for _, list := range listRepo {
		rows = append(rows, entity.List{
			ID:      list.ID,
			Name:    list.Name,
			Status:  list.Status,
			Details: list.Details,
		})
	}

	sp.AddEvent("Get Lists Success") // เพิ่มเหตุการณ์

	return rows, nil
}

func (r todoRepositoryDB) GetListByID(ctx context.Context, id string) (entity.List, error) {
	ctx, sp := tracer.Start(ctx, "repositories.GetListByID") 
	defer sp.End()                                           

	listRepo := TodoModel{}
	result := r.db.WithContext(ctx).Where("id = ?", id).Limit(1).Find(&listRepo) // ดึงข้อมูลจากฐานข้อมูลตาม ID
	if result.Error != nil {
		return entity.List{}, nil 
	}

	sp.AddEvent("Get List By ID Success") // เพิ่มเหตุการณ์

	return entity.List{
		ID:      listRepo.ID,
		Name:    listRepo.Name,
		Status:  listRepo.Status,
		Details: listRepo.Details,
	}, nil
}

func (r todoRepositoryDB) CreateList(ctx context.Context, list entity.List) (entity.List, error) {
	ctx, sp := tracer.Start(ctx, "repositories.CreateList")
	defer sp.End()                                    

	result := r.db.WithContext(ctx).Create(&TodoModel{
		Name:    list.Name,
		Status:  list.Status,
		Details: list.Details,
	}) // สร้างรายการใหม่ในฐานข้อมูล

	if result.Error != nil {
		slog.Error("query error") 
		return list, result.Error
	}

	lastInsertedID := 0
	
    // ดึง ID ที่เพิ่งถูกอัปเดตล่าสุด (descending order)
	r.db.Table("todo_models").Select("id").Order("id desc").Limit(1).Row().Scan(&lastInsertedID) // ดึง ID ที่เพิ่งถูกสร้างขึ้นล่าสุด
	list.ID = lastInsertedID

	sp.AddEvent("Create List Success") // เพิ่มเหตุการณ์

	return list, nil
}

func (r todoRepositoryDB) UpdateList(ctx context.Context, list entity.List, id string) (entity.List, error) {
	ctx, sp := tracer.Start(ctx, "repository_UpdateList") 
	defer sp.End()                                   

	listRepo := TodoModel{}
	result := r.db.WithContext(ctx).Where("id = ?", id).Limit(1).Find(&listRepo) // ดึงข้อมูลจากฐานข้อมูลตาม ID
	if result.Error != nil {
		slog.Error("query error")
		return entity.List{}, result.Error
	}

	listRepo = TodoModel{
		ID:      list.ID,
		Name:    list.Name,
		Status:  list.Status,
		Details: list.Details,
	}
	result = r.db.WithContext(ctx).Where("id = ?", id).Updates(&listRepo) // อัปเดตข้อมูลในฐานข้อมูล
	if result.Error != nil {
		slog.Error("query error") 
		return list, result.Error
	}

	lastInsertedID := 0

    // ดึง ID ที่เพิ่งถูกอัปเดตล่าสุด (descending order)
	r.db.Table("todo_models").Select("id").Order("id desc").Limit(1).Row().Scan(&lastInsertedID) 
	list.ID = lastInsertedID

	sp.AddEvent("Update List Success") // เพิ่มเหตุการณ์

	return list, nil
}

func (r todoRepositoryDB) PatchList(ctx context.Context, list entity.List, id string) (entity.List, error) {
	ctx, sp := tracer.Start(ctx, "repository_PatchList") 
	defer sp.End()                                     

	listRepo := TodoModel{}
	result := r.db.WithContext(ctx).Where("id = ?", id).Find(&listRepo) // ดึงข้อมูลจากฐานข้อมูลตาม ID
	if result.Error != nil {
		slog.Error("query error") 
		return entity.List{}, result.Error
	}

	if list.Name != "" {
		listRepo.Name = list.Name
	}
	if list.Status != "" {
		listRepo.Status = list.Status
	}
	if list.Details != "" {
		listRepo.Details = list.Details
	}

	result = r.db.WithContext(ctx).Save(&listRepo) // บันทึกข้อมูลในฐานข้อมูล
	if result.Error != nil {
		slog.Error("query error")
		return entity.List{}, result.Error
	}

	list = entity.List{
		ID:      listRepo.ID,
		Name:    listRepo.Name,
		Status:  listRepo.Status,
		Details: listRepo.Details,
	}

	sp.AddEvent("Patch List Success") // เพิ่มเหตุการณ์

	return list, nil
}

// DeleteList เป็นเมธอดสำหรับลบรายการ Todo จากฐานข้อมูลตาม ID
func (r todoRepositoryDB) DeleteList(ctx context.Context, id string) error {
	ctx, sp := tracer.Start(ctx, "repository_DeleteList") 
	defer sp.End()                                      

	deleteList := TodoModel{}
	result := r.db.WithContext(ctx).Where("id = ?", id).Delete(&deleteList) // ลบข้อมูลจากฐานข้อมูลตาม ID
	if result.Error != nil {
		slog.Error("query error") 
		return result.Error
	}
	sp.AddEvent("Delete List Success") // เพิ่มเหตุการณ์

	return nil
}

func (r todoRepositoryDB) SortListsByID(ctx context.Context) ([]entity.List, error) {
	ctx, sp := tracer.Start(ctx, "repository_SortList") 
	defer sp.End()                                     

	lists := []TodoModel{}
	result := r.db.WithContext(ctx).Order("id").Find(&lists) // ดึงข้อมูลจากฐานข้อมูลและเรียงลำดับตาม ID
	if result.Error != nil {
		slog.Error("query error") 
		return nil, result.Error
	}

	var rows []entity.List
	for _, list := range lists {
		rows = append(rows, entity.List{
			ID:      list.ID,
			Name:    list.Name,
			Status:  list.Status,
			Details: list.Details,
		})
	}

	sp.AddEvent("Sort List Success") // เพิ่มเหตุการณ์

	return rows, nil
}
