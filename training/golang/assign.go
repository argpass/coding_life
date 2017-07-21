package main

import "fmt"

func mockValue() (value int, err error)  {
	return 99, nil
}

// 测试":"赋值
func testColonAssignment()  {
	var value int
	fmt.Printf("value1 %p\n", &value)
	defer func(){
		fmt.Printf("value in defer %p\n", &value)
	}()
	// 还是旧的value；同时创建新的err
	value, err := mockValue()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("value2 %p\n", &value)
	
	// 新作用于创建了新value,类型随意
	for value, _:=range []int64{1,} {
		fmt.Printf("value3 %p\n", &value)
	}
	
	// 还是旧的value
	fmt.Printf("value4 %p\n", &value)
}

func main()  {
	testColonAssignment()
}
