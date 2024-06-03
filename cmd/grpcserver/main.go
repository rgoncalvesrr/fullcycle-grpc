package main

import (
	"database/sql"
	"log"
	"net"

	_ "github.com/mattn/go-sqlite3"
	"github.com/rgoncalvesrr/fullcycle-grpc/internal/database"
	"github.com/rgoncalvesrr/fullcycle-grpc/internal/pb"
	"github.com/rgoncalvesrr/fullcycle-grpc/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	categoryDB := database.NewCategory(db)

	categoryService := service.NewCategoryService(*categoryDB)

	grpcServer := grpc.NewServer()
	pb.RegisterCategoryServiceServer(grpcServer, categoryService)
	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}

}
