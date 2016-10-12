package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func main() {
	var s string = `{"name":"argpass", "age":99}`
	var data = map[string]interface{}{}
	json.Unmarshal(([]byte)(s), &data)
	fmt.Println(data)
	age, _ := data["age"]

	// the integer values will be converted to float64
	fmt.Println(reflect.ValueOf(age).Kind() == reflect.Float64)

	// to convert int to json is ok
	var dMap = map[string]interface{}{
		"name": "ak",
		"age":  99,
	}
	dJson, _ := json.Marshal(dMap)
	fmt.Println((string)(dJson))
}
