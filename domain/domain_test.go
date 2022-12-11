package domain

import (
	"testing"
)

func TestDuty(t *testing.T) {
	tests := []struct {
		name               string
		inputURLs          []string
		inputParallelCount int
		wantURLs           []string
		wantParallelCount  int
	}{
		{
			name:               "happy path",
			inputURLs:          []string{"test.com"},
			inputParallelCount: 10,
			wantURLs:           []string{"test.com"},
			wantParallelCount:  10,
		},
	}

	for _, tt := range tests {
		d := NewDuty(tt.inputParallelCount, tt.inputURLs)
		if d.ParallelCount() != tt.wantParallelCount {
			t.Errorf("got %d want %d", d.ParallelCount(), tt.wantParallelCount)
		}
		if d.URLs()[0] != tt.wantURLs[0] {
			t.Errorf("got %s want %s", d.URLs()[0], tt.wantURLs[0])
		}
	}
}

func TestHash(t *testing.T) {
	tests := []struct {
		name  string
		input []byte
		want  string
	}{
		{
			name:  "happy path",
			input: []byte("test"),
			want:  "098f6bcd4621d373cade4e832627b4f6",
		},
	}
	for _, tt := range tests {
		h := Hash(tt.input)
		if h != tt.want {
			t.Errorf("test: %s got %s want %s", tt.name, h, tt.want)
		}
	}
}
