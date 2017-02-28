mkfile_path := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

ANGULAR_CLI=ng
CMD_DIRS=$(wildcard cmd/*)
PKG_DIRS=$(wildcard pkg/*)

.PHONY: force

all: fe be all-scripts

fe:
	cd ./frontend && npm install && $(ANGULAR_CLI) build

fe-watch:
	cd ./frontend && $(ANGULAR_CLI) build --watch

all-scripts:
	$(MAKE) -C ./scripts

be: $(PKG_DIRS) $(CMD_DIRS)

$(CMD_DIRS): force
	cd $@ && export GOOS="linux" && go get && go build -o $(mkfile_path)/build/$(notdir $@)

$(PKG_DIRS): force
	cd $@ && go get -t && go test -short && export GOOS="linux" && go build

clean:
	rm -rf build && mkdir build && touch build/.gitkeep
