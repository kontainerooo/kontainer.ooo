# kontainer.ooo

|Master|Develop|
|---|---|
|[![Build Status](https://travis-ci.org/kontainerooo/kontainer.ooo.svg?branch=master)](https://travis-ci.org/kontainerooo/kontainer.ooo)|[![Build Status](https://travis-ci.org/kontainerooo/kontainer.ooo.svg?branch=develop)](https://travis-ci.org/kontainerooo/kontainer.ooo)|

## Building and running kontainer.ooo (Development environment)
The build consists of two build processes.

1. Building the go backend
1. Building the angular frontend

It is recommended that you build the project using the provided vagrant machine or at least a linux machine.

### 1. Building the go backend
The backend requires `golang` and the `protobuf compiler` to be installed. Additionally the go packages `github.com/golang/protobuf/protoc-gen-go` and `github.com/golang/protobuf/proto` need to be installed.

To be able to run **kontainer.ooo** a `postgresql` instance and `iptables` are required.

The `Vagrantfile` already has these dependencies configured.

1. Run `vagrant up` to provision and create the virtual machine.
1. SSH into the vagrant machine with `vagrant ssh`
1. Inside the vagrant machine the repository is mounted in `/var/go/src/github.com/kontainerooo/kontainer.ooo/`
1. Run `make be` inside the directory to build the backend.

### 2. Building the frontend
The frontend can be built independently from the backend. Required are `node` and `npm`.
Additionally the global npm module `@angular/cli` is required.

Please note that if you ever run `npm install` on your host machine then you need to run `npm rebuild` inside the vagrant machine if your host is not a 64bit linux.

After the first time running `vagrant up` please run `nvm install stable` and `npm install -g @angular/cli`

1. Run `npm install` inside `/frontend/`
1. Run `make fe` to build the frontend
1. Run `npm start` to serve the frontend (`npm start` runs `ng serve --host 0.0.0.0` in order to work in the vagrant machine)
