package main

import (
	"fmt"
)

type ParamMap map[string]interface{}

var E_NONE = fmt.Errorf("no data")
var E_INVALID_INT = fmt.Errorf("Invalid %s", "int")

func (m ParamMap) GetInt(key string) (value int, err error) {
	var ok bool
	var v interface{}
	value, err = 0, nil
	v, ok = m[key]
	if !ok {
		err = E_NONE
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		err = E_INVALID_INT
	}
	return value, err
}

func convertCustomType() {
	var err error
	var v int
	var d = map[string]interface{}{"int": 99}
	// convert to ParamMap type
	var c = ParamMap(d)
	_, err = c.GetInt("age")
	fmt.Println(err == E_NONE)

	v, err = c.GetInt("int")
	fmt.Println(v, err)
}

func convertBasicType() {
	var valueFloat float64 = 99
	var valueInt int = 33

	// convert float to int
	valueInt = int(valueFloat)
	fmt.Println(valueInt)
}

func main() {
	convertBasicType()
}
