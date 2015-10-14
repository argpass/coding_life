/**
 * ************************ Dynamic String Type , C style *****************
 */

#ifndef ALG_DSTRING_H
#define ALG_DSTRING_H

#ifndef size_t

typedef char *str;

struct __attribute__ ((__packed__)) str_info_t {
    unsigned long len;
};

struct __attribute__ ((__packed__)) str_t {
    unsigned long len;
};


str str_create(const char *);  // create a str from normal c style string
void str_free(str s);
str str_cat(str sa, str sb);  // concat sa and sb
str str_join(str sep, str str_seq[]);  // concat a str array with sep

#endif //ALG_DSTRING_H
