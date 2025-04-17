package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"go.opentelemetry.io/otel"
)

type WeatherAPIResponse struct {
	Current struct {
		TempC float64 `json:"temp_c"`
	} `json:"current"`
}

func GetTemperatureByCity(ctx context.Context, city string) (float64, error) {
	tr := otel.Tracer("weather-service")
	ctx, span := tr.Start(ctx, "CallWeatherAPI")
	defer span.End()

	apiKey := os.Getenv("WEATHER_API_KEY")
	url := fmt.Sprintf("https://api.weatherapi.com/v1/current.json?key=%s&q=%s", apiKey, city)

	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var data WeatherAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return 0, err
	}

	return data.Current.TempC, nil
}
