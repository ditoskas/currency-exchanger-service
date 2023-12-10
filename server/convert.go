package main

import (
	"context"
	"encoding/json"
	pb "github.com/ditoskas/currency-exchanger-service/server/pb"
	"github.com/redis/go-redis/v9"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Response struct {
	Success   bool               `json:"success"`
	Terms     string             `json:"terms"`
	Privacy   string             `json:"privacy"`
	Timestamp int                `json:"timestamp"`
	Source    string             `json:"source"`
	Quotes    map[string]float64 `json:"quotes"`
}

var ctx = context.Background()

func (s *Server) Convert(ctx context.Context, in *pb.ConvertRequest) (*pb.ConvertResponse, error) {
	log.Printf("Convert was invoked with %v\n", in)

	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDRESS"),
		Password: "",
		DB:       0,
	})
	var redisKey = in.FromCurrency + in.ToCurrency
	val, err := client.Get(ctx, redisKey).Result()
	rate := 0.0
	if err != nil {
		log.Print("Rate not found in Redis execute http request\n")

		url := os.Getenv("BASE_EXCHANGE_URL") + "live?access_key=" + os.Getenv("EXCHANGE_API_KEY") + "&source=" + in.FromCurrency
		log.Printf("HTTP GET\n %v\n", url)
		response, err := http.Get(url)
		if err != nil {
			log.Fatalf("Failed request: %v\n", err)
		}
		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatalf("Failed to read body: %v\n", err)
		}

		log.Printf("Response body is %v\n", string(responseData))

		var responseObject Response
		err = json.Unmarshal(responseData, &responseObject)
		if err != nil {
			log.Fatalf("Failed to Unmarshal: %v\n", err)
		}
		//loop responseObject.Quotes
		for currencyPair, rateValue := range responseObject.Quotes {
			log.Printf("Store Redis [%s] = [%f]\n", currencyPair, rateValue)
			client.Set(ctx, currencyPair, rateValue, 1*time.Hour)
		}
		rate = responseObject.Quotes[redisKey]
		log.Printf("Rate is %v\n", rate)
	} else {
		cachedRate, err := strconv.ParseFloat(val, 64)
		if err != nil {
			log.Fatalf("Failed to float rate: %v\n", err)
		}
		rate = cachedRate
	}

	return &pb.ConvertResponse{
		AmountTo:       in.Amount * rate,
		ToCurrency:     in.ToCurrency,
		AmountOriginal: in.Amount,
		FromCurrency:   in.FromCurrency,
	}, nil
}
