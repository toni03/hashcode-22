
# ***********************************************
### BOXES
# ***********************************************
$os_box = "centos/7"
$os_box_version = "2004.01"
# ***********************************************

# ***********************************************
### VARIABLES
# ***********************************************
$manager_ip = "172.28.200.5"
$hostname = "centos-hashcode-21"
$vm_name = "centos-hashcode-21"
$vm_mem= "2048"
$vm_cpus = 2
# ***********************************************

Vagrant.configure("2") do |config|
  config.vagrant.plugins = "vagrant-vbguest"

  config.vm.boot_timeout = 50000
  config.vm.box = "#{$os_box}"
  if !(defined?($os_box_version)).nil?
    config.vm.box_version = "#{$os_box_version}"
  end
  config.vm.network "private_network", ip: "#{$manager_ip}"

  config.vm.hostname = "#{$hostname}"
  
  config.vm.provider "virtualbox" do |vb|
    # Display the VirtualBox GUI when booting the machine
    vb.gui = false
    vb.name = "#{$vm_name}"
    # Customize the amount of memory on the VM:
    vb.memory = "#{$vm_mem}"
    vb.cpus = $vm_cpus
  end

  config.vbguest.auto_update = true
  config.vbguest.installer_options = { allow_kernel_upgrade: true }

  begin
    config.vm.synced_folder './', "/vagrant", disabled: true
    config.vm.synced_folder './', "/opt/vagrant"
  rescue
    print "Sync Links fail something\n"
    puts 'Sync Links not found'
  end
  
  config.vm.provision "shell", inline: <<-SHELL
    ## Update SO
    sudo yum -y update
    ## Install dependencies
    sudo yum install -y dos2unix nano curl wget unzip yum-utils device-mapper-persistent-data lvm2 \
    screen net-tools 
    ## Install Docker
    sudo yum-config-manager --add-repo https://download.docker.com/linux/centos/docker-ce.repo
    sudo yum-config-manager --enable rhel-7-server-extras-rpms
    sudo yum makecache fast
    sudo yum -y update
    sudo yum -y install docker-ce
    sudo systemctl enable docker
    sudo systemctl start docker
  SHELL

  config.vm.provision "shell", inline: <<-SHELL
    ## VirtualBox Sync Date
    sudo crontab -l > /tmp/date-sync 2>/dev/null
    sudo echo "@reboot VBoxService  --timesync-set-start" >> /tmp/date-sync
    sudo crontab /tmp/date-sync
    sudo VBoxService  --timesync-set-start
  SHELL
  
  config.vm.provision "shell", inline: <<-SHELL
    ## Creating a Swap of 3GB
    SIZE=3
    SIZE_BYTES=$(( 1048576 * $SIZE ))
    dd if=/dev/zero of=/var/tmp/swapfile bs=1024 count=$SIZE_BYTES
    chmod 600 /var/tmp/swapfile
    mkswap /var/tmp/swapfile
    echo "/var/tmp/swapfile swap swap defaults 0 0" >> /etc/fstab
    swapon -a
  SHELL

  config.vm.provision "shell", inline: <<-SHELL
    mkdir /opt/go && cd /opt/go/
    wget https://golang.org/dl/go1.16.linux-amd64.tar.gz

    tar -C /usr/local -xzf go1.16.linux-amd64.tar.gz
    echo "export PATH=\$PATH:/usr/local/go/bin" > /etc/profile.d/go.sh
  SHELL

end
