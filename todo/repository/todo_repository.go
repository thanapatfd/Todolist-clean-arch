package repository

import (
	"errors"
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

type todoRepositoryDB struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) *todoRepositoryDB {
	return &todoRepositoryDB{db: db}
}

func (r todoRepositoryDB) GetLists(name string, status string) ([]entity.List, error) {
	listRepo := []TodoModel{}
	result := r.db
	if name != "" {
		result = result.Where("name LIKE ?", "%"+name+"%")
	}

	if status != "" {
		result = result.Where("status = ?", status)
	}

	result = result.Find(&listRepo)
	if result.Error != nil {
		slog.Error("query error")
		return nil, result.Error
	}

	var rows []entity.List
	for _, list := range listRepo {
		// fmt.Println(list)
		rows = append(rows, entity.List{
			ID:      list.ID,
			Name:    list.Name,
			Status:  list.Status,
			Details: list.Details,
		})
	}

	return rows, nil
}

func (r todoRepositoryDB) GetListByID(id string) (entity.List, error) {
	listRepo := TodoModel{}
	result := r.db.Where("id = ?", id).Limit(1).Find(&listRepo)
	if result.Error != nil {
		return entity.List{}, result.Error
	}

	if result.RowsAffected == 0 {
		return entity.List{}, errors.New("list not found")
	}

	return entity.List{
		ID:      listRepo.ID,
		Name:    listRepo.Name,
		Status:  listRepo.Status,
		Details: listRepo.Details,
	}, nil
}

func (r todoRepositoryDB) CreateList(list entity.List) (entity.List, error) {
	result := r.db.Create(&TodoModel{
		Name:    list.Name,
		Status:  list.Status,
		Details: list.Details,
	})
	if result.Error != nil {
		return list, result.Error
	}

	lastInsertedID := 0
	r.db.Table("todo_models").Select("id").Order("id desc").Limit(1).Row().Scan(&lastInsertedID)
	list.ID = lastInsertedID

	return list, nil
}
func (r todoRepositoryDB) UpdateList(list entity.List, id string) (entity.List, error) {
	listRepo := TodoModel{}
	result := r.db.Where("id = ?", id).Limit(1).Find(&listRepo)
	if result.Error != nil {
		return entity.List{}, result.Error

	}
	listRepo = TodoModel{
		ID:      list.ID,
		Name:    list.Name,
		Status:  list.Status,
		Details: list.Details,
	}
	result = r.db.Where("id = ?", id).Updates(&listRepo)
	if result.Error != nil {
		return list, result.Error
	}
	lastInsertedID := 0
	r.db.Table("todo_models").Select("id").Order("id desc").Limit(1).Row().Scan(&lastInsertedID)
	list.ID = lastInsertedID

	return list, nil

}

func (r todoRepositoryDB) PatchList(list entity.List, id string) (entity.List, error) {
	listRepo := TodoModel{}
	result := r.db.Where("id = ?", id).Find(&listRepo)
	if result.Error != nil {
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

	result = r.db.Save(&listRepo)
	if result.Error != nil {
		return entity.List{}, result.Error
	}

	list = entity.List{
		ID:      listRepo.ID,
		Name:    listRepo.Name,
		Status:  listRepo.Status,
		Details: listRepo.Details,
	}

	return list, nil
}

func (r todoRepositoryDB) DeleteList(id string) error {
	deleteList := TodoModel{}

	result := r.db.Where("id = ?", id).Delete(&deleteList)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r todoRepositoryDB) SortListsByID() ([]entity.List, error) {
	lists := []TodoModel{}
	result := r.db.Order("id").Find(&lists)
	if result.Error != nil {
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
	return rows, nil
}
