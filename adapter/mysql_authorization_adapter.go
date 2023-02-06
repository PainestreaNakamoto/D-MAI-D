package adapter

import "ddd/domain"

type MysqlAuthorizationAdapter struct{}

func (a *MysqlAuthorizationAdapter) Login(username, password string) (string, error) {
	// Your code to query the database for authentication
	return "jwt_token", nil
}

func (a *MysqlAuthorizationAdapter) Register(user domain.User) error {
	// Your code to insert the user into the database
	return nil
}

func (a *MysqlAuthorizationAdapter) GetUserByUsername(username string) (domain.User, error) {
	// Your code to retrieve the user from the database
	return domain.User{}, nil
}

func (a *MysqlAuthorizationAdapter) GetUserByID(id int) (domain.User, error) {
	// Your code to retrieve the user from the database
	return domain.User{}, nil
}
