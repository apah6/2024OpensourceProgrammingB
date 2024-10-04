package main

import (
	"fmt"
	"reflect"
)

func main() {
	i := 13
	//f := 12.9
	var f float64 = 12.9
	c1 := 'Z'
	c2 := '김'

	fmt.Printf("value i : %d, value f : %f\n", i, f)
	fmt.Printf("%d * %f = %f\n", i, f, float64(i)*f)
	fmt.Printf("%d * %f = %d\n", i, f, i*int(f))

	fmt.Println(c1, c2)    //10진 표현
	fmt.Printf("%x\n", c2) //유니코드 '김' 16진수 변환 유니코드 사이트:https://www.unicode.org/
	fmt.Println(reflect.TypeOf(f), reflect.TypeOf(i), reflect.TypeOf(c1), reflect.TypeOf(c2))
}
