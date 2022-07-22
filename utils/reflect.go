package utils

import (
	"fmt"
	"reflect"
)

func reflectInfo(i interface{}) {
	t := reflect.TypeOf(i)
	fmt.Println("Type: ", t.Name())
	v := reflect.ValueOf(i)
	fmt.Println("Fields: ")
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Printf("%6s: %v = %v\n", f.Name, f.Type, val)
	}
}

func getFieldName(t interface{}) {
}
