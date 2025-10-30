package constants

const (
	LongSetupDesc = `Reads the specified YAML configuration file and sets up 
the virtual environment according to the defined parameters.

Example: 
- wormhole setup --config ./path/to/config/file.yaml`

	LongCreateDesc = `Creates the virtual machines specified by name. 
If no names are provided, all virtual machines in the environment will be created.

Example:
- wormhole create
- wormhole create node0 node1`

	LongStopDesc = `Stops the virtual machines specified by name. 
If no names are provided, all running virtual machines in the environment will be stopped.

Example:
- wormhole stop               
- wormhole stop node0 node1`

	LongResumeDesc = `Resumes the virtual machines specified by name. 
If no names are provided, all stopped virtual machines in the environment will be resumed.

Example:
- wormhole resume               
- wormhole resume node0 node1`

	LongReloadDesc = `Reloads the specified virtual machines, applying any updated configuration. 
If no VM names are provided, all virtual machines in the environment will be reloaded.

Example:
- wormhole reload               
- wormhole reload node0 node1`

	LongShutDownDesc = `Stops the specified virtual machines gracefully. 
If no VM names are provided, all virtual machines in the environment will be stopped.

Example:
- wormhole stop
- wormhole stop node0 node1`

	LongDestroyDesc = `Permanently destroy one or more virtual machines and clean up all associated resources. 
If no names are provided, all virtual machines defined in the environment will be destroyed. 
Use with caution, as this operation cannot be undone.

Example:
- wormhole destroy
- wormhole destroy node0 node1`

	LongStartNodesDesc = `Start the star service on one or more virtual machines defined in the environment. 
If no names are provided, the service will be started on all available nodes. 
This command runs the node_start.sh script inside each VM, initializing star and related components.

Example:
- wormhole start-nodes
- wormhole start-nodes node0 node1`

	LongStopNodesDesc = `Stop the star service on one or more virtual machines. 
If no names are provided, the service will be stopped on all running nodes. 
This command reads the star PID from each VM and terminates the process safely running the node_stop.sh.

Example:
- wormhole stop-nodes
- wormhole stop-nodes node0 node1`
)
