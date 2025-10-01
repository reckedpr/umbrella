package main

import (
	"github.com/reckedpr/umbrella/internal/format"
	"github.com/reckedpr/umbrella/internal/weather"

	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiKey := os.Getenv("WEATHER_API_KEY")
	queryLocation := os.Getenv("DEFAULT_LOCATION")

	if len(os.Args) >= 2 {
		queryLocation = os.Args[1]
	}

	weather, _ := weather.FetchForecast(apiKey, queryLocation)

	format.DisplayWeather(weather)
}
