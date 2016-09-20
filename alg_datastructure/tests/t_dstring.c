#include <string.h>
#include "stdlib.h"
#include "stdio.h"
#include "dstring.h"


int main(void){
    /* ***************** init *********************/
    str s = strNew("hello");
    str b = strNewLen("go", 0);
    printf("init s addr:>%d<\n", s);
    printf("init s:>%s<\n", s);
    printf("init b:>%s<\n", b);
    printf("len of b is:%d\n", strLength(b));
    printf("len of s is:%d\n", strLength(s));
    printf("cap of b is:%d\n", strCap(b));
    printf("cap of s is:%d\n", strCap(s));

    /******************* dup **********************/
    str c = strDup(s);
    printf("dup c addr:>%d<\n", c);
    printf("addr of c not equal s:%d\n", s != c);
    printf("dup c:>%s<\n", c);
    printf("value of c is same with s's :%d\n", strEqual(c, s));

    /****************** modify *********************/
    s = strConcat(s, "!");
    printf("after concat s addr:>%d<\n", s);
    printf("after concat s:>%s<\n", s);
    printf("after concat len of s:>%d<\n", strLength(s));
    printf("after concat cap of s:>%d<\n", strCap(s));

    /* **************** join *********************/
    char *seq[] = {"abc", "efggggggggggggggggggggggggggg"};
    printf("sizeof seq:%d\n", sizeof(seq)/sizeof(char *));
    str d = strJoin(",", seq, sizeof(seq)/sizeof(char *));
    printf("joind created str:%s\n", d);

    /****************** release ******************/
    printf("this is a test string");
    strFree(s);
    strFree(b);
    strFree(c);
    strFree(d);
    return 0;
}
