package main

import (
	"mars_rover/src/infra/apiServer"
	"mars_rover/src/infra/webServer"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go webServer.RunOn(":6969", &wg)
	go apiServer.RunOn(":4242", &wg)

	wg.Wait()
}
