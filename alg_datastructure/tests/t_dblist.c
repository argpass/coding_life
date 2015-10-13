//
// Created by akun on 15-10-13.
//

#include "t_dblist.h"
#include "stdio.h"

int main(void){
    int *pInt = malloc(sizeof(int));
    if(pInt && 1)
        printf("hello world%d\n", pInt || 8);
    free(NULL);
}
