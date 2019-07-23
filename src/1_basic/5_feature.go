package main

import "fmt"

func main() {
	//모호하거나 애매한 문법 금지
	//후치 연산자 허용, 전치 연산자 비허용
	//증감연산 반환 값 없음
	//포인터(Pointer 허용, 연산 비허용)
	//주석 (//, /**/)

	sum, i := 0, 0

	for i <= 100 {
		//sum += i++ //예외 발생(증감 연산 반환 값 X)
		sum += i
		i++
		//++i //예외 발생(전치 연산자 비허용)
	}
	fmt.Println("ex1 : ", sum)

	//문장 끝 세미콜론(;) 주의
	//자동으로 끝에 세미콜론 삽입
	//두 문장을 한 문장에 표현할 경우 명시적으로 세미콜론 사용
	//반복문 및 제어문(if, for) 사용할 경우 주의

	for i := 0; i <= 10; i++ {
		fmt.Print("ex2 : ")
		fmt.Println(i)
	}

	//코드 서식 지정
	//한 사람이 코딩 한 것 같은 일관성 유지
	//코드 스타일 유지
	//gofmt -h : 사용법
	//gofmt -w : 원본파일에 반영

}
