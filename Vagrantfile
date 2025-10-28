Vagrant.configure("2") do |config|
    config.vm.box = "ubuntu/jammy64"
    config.vm.box_version = "20241002.0.0"

    # Resource configuration 
    config.vm.provider "virtualbox" do |vb|
        vb.memory = 4096        
        vb.cpus = 4 
        vb.gui = false             
    end

    # Common provisioning script for all VMs
    config.vm.provision "shell", name: "install-dependecies", path: "install-dependencies.sh"

    # Create node VMs
    config.vm.define "node0" do |node|
        node.vm.hostname = "node0"
        node.vm.network "private_network", ip: "192.168.56.100"
        node.vm.network "forwarded_port", guest: 6739, host: 11000
    end
end