package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name     string
		headers  http.Header
		expected string
		wantErr  bool
	}{
		{
			name: "gets api key from authorization header",
			headers: http.Header{
				"Authorization": []string{"ApiKey test-key"},
			},
			expected: "test-key",
			wantErr:  false,
		},
		{
			name:     "missing authorization header",
			headers:  http.Header{},
			expected: "",
			wantErr:  true,
		},
		{
			name: "malformed authorization header",
			headers: http.Header{
				"Authorization": []string{"Bearer test-key"},
			},
			expected: "",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAPIKey(tt.headers)

			if tt.wantErr {
				if err == nil {
					t.Fatal("expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}

			if got != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, got)
			}
		})
	}
}
