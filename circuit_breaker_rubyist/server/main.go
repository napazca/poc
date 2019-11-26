package main

import (
	"fmt"
	"log"
	"net/http"

	circuit "github.com/rubyist/circuitbreaker"
)

var i int

type CircuitBreaker struct {
	P  *circuit.Panel
	CB *circuit.Breaker
}

func (cb *CircuitBreaker) handler(w http.ResponseWriter, r *http.Request) {
	i++
	if cb.CB.Ready() {
		// Breaker is not tripped, proceed
		if i%3 == 0 {
			fmt.Print(i, ". success: ")
			cb.CB.Success()
		} else {
			fmt.Print(i, ". process with failed: ")
			cb.CB.Fail()
		}
	} else {
		// Breaker is in a tripped state.
		fmt.Print(i, ". tripped: ")
	}
	fmt.Println(i, "success", cb.CB.Successes(),
		", failed", cb.CB.Failures(),
		", consec failure", cb.CB.ConsecFailures(),
		", error rate", cb.CB.ErrorRate())
}

func (cb *CircuitBreaker) panelHandler(w http.ResponseWriter, r *http.Request) {
	i++
	if cb.P.Circuits["b1"].Ready() {
		// Breaker is not tripped, proceed
		if i%3 == 0 {
			fmt.Print(i, ". b1 success: ")
			cb.P.Circuits["b1"].Success()
		} else {
			fmt.Print(i, ". b1 process with failed: ")
			cb.P.Circuits["b1"].Fail()
		}
	} else {
		// Breaker is in a tripped state.
		fmt.Print(i, ". b1 tripped: ")
	}
	fmt.Println(i, "success", cb.CB.Successes(),
		", failed", cb.CB.Failures(),
		", consec failure", cb.CB.ConsecFailures(),
		", error rate", cb.CB.ErrorRate())

	if cb.P.Circuits["b2"].Ready() {
		// Breaker is not tripped, proceed
		if i%3 == 0 {
			fmt.Print(i, ". b2 success: ")
			cb.P.Circuits["b2"].Success()
		} else {
			fmt.Print(i, ". b2 process with failed: ")
			cb.P.Circuits["b2"].Fail()
		}
	} else {
		// Breaker is in a tripped state.
		fmt.Print(i, ". b2 tripped: ")
	}
	fmt.Println(i, "success", cb.CB.Successes(),
		", failed", cb.CB.Failures(),
		", consec failure", cb.CB.ConsecFailures(),
		", error rate", cb.CB.ErrorRate())
}

func main() {
	p := circuit.NewPanel()
	b1 := circuit.NewRateBreaker(0.50, 10) // max rate: 1
	b2 := circuit.NewRateBreaker(0.70, 10)

	p.Add("b1", b1)
	p.Add("b2", b2)

	// Creates a circuit breaker that will trip if the function fails 10 times
	cb := circuit.NewRateBreaker(0.50, 10)

	events := cb.Subscribe()
	go func() {
		for {
			e := <-events
			// Monitor breaker events like BreakerTripped, BreakerReset, BreakerFail, BreakerReady
			switch e {
			case circuit.BreakerTripped:
				//log.Println("breaker tripped")
			//case circuit.BreakerReset:
			//	log.Println("breaker reset")
			//case circuit.BreakerFail:
			//	log.Println("breaker fail")
			//case circuit.BreakerReady:
			//	log.Println("breaker ready")
			}
		}
	}()

	circuitBreaker := CircuitBreaker{
		P:  p,
		CB: cb,
	}

	http.HandleFunc("/", circuitBreaker.handler)
	http.HandleFunc("/panel", circuitBreaker.panelHandler)
	log.Fatal(http.ListenAndServe(":50000", nil))

}
