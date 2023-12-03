package usecase

import (
	"my-favorite-pokemon-rest-api/model"
	"my-favorite-pokemon-rest-api/repository"
	"my-favorite-pokemon-rest-api/validator"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// インターフェース
type IUserUsecase interface {
	SingUp(user model.User) (model.UserResponse, error)
	Login(user model.User) (string, error)
}

// インターフェースを実装するstruct

type userUsecase struct {
	ur repository.IUserRepository
	uv validator.IUserValidation
}

// コンストラクタ
func NewUserUsecase(ur repository.IUserRepository, uv validator.IUserValidation) IUserUsecase {
	return &userUsecase{ur: ur, uv: uv}
}

// メソッドの内容をもりもり
func (uu *userUsecase) SingUp(user model.User) (model.UserResponse, error) {
	if err := uu.uv.UserValidate(user); err != nil {
		return model.UserResponse{}, err
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return model.UserResponse{}, err
	}
	newUser := model.User{
		Email:    user.Email,
		Password: string(hash),
	}
	if err := uu.ur.Create(&newUser); err != nil {
		return model.UserResponse{}, err
	}
	resUser := model.UserResponse{
		ID:    newUser.ID,
		Email: newUser.Email,
	}
	return resUser, nil
}

func (uu *userUsecase) Login(user model.User) (string, error) {
	if err := uu.uv.UserValidate(user); err != nil {
		return "", err
	}
	storedUser := model.User{}
	if err := uu.ur.GetUserByEmail(&storedUser, user.Email); err != nil {
		return "", nil
	}
	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password)); err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": storedUser.ID,
		"exp":     time.Now().Add(time.Hour * 12).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
