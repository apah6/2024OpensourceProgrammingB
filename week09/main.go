package main

import (
	"fmt"
	"math/rand"
	"reflect"
)

func main() {
	a := fmt.Sprintf("%d\n", rand.Intn(6)+1)
	fmt.Println(reflect.TypeOf(a))
	fmt.Printf("%T\n", a)

	i := 1

	for i <= 100 { //while
		fmt.Printf("%d점수\n", i)
		i = i + 1
	}
}
