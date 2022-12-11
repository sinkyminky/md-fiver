package application

import (
	"mdFiver/domain"
	"testing"
)

func TestPrepareURL(t *testing.T) {

	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "without scheme",
			input: "test.com",
			want:  "http://test.com",
		},
		{
			name:  "with http scheme",
			input: "http://test.com",
			want:  "http://test.com",
		},
		{
			name:  "with https scheme",
			input: "https://test.com",
			want:  "https://test.com",
		},
		{
			name:  "with www",
			input: "www.test.com",
			want:  "http://www.test.com",
		},
	}

	for _, tt := range tests {
		pu := prepareURL(tt.input)
		if pu != tt.want {
			t.Errorf("test: %s got %s want %s", tt.name, pu, tt.want)
		}
	}
}

func TestPrepareChunks(t *testing.T) {
	tests := []struct {
		name  string
		input *domain.Duty
		want  [][]string
	}{
		{
			name:  "parallel count bigger than url count",
			input: domain.NewDuty(10, []string{"test0.com", "test1.com"}),
			want:  [][]string{{"test0.com", "test1.com"}},
		},
		{
			name:  "parallel count lower than url count",
			input: domain.NewDuty(2, []string{"test0.com", "test1.com", "test2.com", "test3.com", "test4.com"}),
			want:  [][]string{{"test0.com", "test1.com"}, {"test2.com", "test3.com"}, {"test4.com"}},
		},
		{
			name:  "parallel count equal to url count",
			input: domain.NewDuty(5, []string{"test0.com", "test1.com", "test2.com", "test3.com", "test4.com"}),
			want:  [][]string{{"test0.com", "test1.com", "test2.com", "test3.com", "test4.com"}},
		},
	}

	for _, tt := range tests {
		chunks := prepareChunks(tt.input)
		for i, chunk := range chunks {
			for v, url := range chunk {
				if tt.want[i][v] != url {
					t.Errorf("test: %s got %s want %s at %dth chunk %dth index", tt.name, chunks[i][v], tt.want[i][v], i, v)
				}
			}
		}
	}
}
