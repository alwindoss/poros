package main

import (
	"log"

	"github.com/alwindoss/poros"
	"github.com/alwindoss/poros/internal/engine"
)

func main() {
	cfg := &poros.Config{}
	log.Fatal(engine.Run(cfg))
}
