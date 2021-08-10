package Users

import (
	"time"
)

type RegisterInput struct {
	Name     string    `json:"name" binding:"required"`
	Email    string    `json:"email" binding:"required,email"`
	BOD      time.Time `json:"bod" binding:"required"`
	Password string    `json:"password" binding:"required"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
