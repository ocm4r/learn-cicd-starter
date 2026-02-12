package auth

import (
	"errors"
	"net/http"
	"testing"
)

// Unit tests for GetAPIKey
func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name           string
		headers        http.Header
		expectedAPIKey string
		expectedErr    error
	}{
		{
			name:           "No Authorization header",
			headers:        http.Header{},
			expectedAPIKey: "",
			expectedErr:    ErrNoAuthHeaderIncluded,
		},
		{
			name: "Malformed Authorization header",
			headers: http.Header{
				"Authorization": []string{"Bearer some-token"},
			},
			expectedAPIKey: "",
			expectedErr:    errors.New("malformed authorization header"),
		},
		{
			name: "Correct Authorization header",
			headers: http.Header{
				"Authorization": []string{"ApiKey valid-api-key"},
			},
			expectedAPIKey: "valid-api-key",
			expectedErr:    nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			apiKey, err := GetAPIKey(tt.headers)

			if apiKey != tt.expectedAPIKey || (err != nil && err.Error() != tt.expectedErr.Error()) {
				t.Errorf("expected (%s, %v), got (%s, %v)", tt.expectedAPIKey, tt.expectedErr, apiKey, err)
			}
		})
	}
}
