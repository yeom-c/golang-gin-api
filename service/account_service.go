package service

import (
	"github.com/yeom-c/golang-gin-api/model"
	"github.com/yeom-c/golang-gin-api/repository"
)

type AccountService struct {
	repo       *repository.AccountRepository
	dynamoRepo *repository.AccountDynamoRepository
}

func NewAccountService(repo *repository.AccountRepository, dynamoRepo *repository.AccountDynamoRepository) *AccountService {
	return &AccountService{repo, dynamoRepo}
}

func (s *AccountService) GetAccountById(id int) (*model.Account, error) {
	return s.repo.FindById(id)
}

func (s *AccountService) GetDynamoAccountById(id int) (*model.DynamoAccount, error) {
	return s.dynamoRepo.FindById(id)
}
