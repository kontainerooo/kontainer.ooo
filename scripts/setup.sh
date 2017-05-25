#!/bin/sh

USERDIR=/home/vagrant
GOPATH=/var/go/src

# Install postgres
add-apt-repository "deb http://apt.postgresql.org/pub/repos/apt/ trusty-pgdg main"
wget --quiet -O - https://www.postgresql.org/media/keys/ACCC4CF8.asc | apt-key add -
apt-get update -y
apt-get install postgresql-9.6 git -y

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

# Install netns
mkdir -p /var/lib/kontainerooo/images
curl [REDACTED] -O /var/lib/kontainerooo/images/rootfs.tar

mkdir -p /var/lib/kontainerooo/kmi
cp /var/go/src/github.com/kontainerooo/kontainer.ooo/pkg/kmi/test.kmi /var/lib/kontainerooo/kmi

go get github.com/jessfraz/netns

# Copy config file
cp /var/go/src/github.com/kontainerooo/kontainer.ooo/config.json.sample /var/lib/kontainerooo/config.json

chown -R vagrant /var/go
chown -R vagrant /home/vagrant/.nvm
chown -R vagrant /var/lib/kontainerooo
