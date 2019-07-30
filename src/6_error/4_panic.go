package main

import (
	"fmt"
)

//Go Recover : 에러 복구 가능
//Panic으로 발생한 에러를 복구 후 코드 재실행(프로그램 종료 안 됨)
//즉, 코드 흐름을 정상상태로 복구하는 기능
//Panic에서 설정한 메시지를 받아올 수 있다.
func runFunc() {
	defer func() {
		fmt.Println("<< After >>")
		s := recover()
		fmt.Println("Error Message : ", s)
	}()

	// fmt.Println("<< Before >>")
	// panic("Error occurred!")
	// fmt.Println("test") //호출 안됨

	a := [3]int{1, 2, 3}

	for i := 0; i < 5; i++ {
		fmt.Printf("a[%d] : %d\n", i, a[i]) //에러 발생(index out of range)
	}
}

func main() {
	//Go panic : 사용자가 에러 발생 가능
	//Panic 함수는 호출 즉시, 해당 메소드를 중지시키고 defer 함수를 호출.
	//자기자신을 호출한 곳으로 리턴.
	//런타임 이외에 사용자가 코드 흐름에 따라 에러를 발생 시킬 때 중요!

	runFunc()

	fmt.Println("Hello Golang !")
}
