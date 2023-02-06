package main

import (
	"ddd/adapter"
	"ddd/domain"
	"ddd/usecase"
	"fmt"
)

func main() {
	mysqlAuthorizationAdapter := &adapter.MysqlAuthorizationAdapter{}
	useCases := usecase.NewUseCases(mysqlAuthorizationAdapter)

	jwt, err := useCases.Login("username", "password")
	if err != nil {
		panic(err)
		// Handle error
	}
	fmt.Println(jwt)
	// Use the JWT token

	err = useCases.Register(domain.User{
		Username: "new_username",
		Password: "new_password",
	})
	if err != nil {
		panic(err)
		// Handle error
	}
}
