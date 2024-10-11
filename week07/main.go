package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Print("이름 입력: ")
	r := bufio.NewReader(os.Stdin)
	i, err := r.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(i)
	}

}
