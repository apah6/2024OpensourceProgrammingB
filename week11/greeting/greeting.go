package greeting

// 소문자로 시작하는 함수는 외부 접근 불가
// 단 ,내부에서는 접근 가능
import "fmt"

func hi(name string) {
	fmt.Printf("Hi %s!\n", name)
}

func hello(name string) {
	fmt.Printf("Hello %s~\n", name)
}

func EnglishGreetings(name string) {
	hello(name)
	hi(name)
}
