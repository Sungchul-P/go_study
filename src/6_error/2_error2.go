package main

import (
	"errors"
	"fmt"
	"log"
)

func notZero(n int) (string, error) {
	if n != 0 {
		s := fmt.Sprint("Hello Golang : ", n)
		return s, nil
	}

	// return "", fmt.Errorf("%d를 입력했습니다. 에러 발생!", n)
	return "", errors.New("0을 입력했습니다. 에러 발생!")

}

func main() {
	a, err := notZero(1)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("a : ", a)

	b, err := notZero(0) // 0넣으면 에러 발생

	if err != nil {
		log.Fatal(err)
	}

	//Fatal 이후 프로그램 종료. 아래 코드는 실행 안됨.
	fmt.Println("b : ", b)
	fmt.Println("End Error Handling!")

}
