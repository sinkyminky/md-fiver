package application

import (
	"fmt"
	"mdFiver/domain"
	"mdFiver/infrastructure/http"
	"strings"
	"sync"
)

type service struct {
	httpClient http.Client
}

type Service interface {
	Process(duty *domain.Duty)
}

func NewService(hc http.Client) Service {
	return &service{
		httpClient: hc,
	}
}

func (s service) Process(duty *domain.Duty) {
	chunks := prepareChunks(duty)
	for _, chunk := range chunks {
		var wg sync.WaitGroup
		for _, url := range chunk {
			wg.Add(1)
			go func(url string) {
				defer wg.Done()
				nu := prepareURL(url)
				r, err := s.httpClient.Get(nu)
				if err != nil {
					fmt.Println("url: ", url, " error: ", err)
					return
				}
				fmt.Println(nu, " ", domain.Hash(r))
			}(url)
		}
		wg.Wait()
	}
}

func prepareChunks(duty *domain.Duty) [][]string {
	lenURLs := len(duty.URLs())
	parallelCount := duty.ParallelCount()
	chunkCount := lenURLs / parallelCount
	chunks := make([][]string, 0)
	for i := 0; i <= chunkCount; i++ {
		if i*parallelCount == lenURLs {
			break
		}
		if i == chunkCount {
			chunks = append(chunks, duty.URLs()[i*parallelCount:])
		} else {
			chunks = append(chunks, duty.URLs()[i*parallelCount:(i+1)*parallelCount])
		}
	}
	return chunks
}

func prepareURL(s string) string {
	if !strings.HasPrefix(s, "http://") && !strings.HasPrefix(s, "https://") {
		s = "http://" + s
	}
	return s
}
