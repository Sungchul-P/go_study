package main

import "fmt"

func main() {
	//열거형
	//상수를 사용하는 일정한 규칙에 따라 숫자를 증가시키는 묶음
	// const (
	// 	Jan = 1
	// 	Feb = 2
	// 	Mar = 3
	// 	Apr = 4
	// 	May = 5
	// 	Jun = 6
	// )
	// fmt.Println("Jan : ", Jan)
	// fmt.Println("Feb : ", Feb)
	// fmt.Println("Mar : ", Mar)
	// fmt.Println("Apr : ", Apr)
	// fmt.Println("May : ", May)
	// fmt.Println("Jun : ", Jun)

	const (
		_ = iota
		A
		B
		C
	)

	const (
		Jan = iota + 1 //계산식 가능
		Feb
		Mar
		Apr
		May
		Jun
	)

	fmt.Println("A : ", A)
	fmt.Println("B : ", B)
	fmt.Println("C : ", C)

	fmt.Println("Jan : ", Jan)
	fmt.Println("Feb : ", Feb)
	fmt.Println("Mar : ", Mar)
	fmt.Println("Apr : ", Apr)
	fmt.Println("May : ", May)
	fmt.Println("Jun : ", Jun)

	const (
		_ = iota + 0.75*2
		DEFAULT
		SILVER
		GOLD
		PLATINUM
	)

	fmt.Println("DEFAULT : ", DEFAULT)
	fmt.Println("SILVER : ", SILVER)
	fmt.Println("GOLD : ", GOLD)
	fmt.Println("PLATINUM : ", PLATINUM)
}
