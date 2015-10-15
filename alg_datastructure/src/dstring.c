//
// Created by akun on 15-10-14.
//

#include "dstring.h"
#include "string.h"
#include "stdlib.h"
#include "stdio.h"

/* create a str from normal c style string */
str strNewLen(const char *origin, size_t len){
    str s;
    char *space;
    size_t head_size = sizeof(struct str_t);

    if(len == 0){
        if(origin == NULL)
            len = 0;
        else
            len = strlen(origin);
    }
    /* init space */
    space =(char *)malloc(sizeof(char) * (head_size + len + 1));
    s = space + head_size;
    /* init head of str */
    ((struct str_t*)space)->len = len;
    ((struct str_t*)space)->cap = len;
    /* init data field */
    memset(s, 0, len);
    memcpy(s, origin, len);
    s[len] = '\0';

    return s;
}

str strAddCap(str s, size_t added_len){
    char *p;
    p = (char *)malloc(sizeof(struct str_t) + strLength(s) + 1 + added_len);
    if(p == NULL)
        return NULL;
     // TODO
    memset(s + strLength(s) + 1, 0, added_len);
    strHead(s)->cap += added_len;
    free(p);
    return STR_OK;
}

/* Append char * sb to sa */
str strConcat(str sa, const char *sb){
    size_t added_len;

    if(sb == NULL)
        return sa;
    added_len = strlen(sb);
    /* ensure available space more than added_len */
    if((strCap(sa)-strLength(sa)) < added_len){
        /* add cap , return NULL if fail */
        if(strAddCap(sa, added_len) != STR_OK)
            return NULL;
    }
    /* copy data to sa and update len info */
    memcpy(sa + strLength(sa), sb, added_len);
    strHead(sa)->len = strLength(sa) + added_len;
    sa[strLength(sa)] = '\0';
    return sa;
}

/*
 * check if content of in sa's data field equal with sb'sa,
 * different with '=='.
 */
bool strEqual(str sa, const char *sb){
    size_t i;

    if(strlen(sb) == strLength(sa)){
        for(i = 0; i < strLength(sa); i++){
            if(*(sa + i) != *(sb + i))
                return false;
        }
        return true;
    } else {
        return false;
    }
};
