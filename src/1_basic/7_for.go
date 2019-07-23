package main

import "fmt"

func main() {
	//반복문 - for
	//Go에서 유일한 반복문

	//ex1) for 기본형식
	for i := 0; i < 5; i++ {
		fmt.Println("ex1 : ", i)
	}

	//무한 루프
	/*
		for {
			fmt.Println("Infinite Loop !")
		}
	*/

	//ex2) while문 형식
	sum1, i := 0, 0
	for i <= 10 {
		sum1 += i
		i++
	}
	fmt.Println("ex2 : ", sum1)

	//ex3) Range 사용
	loc := []string{"Seoul", "Busan", "Incheon"}
	for index, name := range loc {
		fmt.Println("ex3 : ", index, name)
	}

	//ex4) break 사용
	sum2, i := 0, 0
	for {
		if i > 100 {
			break
		}
		sum2 += i
		i++
	}
	fmt.Println("ex4 : ", sum2)

	//ex5)
	for i, j := 0, 0; i <= 10; i, j = i+1, j+10 {
		fmt.Println("ex5 : ", i, j)
	}

	//ex6) Loop 레이블을 이용한 제어
Loop1:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 2 && j == 4 {
				break Loop1
			}
			fmt.Println("ex6 : ", i, j)
		}
	}

	//ex7) continue 사용
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("ex7 : ", i)
	}
}
