package main

import (
	"fmt"
	"reflect"
)

type Per struct {
	Name []string
}

func main() {
	names := []string{"2", "3"}
	p := Per{Name: names}
	rv := reflect.ValueOf(p)
	fmt.Println(rv)
	fmt.Println(rv.Field(0).Interface().([]string)[0])
}
