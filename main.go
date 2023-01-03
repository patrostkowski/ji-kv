package main

import (
	"log"

	"github.com/patrostkowski/ji-kv/api"
)

func main() {
	log.Println("start")
	httpApi := api.NewServer("0.0.0.0:9991")
	log.Fatal(httpApi.Start())

}
