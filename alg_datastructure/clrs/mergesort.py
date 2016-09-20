def merge(A, p, m, r):
    L = A[p: m + 1]
    R = A[m + 1: r + 1]
    L.append(1<<64)
    R.append(1<<64)
    j, k = 0, 0
    for i in range(p, r+1, 1):
        if L[j] < R[k]:
            A[i] = L[j]
            j += 1
        else:
            A[i] = R[k]
            k += 1

        
def merge_sort(A, p, r):
    if p < r:
        m = (p + r) / 2
        merge_sort(A, p, m)
        merge_sort(A, m + 1, r)
        merge(A, p, m, r)


if __name__ == "__main__":
    import random as rnd
    A = [rnd.randint(1, 100) for i in xrange(200)]
    print "before:", A
    merge_sort(A, 0, len(A) - 1)
    print "after:", A
