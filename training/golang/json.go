package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type dict map[string]interface{}

// 1.numeric => float64
// 2.array => []interface{}
// 3.object => map[string]interface{}
// 4.不要再json中加"\xx", 如果要"\"字符 应该"\\xx"
func testJsonMapping() {
	var s string = `{"name":"argpass", "age":99, "paths": ["pa", "p2"], "pattern": "\\["}`
	// var data map[string]interface{}
	var data dict
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

func testCustomizedType() {
	type Flag int
	type DataType struct {
		Name string `json:"name"`
		Flag Flag   `json:"flag"`
	}
	data := DataType{Flag: 123}
	s, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("json data:%s\n", s)
	data.Flag = 99
	err = json.Unmarshal(s, &data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("data:%+v\n", data)
}

// 组合类型在json操作时是扁平化的 (如果嵌入的域没有设置json tag)
// 非组合嵌入或者嵌入域有json tag则不是扁平化的
func testEmbbedType() {
	type InnerType struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	type InnerType2 struct {
		Name string `json:"name"`
	}

	type unExportedType struct {
		Name string `json:"name"`
	}

	type DataType struct {
		InnerType
		Name string `json:"name"`
		Tag  string
		UnT  unExportedType `json:"unt"`
	}
	dt := DataType{InnerType: InnerType{Name: "inner"}}
	data, err := json.Marshal(dt)
	if err != nil {
		panic(err)
	}
	fmt.Printf("data is:%s\n", data)
	js := `
    {"age":88, "name":"innerName","Tag":"tag-hello", "unt": {"name": "kk"}}
	`
	dt = DataType{}
	err = json.Unmarshal([]byte(js), &dt)
	if err != nil {
		panic(err)
	}
	fmt.Printf("dt unmarshal as :%+v\n", dt)
}

func testValueTypeAdaptation() {
	type MyData struct {
		Lat  float32 `json:"lat"`
		Lng  float64 `json:"__lng"`
		Type int     `json:"type"`
	}
	js := `
	{"lat": 23, "__lng": 22.9, "type": 9}
	`
	data := MyData{}
	err := json.Unmarshal([]byte(js), &data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("data:%+v\n", data)
}

func testAnonymous() {
	js := `
	{"name": "akun", "error": {"code": 99}}
	`
	var v struct {
		Name string `json:"name"`
		Err  struct {
			Code int `json:"code"`
		} `json:"error"`
	}
	err := json.Unmarshal([]byte(js), &v)
	if err != nil {
		panic(err)
	}
	fmt.Printf("v:%+v\n", v)
	data, _ := json.Marshal(v)
	fmt.Printf("data:%s\n", data)

}

func main() {
	testAnonymous()
}
