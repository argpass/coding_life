/**
 * Desc:确定当前系统的字节序是大端对齐还是小端对齐
 */

#include <sys/errno.h>
#include "stdio.h"


typedef union {
    short short_value;
    char ch[sizeof(short)];
} short_un;

int main(void){
    short_un un;
    un.short_value = 0x0102;

    if(un.ch[0] == 1){
        printf("big-endian\n");
    } else {
        printf("little-endian\n");
    }
    printf("sizeof 0x0102:%d", (int)(sizeof(errno)));
}