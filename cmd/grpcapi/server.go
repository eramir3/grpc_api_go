package main

import (
	"fmt"
	"grpcapi/internals/api/handlers"
	"grpcapi/internals/repositories/mongodb"
	"log"
	"net"
	"os"

	pb "grpcapi/proto/gen"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	mongodb.CreateMongoClient()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	server := grpc.NewServer()

	pb.RegisterExecsServiceServer(server, &handlers.Server{})
	pb.RegisterStudentsServiceServer(server, &handlers.Server{})
	pb.RegisterTeachersServiceServer(server, &handlers.Server{})

	reflection.Register(server)

	port := os.Getenv("SERVER_PORT")

	fmt.Println("gRPC Server is running on port:", port)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("Error listening on specified port:", err)
	}

	err = server.Serve(lis)
	if err != nil {
		log.Fatal("Failed to serve:", err)
	}
}
