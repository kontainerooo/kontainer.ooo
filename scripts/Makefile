OBJECTS = docker-shell
CC = gcc
CFLAGS = -Wall -Wextra -Werror

all: $(OBJECTS)

docker-shell:
	$(CC) $(CFLAGS) docker-shell.c -o docker-shell
	sudo chown root docker-shell
	sudo chmod u+s docker-shell

clean:
	rm -rf $(OBJECTS)
