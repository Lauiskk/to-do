package domain

import (
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
}

type ToDoFindFilter struct {
}

func (t *ToDoFindFilter) ApplyToDoFilter(db *gorm.DB) (*gorm.DB, []response.ValidationError) {
	var errList []response.ValidationError

	if c.UserId != 0 {
		db = db.Joins("JOIN users ON users.company_id = companies.id").
			Where("users.id = ?", c.UserId)
	}

	if c.CompanyType != "" {
		likeQuery := "%" + c.CompanyType + "%"
		db = db.Where("(companies.id = ? AND companies.company_type ILIKE ?) OR (companies.company_type ILIKE ?)", c.CompanyID, likeQuery, likeQuery)
	}

	if c.CreatedGreaterOrEqualThan != "" {
		createdGreaterOrEqualThan, err := time.ParseInLocation("2006-01-02", c.CreatedGreaterOrEqualThan, time.Local)
		if err != nil {
			errList = append(errList, response.NewValidationError("created_greater_or_equal_than", response.ErrMsgInvalidFormat, nil, nil))
		} else {
			db = db.Where("(companies.id = ? AND companies.created_at >= ?) OR (companies.created_at >= ?)", c.CompanyID, createdGreaterOrEqualThan, createdGreaterOrEqualThan)
		}
	}

	if c.CreatedLessOrEqualThan != "" {
		createdLessOrEqualThan, err := time.ParseInLocation("2006-01-02", c.CreatedLessOrEqualThan, time.Local)
		if err != nil {
			errList = append(errList, response.NewValidationError("created_less_or_equal_than", response.ErrMsgInvalidFormat, nil, nil))
		} else {
			db = db.Where("(companies.id = ? AND companies.created_at < ?) OR (companies.created_at < ?)", c.CompanyID, createdLessOrEqualThan, createdLessOrEqualThan)
		}
	}

	if c.Name != "" {
		likeQuery := "%" + c.Name + "%"
		db = db.Where("(companies.id = ? AND companies.name ILIKE ?) OR (companies.name ILIKE ?)", c.CompanyID, likeQuery, likeQuery)
	}

	if c.Email != "" {
		if err := Validator.Struct(c); err != nil {
			errList = append(errList, response.NewValidationError("email", response.ErrMsgInvalidFormat, nil, nil))
		}
		db = db.Where("(companies.id = ? AND companies.email = ?) OR (companies.email = ?)", c.CompanyID, c.Email, c.Email)
	}

	if c.TradeName != nil {
		likeQuery := "%" + *c.TradeName + "%"
		db = db.Where("(companies.id = ? AND companies.trade_name ILIKE ?) OR (companies.trade_name ILIKE ?)", c.CompanyID, likeQuery, likeQuery)
	}

	if c.TaxNumber != "" {
		taxNumber, err := validators.ValidateTaxNumber(c.TaxNumber)
		if err != nil {
			errList = append(errList, response.NewValidationError("tax_number", response.ErrMsgInvalidFormat, nil, nil))
		} else {
			c.TaxNumber = taxNumber
			db = db.Where("(companies.id = ? AND companies.tax_number = ?) OR (companies.tax_number = ?)", c.CompanyID, c.TaxNumber, c.TaxNumber)
		}
	}

	if c.Phone != "" {
		phone, err := validators.IsValidPhoneNumber(c.Phone)
		if err != nil {
			errList = append(errList, response.NewValidationError("phone", response.ErrMsgInvalidFormat, nil, nil))
		} else {
			c.Phone = phone
			db = db.Where("(companies.id = ? AND companies.phone = ?) OR (companies.phone = ?)", c.CompanyID, c.Phone, c.Phone)
		}
	}

	return db, errList
}
