package main

import (
	"log"

	"github.com/KadirbekSharau/Byte/configs"
)

func main() {
	if err := configs.Init(); err != nil {
		log.Fatal("%s", err.Error())
	}

	
}