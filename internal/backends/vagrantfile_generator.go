package backends

import (
	_ "embed"
	"fmt"
	"os"
	"text/template"

	"github.com/c12s/wormhole/internal/config"
)

//go:embed vagrantfile_template.rb
var vagrantfileTemplate string

const ipOffset = 100
const vagrantfileDir = "."

type VM struct {
	Name      string
	Hostname  string
	IP        string
	GuestPort int
	HostPort  int
}

type TemplateData struct {
	Box        string
	BoxVersion string
	Memory     int
	CPUs       int
	GUI        bool
	VMs        []VM
}

func newTemplateDataFromConfig(conf *config.Config) *TemplateData {
	vms := make([]VM, conf.VMCount)

	for i := 0; i < conf.VMCount; i++ {
		vms[i] = VM{
			Name:      fmt.Sprintf("%s%d", conf.NameBase, i),
			Hostname:  fmt.Sprintf("%s%d", conf.NameBase, i),
			IP:        fmt.Sprintf("%s%d", conf.IPBase, ipOffset+i),
			GuestPort: conf.GuestPort,
			HostPort:  conf.HostPortBase + i,
		}
	}

	return &TemplateData{
		Box:        conf.OSDistro,
		BoxVersion: conf.OSVersion,
		Memory:     conf.Memory,
		CPUs:       conf.CPUs,
		GUI:        conf.GUI,
		VMs:        vms,
	}
}

func GenerateVagrantfile(conf *config.Config) error {
	var tmpl = template.Must(template.New("vagrantfile").Parse(vagrantfileTemplate))

	var path = fmt.Sprintf("%s/Vagrantfile", vagrantfileDir)
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("create Vagrantfile: %w", err)
	}
	defer file.Close()

	data := newTemplateDataFromConfig(conf)

	err = tmpl.Execute(file, data)
	if err != nil {
		return fmt.Errorf("populate Vagrantfile template: %w", err)
	}

	return nil
}
