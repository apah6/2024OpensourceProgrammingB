package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("이름 입력: ")
	r := bufio.NewReader(os.Stdin)
	i, err := r.ReadString('\n')
	fmt.Println(i)
	fmt.Println(err)
}
