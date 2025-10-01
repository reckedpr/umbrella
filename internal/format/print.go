package format

import (
	"fmt"
	"time"

	"github.com/fatih/color"
	"github.com/reckedpr/umbrella/internal/model"
)

var (
	yellowBg = color.New(color.FgBlack, color.BgYellow).Add(color.Bold)
	redBg    = color.New(color.FgBlack, color.BgRed).Add(color.Bold)
	greenBg  = color.New(color.FgBlack, color.BgGreen).Add(color.Bold)
	muteFg   = color.New(color.FgBlack)
	muteBg   = color.New(color.FgHiBlack)
)

func DisplayWeather(w model.Weather) {
	hours := append(w.Forecast.ForecastDay[0].Hour, w.Forecast.ForecastDay[1].Hour...)

	fmt.Printf(
		"%s, %s\n\n",
		w.Location.Name,
		w.Location.Country,
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
