mkfile_path := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

ANGULAR_CLI=ng
CMD_DIRS=$(wildcard cmd/*)
PKG_DIRS=$(wildcard pkg/*)

PROTOC=protoc
PROTOC_OPTS="-Imessages/"
PROTOC_DIRS=$(wildcard messages/*)

.PHONY: force

all: fe-test fe proto be all-scripts

fe:
	cd ./frontend && npm install && $(ANGULAR_CLI) build

fe-watch:
	cd ./frontend && $(ANGULAR_CLI) build --watch

fe-test:
	cd ./frontend && karma start karma.conf.js --single-run

all-scripts:
	$(MAKE) -C ./scripts

be: proto $(PKG_DIRS) $(CMD_DIRS)

$(CMD_DIRS): force
	cd $@ && export GOOS="linux" && go get && go build -o $(mkfile_path)/build/$(notdir $@)

$(PKG_DIRS): force
	cd $@ && export GOOS="linux" && go get -t && go test -short && go build

proto: $(PROTOC_DIRS)

$(PROTOC_DIRS): force
	$(PROTOC) $(PROTOC_OPTS) --go_out=plugins=grpc,Mkmi.proto=github.com/kontainerooo/kontainer.ooo/pkg/kmi/pb:pkg/$(basename $(notdir $@))/pb ./messages/$(basename $(notdir $@)).proto

clean:
	rm -rf build && mkdir build && touch build/.gitkeep
