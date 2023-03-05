package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"user/injector"
	"user/pb/user"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	reflection.Register(server)

	userHandler := injector.NewUserInjector()
	user.RegisterUserServiceServer(server, userHandler)

	fmt.Printf("⚡️[server]: gRPC Server is running on port %s\n", port)
	log.Fatal(server.Serve(lis))
}
