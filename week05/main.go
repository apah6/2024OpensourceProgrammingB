package main

import (
	"fmt" //C 언어의 #include <stdio.h>와 동일
	"math"
	"reflect"
	"strings"
)

func main() {
	// var i int
	// i = 55

	//var i int = 55
	//f := 1.92
	i := "55" //단축연산자 추론 하여 타입을 지정

	var f float64 = 1.92

	fmt.Println(strings.Title("kim inha"))
	fmt.Printf("%f\n", math.Ceil(f))
	fmt.Println(reflect.TypeOf(f), reflect.TypeOf(i))
	fmt.Println("f is", f)
	fmt.Println("i is", i)
	fmt.Print("i is", i, "\n")
	fmt.Println("i is", i)
	//fmt.Printf("i is %d\n", i)
	fmt.Printf("i is %s\n", i)
}
