package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"go.opentelemetry.io/otel"
)

type CEPRequest struct {
	CEP string `json:"cep"`
}

func HandleCEPInput(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	tr := otel.Tracer("input-service")
	ctx, span := tr.Start(ctx, "HandleCEPInput")
	defer span.End()

	var req CEPRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	if !isValidCEP(req.CEP) {
		http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
		return
	}

	serviceBURL := os.Getenv("SERVICE_B_URL")
	resp, err := http.Get(serviceBURL + "/weather?cep=" + req.CEP)
	if err != nil {
		http.Error(w, "service B unavailable", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}
