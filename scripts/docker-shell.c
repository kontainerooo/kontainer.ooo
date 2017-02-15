/*
 * docker-shell.c
 *
 * This program is executed as a user's shell on log in. A username
 * comprises of the user's name and the particular container's name.
 * The program attaches to the container.
 *
 * The container has to have /bin/bash installed as it is used as entry
 * point.
 *
 * As the docker daemon currently runs as root user this has to be
 * a setuid binary.
 *
 * $ gcc -Wall -Werror docker-shell.c -o docker-shell.c
 * $ sudo chown root docker-shell
 * $ sudo chmod u+s docker-shell
 *
 */

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/types.h>
#include <unistd.h>
#include <pwd.h>
#include <fcntl.h>
#include <time.h>
#include <stdarg.h>

#define LOG_FILE "/var/log/container-ssh.log"

int log_fd;

void exit_error()
{
    printf("%s\n", "Could not connect to container.");
    exit(EXIT_FAILURE);
}

void open_logfile()
{
    log_fd = open(LOG_FILE, O_WRONLY | O_CREAT | O_APPEND);

    if(log_fd == -1)
        exit_error();
}

void close_logfile()
{
    close(log_fd);
}

void _log(char *message)
{
    char log_buf[250];
    char time_buf[20];
    struct tm *_time;

    time_t now = time(NULL);
    _time = gmtime(&now);

    strftime(time_buf, sizeof(time_buf), "%Y-%m-%d %H:%M:%S", _time);
    sprintf(log_buf, "%s %s\n", time_buf, message);

    write(log_fd, log_buf, sizeof(log_buf));
}

void log_fmt(char *fmt, ...)
{
    char buf[150];

    va_list vl;
    va_start(vl, fmt);

    vsnprintf(buf, sizeof(buf), fmt, vl);

    va_end(vl);

    _log(buf);
}

int main()
{
	struct passwd *pw;
	uid_t uid;

	char cmd[150];

	uid = getuid();
	pw = getpwuid(uid);

	if (pw) {
		// TODO: don't run docker as root
		setuid(0);

        open_logfile();
		log_fmt("START - container: %s\n", pw->pw_name);
        close_logfile();

		sprintf(cmd, "docker exec -it %.128s /bin/bash", pw->pw_name);
		system(cmd);

        open_logfile();
        log_fmt("STOP - container: %s\n", pw->pw_name);
        close_logfile();

		exit(EXIT_SUCCESS);
	}

	exit(EXIT_FAILURE);
}
