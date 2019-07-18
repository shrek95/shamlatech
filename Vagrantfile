$shell = <<-SHELL

echo "SHELL IS HIT"
apt-get update

SHELL


Vagrant.configure("2") do |config|
 config.vm.define "shreyubuntu16" do |shreyubuntu16|
    shreyubuntu16.vm.box = "bento/ubuntu-16.04"
    shreyubuntu16.ssh.username = 'vagrant'
    shreyubuntu16.ssh.password = 'vagrant'
    shreyubuntu16.ssh.insert_key = 'true' 
    shreyubuntu16.vm.network "forwarded_port", guest: 80, host: 8080, host_ip: "127.0.0.1"
    shreyubuntu16.vm.network "private_network", ip: "10.1.1.11", network: "24"
    shreyubuntu16.vm.provision "shell", inline: $shell
    shreyubuntu16.vm.provider "virtualbox" do |vb|
      vb.gui = false
      vb.memory = "2048"
      vb.cpus = "1"
    end 
 end
end
