package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// 1.numeric => float64
// 2.array => []interface{}
// 3.object => map[string]interface{}
// 4.不要再json中加"\xx", 如果要"\"字符 应该"\\xx"
func main() {
	var s string = `{"name":"argpass", "age":99, "paths": ["pa", "p2"], "pattern": "\\["}`
	var data = map[string]interface{}{}
	json.Unmarshal(([]byte)(s), &data)
	fmt.Println(data)
	age := data["age"]
	paths := data["paths"]
	pattern := data["pattern"]
	fmt.Println("pattern:", pattern)

	fmt.Println(reflect.ValueOf(age).Kind() == reflect.Float64)
	fmt.Println(reflect.ValueOf(paths).Kind() == reflect.Slice)

	// to convert int to json is ok
	var dMap = map[string]interface{}{
		"name": "ak",
		"age":  99,
	}
	dJson, _ := json.Marshal(dMap)
	fmt.Println((string)(dJson))
}
