package domain

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"time"
)

type Default struct {
	ID        uint            `json:"id,omitempty" gorm:"primaryKey"`
	CreatedAt *time.Time      `json:"created_at"`
	UpdatedAt *time.Time      `json:"updated_at"`
	DeletedAt *gorm.DeletedAt `json:"-" gorm:"index"`
}

var Validator = validator.New()
