#include <string.h>
#include "stdlib.h"
#include "stdio.h"

long hash_key(void *key, long mask){
    int *p;

    p = (int *)key;
    printf("(%d),", *p);
    return (*p) & mask;
}

struct my_t{
    unsigned len;
};


int main(){
    int a = 9999;
    int b = 55;
    int c = 9999;
    printf("hash:%d, %d, %d\n", hash_key(&a, 255), hash_key(&b, 255), hash_key(&c, 255));
    printf("hash:%d, %d, %d\n", hash_key("abc", 255), hash_key("a", 255), hash_key("eciyyyyk", 255));

    struct my_t t1;
    t1.len = 99998;
    struct my_t t2;
    t2.len = 8;
    struct my_t t3;
    t3.len = 99998;
    printf("hash:%d, %d, %d\n", hash_key(&t1, 255), hash_key(&t2, 255), hash_key(&t3, 255));
    return 0;
}
