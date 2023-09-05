package repository

import (
	"github.com/yeom-c/golang-gin-api/model"
	"gorm.io/gorm"
)

type AccountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) *AccountRepository {
	return &AccountRepository{db}
}

func (r *AccountRepository) FindById(id int) (*model.Account, error) {
	var acc model.Account
	err := r.db.First(&acc, id).Error
	return &acc, err
}
