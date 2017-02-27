mkfile_path := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

OBJECTS = docker-shell
CC = gcc
CFLAGS = -Wall -Wextra -Werror

all: $(OBJECTS)

docker-shell:
	$(CC) $(CFLAGS) docker-shell.c -o ../build/docker-shell
	sudo chown root ../build/docker-shell
	sudo chmod u+s ../build/docker-shell
