package domain

import (
	"ProjectsGo/pkg/utils/response"
	"gorm.io/gorm"
	"time"
)

type ToDo struct {
	Default
	Title       string     `json:"title" validate:"required" gorm:"type:varchar(255);not null"`
	Description string     `json:"description" validate:"omitempty" gorm:"type:varchar(255)"` // Não obrigatorio
	Priority    string     `json:"priority" validate:"oneof=LOW MEDIUM HIGH" gorm:"type:enum('LOW', 'MEDIUM', 'HIGH');default:'LOW'"`
	Category    string     `json:"category" validate:"omitempty" gorm:"type:varchar(255)"`
	IsCompleted bool       `json:"is_completed" validate:"omitempty" gorm:"default:false"`
	UserID      uint       `json:"user_id" validate:"omitempty" gorm:"not null"`
	Deadline    *time.Time `json:"deadline" validate:"required" gorm:"type:datetime"`
}

type ToDoUpdateRequest struct {
	//TODO create UPDATE
}

type ToDoFindFilter struct {
	//TODO create FIND
}

func (t *ToDoFindFilter) ApplyToDoFilter(db *gorm.DB) (*gorm.DB, []response.ValidationError) {
	var errList []response.ValidationError

	return db, errList
}
