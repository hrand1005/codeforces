#include <stdio.h>
#include <stdint.h>
#include <stdlib.h>
#include <string.h>

#define MAXLINE 150

void solve(unsigned long long n, unsigned long long m, unsigned long long a) {
    unsigned long long per_row = (n % a == 0) ? (n / a) : (n / a) + 1;
    unsigned long long per_column = (m % a == 0) ? (m / a) : (m / a) + 1;
    printf("%llu\n", per_row * per_column);
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

int main(void) {
    unsigned long long arr[3];
    unsigned long long n, m, a;
    read_ull_array(arr, 3);

    n = arr[0];
    m = arr[1];
    a = arr[2];

    solve(n, m, a);
}