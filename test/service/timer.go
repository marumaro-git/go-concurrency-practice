package service

import (
	"fmt"
	"time"
)

func TimerExample() {
	ticker := time.NewTicker(200 * time.Millisecond)
	defer ticker.Stop()

	timeout := time.After(1 * time.Second)
	counter := 0

	fmt.Println("Starting timer example...")

	for {
		select {
		case <-ticker.C:
			counter++
			fmt.Printf("Tick %d\n", counter)
		case <-timeout:
			fmt.Println("Timeout reached, stopping timer example.")
			return
		}
	}
}

func TimerWithTimeoutExample() {
	ticker := time.NewTicker(200 * time.Millisecond)
	defer ticker.Stop()

	globalTimeout := time.After(1 * time.Second)
	counter := 0

	fmt.Println("Starting timer with timeout example...")

	for {
		processingTimeout := time.After(300 * time.Millisecond)
		time.Sleep(400 * time.Millisecond)
		select {
		case <-ticker.C:
			counter++
			fmt.Printf("Tick %d\n", counter)
		case <-processingTimeout:
			fmt.Println("Processing timeout reached, stopping timer with timeout example.")
			// return
		case <-globalTimeout:
			fmt.Println("Timeout reached, stopping timer with timeout example.")
			return
		}
	}
}
