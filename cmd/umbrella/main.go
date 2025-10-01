package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/reckedpr/umbrella/internal/cli"
	"github.com/reckedpr/umbrella/internal/format"
	"github.com/reckedpr/umbrella/internal/weather"
)

func main() {

	args := cli.ParseArgs()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiKey := os.Getenv("WEATHER_API_KEY")
	queryLocation := os.Getenv("DEFAULT_LOCATION")

	weather, _ := weather.FetchForecast(apiKey, queryLocation)

	format.DisplayWeather(weather, args)
}
