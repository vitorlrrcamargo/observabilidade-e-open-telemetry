package handler

import (
	"encoding/json"
	"net/http"
	"regexp"

	"github.com/vitorlrrcamargo/observabilidade-e-open-telemetry/service"
	"go.opentelemetry.io/otel"
)

type WeatherResponse struct {
	City  string  `json:"city"`
	TempC float64 `json:"temp_C"`
	TempF float64 `json:"temp_F"`
	TempK float64 `json:"temp_K"`
}

func GetWeatherByCEP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	tr := otel.Tracer("weather-service")
	ctx, span := tr.Start(ctx, "GetWeatherByCEP")
	defer span.End()

	cep := r.URL.Query().Get("cep")

	if !isValidCEP(cep) {
		http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
		return
	}

	city, err := service.GetCityFromCEP(ctx, cep)
	if err != nil {
		http.Error(w, "can not find zipcode", http.StatusNotFound)
		return
	}

	celsius, err := service.GetTemperatureByCity(ctx, city)
	if err != nil {
		http.Error(w, "error getting temperature", http.StatusInternalServerError)
		return
	}

	response := WeatherResponse{
		City:  city,
		TempC: celsius,
		TempF: celsius*1.8 + 32,
		TempK: celsius + 273,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func isValidCEP(cep string) bool {
	match, _ := regexp.MatchString(`^\d{8}$`, cep)
	return match
}
