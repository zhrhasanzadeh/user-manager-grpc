package main

import (
	"context"
	"github.com/jackc/pgx/v4"
	"google.golang.org/grpc"
	"grpcprj/internal/handler"
	pb "grpcprj/internal/pb"
	"grpcprj/internal/repository/postgres"
	"grpcprj/internal/service"
	"grpcprj/pkg"
	"log"
	"net"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	conn, err := pgx.Connect(context.Background(), c.DatabaseURL)
	if err != nil {
		panic(err)
	}

	db := postgres.NewPostgresRepository(conn)
	newService := service.NewService(db)
	s := grpc.NewServer()
	gHandler := handler.NewGRPCHandler(newService)
	pb.RegisterUserServiceServer(s, gHandler)

	lis, err := net.Listen("tcp", c.Port)
	if err != nil {
		log.Fatalln("failed to listing:", err)
	}

	log.Printf("server listening at %v", lis.Addr())
	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
