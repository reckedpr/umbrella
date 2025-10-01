package model

// args

type Args struct {
	Location string `arg:"-l"`
	Units    string `arg:"-u" default:"c"`
}

// misc

type Error struct {
	Message string `json:"message"`
}

// api

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
	TempF     float64   `json:"temp_f"`
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
	TempF        float64   `json:"temp_f"`
	Condition    Condition `json:"condition"`
	ChanceOfRain float64   `json:"chance_of_rain"`
}

type Condition struct {
	Text string `json:"text"`
}
