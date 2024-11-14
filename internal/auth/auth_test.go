package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name    string
		headers http.Header
		want    string
		wantErr error
	}{
		{
			name: "valid key",
			headers: func() http.Header {
				h := make(http.Header)
				h.Set("Authorization", "ApiKey Test")
				return h
			}(),
			want:    "Test",
			wantErr: nil,
		},
		{
			name: "no key",
			headers: func() http.Header {
				h := make(http.Header)
				return h
			}(),
			want:    "",
			wantErr: errors.New("no authorization header included"),
		}, {
			name: "malformed key",
			headers: func() http.Header {
				h := make(http.Header)
				h.Set("Authorization", "Test")
				return h
			}(),
			want:    "",
			wantErr: errors.New("malformed authorization header"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAPIKey(tt.headers)
			if (err == nil) != (tt.wantErr == nil) ||
				(err != nil && tt.wantErr != nil && err.Error() != tt.wantErr.Error()) {
				t.Errorf("got error %v, want %v", err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
