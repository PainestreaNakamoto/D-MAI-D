package domain

type User struct {
	Username string
	Password string
}

type AuthorizationService interface {
	Login(username, password string) (string, error)
	Register(User) error
	GetUserByUsername(username string) (User, error)
	GetUserByID(id int) (User, error)
}
