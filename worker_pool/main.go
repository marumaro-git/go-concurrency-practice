package main

import (
	"fmt"
	"worker_pool/service"
)

func main() {

	// worker pool example
	fmt.Println("Starting worker pool example...")
	service.Run()
	fmt.Println("Worker pool example finished.")
}
