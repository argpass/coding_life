def insertion_sort(A):
    for i in range(1, len(A), 1):
        key = A[i]
        j = i - 1
        while j >= 0 and key < A[j]:
            A[j + 1] = A[j]
            j -= 1
        A[j + 1] = key


if __name__ == "__main__":
    import random
    A = [random.randint(1, 100) for i in xrange(50)]
    print "before:", A
    insertion_sort(A)
    print "after:", A
