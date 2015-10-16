/*************************************************************************
    > Author: zkchen
    > Mail: zhongkunchen@126.com 
    > Created Time: 2015年10月16日 星期五 13时38分08秒
 ************************************************************************/
#ifndef ALG_DICT_H
#define ALG_DICT_H


#include"stdlib.h"
#include "base.h"


#define makeSizeMask(x) (255 << 8*(x-1) | 255)

struct DictEntry{
    void *key;
    void *value;
    struct DictEntry *next;
};


typedef struct{
    struct DictEntry **slot;
    unsigned long sizeMask;
} *dict;


/* public functions */
dict dictNew(void);
void dictAdd(const void *key, void *value);
void dictFree(dict d);
void *dictGet(dict d, void *key);
bool dictHasKey(dict d, void *key);




#endif
