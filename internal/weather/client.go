package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/reckedpr/umbrella/internal/model"
)

func FetchForecast(apiKey, queryLocation string) (model.Weather, error) {
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
			Error model.Error `json:"error"`
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

	var weather model.Weather
	err = json.Unmarshal(body, &weather)
	if err != nil {
		panic(err)
	}

	return weather, nil
}
