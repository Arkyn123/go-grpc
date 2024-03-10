package service

import "go-grpc/users"

type Users interface {
	CreateUser(user users.User) (users.User, int, error)
	FindAllUsers() (users.Users, int, error)
	FindUserById(id int) (users.User, int, error)
}
