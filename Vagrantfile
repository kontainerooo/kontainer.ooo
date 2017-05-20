# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure(2) do |config|

  # Ubuntu is fine
  config.vm.box = "ubuntu/trusty64"

  # Accessible in private network
  config.vm.network "private_network", ip: "172.30.1.2"

  # GRPC Port
  config.vm.network "forwarded_port", guest: 8082, host: 8082

  # Websocket Port
  config.vm.network "forwarded_port", guest: 8081, host: 8081

  # Docker Port
  # In the future docker will not listen to tcp. Sorry.
  config.vm.network "forwarded_port", guest: 2375, host: 9075

  # Fronted Port
  config.vm.network "forwarded_port", guest: 4200, host: 4200

  # Sync kontainerooo in gopath
  config.vm.synced_folder "./", "/var/go/src/github.com/kontainerooo/kontainer.ooo"

  # Run setup script
  config.vm.provision "shell", path: "./scripts/setup.sh"

  config.vm.provider "virtualbox" do |v|
    v.memory = 2048
    v.cpus = 2
  end

end
