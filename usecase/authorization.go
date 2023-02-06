package usecase

import (
	"ddd/domain"
	_ "ddd/port"
)

type UseCases struct {
	AuthorizationService domain.AuthorizationService
}

func NewUseCases(authorizationService domain.AuthorizationService) *UseCases {
	return &UseCases{
		AuthorizationService: authorizationService,
	}
}

func (u *UseCases) Login(username, password string) (string, error) {
	return u.AuthorizationService.Login(username, password)
}

func (u *UseCases) Register(user domain.User) error {
	return u.AuthorizationService.Register(user)
}
