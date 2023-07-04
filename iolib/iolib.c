#include <stdio.h>
#include <stdlib.h>
#define MAXLINE 20

/* read_int: parse next line of stdin to int */
int read_int() {
    int i;

    char buf[MAXLINE];
    if (fgets(buf, MAXLINE, stdin) == NULL) {
        fprintf(stderr, "failed to read line into buffer");
        exit(1);
    }

    return atoi(buf);
}