package validator

import (
	"my-favorite-pokemon-rest-api/model"

	validation "github.com/go-ozzo/ozzo-validation"
)

type IStarValidation interface {
	StarValidation(star model.Star) error
}

type starValidation struct{}

func NewStarValidation() IStarValidation {
	return &starValidation{}
}

func (sv *starValidation) StarValidation(star model.Star) error {
	return validation.ValidateStruct(
		&star,
		validation.Field(
			&star.Evaluation,
			// Evaluationが数値どうかチェック
			// Evaluationが1~3の値であることをチェック
			validation.Required.Error("Evaluation is Empty"),
			validation.In(uint(1), uint(2), uint(3)).Error("Evaluation must be between 1 and 3 a"),
			validation.Min(uint(1)).Error("Evaluation must be between 1 and 3 b"),
			validation.Max(uint(3)).Error("Evaluation must be between 1 and 3 c"),
		),
	)
}
