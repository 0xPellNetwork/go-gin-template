package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	Name  string `json:"name" binding:"required" gorm:"not null"`
	Email string `json:"email" binding:"required,email" gorm:"uniqueIndex;not null"`
	Age   int    `json:"age" binding:"min=1,max=150"`
	Phone string `json:"phone"`
}

type CreateUserRequest struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	Age   int    `json:"age" binding:"min=1,max=150"`
	Phone string `json:"phone"`
}

type UpdateUserRequest struct {
	Name  *string `json:"name"`
	Email *string `json:"email" binding:"omitempty,email"`
	Age   *int    `json:"age" binding:"omitempty,min=1,max=150"`
	Phone *string `json:"phone"`
}

type GetUsersQuery struct {
	Page     int    `form:"page" binding:"omitempty,min=1"`
	PageSize int    `form:"page_size" binding:"omitempty,min=1,max=100"`
	Name     string `form:"name"`
	Email    string `form:"email"`
}
