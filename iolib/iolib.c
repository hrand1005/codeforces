#include <stdio.h>
#include <stdint.h>
#include <stdlib.h>
#include <string.h>

#define MAXLINE 150

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

/* read_int: parse next line of stdin to int array (space separated) */
void read_int_array(int *arr, int n) {
    char buf[MAXLINE];
    char *tok;

    if (fgets(buf, MAXLINE, stdin) == NULL) {
        fprintf(stderr, "failed to read line into buffer");
        exit(1);
    }
    buf[strcspn(buf, "\r\n")] = 0;

    tok = strtok(buf, " ");
    arr[0] = atoi(tok);
    for (int i = 1; i < n; i++) {
        tok = strtok(NULL, " ");
        arr[i] = atoi(tok);
    }
}

/* read_int: parse next line of stdin to int array (space separated) */
void read_uint64_array(uint64_t *arr, uint64_t n) {
    char buf[MAXLINE];
    char *tok;

    if (fgets(buf, MAXLINE, stdin) == NULL) {
        fprintf(stderr, "failed to read line into buffer");
        exit(1);
    }
    buf[strcspn(buf, "\r\n")] = 0;

    tok = strtok(buf, " ");
    arr[0] = atoi(tok);
    for (int i = 1; i < n; i++) {
        tok = strtok(NULL, " ");
        arr[i] = strtoul(tok, NULL, 10);
    }
}

/* read_int: parse next line of stdin to int array (space separated) */
void read_ull_array(unsigned long long *arr, unsigned long long n) {
    char buf[MAXLINE];
    char *tok;

    if (fgets(buf, MAXLINE, stdin) == NULL) {
        fprintf(stderr, "failed to read line into buffer");
        exit(1);
    }
    buf[strcspn(buf, "\r\n")] = 0;

    tok = strtok(buf, " ");
    arr[0] = atoi(tok);
    for (int i = 1; i < n; i++) {
        tok = strtok(NULL, " ");
        arr[i] = strtoull(tok, NULL, 10);
    }
}

/* read_string: reads string from stdin into buf with max size n */
void read_string(char *buf, int n) {
    if (fgets(buf, n, stdin) == NULL) {
        fprintf(stderr, "failed to read line into buffer");
        exit(1);
    }
    buf[strcspn(buf, "\r\n")] = 0;
}