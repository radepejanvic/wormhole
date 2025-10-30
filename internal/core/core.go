package core

import (
	"fmt"

	"github.com/c12s/wormhole/internal/backends"
	"github.com/c12s/wormhole/internal/config"
)

type Backend interface {
	Setup(conf *config.Config) error
	Create(vms ...string) error
	Stop(vms ...string) error
	Resume(vms ...string) error
	Reload(vms ...string) error
	ShutDown(vms ...string) error
	Destroy(vms ...string) error
	StartNodes(vms ...string) error
	StopNodes(vms ...string) error
}

func NewBackend(name string) (Backend, error) {
	switch name {
	case "vagrant":
		return backends.NewVagrantBackend()
	default:
		return nil, fmt.Errorf("unsupported backend: %s", name)
	}
}
