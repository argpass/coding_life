# C++ Primer Plus 笔记

## 数组
C++和C将数组名解释为第一个元素的指针.&用于数组名时将返回整个数组的地址。
```C++
// 假设int为4字节
int cookies[10];
cookies == &cookies[0]; // array name is the address of the first element.
p = &cookies // int *p [10]
sizeof(cookies) == 10 * 4
```

## 函数

### 函数原型
- 避免使用函数原型的唯一方法是在首次使用前定义它，但这不总是可行的（在头文件中定义非内联函数可能出现“重复定义”问题）

是否可以将函数定义放头文件中，通过使用条件宏(ifndef __header__)来避免“重定义”？

- C++原型与ANSI原型的重要区别:ANSI C中原型是可选的，但在C++中,原型是不可少的;在ANSI C中括号为空意味着不指出参数——意味着将在之后定义参数列表，而C++中不指定参数列表时应该用省略号。

```C++
void say(); // C-style
void say(...); // C++-style
```
### 函数与数组
- 在C++中当且仅当用于函数头或者函数原型中时，int *arr 和 int arr[] 的含义是相同的.即`int sum_arr(int *arr)`,`int sum_arr(int arr[])`是相同的.
- 数组作为参数的两种方法:
  - 传递数组,元素个数
  - 传递数组区间(头尾指针),STL中使用的超尾(标识数组结尾的参数将是指向最后一个元素后面的指针)

#### 使用const
要尽可能使用const:
- 避免无意间修改了数据导致bug。
- 使用const使得函数能够处理const和非const实参.

常量指针、指针常量、常量常指针:
- 常量指针:`const int *p`,不能修改所指向内存数据.
- 指针常量:`int * const p`,不能修改p的指向,但可以修改p指向内存的数据.
- 常量常指针: `const int * const p`

#### 二维数组作参数

```C++
int sum(int (* p)[4]); // 正确的sum函数原型
int data[3][4] = {
  {1,2,3,4},
  {1,2,3,4},
  {1,2,3,4},
};
int total = sum(data);
data[row][col] == *(*(p + row) + col)

```

#### C-style字符串作参数
C-style字符串与常规char数组的重要却别是字符串内置的结束符号`'\0'`,所以作为参数时不必传入长度（为了效率也可以穿入）.

A Example:
```C++
unsigned int count_in_str(const char * str, char ch) {
    unsigned int count = 0;
    while(*str) {
        if (*str == char) {
            count++;
        }
        str++
    }
    return count;
}
```
### 函数与指针

指针数组与数组的指针:
- `int (*pa) [3]`: 3个元素的int数组的指针。
- `int *pa [3]`: 3个指针元素的数组，元素类型`int*`.
- 数值上看两者相同但意义完全不一样.

```C++
// 函数原型与对应的函数指针
double pam(int);
double (*pf) (int);

double *fn(const double ar[], int n);
double *(*pf)(const double *ar, int)
// 创建一个包含3个该函数指针的数组
// 理解的要点是[]优先级高于*,所以pf先与[]结合，说明它是包含3个元素的数组，然后往外看类型是`double *(*pf)(const double *ar, int)`.如果要得该数组的指针则将pf替换为`(*p)`
double *(*pf[3])(const double *ar, int) = {f1,f2,f3};

// typedef来简化
typedef double *(*pft)(const double *ar, int) = {f1,f2,f3};
pft arr[3] = {f1,f2,f3};

// call function pointer
(*pf)(val)
```

### 内联函数
- 内联函数的通常做法是省略原型，将整个定义放在本应该提供原型的地方。
- 内联函数不能递归。
- C中使用宏来提供内联实现,C++中使用`inline`关键字(原型和定义都要使用，也可以不创建原型)。
- 如果一个函数占用多行则将其作为内联函数就不太合适。

### 引用
- 引用必须在声明引用变量时进行初始化。
- 当传递的数据比较大时引用参数将很有用(效率)。但为了避免被修改应使用`const &T`声明。

#### 临时变量
考虑`void fn(double &rv)`函数，`fn(x+3)`会怎样呢？在现代编译器中这是错误的，较古老的会发出警告。
原因：对于`const &T`这样的参数创建临时变量只在调用期间存在，此后编译器可以随时将其删除，所以允许创建临时变量；`&T` 这样的参数则对创建临时变量是拒绝的（古老的一些编译器只是发出警告）。

如果引用参数是const则在以下情况会创建临时变量:
- 实参类型正确但不是左值(注意const 变量是不可修改的左值)
- 实参类型不正确，但可以转换成正确的类型。

```C++
void swap(long &a, long &b);
void swap2(const long &a, const long &b);
int a=8, b=9;
swap(a, b); // 错误:参数类型不正确将导致创建临时变量，但临时变量用于非const引用则被决绝。有的编译器只是警告。
swap2(a, b); // 这种临时变量是OK的。
```

所以，将引用参数声明为const的理由:
- 使用const引用使函数能够正确生成并使用临时变量。
- 避免无意中修改引起的bug。
- 使函数能处理const 和非const实参，否则只能接受非const数据。

### 默认参数
- 必须从右向左添加默认值
- 原型指定了默认值即可，函数定义可以不指定默认值

```C++
char * left(const char * str, int n=1);
char * left(const char * str, int n) {}
```

### 函数重载
函数重载的关键是函数的特征标(参数列表、函数名),不包括返回值类型。

参数匹配总是调用“最匹配”的版本。
```C++
// 与可修改左值匹配
void stove(double &r1);

// 与可修改左值、const左值和右值参数匹配
void stove(const double &r2);

// 与右值参数匹配
void stove(double && r3)
```

### 函数模版
定义:可以用`template <typename T>`或者`template <class T>`(以前的写法) 来定义模版函数。原型也要用一致的声明。

An Example:
```C++
template <typename AnyType>
void Swap(AnyType &a, AnyType &b);

template <typename AnyType>
void Swap(AnyType &a, AnyType &b){
    AnyType temp;
    temp = a;
    a = b;
    b = temp;
}

// 重载
template <typename AnyType>
void Swap(AnyType a[], AnyType b[]) {
    AnyType temp;
    for (int i=0; i < n; i++) {
        temp = a[i];
        a[i]=b[i];
        b[i]=temp;
    }
}
```

#### 具体化、实例化

具体化优先常规模版，具体函数优先具体化。
具体化还是属于模版函数，只是“更具体”的模版函数，模版函数最终生成函数定义时即是实例化。由编译器进行的实例化是隐式实例化，也可进行显示实例化。

**在一个文件（或者转化单元）中使用同一种类型的显示实例化和显示具体化将出错**

```C++
template <typename T>
void swap(T &a, T &b) {
    T temp;
    temp = a;
    a = b;
    b = temp;
}

// 显示具体化int类型的swap
template <> void swap<int>(int &a, int &b);
template <> void swap<int>(int &a, int &b){
    // 函数定义
}

int main() {
    int a = 1, b=9;
    swap(a, b);
    long c=9, d=8;
    // 显示实例化
    swap<long>(c, d);
    return 0;
}

```

#### 编译器选择哪个函数版本
从最佳到最差的匹配顺序:
- 完全匹配，但常规函数优先于模版函数
- 提升转换(char,short 提升为int, float 提升为double)
- 标准转换(如int转char, long转double)
- 用户定义的转换，如类声明中定义的转换。

通常，有两个函数完全匹配是一种错误。

```C++
void recycle(int);         // #1
void recycle(const int);   // #2
void recycle(int &);       // #3
void recycle(const int &); // #4

// const,非const的区别只适用于指针和引用，如果只定义了#1,#2则出现二义性。
int x=9;
recycle(x); // Error
recycle(&x); // OK
```

## 内存模型
在包含头文件时，使用"header.h"和<header.h>的情况有所不同。如果头文件名包含在<>中则编译器在存储标准头文件的目录中进行查找；如果包含在双引号中则首先在当前工作目录或者源码目录中查找最后在标准库存放目录查找。

不要将函数定义或者变量声明放到头文件中，头文件中常包含的内容:
- 函数原型
- 使用#define或者const 定义的符号常量。
- 结构声明
- 类声明
- 模版声明
- 内联函数

### 存储持续性
- 自动存储持续性:自动变量
- 静态存储持续性:函数外定义的变量和static变量
- 线程存储持续性(C++11)
- 动态存储持续性: 堆上申请的

```C++
int global = 1000; // 省略了`extern`的全局变量定义,静态存储持续性、外部链接性

extern char global_ch = 'Z' // 此处为初始化所以是“定义”，静态存储持续性、外部链接性

extern int global_from_outer; // 此处为声明，静态存储持续性、外部链接性

static int one_file=50; // 静态存储持续性、内部链接性
extern static int multi_file=50; // 静态存储持续性、外部链接性

const int type=666; // 静态存储持续性、内部链接性
extern const int global_type=999; // 静态存储持续性、外部链接性

int main(){
    static int count = 0; // 静态存储持续性、无链接性
    
    int global_from_outer = 99; // 隐藏了外部变量
    int val = ::global_from_outer; // 作用域解析运算符来访问被隐藏的外部变量
}

// 函数默认是外部链接性，但可以加static来定义内部链接性的函数
static void inner(){}
```
要使用外部定义的全局变量需要`extern`来在本文件中进行声明;定义全局变量时`extern`是可以省略的;全局const 常量是内部链接性的(可以加`extern`来改变这种默认行为)，所以可以放在头文件中

### 语言链接性
```C++
// 使用C协议
extern "C" void spiff(int);

// 使用C++协议
extern void spiff(int);
extern "C++" void spiff(int);
```
