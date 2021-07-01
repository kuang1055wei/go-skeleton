package model

import (
	"gorm.io/gorm"
)

type Model struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt DateTime       `json:"created_at" swaggertype:"primitive,integer"`
	UpdatedAt DateTime       `json:"updated_at" swaggertype:"primitive,integer"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at" swaggertype:"primitive,integer"`
}
