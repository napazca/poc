package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/afex/hystrix-go/hystrix"
)

var i int

func main() {
	hystrix.ConfigureCommand("my_command", hystrix.CommandConfig{
		// How long to wait for command to complete, in milliseconds
		Timeout: 50000,

		// MaxConcurrent is how many commands of the same type
		// can run at the same time
		MaxConcurrentRequests: 300,

		// VolumeThreshold is the minimum number of requests
		// needed before a circuit can be tripped due to health
		RequestVolumeThreshold: 10,

		// SleepWindow is how long, in milliseconds,
		// to wait after a circuit opens before testing for recovery
		SleepWindow: 1000,

		// ErrorPercentThreshold causes circuits to open once
		// the rolling measure of errors exceeds this percent of requests
		ErrorPercentThreshold: 50,
	})

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":50000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	i++
	resultCh := make(chan []byte)
	errCh := hystrix.Go("my_command", func() error {
		if i%3 == 0 {
			resultCh <- []byte("success")
			return nil
		} else {
			return fmt.Errorf("%d. error", i)
		}
	}, nil)

	select {
	case res := <-resultCh:
		log.Println("success to get response from sub-system:", string(res))
		w.WriteHeader(http.StatusOK)
	case err := <-errCh:
		log.Println("failed to get response from sub-system:", err.Error())
		w.WriteHeader(http.StatusServiceUnavailable)
	}
}
