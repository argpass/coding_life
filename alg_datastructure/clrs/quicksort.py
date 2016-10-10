def partition(A, p, r):
    i = p
    x = A[r]
    for j in xrange(p, r):
        if A[j] <= x:
            A[i], A[j] = A[j], A[i]
            i += 1
    A[r], A[i] = A[i], A[r]
    return i

def quick_sort(A, p, r):
    if r > p:
        i = partition(A, p, r)
        quick_sort(A, p, i-1)
        quick_sort(A, i + 1, r)


if __name__ == "__main__":
    seq = [3,8,4,5,7,1]
    quick_sort(seq, 0, len(seq) - 1)
    print seq
