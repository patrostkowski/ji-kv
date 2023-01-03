package main

import (
	"log"
	"sync"

	"github.com/patrostkowski/ji-kv/api"
	"github.com/patrostkowski/ji-kv/quorum"
)

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(2)

	httpApi := api.NewServer("0.0.0.0:9991")
	quorumServer := quorum.NewServer("0.0.0.0:9992")

	go func() {
		log.Fatal(httpApi.Start())
		wg.Done()
	}()

	go func() {
		log.Fatal(quorumServer.Start())
		wg.Done()
	}()

	wg.Wait()
}
