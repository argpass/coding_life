#include "stdlib.h"
#include "stdio.h"

template <typename T>
void swap(T &a, T &b) {
    T temp;
    temp = a;
    a = b;
    b = temp;
}

template <> void swap<int>(int &a, int &b) {
}

int main() {
    long a = 1, b=9;
    swap<long>(a, b);
    return 0;
}
