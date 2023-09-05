package service

type AccountService struct {
}

func NewAccountService() *AccountService {
	return &AccountService{}
}

func (s *AccountService) Get() (string, error) {
	return "account", nil
}
