package main

import "fmt"

func main() {
	//배열은 용량, 길이가 항상 같다.
	//배열 vs. 슬라이스
	//길이고정 | 길이 가변
	//값 타입  | 참조 타입
	//복사 전달 | 참조 값 전달
	//전체 비교연산자 사용 가능 | 비교 연산자 사용 불가
	//슬라이스를 많이 사용한다.

	//cap() : 배열, 슬라이스 용량
	//len() : 배열, 슬라이스 개수

	var arr1 [5]int
	arr2 := [5]int{1, 2, 3} //나머지 0 초기화
	arr3 := [...]int{1, 2, 3, 4, 5}
	arr4 := [5][5]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10}, //콤마 주의
	}
	arr5 := [...]string{"Kim", "Lee", "Park"}

	arr1[2] = 5 //값 삽입

	fmt.Printf("%T %d %v\n", arr1, len(arr1), arr1)
	fmt.Printf("%T %d %v\n", arr2, len(arr2), arr2)
	fmt.Printf("%T %d %v\n", arr3, len(arr3), arr3)
	fmt.Printf("%T %d %v\n", arr4, len(arr4), arr4)
	fmt.Printf("%T %d %v\n", arr5, len(arr5), arr5)

	fmt.Println("=================================")
	//배열 순회
	for i := 0; i < len(arr1); i++ {
		fmt.Printf("%d \t", arr1[i])
	}
	fmt.Println()

	//range 사용
	for i, v := range arr2 {
		fmt.Printf("%d %v \t | ", i, v)
	}
	fmt.Println()

	//인덱스 생략
	for _, v := range arr3 {
		fmt.Printf("%v \t | ", v)
	}
	fmt.Println()

	//배열 복사
	//값만 복사되고 새로 생성된다. 확인 중요 !

	arrCp1 := [5]int{1, 10, 100, 1000, 10000}
	arrCp2 := arrCp1

	fmt.Println("=================================")
	fmt.Printf("arrCp1 : %v %p\n", arrCp1, &arrCp1)
	fmt.Printf("arrCp2 : %v %p\n", arrCp2, &arrCp2)
}
