def partition(A, p, r):
    key = A[r]    
    j = p
    for i in range(p, r, 1):
        if A[i] < key:
            A[i], A[j] = A[j], A[i] 
            j += 1
    A[j], A[r] = A[r], A[j]
    return j


def quicksort(A, p, r):
    if p < r:
        q = partition(A, p, r)
        quicksort(A, p, q - 1)
        quicksort(A, q + 1, r)


if __name__ == "__main__":
    import random as rnd
    A = [rnd.randint(1, 100) for i in xrange(200)]
    print "before:", A
    quicksort(A, 0, len(A) - 1)
    print "after:", A
