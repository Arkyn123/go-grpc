package handler

import (
	"context"

	"go-grpc/pkg/service"
	"go-grpc/users"
	"math/rand"
)

type GRPCHandler struct {
	users.UnimplementedUsersServiceServer
	services *service.Service
}

func NewHandler(services *service.Service) *GRPCHandler {
	return &GRPCHandler{services: services}
}

func (s *GRPCHandler) Create(ctx context.Context, req *users.CreateRequest) (*users.User, error) {


	return &users.User{Id: rand.Int31(), Name: req.Name}, nil
}

func (s *GRPCHandler) FindAll(ctx context.Context, req *users.Empty) (*users.Users, error) {
	userList := []*users.User{
		{Id: rand.Int31(), Name: "User 1"},
		{Id: rand.Int31(), Name: "User 2"},
	}

	return &users.Users{Users: userList}, nil
}

func (s *GRPCHandler) FindById(ctx context.Context, req *users.FindByIdRequest) (*users.User, error) {
	return &users.User{Id: req.Id, Name: "sueta"}, nil
}
