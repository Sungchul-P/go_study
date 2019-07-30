package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	//에러 처리
	//Go에서는 기본적으로 error 패키지에서 error 인터페이스를 제공
	//type error interface { Error() string }
	//기본적으로 메서드 마다 리턴 타입 2개(리턴 값, 에러)
	//주로 1. Errorf(에러 타입 리턴) 메소드 2. Fatal(프로그램 종료) 메소드를 통해서 에러 출력

	//기본적인 메서드 에러 처리 예제
	f, err := os.Open("unnamedfile") //err
	if err != nil {
		log.Fatal(err.Error()) // 방법1
		// log.Fatal(err) // 방법2
	}
	fmt.Println("==================")
	fmt.Println("f.Name() : ", f.Name())
}
