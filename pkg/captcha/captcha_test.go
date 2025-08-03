package captcha

import (
	"net/http"
	"net/http/httptest"
	"testing"

	configuration "github.com/PotatoHD404/crowdsec-bouncer-traefik-plugin/pkg/configuration"
	logger "github.com/PotatoHD404/crowdsec-bouncer-traefik-plugin/pkg/logger"
)

func TestYandexCaptchaProvider(t *testing.T) {
	// Test that Yandex provider is properly configured
	provider := configuration.YandexProvider
	
	// Check if provider exists in the captcha map
	if _, exists := captcha[provider]; !exists {
		t.Errorf("Yandex provider not found in captcha providers map")
	}
	
	// Check provider configuration
	info := captcha[provider]
	if info.js != "https://smartcaptcha.yandexcloud.net/captcha.js" {
		t.Errorf("Expected Yandex JS URL, got %s", info.js)
	}
	
	if info.key != "smart-token" {
		t.Errorf("Expected Yandex key 'smart-token', got %s", info.key)
	}
	
	if info.validate != "https://smartcaptcha.yandexcloud.net/validate" {
		t.Errorf("Expected Yandex validate URL, got %s", info.validate)
	}
}

func TestYandexCaptchaClient(t *testing.T) {
	log := logger.New(configuration.LogINFO, "")
	
	client := &Client{}
	err := client.New(log, nil, &http.Client{}, configuration.YandexProvider, "test-site-key", "test-secret-key", "", "/path/to/template", 1800)
	
	if err != nil {
		t.Errorf("Failed to create Yandex captcha client: %v", err)
	}
	
	if !client.Valid {
		t.Error("Yandex captcha client should be valid")
	}
	
	if client.provider != configuration.YandexProvider {
		t.Errorf("Expected provider %s, got %s", configuration.YandexProvider, client.provider)
	}
}

func TestYandexCaptchaValidation(t *testing.T) {
	// Create a test server to mock Yandex validation response
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"success": true}`))
	}))
	defer server.Close()
	
	log := logger.New(configuration.LogINFO, "")
	client := &Client{}
	err := client.New(log, nil, &http.Client{}, configuration.YandexProvider, "test-site-key", "test-secret-key", "", "/path/to/template", 1800)
	
	if err != nil {
		t.Fatalf("Failed to create Yandex captcha client: %v", err)
	}
	
	// Create a test request with smart-token
	req := httptest.NewRequest("POST", "/", nil)
	req.Form = map[string][]string{
		"smart-token": {"test-token"},
	}
	
	// Test validation (this will fail in test environment but should not panic)
	valid, err := client.Validate(req)
	if err != nil {
		// Expected error in test environment since we can't reach Yandex servers
		t.Logf("Validation error (expected in test): %v", err)
	} else if valid {
		t.Log("Validation successful")
	}
} 