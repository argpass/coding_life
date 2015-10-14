//
// Created by akun on 15-10-14.
//

#include "dstring.h"
#include "stdlib.h"

/* create a str from normal c style string */
str str_create(const char *origin){
    struct str_info_t *info_t;
    char *str_space;
    unsigned long len;

    len = (unsigned) strlen(origin);
    str_space = malloc(sizeof(char) * (sizeof(struct str_info_t) + len + 1));

}

str str_cat(str sa, str sb);  // concat sa and sb
str str_join(str sep, str str_seq[]);  // concat a str array with sep
