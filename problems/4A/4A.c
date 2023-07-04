#include <stdio.h>
#include <stdlib.h>
#define MAXLINE 20

int main() {
    int i;

    char buf[MAXLINE];
    if (fgets(buf, MAXLINE, stdin) == NULL) {
        fprintf(stderr, "failed to read line into buffer");
        return 1;
    }

    i = atoi(buf);
    fprintf(stdout, "%s\n", (i % 2 == 0 && i != 2) ? "YES" : "NO");
    return 0;
}