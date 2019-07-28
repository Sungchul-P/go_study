package main

import "fmt"

func exF1() {
	fmt.Println("f1 : start")
	defer exF2() //마지막에 호출
	fmt.Println("f1 : end")
}

func exF2() {
	fmt.Println("f2 : called")
}

func sayHello(msg string) {
	defer func() {
		fmt.Println(msg)
	}()

	func() {
		fmt.Print("Hi ")
	}()
}

func stack() {
	for i := 1; i <= 10; i++ {
		defer fmt.Println("stack : ", i)
	}
}

func start(t string) string {
	fmt.Println("start : ", t)
	return t
}

func end(t string) {
	fmt.Println("end : ", t)
}

func call() {
	defer end(start("b"))
	fmt.Println("in call()")
}

func main() {
	//Defer : 함수 지연 실행
	//Defer를 호출한 함수가 종료되기 직전에 호출된다.
	//타 언어의 Finally 문과 비슷
	//주로 리소스 반환 등에 사용

	exF1()
	fmt.Println()
	sayHello("Golang!")
	fmt.Println()
	stack()
	fmt.Println()
	call() // start() -> in call -> end()
}
