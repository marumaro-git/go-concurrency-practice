package service

import (
	"fmt"
	"time"

	"github.com/alitto/pond/v2"
)

func PondRun() {
	// Create a pool with limited concurrency
	pool := pond.NewPool(10)

	// Submit 10 tasks
	for i := 0; i < 50; i++ {
		i := i
		pool.Submit(func() {
			time.Sleep(time.Millisecond * 500)
			fmt.Printf("Running task #%d\n", i)
		})
	}

	// Stop the pool and wait for all submitted tasks to complete
	pool.StopAndWait()
}
