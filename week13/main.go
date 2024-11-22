package main

import (
	"fmt"
)

// func test(strs ...string, i int) { err
func test(i int, strs ...string) {
	fmt.Println(i, strs)
}

func main() {
	//fmt.Println(os.Args, len(os.Args), reflect.TypeOf(os.Args))
	//slices := os.Args[1:]
	//fmt.Println(slices, slices[2])
	test(1, "123")
	test(1, "123", "asd")
	test(1)
	test(1, "123", "asd", "1234")
}
