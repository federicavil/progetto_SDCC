package utils

import (
	"fmt"
	"github.com/sony/gobreaker"
	"sync"
	"time"
)

var once sync.Once

var (
	instance *gobreaker.CircuitBreaker
)

func GetCircuitBreaker() *gobreaker.CircuitBreaker {

	if instance == nil {

		once.Do(func() { // <-- atomic, does not allow repeating
			fmt.Println("CREATE CB INSTANCE")
			instance = gobreaker.NewCircuitBreaker(
				gobreaker.Settings{
					Name:        "notification-circuit-breaker",
					MaxRequests: 3,
					Timeout:     40 * time.Second,
					Interval:    100 * time.Second,
					ReadyToTrip: func(counts gobreaker.Counts) bool {
						return counts.ConsecutiveFailures > 3
					},
					OnStateChange: func(name string, from gobreaker.State, to gobreaker.State) {
						fmt.Printf("CircuitBreaker '%s' changed from '%s' to '%s'\n", name, from, to)
					},
				},
			) // <-- thread safe

		})
	}

	return instance
}
