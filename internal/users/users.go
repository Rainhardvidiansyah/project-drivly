package users

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type Users struct {
	ID          int          `gorm:"primaryKey;autoIncrement"`
	Name        string       `json:"name" gorm:"type:varchar(100)"`
	Email       string       `json:"email" gorm:"type:varchar(100);uniqueIndex;not null"`
	Password    string       `json:"-" gorm:"not null"`
	IsActive    bool         `json:"isActive" gorm:"default:false"`
	ActivatedAt sql.NullTime // Uses sql.NullTime for nullable time fields
	CreatedAt   time.Time    // Automatically managed by GORM for creation time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
