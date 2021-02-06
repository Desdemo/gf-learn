package main

import (
	"fmt"
	"github.com/gogf/gf/os/gtime"
	"reflect"
	"time"
)

type Person struct {
	Id          int64  `json:"id"`
	StudentName string `json:"name" filter:"like"`
	TTime       []gtime.Time
	StringTime  []string
}

type List []interface{}

//
//func NewList(interface{}) interface{}  {
//
//}

func main() {
	sTime := gtime.Now()
	s2Time := sTime.Add(time.Second * 100)
	sliTime := make([]gtime.Time, 0)
	sliTime = append(sliTime, *sTime)
	sliTime = append(sliTime, *s2Time)
	p := &Person{
		Id:          666,
		StudentName: "dongdong",
		TTime:       sliTime,
		StringTime:  []string{"1", "2"},
	}

	// 获取类型名称
	v := reflect.ValueOf(*p)
	//ty := reflect.TypeOf(*p)
	//fmt.Println(ty.Kind())

	for i := 0; i < v.NumField(); i++ {
		filed := v.Field(i)
		//tagName := filed.Tag
		//fmt.Println(tagName)
		s := filed.Kind()
		if s == reflect.Slice {
			//rrv := reflect.ValueOf(filed)
			//rrvT := rrv.Kind()
			//fmt.Println(".........")
			//fmt.Println(rrvT)
		}
		fmt.Println(s)
		fmt.Println("****************")
		tyN := v.Field(i).Type().String()
		fmt.Println(tyN)
		if tyN == "[]string" {
			fmt.Println("ok")
			rrv := reflect.ValueOf(filed)
			fmt.Println(rrv.Field(0).Interface().(string))
			//value := rrv.Index(0)
			//fmt.Println(value)
			//for x := 0; x < value.NumField(); x ++ {
			//	//st := rrv.Field(i).Interface().(string)
			//	//fmt.Println(st)
			//	file := value.Field(x)
			//	fmt.Println("----------------")
			//	fmt.Println(file)
			//}
		}
		if tyN == "[]gtime.Time" {
			rrv := reflect.ValueOf(v)
			for x := 0; x < rrv.NumField(); x++ {
				//st := rrv.Field(i).Interface().(string)
				//fmt.Println(st)
			}
		}
		//
		//	// 字段名称
		//	fileName := filed.Name
		//	fmt.Println(fileName)
		//	//获取字段具体的值
		//	fmt.Println("*******************")
		//	value:=v.Field(i).Interface()
		//	fmt.Println(value)
		//	fmt.Println("*******************")
		//
		//	fmt.Println(tagName.Get("filter"))
		//
		//
	}
	//fmt.Println("*******************")
	//fmt.Println(v.Kind().String())

	fmt.Println(RefType(p))

}

// 类型判断
func RefType(value interface{}) bool {
	if value == nil {
		return false
	}
	rv := reflect.ValueOf(value)
	switch rv.Kind() {
	case reflect.Struct:
		return true
	case reflect.Ptr:
		return RefType(rv.Elem())
	default:
		return false
	}
}
