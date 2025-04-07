#include <stdlib.h>
#include <string.h>

int get_string(char **out) {
    const char *msg = "Hello from C!";
    *out = malloc(strlen(msg) + 1);
    if (*out == NULL) return -1;
    strcpy(*out, msg);
    return 0;
}