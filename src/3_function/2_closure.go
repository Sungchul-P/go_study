package main

import "fmt"

func increaseCnt() func() int {
	n := 0 //지역변수
	return func() int {
		n++
		return n
	}
}

func main() {
	//클로저(Closure)
	//익명함수 사용할 경우 함수를 변수에 할당해서 사용가능
	//이 때, 함수는 일급 객체이므로 변수의 값으로 사용 가능
	//현재 범위에 있는 변수의 값을 캡처 후 호출 할 때 변수 사용 가능(선언 시점 기준)

	multiply := func(x, y int) int {
		return x * y
	}

	r1 := multiply(5, 10)
	fmt.Println("r1 : ", r1)

	m, n := 10, 5 //변수를 캡처
	sum := func(c int) int {
		return m + n + c // 지역 변수는 소멸되지 않는다.(함수 호출 시 마다 사용 가능)
	}

	r2 := sum(10)
	fmt.Println("r2 : ", r2)

	cnt := increaseCnt()

	for i := 0; i < 5; i++ {
		fmt.Println("cnt : ", cnt())
	}
}
