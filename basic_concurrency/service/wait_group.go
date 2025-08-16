package service

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func WaitGroupSample() {
	fmt.Println("[WaitGroup]Random number is ...")

	var wg sync.WaitGroup
	wg.Add(1)
	f := func() {
		defer wg.Done()
		rand.New(rand.NewSource(time.Now().UnixNano()))

		num := rand.Intn(10)
		fmt.Printf("[WaitGroup]Random number is: %d\n", num)
	}

	go f()

	wg.Wait()
}
