package model

import (
	"github.com/golang-module/carbon"

	"gorm.io/gorm"
)

type Model struct {
	ID        uint                    `gorm:"primarykey" json:"id"`
	CreatedAt carbon.ToDateTimeString `json:"created_at"`
	UpdatedAt carbon.ToDateTimeString `json:"updated_at"`
	DeletedAt gorm.DeletedAt          `gorm:"index" json:"deleted_at"`
}
