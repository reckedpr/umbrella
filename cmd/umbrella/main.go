package main

import (
	"github.com/reckedpr/umbrella/internal/weather"

	"fmt"
	"log"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

var (
	yellowBg = color.New(color.FgBlack, color.BgYellow).Add(color.Bold)
	redBg    = color.New(color.FgBlack, color.BgRed).Add(color.Bold)
	greenBg  = color.New(color.FgBlack, color.BgGreen).Add(color.Bold)
	muteFg   = color.New(color.FgBlack)
	muteBg   = color.New(color.FgHiBlack)
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

	combinedHours := append(weather.Forecast.ForecastDay[0].Hour, weather.Forecast.ForecastDay[1].Hour...)

	location, _, hours := weather.Location, weather.Current, combinedHours

	fmt.Printf(
		"%s, %s\n\n",
		location.Name,
		location.Country,
	)

	lastText := ""

	for _, hour := range hours {
		var condText string
		date := time.Unix(hour.TimeEpoch, 0)

		if date.Before(time.Now()) || date.After(time.Now().Add(time.Hour*7)) {
			continue
		}

		dateStr := muteBg.Sprint(date.Format("15:04"))

		ChanceOfRainStr := fmt.Sprintf(" %.0f%% ", hour.ChanceOfRain)
		ChanceOfRainStr = fmt.Sprintf("%-6s", ChanceOfRainStr)
		TempCStr := fmt.Sprintf("%.0f󰔄", hour.TempC)

		if hour.ChanceOfRain >= 90 {
			ChanceOfRainStr = redBg.Sprint(ChanceOfRainStr)
		} else if hour.ChanceOfRain >= 60 {
			ChanceOfRainStr = yellowBg.Sprint(ChanceOfRainStr)
		} else {
			ChanceOfRainStr = greenBg.Sprint(ChanceOfRainStr)
		}

		if lastText == hour.Condition.Text {
			condText = muteFg.Sprint("^")
		} else {
			condText = hour.Condition.Text
		}

		lastText = hour.Condition.Text

		message := fmt.Sprintf(
			" %s %-3s %s %s\n",
			dateStr,
			TempCStr,
			ChanceOfRainStr,
			condText,
		)

		fmt.Print(message)
	}
}
