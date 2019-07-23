package main

import "fmt"

func main() {
	//제어문 - IF
	//반드시 Boolean 식으로 검사 -> 1, 0(자동 형 변환 없음)
	//소괄호 사용 안함

	var a int = 20
	b := 23

	if a >= 15 {
		fmt.Println("15 이상")
	} else {
		fmt.Println("15 미만")
	}

	if b >= 25 {
		fmt.Println("25 이상")
	} else if b >= 20 && b < 25 {
		fmt.Println("20 이상 25 미만")
	} else {
		fmt.Println("20 미만")
	}

	if c := 40; c >= 35 {
		fmt.Println("35 이상")
	}
}
