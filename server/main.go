package main

import (
	pb "github.com/ditoskas/currency-exchanger-service/server/pb"
	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

type Server struct {
	pb.ConvertServiceServer
}

func main() {
	address := os.Getenv("LISTEN_ADDRESS")
	lis, err := net.Listen("tcp", address)

	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}

	log.Printf("Listening on %s\n", address)

	server := grpc.NewServer()
	pb.RegisterConvertServiceServer(server, &Server{})
	if err = server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
