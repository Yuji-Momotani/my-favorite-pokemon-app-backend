package validator

import (
	"my-favorite-pokemon-rest-api/model"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

// interface
type IUserValidation interface {
	UserValidate(user model.User) error
}

// struct
type userValidation struct{}

// コンストラクタ
func NewUserValidation() IUserValidation {
	return &userValidation{}
}

// バリデーション処理
func (uv *userValidation) UserValidate(user model.User) error {
	return validation.ValidateStruct(
		&user,
		validation.Field(
			&user.Email,
			validation.Required.Error("Email is Empty"),
			is.Email,
			validation.Length(1, 50).Error("Email must be between 1 and 50 characters"),
		),
		validation.Field(
			&user.Password,
			validation.Required.Error("Password is Empty"),
			validation.Length(5, 15).Error("Password must be between 5 and 15 characters"),
		),
	)
}
