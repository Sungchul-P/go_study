package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//제어문 - Switch
	//switch 뒤 표현식(Expression) 생략 가능
	//case 뒤 표현식(Expression) 사용 가능
	//자동 break 때문에 fallthrough 존재
	//Type 분기 -> 값이 아닌 변수 Type 으로 분기 가능

	a := 30 / 15
	switch a {
	case 2, 4, 6:
		fmt.Println("a -> ", a, "짝수")
	case 1, 3, 5:
		fmt.Println("a -> ", a, "홀수")
	}

	switch b := 27; {
	case b < 0:
		fmt.Println(b, "는 음수")
	case b == 0:
		fmt.Println(b, "는 0")
	case b > 0:
		fmt.Println(b, "는 양수")
	}

	switch c := "go"; c {
	case "go":
		fmt.Println("Go!")
		fallthrough
	case "java":
		fmt.Println("Java!")
		fallthrough
	case "python":
		fmt.Println("Python!")
		fallthrough
	case "ruby":
		fmt.Println("Ruby!")
		// default:
		// 	fmt.Println("일치하는 값 없음")
	}

	rand.Seed(time.Now().UnixNano())

	switch i := rand.Intn(100); {
	case i >= 50 && i < 100:
		fmt.Println("i -> ", i, " 50 이상 100미만")
	case i >= 25 && i < 50:
		fmt.Println("i -> ", i, " 25 이상 50미만")
	default:
		fmt.Println("i -> ", i, " 기본 값")
	}

}
