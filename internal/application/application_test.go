package application_test

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Gn1daa/calc_go/internal/application"
)

func TestCalcHandlerSuccessCase(t *testing.T) {
	expected := `{"result": "4.00"}`
	req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBufferString(`{"expression":"2+2"}`))
	w := httptest.NewRecorder()
	application.CalcHandler(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if string(data) != expected {
		t.Errorf("Expected %s but got %v", expected, string(data))
	}
}
func TestCalcHandlerInvalidExpressionCase(t *testing.T) {
	expected := `{"error": "Expression is not valid"}`
	req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBufferString(`{"expression":"2+2("}`))
	w := httptest.NewRecorder()
	application.CalcHandler(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if string(data) != expected {
		t.Errorf("Expected %s but got %v", expected, string(data))
	}
}
func TestCalcHandlerDividionByZeroCase(t *testing.T) {
	expected := `{"error": "Expression is not valid"}`
	req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBufferString(`{"expression":"2/0"}`))
	w := httptest.NewRecorder()
	application.CalcHandler(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if string(data) != expected {
		t.Errorf("Expected %s but got %v", expected, string(data))
	}
}
