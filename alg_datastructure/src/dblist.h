/**
 * ******************* double linked list ************************
 */

#ifndef ALG_T_DBLIST_H
#define ALG_T_DBLIST_H

#include "stdlib.h"
#include "base.h"


/*
 * Node, List, and Iterator define
 */
struct ListNode {
    struct ListNode *prev;
    struct ListNode *next;
    void *value;
};



typedef void *(*func_copy_value)(const void *);
typedef void (*func_free_value)(void *);
typedef int (*func_cmp_value)(const void *value_a, const void *value_b);

typedef struct List {
    struct ListNode *head;
    struct ListNode *tail;
    unsigned long len;

    // Copy value to a new one
    func_copy_value copy_value;

    // Free value point
    func_free_value free_value;

    // Compare two value
    func_cmp_value cmp_value;
};


typedef int Direction;

typedef struct ListIterator {
    struct ListNode *next;
    Direction direction;
};


/* Functions implements as Macros */
#define install_copy_method(l,m) ((l)->copy_value = (m))
#define install_free_method(l,m) ((l)->free_value = (m))
#define install_comp_method(l,m) ((l)->cmp_value = (m))

#define get_copy_method(l) ((l)->copy_value)
#define get_free_method(l) ((l)->free_value)
#define get_comp_method(l) ((l)->cmp_value)

#define list_len(l) ((l)->len)
#define is_empty(l) (((l)->len == 0))

#define first(l) ((l)->head)
#define last(l) ((l)->tail)


/* Functions Definition */
struct List *list_create(func_copy_value cp, func_cmp_value cmp, func_free_value free_value);
struct ListNode *list_create_node(void *value);

struct List *list_tail_append(struct List *, void *value);
struct List *list_head_append(struct List *, void *value);
struct List *list_del_node(struct List *, struct ListNode *);
struct List *list_insert(struct List *, struct ListNode *position, void *value, int after);

struct ListNode *list_search(struct List *, void *key);
struct ListNode *list_locate(struct List *, long index);
long list_index(struct List *, void *key);

struct ListIterator *list_iter_create(struct List *, Direction d);
struct ListNode *list_iter_next(struct ListIterator *);
void list_iter_destory(struct ListIterator *);

void list_rotate(struct List *);

void list_destory(struct List *);


/* Directions for ListIterator */
#define TAIL_START 0
#define HEAD_START 1


#endif //ALG_T_DBLIST_H
