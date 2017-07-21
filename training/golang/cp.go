package main

import "fmt"

func main()  {
	type MyData struct {
		Core map[string]interface{}
	}
	dataA := MyData{Core:map[string]interface{}{"abc": 99}}
	dataB := dataA
	// 会修改原来的字典, 说明是浅拷贝
	dataB.Core["abc"] = 88
	fmt.Printf("值拷贝后原字典为:%v\n", dataA.Core)
	
	dataC := *(&dataA)
	// 会修改原来的字典, 说明是浅拷贝
	dataC.Core["abc"] = 77
	fmt.Printf("值拷贝后原字典为:%v\n", dataA.Core)
}
