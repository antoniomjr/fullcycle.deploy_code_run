package main

import (
	"crypto/tls"
	"encoding/json"
	"net/http"
	"os"
	"testing"
)

func TestGetLocation(t *testing.T) {
	// Mock HTTP client
	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	// Mock response
	resp, err := httpClient.Get("https://viacep.com.br/ws/13064722/json/")
	if err != nil {
		t.Fatalf("Failed to get location: %v", err)
	}
	defer resp.Body.Close()

	var viaCEP ViaCEPResponse
	if err := json.NewDecoder(resp.Body).Decode(&viaCEP); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if viaCEP.Localidade != "Campinas" {
		t.Errorf("Expected Campinas, got %s", viaCEP.Localidade)
	}
}

func TestGetTemperature(t *testing.T) {
	// Set environment variable for API key
	os.Setenv("WEATHER_API_KEY", "677f4ab826e94d0fbec120112241809")

	temp, err := getTemperature("Campinas")
	if err != nil {
		t.Fatalf("Failed to get temperature: %v", err)
	}

	if temp == 0 {
		t.Errorf("Expected non-zero temperature, got %f", temp)
	}
}
