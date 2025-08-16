package service

import (
	"fmt"
	"math/rand"
	"time"
)

func ChannelSample() {
	fmt.Println("[Channel]Random number is ...")

	c := make(chan int)

	f := func(c chan<- int) {
		rand.New(rand.NewSource(time.Now().UnixNano()))

		num := rand.Intn(10)
		c <- num
	}

	go f(c)
	num := <-c

	fmt.Printf("[Channel]Random number is: %d\n", num)
	close(c)
}

func ChannelLoopSample() {

	src := []int{1, 2, 3, 4, 5}
	dst := make([]int, len(src))

	c := make(chan int)

	for _, v := range src {

		go func(v int) {
			result := v * 2
			c <- result
		}(v)
	}

	for i := range src {
		result := <-c
		dst[i] = result
	}

	fmt.Println("Results:", dst)
	close(c)
}
