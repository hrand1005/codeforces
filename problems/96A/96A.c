#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define MAXLINE 110

void solve(char *team, int n) {
    int consecutive = 1;
    for (int i = 1; i < n; i++) {
        if (team[i] == team[i-1])
            consecutive++;
        else
            consecutive = 1;
        if (consecutive >= 7) {
            printf("YES\n");
            return;
        }
    }
    printf("NO\n");
}

/* read_string: reads string from stdin into buf with max size n */
char *read_string(char *buf, int n) {
    if (fgets(buf, n, stdin) == NULL) {
        fprintf(stderr, "failed to read line into buffer");
        exit(1);
    }
    buf[strcspn(buf, "\r\n")] = 0;

    return buf;
}

int main(void) {
    char buf[MAXLINE];
    read_string(buf, MAXLINE - 1);
    solve(buf, strlen(buf));
    return 0;
}