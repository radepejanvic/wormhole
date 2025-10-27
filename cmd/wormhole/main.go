package main

import (
	"fmt"
	"log"

	config "github.com/c12s/wormhole/internal/config"
)

func main() {
	conf, err := config.NewFromYaml("test.yaml")
	if err != nil {
		log.Fatalf("%v", err)
	}

	fmt.Printf("%q", conf.OSDistro)
}
