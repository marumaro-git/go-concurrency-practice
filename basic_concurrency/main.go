package main

import (
	"basic_concurrency/service"
	"fmt"
)

func main() {

	// cannel example
	fmt.Println("Channel Sample:")
	service.ChannelSample()
	service.ChannelLoopSample()
	fmt.Println()

	// wait group example
	fmt.Println("WaitGroup Sample:")
	service.WaitGroupSample()
	fmt.Println()

	// once example
	fmt.Println("Once Sample:")
	service.OnceSample()
	fmt.Println()

	// timer example
	fmt.Println("Timer Example:")
	service.TimerExample()
	service.TimerWithTimeoutExample()
	fmt.Println()
}
