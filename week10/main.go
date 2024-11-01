package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("정수 입력: ")
	in := bufio.NewReader(os.Stdin)
	number, err := in.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	number = strings.TrimSpace(number)
	n, err := strconv.Atoi(number)

	if err != nil {
		log.Fatal(err)
	}

	count := 0
	i := 2

	for i < n {
		if n%i == 1 {
			count = count + 1
		}
		i++
	}

	if count == 0 {
		fmt.Printf("%d는(은) 소수입니다.\n", n)
	} else {
		fmt.Printf("%d는(은) 소수가 아닙니다!\n", n)
	}

}
