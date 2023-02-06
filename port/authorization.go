package port

import "ddd/domain"

type AuthorizationPort interface {
	Login(username, password string) (string, error)
	Register(domain.User) error
	GetUserByUsername(username string) (domain.User, error)
	GetUserByID(id int) (domain.User, error)
}
