package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/vitorlrrcamargo/observabilidade-e-open-telemetry/handler"
	"github.com/vitorlrrcamargo/observabilidade-e-open-telemetry/otelsetup"
)

func main() {
	// Inicializa o Tracer OTEL com Zipkin
	shutdown := otelsetup.InitTracer("http://zipkin:9411/api/v2/spans")
	defer shutdown(context.Background())

	mode := os.Getenv("SERVICE_MODE")

	switch mode {
	case "input":
		http.HandleFunc("/cep", handler.HandleCEPInput)
	case "weather":
		http.HandleFunc("/weather", handler.GetWeatherByCEP)
	default:
		log.Fatal("Invalid SERVICE_MODE. Must be 'input' or 'weather'")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Listening on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
