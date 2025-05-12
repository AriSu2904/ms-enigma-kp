package repositories

import (
	"awesomeProject/models"
	"errors"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *models.User) error
	FindByNik(nik string) (*models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) FindByNik(nik string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("nik = ?", nik).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
