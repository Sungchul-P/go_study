package main

import "fmt"
import "sort"

func main() {
	//슬라이스(길이 & 용량 개념)
	//배열과 같지만 크기가 동적 할당 가능

	//추가 및 병합
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{8, 9, 10, 11, 12}
	s3 := []int{13, 14, 15, 16, 17}
	s1 = append(s1, 6, 7)
	s2 = append(s1, s2...)      //슬라이스를 삽입할 경우 '...' 사용
	s3 = append(s2, s3[0:3]...) // 추출 후 병합

	fmt.Println("s1 : ", s1)
	fmt.Println("s2 : ", s2)
	fmt.Println("s3 : ", s3)

	s4 := make([]int, 0, 5)
	for i := 0; i < 15; i++ {
		s4 = append(s4, i)
		//길이 및 용량 자동 증가(용량 : 2배)
		//용량(capacity)을 초과하는 경우 현재 용량의 2배에 해당하는 Underlying array를 생성
		//기존 배열 값들을 모두 새 배열에 복제한 후 다시 슬라이스를 할당한다.
		fmt.Printf("len : %d, cap : %d, value : %v\n", len(s4), cap(s4), s4)
	}

	fmt.Println("======================================")

	//정렬
	//sort 패키지 : https://golang.org/pkg/sort/
	slice1 := []int{3, 6, 10, 9, 1, 4, 5, 8, 2, 7}
	slice2 := []string{"b", "d", "f", "a", "c", "e"}
	fmt.Println("slice1 : ", sort.IntsAreSorted(slice1)) // false
	sort.Ints(slice1)
	fmt.Println("slice1 : ", slice1)                     // Int형 정렬
	fmt.Println("slice1 : ", sort.IntsAreSorted(slice1)) // true

	fmt.Println()

	fmt.Println("slice2 : ", sort.StringsAreSorted(slice2)) // false
	sort.Strings(slice2)
	fmt.Println("slice2 : ", slice2)                        // String형 정렬
	fmt.Println("slice2 : ", sort.StringsAreSorted(slice2)) // true

	fmt.Println("======================================")
	//슬라이스 복사
	//copy(복사 대상, 원본)
	//make로 공간을 할당 후 복사 해야 한다.

	a := []int{1, 2, 3, 4, 5}
	b := make([]int, 5)
	copy(b, a)
	b[0] = 7
	b[4] = 10

	fmt.Println("a : ", a) // 원본 유지
	fmt.Println("b : ", b) // 복사 된 대상 변경

	c := a[0:3] // 부분 추출은 참조가 유지되므로, 원본이 변경된다.
	c[1] = 7
	fmt.Println("a : ", a) // 원본 변경
	fmt.Println("c : ", c)

}
