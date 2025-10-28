package backends

import (
	"fmt"

	"github.com/bmatcuk/go-vagrant"
	"github.com/c12s/wormhole/internal/config"
)

type VagrantBackend struct {
	client *vagrant.VagrantClient
}

func NewVagrantBackend() (*VagrantBackend, error) {
	client, err := vagrant.NewVagrantClient(vagrantfileDir)
	if err != nil {
		return nil, fmt.Errorf("initialize Vagrant client: %w", err)
	}

	return &VagrantBackend{client: client}, nil
}

func (v *VagrantBackend) Setup(conf *config.Config) error {
	err := GenerateVagrantfile(conf)
	if err != nil {
		return fmt.Errorf("setup VMs: %w", err)
	}

	return nil
}

func (v *VagrantBackend) Create(vms ...string) error {
	cmd := v.client.Up()
	cmd.Verbose = true
	cmd.AdditionalArgs = vms

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("create VMs: %w", err)
	}
	if cmd.Error != nil {
		return fmt.Errorf("vagrant error: %w", cmd.Error)
	}

	return nil
}

func (v *VagrantBackend) Stop(vms ...string) error {
	cmd := v.client.Suspend()
	cmd.Verbose = true
	cmd.AdditionalArgs = vms

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("stop VMs: %w", err)
	}
	if cmd.Error != nil {
		return fmt.Errorf("vagrant error: %w", cmd.Error)
	}

	return nil
}

func (v *VagrantBackend) Resume(vms ...string) error {
	cmd := v.client.Resume()
	cmd.Verbose = true
	cmd.AdditionalArgs = vms

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("resume VMs: %w", err)
	}
	if cmd.Error != nil {
		return fmt.Errorf("vagrant error: %w", cmd.Error)
	}

	return nil
}

func (v *VagrantBackend) Reload(vms ...string) error {
	cmd := v.client.Reload()
	cmd.Verbose = true
	cmd.AdditionalArgs = vms

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("resume VMs: %w", err)
	}
	if cmd.Error != nil {
		return fmt.Errorf("vagrant error: %w", cmd.Error)
	}

	return nil
}

func (v *VagrantBackend) ShutDown(vms ...string) error {
	cmd := v.client.Halt()
	cmd.Verbose = true
	cmd.AdditionalArgs = vms

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("shut down VMs: %w", err)
	}
	if cmd.Error != nil {
		return fmt.Errorf("vagrant error: %w", cmd.Error)
	}

	return nil
}

func (v *VagrantBackend) Destroy(vms ...string) error {
	cmd := v.client.Destroy()
	cmd.Verbose = true
	cmd.AdditionalArgs = vms

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("destroy VMs: %w", err)
	}
	if cmd.Error != nil {
		return fmt.Errorf("vagrant error: %w", cmd.Error)
	}

	return nil
}
