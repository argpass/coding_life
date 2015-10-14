//
// Created by akun on 15-10-13.
//

#include "dblist.h"
#include "stdio.h"


typedef int DataType;
void *copy_value(const void * value){
    DataType *data;

    if((data = malloc(sizeof(DataType))) == NULL)
        return NULL;
    *data = *((DataType *) value);
    return data;
}

int cmp_data(const void *va, const void *vb){
    DataType *_a;
    DataType *_b;

    _a = (DataType *) va;
    _b = (DataType *) vb;
    if(*_a == *_b)
        return 0;
    else
        return -1;
}


int main(void){
    struct ListIterator *iterator;
    struct ListNode *current;
    struct List *list;
    int *p;    // tmp data pointer
    int i;     // tmp variable for iterator
    func_copy_value f_copy;

    // --------------- create list and test basic functions ----------//
    f_copy = copy_value;
    if((list=list_create(f_copy, NULL, NULL)) == NULL)
        return 127;
    install_comp_method(list, cmp_data);
    printf("create done\n");
    int a,b;
    a = 99;
    b = 98;
    printf("compare 2 value:%d\n", list->cmp_value(&a, &b));


    // ---------------- init list (add on head)----------------- //
    for(i=0; i < 10; i++){
        p = malloc(sizeof(int));
        *p = i;
        if((list=list_head_append(list, p)) == NULL)
            return 127;
    }
    printf("len of list:%d\n", list_len(list));


    // ---------------- init list (append at tail)----------------- //
    for(i=10; i < 20; i++){
        p = malloc(sizeof(int));
        *p = i;
        if((list=list_tail_append(list, p)) == NULL)
            return 127;
    }
    printf("len of list:%d\n", list_len(list));


    // ------------------- insert value ---------------------- //
    current = list_locate(list, 18);
    if(current == NULL)
        printf("can't locate to index %d\n", 18);
    else
        printf("locate to 18 and value is %d\n", *(int *)(current->value));
    i = 999;
    if(list_insert(list, current, copy_value(&i), 0) == NULL)
        printf("insert before fail\n");
    i = 888;
    if(list_insert(list, current, copy_value(&i), 1) == NULL)
        printf("insert after fail\n");
    printf("len of list after insert 2 node:%d\n", list_len(list));


    // --------------- search related functions ---------------- //
    i = 18;
    current = list_search(list, &i);
    if(current !=NULL)
        printf("search 18 and node addr:%d\n", current);
    else
        printf("search 18 fail\n");
    printf("index of i is :%d\n", list_index(list, &i));


    // --------------- iter list from head to tail----------------- //
    if((iterator = list_iter_create(list, HEAD_START)) == NULL)
        return 127;
    printf("list data iter from head:\n");
    while ((current = list_iter_next(iterator))){
        p = (int *)current->value;
        printf("%d ", *p);
    }
    printf("\n");
    list_iter_destory(iterator);


    // --------------- iter list from tail to start----------------- //
    if((iterator = list_iter_create(list, TAIL_START)) == NULL)
        return 127;
    printf("list data iter from head:\n");
    while ((current = list_iter_next(iterator))){
        p = (int *)current->value;
        printf("%d ", *p);
    }
    printf("\n");
    list_iter_destory(iterator);

    list_destory(list);
    return 0;
}
