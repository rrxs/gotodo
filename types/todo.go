package types

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Title  string
	IsDone bool
}

type TodoResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deteledAt,omitempty"`
	Title     string    `json:"description"`
	IsDone    bool      `json:"isDone"`
}
