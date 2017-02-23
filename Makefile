ANGULAR_CLI=ng


all: fe be all-scripts

fe:
	cd ./frontend && $(ANGULAR_CLI) build

fe-watch:
	cd ./frontend && $(ANGULAR_CLI) build --watch

all-scripts:
	$(MAKE) -C ./scripts

be: all-services cmds

all-services: # empty as of now

cmds: # empty as of now
