package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	expected := "ValidKeyHere"
	value := "ApiKey " + expected

	header := make(http.Header)
	header.Set(Authorization, value)

	recieved, err := GetAPIKey(header)
	if err != nil {
		t.Errorf("expected no errors, got=%s", err.Error())
	}

	if recieved != expected {
		t.Errorf("expected api key to be=%s, got=%s", expected, recieved)
	}
}

func TestGetAPIKeyErrors(t *testing.T) {
	tests := []struct {
		input    string
		expected error
	}{
		{"", ErrNoAuthHeaderIncluded},
		{"somekeyhere", ErrMalformedAuthHeader},
		{" emptyfirstvalue", ErrMalformedAuthHeader},
	}

	for _, tt := range tests {
		header := make(http.Header)
		header.Set(Authorization, tt.input)

		key, err := GetAPIKey(header)
		if err == nil {
			t.Errorf("expected err=%s, got=nil", tt.expected.Error())
		}

		if key != "" {
			t.Errorf("expected key='', got=%s", key)
		}

		if err != tt.expected {
			t.Errorf("expected err=%s, got=%s", tt.expected.Error(), err.Error())
		}
	}
}
