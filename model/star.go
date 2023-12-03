package model

import "time"

type Star struct {
	ID         uint      `json:"id" gorm:"primaryKey"` //intのprimaryKeyでGormが自動でautoincrementを付加
	PokemonID  uint      `json:"pokemon_id" gorm:"not null"`
	Evaluation uint      `json:"evaluation" gorm:"not null"`
	UserID     uint      `json:"user_id" gorm:"not null"`
	User       User      `json:"user" gorm:"foreignKey:user_id; constraint:OnDelete:CASCADE"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type StarResponse struct {
	PokemonID  uint `json:"pokemon_id"`
	Evaluation uint `json:"evaluation"`
	UserID     uint `json:"user_id"`
}
