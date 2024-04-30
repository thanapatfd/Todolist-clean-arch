package repository

import (
	"errors"

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

func (r todoRepositoryDB) GetLists() ([]entity.List, error) {
	lists := []TodoModel{}
	result := r.db.Find(&lists)
	if result.Error != nil {
		return nil, result.Error
	}

	var rows []entity.List
	for _, list := range lists {
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
	checkList := TodoModel{}
	result := r.db.Where("id = ?", id).Limit(1).Find(&checkList)
	if result.Error != nil {
		return entity.List{}, result.Error
	}

	if result.RowsAffected == 0 {
		return entity.List{}, errors.New("list not found")
	}

	return entity.List{
		ID:      checkList.ID,
		Name:    checkList.Name,
		Status:  checkList.Status,
		Details: checkList.Details,
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
	checkList := TodoModel{}
	result := r.db.Where("id = ?", id).Limit(1).Find(&checkList)
	if result.Error != nil {
		return entity.List{}, result.Error
	} else {
		updateList := TodoModel{
			ID:      list.ID,
			Name:    list.Name,
			Status:  list.Status,
			Details: list.Details,
		}
		result := r.db.Where("id = ?", id).Updates(&updateList)
		if result.Error != nil {
			return list, result.Error
		}
		lastInsertedID := 0
		r.db.Table("todo_models").Select("id").Order("id desc").Limit(1).Row().Scan(&lastInsertedID)
		list.ID = lastInsertedID

		return list, nil
	}

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
