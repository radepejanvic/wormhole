package main

import (
	"log"

	"github.com/c12s/wormhole/internal/backends"
	"github.com/c12s/wormhole/internal/core"
)

func main() {
	conf, err := core.LoadFromYaml("config.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	backend, err := backends.NewVagrantBackend()
	if err != nil {
		log.Fatalf("Vagrant Backend initialization failed: %v", err)
	}

	if err := backend.Setup(conf); err != nil {
		log.Fatalf("Setup failed: %v", err)
	}

	if err := backend.Create(); err != nil {
		log.Fatalf("Create failed: %v", err)
	}
}
