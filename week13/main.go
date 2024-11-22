package main

import (
	"fmt"
)

func main() {
	var emptySlice []bool
	//var emptySlice = make([]bool,5)
	fmt.Printf("%#v %d\n", emptySlice, len(emptySlice))
	if len(emptySlice) == 0 {
		emptySlice = append(emptySlice, true)
	}
	fmt.Printf("%#v %d\n", emptySlice, len(emptySlice))

	gpas := [5]float64{3.5, 4.1, 4.5, 3.9, 4.23} // array := array literal
	gpas_Slice := gpas[1:4]
	//gpas_Slice[1] = 2.71
	gpas[2] = 2.71
	//gpas_Slice = append(gpas_Slice, 4.3)
	gpas_Slice = append(gpas_Slice, 4.3, 5.55)
	fmt.Println(len(gpas_Slice), gpas_Slice, gpas)
}
