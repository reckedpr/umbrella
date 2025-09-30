package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Error struct {
	Message string `json:"message"`
}

type Weather struct {
	Location Location `json:"location"`
	Current  Current  `json:"current"`
	Forecast Forecast `json:"forecast"`
}

type Location struct {
	Name    string `json:"name"`
	Country string `json:"country"`
}

type Current struct {
	TempC     float64   `json:"temp_c"`
	Condition Condition `json:"condition"`
}

type Forecast struct {
	ForecastDay []ForecastDay `json:"forecastday"`
}

type ForecastDay struct {
	Hour []Hour `json:"hour"`
}

type Hour struct {
	TimeEpoch    int64     `json:"time_epoch"`
	TempC        float64   `json:"temp_c"`
	Condition    Condition `json:"condition"`
	ChanceOfRain float64   `json:"chance_of_rain"`
}

type Condition struct {
	Text string `json:"text"`
}

func FetchForecast(apiKey, queryLocation string) (Weather, error) {
	res, err := http.Get(fmt.Sprintf("http://api.weatherapi.com/v1/forecast.json?key=%s&q=%s&days=2&aqi=no&alerts=no", apiKey, queryLocation))
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	if res.StatusCode == 400 {
		var apiError struct {
			Error Error `json:"error"`
		}

		err = json.Unmarshal(body, &apiError)
		if err != nil {
			panic(err)
		}

		panic(apiError.Error.Message)
	}

	if res.StatusCode != 200 {
		panic("Weather api not available")
	}

	var weather Weather
	err = json.Unmarshal(body, &weather)
	if err != nil {
		panic(err)
	}

	return weather, nil
}
