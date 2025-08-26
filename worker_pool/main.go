package main

import (
	"fmt"
	"worker_pool/service"
)

func main() {

	// worker pool example
	fmt.Println("Starting worker pool example...")
	service.SimpleWorkerRun()
	fmt.Println("Worker pool example finished.")

	// pond example
	fmt.Println("Starting pond example...")
	service.PondRun()
	fmt.Println("Pond example finished.")

}
