package main

import (
	"basic_concurrency/service"
)

func main() {

	// cannel example
	service.ChannelSample()
	service.ChannelLoopSample()

	// wait group example
	service.WaitGroupSample()

}
