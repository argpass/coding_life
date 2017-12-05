package main

import "fmt"

func testStructRefField()  {
	type Location struct {
		Param map[string]interface{}
		Data []string
	}
	loc := Location{}
	// map should be made
	//loc.Param["name"] = "bc"
	if loc.Param == nil {
		fmt.Println("param (map) is nil")
	}
	if loc.Data == nil {
		fmt.Println("data is nil")
	}
	fmt.Printf("loc:%+v\n", loc)
}

type Holder struct {
	status int
	result map[string]interface{}
}

// 传值不能修改status
func (h Holder) ChangeValue(v int)  {
	h.status = v
}

func testReceiver()  {
	h := Holder{status:99}
	h.ChangeValue(88)
	fmt.Printf("after change value:%+v\n", h)
	
}

func main()  {
	testReceiver()
}
