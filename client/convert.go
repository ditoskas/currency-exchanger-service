package main

import (
	"context"
	pb "github.com/ditoskas/currency-exchanger-service/server/pb"
	_ "github.com/joho/godotenv/autoload"
	"log"
)

func convert(c pb.ConvertServiceClient, amount float64, fromCurrency string, toCurrency string) {
	log.Println("Convert was invoked")
	res, err := c.Convert(context.Background(), &pb.ConvertRequest{
		Amount:       amount,
		FromCurrency: fromCurrency,
		ToCurrency:   toCurrency,
	})

	if err != nil {
		log.Fatalf("Failed to convert: %v\n", err)
	}

	log.Printf("Convert: %f %s\n", res.AmountTo, res.ToCurrency)
}
