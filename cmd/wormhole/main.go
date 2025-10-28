package main

import (
	"log"

	"github.com/c12s/wormhole/internal/core"
)

func main() {
	conf, err := core.LoadFromYaml("config.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	backend, err := core.NewBackend("vagrant")
	if err != nil {
		log.Fatalf("Vagrant Backend initialization failed: %v", err)
	}

	if err := backend.Setup(conf); err != nil {
		log.Fatalf("Setup failed: %v", err)
	}

	if err := backend.Create(); err != nil {
		log.Fatalf("Create failed: %v", err)
	}

	if err := backend.Stop("node0"); err != nil {
		log.Fatalf("Stop failed: %v", err)
	}

	if err := backend.Resume("node0"); err != nil {
		log.Fatalf("Resume failed: %v", err)
	}

	if err := backend.Reload("node0"); err != nil {
		log.Fatalf("Reload failed: %v", err)
	}

	if err := backend.ShutDown("node0"); err != nil {
		log.Fatalf("Shut down failed: %v", err)
	}

	if err := backend.Destroy("node0"); err != nil {
		log.Fatalf("Destroy failed: %v", err)
	}
}
