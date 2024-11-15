package main

import (
	"fmt"
	"time"
)

func main() {
	// var dates [3]time.Time
	// dates[2] = time.Unix(1708012346, 0)
	// dates[0] = time.Unix(0, 0)
	// fmt.Println(dates[0], dates[1], dates[2]) // unix time, zero value, unixTime(1708012346)

	//var dates [3]time.Time = [3]time.Time{time.Unix(0, 0), time.Unix(1, 0), time.Unix(1708012346, 0)}
	// dates := [3]time.Time{time.Unix(0, 0), time.Unix(1, 0), time.Unix(1708012346, 0)}
	// fmt.Println(dates[0], dates[1], dates[2]) // unix time, zero value, unixTime(1708012346)

	// dates := [3]time.Time{
	// 	time.Unix(0, 0),
	// 	time.Unix(1, 0),
	// 	time.Unix(1708012346, 0),
	// }
	// fmt.Println(dates[0], dates[1], dates[2])

	dates := [3]time.Time{
		time.Unix(0, 0),
		time.Unix(1, 0),
		time.Unix(1708012346, 0)}

	fmt.Println(dates[0], dates[1], dates[2])
	fmt.Println(dates)         //arry
	fmt.Printf("%#v\n", dates) //arry literal
}
