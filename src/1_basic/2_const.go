package main

import "fmt"

func main() {
	//상수
	//const 사용 초기화, 한 번 선언 후 값 변경 금지 , 고정 된 값 관리 용

	const a string = "test1"
	const b = "test2"
	const c int32 = 10 * 10
	//const d = getHeight()
	const e = 35.6
	const f = false
	/*
		에러 발생
		const g string
		g = "test3"
	*/

	const g, h int = 10, 99
	const i, j, k = true, 0.84, "test"
	const (
		x, y  int16 = 50, 90
		z, z2       = "Data", 7776
	)

	fmt.Println("a : ", a, "b : ", b, "c : ", c, "e : ", e, "f : ", f)
	fmt.Println("g: ", g, " h : ", h)
	fmt.Println("i : ", i, "j : ", j, "k : ", k)
	fmt.Println("x : ", x, "y : ", y)
	fmt.Println("z : ", z, "z2 : ", z2)

}
