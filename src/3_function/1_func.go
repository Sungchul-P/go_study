package main

import "fmt"

func sum(i int, f func(int, int)) {
	f(i, 10)
}

func add(a, b int) {
	fmt.Println("add : ", a+b)
}

func multiValue(i int) {
	i = i * 10
}

func multiReference(i *int) {
	*i *= 10
}

func multiply(x int, y int) (int, int) {
	return x * 10, y * 20
}

//리턴 변수 지정
func multiply2(x int, y int) (r1 int, r2 int) {
	r1 = x * 10
	r2 = y * 20
	return
}

func arrayMultiply(a, b, c, d, e int) (int, int, int, int, int) {
	return a * 1, b * 2, c * 3, d * 4, e * 5
}

func main() {
	//함수 선언 : func 키워드로 선언
	//타 언어와 달리 반환 값(return value) 여러 개 가능

	//콜백 호출
	sum(10, add) // 20 출력

	//call by value
	a := 100
	multiValue(a)
	fmt.Println("call by value : ", a)

	//call by reference
	b := 100
	multiReference(&b)
	fmt.Println("call by reference : ", b)

	x1, x2, x3, x4, x5 := arrayMultiply(1, 2, 3, 4, 5)
	y1, _, y3, _, y5 := arrayMultiply(1, 2, 3, 4, 5)
	fmt.Println("arrayMultiply : ", x1, x2, x3, x4, x5)
	fmt.Println("arrayMultiply : ", y1, y3, y5)

	m1, m2 := multiply(10, 5)
	fmt.Println("multiply : ", m1, m2)
	m3, m4 := multiply2(10, 5)
	fmt.Println("multiply2 : ", m3, m4)
}
