package main

import "fmt" //C 언어의 #include <stdio.h>와 동일

func main() {
	// var i int
	// i = 55

	//var i int = 55

	i := 55 //단축연산자 추론 하여 타입을 지정

	var f float32 = 1.92

	fmt.Print("i is", f)
	fmt.Print("i is", i, "\n")
	fmt.Println("i is", i)
	fmt.Printf("i is %d\n", i)
}
