package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

// Todo structure
type Todo struct {
	ID           uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Title        string    `gorm:"size:255;not null;unique" json:"title"`
	Content      string    `gorm:"size:255;not null;" json:"content"`
	Author       User      `json:"author"`
	AuthorID     uint32    `gorm:"not null" json:"author_id"`
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	HasCompleted bool      `gorm:"default:false" json:"has_completed"`
}

// Prepare the todo
func (t *Todo) Prepare() {
	t.ID = 0
	t.Title = html.EscapeString(strings.TrimSpace(t.Title))
	t.Content = html.EscapeString(strings.TrimSpace(t.Content))
	t.Author = User{}
	t.CreatedAt = time.Now()
	t.UpdatedAt = time.Now()
	t.HasCompleted = false
}

// Validate enforces validation on todo
func (t *Todo) Validate() error {
	if t.Title == "" {
		return errors.New("Title is required")
	}
	if t.Content == "" {
		return errors.New("Content is required")
	}
	if t.AuthorID < 1 {
		return errors.New("Author is required")
	}
	return nil
}

// SaveTodo write to db
func (t *Todo) SaveTodo(db *gorm.DB) (*Todo, error) {
	var err error
	err = db.Debug().Model(&Todo{}).Create(&t).Error
	if err != nil {
		return &Todo{}, err
	}
	if t.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", t.AuthorID).Take(&t.Author).Error
		if err != nil {
			return &Todo{}, err
		}
	}
	return t, nil
}

// FindAllTodos gets all todos from db
func (t *Todo) FindAllTodos(db *gorm.DB) (*[]Todo, error) {
	var err error
	todos := []Todo{}
	err = db.Debug().Model(&Todo{}).Limit(100).Find(&todos).Error
	if err != nil {
		return &[]Todo{}, err
	}
	if len(todos) > 0 {
		for i := range todos {
			err := db.Debug().Model(&User{}).Where("id = ?", todos[i].AuthorID).Take(&todos[i].Author).Error
			if err != nil {
				return &[]Todo{}, err
			}
		}
	}
	return &todos, nil
}

// FindTodoByID gets a todo with tid from db
func (t *Todo) FindTodoByID(db *gorm.DB, tid uint64) (*Todo, error) {
	var err error
	err = db.Debug().Model(&Todo{}).Where("id = ?", tid).Take(&t).Error
	if err != nil {
		return &Todo{}, err
	}
	if t.ID != 0 {
		err = db.Debug().Model(
			&User{}).Where("id = ?", t.AuthorID).Take(&t.Author).Error
		if err != nil {
			return &Todo{}, err
		}
	}
	return t, nil
}

// UpdateATodo updates a todo object
func (t *Todo) UpdateATodo(db *gorm.DB) (*Todo, error) {

	var err error

	err = db.Debug().Model(&Todo{}).Where("id = ?", t.ID).Updates(Todo{Title: t.Title, Content: t.Content, UpdatedAt: time.Now(), HasCompleted: t.HasCompleted}).Error
	if err != nil {
		return &Todo{}, err
	}
	if t.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", t.AuthorID).Take(&t.Author).Error
		if err != nil {
			return &Todo{}, err
		}
	}
	return t, nil
}

// DeleteATodo delete a todo from db
func (t *Todo) DeleteATodo(db *gorm.DB, tid uint64, uid uint32) (int64, error) {
	db = db.Debug().Model(&Todo{}).Where("id = ? and author_id = ?", tid, uid).Take(&Todo{}).Delete(&Todo{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Todo not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
