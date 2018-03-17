# 单例模式 Singleton

## 意图

>保证一个类只有一个实例，并提供一个全局访问点

我的理解：
- 有的对象在需要作为进程的全局对象来使用；
- 创建该对象是比较耗费资源（昂贵)的;
- 保证该对象只被创建一次，使得进程中任意位置访问该对象都是同一对象

## 结构
## 实现
使用C++实现：
---------------
```c++

namespace pattern {

    class Singleton {
        typedef int DataType;
    public:
        DataType data;
        static Singleton * _instance;

        Singleton(DataType data){
            this->data = data;
        }
        static Singleton * instance(DataType data){
            if(_instance == nullptr){
                _instance = new Singleton(data);
            }
            return _instance;
        }
    };

    Singleton * Singleton::_instance = 0;


    // extend
    class MySingleton:public Singleton {
    };

}
```
