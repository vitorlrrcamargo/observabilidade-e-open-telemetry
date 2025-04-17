package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/vitorlrrcamargo/observabilidade-e-open-telemetry/handler"
)

func TestInvalidCEP(t *testing.T) {
	req, _ := http.NewRequest("GET", "/weather?cep=1234", nil)
	rr := httptest.NewRecorder()
	handler.GetWeatherByCEP(rr, req)

	if rr.Code != http.StatusUnprocessableEntity {
		t.Errorf("expected 422, got %d", rr.Code)
	}
}

func TestValidCEPButNotFound(t *testing.T) {
	req, _ := http.NewRequest("GET", "/weather?cep=00000000", nil)
	rr := httptest.NewRecorder()
	handler.GetWeatherByCEP(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Errorf("expected 404, got %d", rr.Code)
	}
}
