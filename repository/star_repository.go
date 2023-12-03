package repository

import (
	"errors"
	"my-favorite-pokemon-rest-api/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// interface
// interfaceを実装するstruct
// コンストラクタ
// 実装部

type IStarRepository interface {
	GetAllStarByUserID(storedStars *[]model.Star, user_id uint) error
	CreateStar(star *model.Star) error
	UpdateStar(star *model.Star, pokemon_id uint, user_id uint) error
	DeleteStar(pokemon_id uint, user_id uint) error
}
type starRepository struct {
	db *gorm.DB
}

func NewStarRepository(db *gorm.DB) IStarRepository {
	return &starRepository{db: db}
}

func (sr *starRepository) GetAllStarByUserID(storedStars *[]model.Star, user_id uint) error {
	if err := sr.db.Where("user_id = ?", user_id).Order("id").Find(storedStars).Error; err != nil {
		return err
	}
	return nil
}
func (sr *starRepository) CreateStar(star *model.Star) error {
	if err := sr.db.Create(star).Error; err != nil {
		return err
	}
	return nil
}
func (sr *starRepository) UpdateStar(star *model.Star, pokemon_id uint, user_id uint) error {
	result := sr.db.Model(star).Clauses(clause.Returning{}).Where("pokemon_id = ? AND user_id = ?", pokemon_id, user_id).Update("evaluation", star.Evaluation)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return errors.New("not found")
	}
	return nil
}
func (sr *starRepository) DeleteStar(pokemon_id uint, user_id uint) error {
	result := sr.db.Where("pokemon_id = ? AND user_id = ?", pokemon_id, user_id).Delete(&model.Star{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return errors.New("not found")
	}
	return nil
}
