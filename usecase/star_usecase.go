package usecase

import (
	"my-favorite-pokemon-rest-api/model"
	"my-favorite-pokemon-rest-api/repository"
	"my-favorite-pokemon-rest-api/validator"
)

// interface
type IStarUsecase interface {
	GetAllStars(user_id uint) ([]model.StarResponse, error)
	CreateStar(star model.Star) (model.StarResponse, error)
	UpdateStar(star model.Star, pokemon_id uint, user_id uint) (model.StarResponse, error)
	DeleteStar(pokemon_id uint, user_id uint) error
}

// interfaceを実装するstruct
type starUsecase struct {
	sr repository.IStarRepository
	sv validator.IStarValidation
}

// コンストラクタ
func NewStarUsecase(sr repository.IStarRepository, sv validator.IStarValidation) IStarUsecase {
	return &starUsecase{sr: sr, sv: sv}
}

// 実装部
func (su *starUsecase) GetAllStars(user_id uint) ([]model.StarResponse, error) {
	storedStar := []model.Star{}
	if err := su.sr.GetAllStarByUserID(&storedStar, user_id); err != nil {
		return []model.StarResponse{}, err
	}
	resStars := []model.StarResponse{}
	for _, v := range storedStar {
		sr := model.StarResponse{
			PokemonID:  v.PokemonID,
			Evaluation: v.Evaluation,
			UserID:     v.UserID,
		}
		resStars = append(resStars, sr)
	}
	return resStars, nil
}

func (su *starUsecase) CreateStar(star model.Star) (model.StarResponse, error) {
	if err := su.sv.StarValidation(star); err != nil {
		return model.StarResponse{}, err
	}
	if err := su.sr.CreateStar(&star); err != nil {
		return model.StarResponse{}, err
	}
	resStar := model.StarResponse{
		PokemonID:  star.PokemonID,
		Evaluation: star.Evaluation,
		UserID:     star.UserID,
	}
	return resStar, nil
}

func (su *starUsecase) UpdateStar(star model.Star, pokemon_id uint, user_id uint) (model.StarResponse, error) {
	if err := su.sv.StarValidation(star); err != nil {
		return model.StarResponse{}, err
	}
	if err := su.sr.UpdateStar(&star, pokemon_id, user_id); err != nil {
		return model.StarResponse{}, err
	}
	resStar := model.StarResponse{
		PokemonID:  star.PokemonID,
		Evaluation: star.Evaluation,
		UserID:     star.UserID,
	}
	return resStar, nil
}

func (su *starUsecase) DeleteStar(pokemon_id uint, user_id uint) error {
	if err := su.sr.DeleteStar(pokemon_id, user_id); err != nil {
		return err
	}
	return nil
}
