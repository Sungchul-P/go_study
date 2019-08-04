package main

import (
	"7_package/lib"
	testlib "7_package/lib2" //별칭 사용
	"fmt"
)

func main() {
	//패키지 접근 제어
	//변수, 상수, 함수, 메서드, 구조체 등 식별자
	//대문자 : 패키지 외부 접근 가능
	//소문자 : 패키지 외부 접근 불가(패키지 내에서만 접근 가능)

	fmt.Println("10 보다 큰 수 ? : ", lib.CheckNum(15))
	fmt.Print("100 보다 큰 수 ? : ", testlib.CheckNum1(101))
	// fmt.Print("1000보다 큰 수 ? : ", testlib.checkNum2(1001)) // 접근 불가능
}
