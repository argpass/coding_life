package main

import "fmt"

func testMapArg()  {
	var result = make(map[string]interface{})
	func(result map[string]interface{}){
		result = make(map[string]interface{})
		result["hello"] = "world"
	}(result)
	fmt.Printf("after call func result:%v\n", result)
}

func main()  {
	testMapArg()
}
