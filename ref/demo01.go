package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Id int64 `json:"id"`
	StudentName string `json:"name" filter:"like"`
}


func main()  {
	p := &Person{
		Id:   666,
		StudentName: "dongdong",
	}

	// 获取类型名称
	v := reflect.ValueOf(*p)
	ty := reflect.TypeOf(*p)

	for i :=0 ; i < ty.NumField(); i ++ {
		filed := ty.Field(i)
		tagName := filed.Tag
		fmt.Println(tagName)

		// 字段名称
		fileName := filed.Name
		fmt.Println(fileName)
		//获取字段具体的值
		fmt.Println("*******************")
		value:=v.Field(i).Interface()
		fmt.Println(value)
		fmt.Println("*******************")

		fmt.Println(tagName.Get("filter"))


	}
	fmt.Println("*******************")
	fmt.Println(v.Kind().String())

}
