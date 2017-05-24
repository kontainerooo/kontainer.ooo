#!/bin/sh

USERDIR=/home/vagrant
GOPATH=/var/go/src

# Move daemon config
mkdir /etc/docker
cp /var/go/src/github.com/kontainerooo/kontainer.ooo/scripts/daemon.json /etc/docker/daemon.json

# Add docker repo
add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"

# Add docker gpg key
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | apt-key add -

# Install go and docker dependencies
apt-get update -y
apt-get install linux-headers-3.13.0-116-generic apt-transport-https ca-certificates curl \
  software-properties-common linux-image-extra-$(uname -r) \
  linux-image-extra-virtual docker-ce zip unzip postgresql-client -y
sudo apt-get --no-install-recommends install -y virtualbox-guest-utils

# Install postgres
add-apt-repository "deb http://apt.postgresql.org/pub/repos/apt/ trusty-pgdg main"
wget --quiet -O - https://www.postgresql.org/media/keys/ACCC4CF8.asc | apt-key add -
apt-get update -y
apt-get install postgresql-9.6 -y

# Install go
curl -O https://storage.googleapis.com/golang/go1.8.linux-amd64.tar.gz
tar xvf go1.8.linux-amd64.tar.gz
rm go1.8.linux-amd64.tar.gz
chown -R root:root go
mv go /usr/local
echo "export GOPATH=/var/go
export PATH=\$PATH:/usr/local/go/bin:\$GOPATH/bin" >> $USERDIR/.profile

export GOPATH=/var/go
export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin

# Install protoc
curl -Lo protoc.zip https://github.com/google/protobuf/releases/download/v3.0.0/protoc-3.0.0-linux-x86_64.zip
unzip protoc.zip -d /
rm protoc.zip
go get -u github.com/golang/protobuf/protoc-gen-go
go get -u github.com/golang/protobuf/proto
chown vagrant /bin/protoc

# Install npm
rm -rf $USERDIR/.nvm && git clone https://github.com/creationix/nvm.git $USERDIR/.nvm && (cd $USERDIR/.nvm && git checkout `git describe --abbrev=0 --tags`) && . $USERDIR/.nvm/nvm.sh
echo ". ~/.nvm/nvm.sh" >> $USERDIR/.profile

# Set localhost to have hostname postgres (will change later)
echo "127.0.0.1 postgres" >> /etc/hosts

# Create postgres user
sudo -u postgres bash -c "psql -c \"CREATE USER kroo WITH PASSWORD 'kroo';\""

# Copy config file
cp /var/go/src/github.com/kontainerooo/kontainer.ooo/config.json.sample /var/lib/kontainerooo/config.json

chown -R vagrant /var/go
chown -R vagrant /home/vagrant/.nvm
chown -R vagrant /var/lib/kontainerooo
