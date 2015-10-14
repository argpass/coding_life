//
// Created by akun on 15-10-14.
//

#include <string.h>
#include "stdlib.h"
#include "stdio.h"

struct __attribute__ ((__packed__)) my16 {
    u_int64_t a;
    u_int32_t b;
};

int main(void){
    char *p = malloc(sizeof(100));
    strcpy(p, "abcdefg");
    p[2] = '\0';
    printf("len %d\n", strlen(p));
    return 0;
}

