package main

import (
	"fmt"
	"go-grpc/pkg/handler"
	"go-grpc/pkg/repository"
	"go-grpc/pkg/service"
	"go-grpc/users"
	"net"
	"os"
	"os/signal"
	"syscall"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

func main() {
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     ("localhost"),
		Port:     "5432",
		Username: "postgres",
		DBName:   "go_base",
		SSLMode:  "disable",
		Password: "root",
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	srv := grpc.NewServer()
	repository := repository.NewRepository(db)
	service := service.NewService(repository)
	handler := handler.NewHandler(service)

	users.RegisterUsersServiceServer(srv, handler)
	fmt.Println(srv.GetServiceInfo())

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Printf("Failed to listen: %v", err)
	}

	if err := srv.Serve(lis); err != nil {
		fmt.Printf("Failed to serve: %v", err)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	fmt.Println("GRPC Shutting Down")

	srv.GracefulStop()
}
