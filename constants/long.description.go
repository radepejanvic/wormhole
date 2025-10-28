package constants

const (
	LongSetupDesc = `Reads the specified YAML configuration file and sets up 
the virtual environment according to the defined parameters.

Example: 
- wormhole setup --config ./path/to/config/file.yaml`

	LongCreateDesc = `Creates the virtual machines defined in the YAML.

If no VM names are provided, all VMs will be created. 
Otherwise, only the specified VMs will be created.

Example:
- wormhole create
- wormhole create node0 node1`
)
