Vagrant.configure("2") do |config|
    config.vm.box = "{{ .Box }}"
    config.vm.box_version = "{{ .BoxVersion }}"

    # Resource configuration 
    config.vm.provider "virtualbox" do |vb|
        vb.memory = {{ .Memory }}        
        vb.cpus = {{ .CPUs }} 
        vb.gui = {{ .GUI }}             
    end

    # Common provisioning script for all VMs
    config.vm.provision "shell", name: "install-dependecies", path: "install-dependencies.sh"

    # Create node VMs
    {{- range .VMs}}
    config.vm.define "{{ .Name }}" do |node|
        node.vm.hostname = "{{ .Hostname }}"
        node.vm.network "private_network", ip: "{{ .IP }}"
        node.vm.network "forwarded_port", guest: {{ .GuestPort }}, host: {{ .HostPort }}
    end
    {{- end }}
end