# Wormhole CLI Documentation
Tool for starting local development with VM-s running agents etc. and control plane communicating with with nodes.
## Features
- VM lifecycle management: setup, create, stop, resume, reload, shutdown, destroy
- Node management: start-nodes, stop-nodes
## Prerequisites
Ensure you have the following installed: 
- [Go](https://go.dev/doc/install) (version 1.16 or higher)
- [Git](https://git-scm.com/install/)
- [Docker](https://docs.docker.com/engine/install/) (for running the complete environment)
- [Vagrant](https://developer.hashicorp.com/vagrant/docs/installation)
- [VirtualBox](https://www.virtualbox.org/wiki/Downloads)
Additionally, you need to clone and set up the tools repository from [https://github.com/c12s/tools](https://github.com/c12s/tools), which includes scripts to pull all other repositories and start Docker containers.
## Installation
1. Clone the tools repository: 
```bash
git clone https://github.com/c12s/tools.git
```
2. Navigate to the tools directory and follow the setup instructions in the README.md file of the tools repository to pull all necessary repositories and start Docker containers:
```bash
cd tools
./install.sh
./control_plane_start.sh
```
3. Clone the cockpit repository in the same parent directory where the tools repository is located:
```bash
cd ..
git clone https://github.com/radepejanvic/wormhole.git
```
4. Navigate to the cockpit project directory:
```bash
cd wormhole
```
5. Build the CLI: 
```bash
go build -o wormhole
```
6. Add the executable to your PATH:
```bash
export PATH=$PATH:$(pwd)
```
## Command reference
### VM lifecycle management
Setup - reads the YAML and generates Vagrantfile
```bash
wormhole setup --config ./config.yaml
```
Create - creates VMs (one, more or all if nothing specified)
```bash 
wormhole create 
wormhole create node0
```
Stop - suspends VMs (one, more or all if nothing specified)
```bash
wormhole stop
wormhole stop node0
```
Resume - resumes VMs (one, more or all if nothing specified)
```bash 
wormhole resume
wormhole resume node0
```
Reload - recreates VM reading Vagrantfile again (one, more or all if nothing specified)
```bash
wormhole reload
wormhole reload node0
```
Shutdown - gracefully powers off VMs (one, more or all if nothing specified)
```bash
wormhole shutdown
wormhole shutdown node0
```
Destroy - deletes VMs and all of their data (one, more or all if nothing specified)
```
wormhole destroy
wormhole destroy node0 
```
### Node management
Start nodes - runs the `/vagrant/node_start.sh` remotely for specified VMs
```bash 
wormhole start-nodes node0 node1
```
Stop nodes - runs the `/vagrant/nodes_stop.sh` remotely for specified VMs
```bash
wormhole stop-nodes node0 node1
```
