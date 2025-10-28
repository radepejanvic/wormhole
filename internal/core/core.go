package core

import "github.com/c12s/wormhole/internal/types"

type Backend interface {
	Setup(conf *types.Config) error
	Create(vms ...string) error
	// Stop(vms ...string) error
	// Resume(vms ...string) error
	// Reload(vms ...string) error
	// Destroy(vms ...string) error
}
