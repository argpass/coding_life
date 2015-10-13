//
// Created by akun on 15-10-13.
//

#include "dblist.h"

/* default util function for list */
int list_default_cmp_value(void *va, void *vb){
    return va == vb;
}

/* Functions Implements */

/* create a list instance */
struct List *list_create(void){
    struct List *list;

    if((list = malloc(sizeof(struct List))) == NULL)
        return NULL;
    list->copy_value = list->free_value = NULL;
    list->cmp_value = list_default_cmp_value;
    list->len = 0;
    list->head = list->tail = NULL;
    return list;
}

/* create a node to list */
struct ListNode *list_create_node(void *value){
    struct ListNode *node;

    if((node = malloc(sizeof(*node))) == NULL)
        return NULL;
    node->value = value;
    node->next = node->prev = NULL;
    return node;
}

/* append a data to list on tail */
struct List *list_tail_append(struct List *list, void *value){
    struct ListNode *node;

    node = list_create_node(value);
    if(list_len(list) == 0){
        list->head = list->tail = node;
        node->next = node->prev = NULL;
        list->len++;
    } else {
        node->prev = list->tail;
        node->next = NULL;
        list->tail->next = node;
        list->tail = list->tail->next;
        list->len++;
    }
    return list;
}

/* delete tail node and return deleted value */
struct List *list_tail_pop(struct List *list, void *value){
    struct ListNode *temp;

    if(list_len(list) == 0){
        value = NULL;
    } else {
        temp = list->tail;
        *value = *temp;
        list->tail = temp->prev;
        list->len--;
    }
    free(temp);
    return list;
}


struct List *list_head_append(struct List *list, void *value){
    struct ListNode *node;

    if((node = list_create_node(value)) == NULL)
        return NULL;
    if(is_empty(list)){
        list->tail = list->head = node;
        node->next = node->prev = NULL;
    } else {
        node->next = list->head;
        node->next->prev = node;
        list->head = node;
    }
    return list;
}

struct List *list_head_pop(struct List *list, void *value){
    list->
}

/* delete a node from list , node must be in list */
struct List *list_del_node(struct List *list, struct ListNode *node){

    if(is_empty((list)))
        return NULL;

    if(list_len(list) == 1){
        list->head = list->tail = NULL;
    } else {
        node->prev->next = node->next;
        node->next->prev = node->prev;
    }

    list->len--;
    list->free_value(node->value);
    free(node);
    return list;
}

/* insert a value to list position must not be null */
struct List *list_insert(struct List *list, struct ListNode *position, void *value, int after){
    struct ListNode *node;

    if((node == list_create_node(value)) == NULL)
        return NULL;
    if(position == NULL)
        return NULL;
    if(after){
        if(position == list->tail){
            node->prev = list->tail;
            node->prev->next = node;
            list->tail = node;
        } else {
            node->next = position->next;
            node->prev = position;
            node->prev->next = node;
            node->next->prev = node;
        }
    } else {
        if(position == list->head){
            node->next = list->head;
            node->next->prev = node;
            list->head = node;
        } else {
            node->next = position;
            node->prev = position->prev;
            node->prev->next = node;
            node->next->prev = node;
        }
    }
    list->len++;
    return list;
}

struct ListNode *list_search(struct List *list, void *key){
    struct ListNode *node;
    struct ListIterator *iter;

    if(is_empty(list))
        return NULL;

    if((iter = list_iter_create(list, HEAD_START)) == NULL)
        return NULL;
    while ((node = list_iter_next(iter))){
        if(list->cmp_value(key, node->value) == 0){
            list_iter_destory(iter);
            return node;
        }
    }

    list_iter_destory(iter);
    return NULL;
}

struct ListNode *list_locate(struct List *list, long index){
    struct ListNode *current;
    struct ListIterator *iterator;
    long i = -1;

    if(index < 0 || index >= list_len(list))
        return NULL;

    if((iterator = list_iter_create(list)) == NULL)
        return NULL;
    while((current = list_iter_next(iterator))){
        i++;
        if(i == index)
            break;
    }
    list_iter_destory(iterator);
    return current;

}

long list_index(struct List *list, void *key){
    struct ListNode *current;
    struct ListIterator *iterator;
    long err_ret = -1;
    long i = -1;

    if(is_empty(list))
        return err_ret;

    if((iterator = list_iter_create(list)) == NULL)
        return -1;
    while((current = list_iter_next(iterator))){
        i++;
        if(list->cmp_value(current->value, key) == 0){
            list_iter_destory(iterator);
            return i;
        }
    }
    list_iter_destory(iterator);
    return err_ret;
}

struct ListIterator *list_iter_create(struct List *list, Direction d){
    struct ListIterator *iter;

    if((iter = malloc(sizeof(*iter))) == NULL)
        return NULL;
    iter->direction = d;
    if(d == HEAD_START)
        iter->next = list->head;
    else
        iter->next = list->tail;
    return iter;
}

struct ListNode *list_iter_next(struct ListIterator *iter){
    struct ListNode *node;

    if(iter->next == NULL)
        return NULL;

    node = iter->next;
    if(iter->direction == HEAD_START)
        iter->next = node->next;
    else
        iter->next = node->prev;
    return node;
}

void list_iter_destory(struct ListIterator *iterator){
    free(iterator);
}

/* rotate the tail to head */
void list_rotate(struct List *list){
    struct ListNode *node;

    if(is_empty(list))
        return;

    node = list->tail;
    node->next = list->head;
    node->next->prev = node;
    list->head = node;
}

void list_destory(struct List *list){
    free(list);
}
