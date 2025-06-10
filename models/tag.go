package models

import (
	"time"
)

type TagPosition struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	Scope     string    `json:"scope" gorm:"index"`
	ScopeType string    `json:"scope_type" gorm:"index"` // project/file/function/class
	Subject   string    `json:"subject" gorm:"index"`
	KeyCode   string    `json:"key_code" gorm:"index"`
	CreatedAt time.Time `json:"created_at"`
}

type Tag struct {
	ID        uint              `gorm:"primarykey" json:"id"`
	Position  TagPosition       `gorm:"embedded" json:"position"`
	Pairs     map[string]string `json:"pairs" gorm:"serializer:json"`
	CreatedAt time.Time         `json:"created_at"`
}
