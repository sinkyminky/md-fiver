package http

import (
	"net/http"
	"testing"
	"time"
)

func TestGet(t *testing.T) {
	tests := []struct {
		name string
		url  string
		err  bool
	}{
		{
			name: "happy path",
			url:  "http://google.com",
			err:  false,
		},
		{
			name: "err with invalid url",
			url:  "google.com",
			err:  true,
		},
	}

	hc := NewClient(&http.Client{Timeout: 3 * time.Second})
	for _, tt := range tests {
		if _, err := hc.Get(tt.url); (err != nil) != tt.err {
			t.Errorf("test: %s got %v want %v", tt.name, err, tt.err)
		}
	}
}
