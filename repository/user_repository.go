package repository

import (
	"my-favorite-pokemon-rest-api/model"

	"gorm.io/gorm"
)

// インターフェース
type IUserRepository interface {
	GetUserByEmail(user *model.User, email string) error
	Create(user *model.User) error
}

// インターフェースを実装するstruct
type userRepository struct {
	db *gorm.DB
}

// コンストラクタ
func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db: db}
}

// 処理実装部
func (ur *userRepository) GetUserByEmail(user *model.User, email string) error {
	if err := ur.db.Where("email=?", email).First(user).Error; err != nil {
		// Firstは合致するデータがない場合もエラーを返す（ErrRecordNotFound）
		return err
	}
	return nil
}

func (ur *userRepository) Create(user *model.User) error {
	if err := ur.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}
