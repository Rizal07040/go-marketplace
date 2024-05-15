package repositories

import (
	"go-marketplace/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	Db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{Db: db}
}

func (t UserRepository) FindAll() (user []models.User, err error) {
	results := t.Db.Find(&user)
	if results.Error != nil {
		return nil, results.Error
	}

	return user, nil
}
