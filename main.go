package main

import (
	"database/sql"
	"net"

	"github.com/black-dev-x/go-grpc/database"
	"github.com/black-dev-x/go-grpc/internal/pb"
	"github.com/black-dev-x/go-grpc/services"
	"google.golang.org/grpc"
)

func main() {
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	categoryDB := database.NewCategoryDB(db)
	categoryService := services.NewCategoryService(*categoryDB)

	grpcServer := grpc.NewServer()
	pb.RegisterCategoryServiceServer(grpcServer, categoryService)

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	if err := grpcServer.Serve(listener); err != nil {
		panic(err)
	}
}
