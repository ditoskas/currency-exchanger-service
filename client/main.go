package main

import (
	pb "github.com/ditoskas/currency-exchanger-service/server/pb"
	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
	"strconv"
)

func main() {
	address := os.Getenv("LISTEN_ADDRESS")
	var args = os.Args[1:]

	if len(args) != 3 {
		log.Fatalf("Please provide the correct arguments, amount, fromCurrency, toCurrency")
	}
	con, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}
	log.Printf("Connected...\n")
	defer func(con *grpc.ClientConn) {
		err := con.Close()
		if err != nil {
			log.Fatalf("Failed to close: %v\n", err)
		}
	}(con)
	c := pb.NewConvertServiceClient(con)

	f, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
	convert(c, f, args[1], args[2])
}
