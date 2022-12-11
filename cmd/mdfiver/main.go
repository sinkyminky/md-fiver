package main

import (
	"flag"
	"mdFiver/application"
	"mdFiver/domain"
	"mdFiver/infrastructure/http"
	http2 "net/http"
	"time"
)

func main() {
	parallelCount := flag.Int("parallel", 10, "parallel count")
	flag.Parse()
	args := flag.Args()
	duty := domain.NewDuty(*parallelCount, args)
	hc := http.NewClient(&http2.Client{Timeout: time.Second * 3})
	s := application.NewService(hc)
	s.Process(duty)
}
