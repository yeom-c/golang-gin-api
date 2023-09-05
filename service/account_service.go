package service

import (
	"github.com/yeom-c/golang-gin-api/model"
	"github.com/yeom-c/golang-gin-api/repository"
)

type AccountService struct {
	repo *repository.AccountRepository
}

func NewAccountService(repo *repository.AccountRepository) *AccountService {
	return &AccountService{repo}
}

func (s *AccountService) GetAccountById(id int) (*model.Account, error) {
	return s.repo.FindById(id)
}
