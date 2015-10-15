
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
    char *head = (char *)strHead(s);
    size_t head_len = sizeof(struct str_t);
    size_t old_len = strLength(s);
    size_t _raw_str_space_len = head_len + old_len + 1;

    /* allocate new space */
    p = (char *)realloc(strHead(s), _raw_str_space_len + added_len);
    if(p == NULL)
        return NULL;
    if(p != head){
        printf("new space and free old\n");
        // TODO why>if not to free result in mem leak ,but if do follow cast err
//        free(head);
    }

    /* let s point to new space's data field
     * from now on the s is a new str
     */
    s = p + sizeof(struct str_t);
    /* initial unused field */
    memset(s + strLength(s) + 1, 0, added_len);
    strHead(s)->cap += added_len;
    return s;
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
        if(strAddCap(sa, added_len) == NULL)
            return NULL;
    }
    printf("->:%s\n", sa);
    printf("len->:%d\n", strLength(sa));
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
