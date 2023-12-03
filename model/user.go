package model

import "time"

type User struct {
	// Gormが自動で複数形 usersでテーブル作成してくれる
	ID        uint      `json:"id" gorm:"primaryKey"` // intのprimaryKeyでAutoIncrementもGormが自動で付与
	Email     string    `json:"email" gorm:"unique; not null"`
	Password  string    `json:"password" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserResponse struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
}
