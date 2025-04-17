package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"go.opentelemetry.io/otel"
)

type ViaCEPResponse struct {
	Localidade string `json:"localidade"`
}

func GetCityFromCEP(ctx context.Context, cep string) (string, error) {
	tr := otel.Tracer("weather-service")
	ctx, span := tr.Start(ctx, "CallViaCEP")
	defer span.End()

	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var data ViaCEPResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return "", err
	}

	if data.Localidade == "" {
		return "", fmt.Errorf("not found")
	}

	return data.Localidade, nil
}
