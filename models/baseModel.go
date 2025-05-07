package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID uuid.UUID `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (b *BaseModel) BeforeCreate(tx *gorm.DB) (err error)  {
	b.ID = uuid.New()
	return
}