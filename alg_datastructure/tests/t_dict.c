#include "stdio.h"
#include "dict.h"
#include "base.h"


int main(){
    dict di;

    di = dictNew();
    printf("init dict and adr:%d\n", di);
    int data = 99;
    dictAdd("name", "akun");
    dictAdd("age", &data);
    void *rvalue = dictGet(di, "age");

    dictFree(di);
    return 0;
}

