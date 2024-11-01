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

	//count := 0

	var isPrime bool = true //int -> bool counts -> isprime 메모리효율증가

	if n <= 1 {
		isPrime = false //가독성
	} else {

		i := 2

		for i < n {
			if n%i == 0 {
				isPrime = false
				break // 1과 자기 자신을 제외한 첫 번째 약수가 발견 되면 반복문 종료
			}
			fmt.Printf("%d ", i)
			i++
		}
	}
	//if count == 0 {
	if isPrime { //비교연산 제거
		fmt.Printf("%d는(은) 소수입니다.\n", n)
	} else {
		fmt.Printf("%d는(은) 소수가 아닙니다!\n", n)
	}

}
