package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	//문자열 선언 : "" , ``
	//문자 char타입 존재하지 않음 -> rune(int32)로 문자 코드 값으로 표현
	//문자 선언 : ''

	var str1 string = "Hello, world!"
	var str2 string = `안녕하세요`

	//길이(바이트 수)
	fmt.Println("len(str1) : ", len(str1))
	fmt.Println("len(str2) : ", len(str2))

	//길이(실제 문자 수)
	fmt.Println("RuneCountInString(str1) : ", utf8.RuneCountInString(str1))
	fmt.Println("RuneCountInString(str2) : ", utf8.RuneCountInString(str2))
	fmt.Println("RuneCountInString(str2) : ", len([]rune(str2)))

	//문자열 추출
	var str3 string = "Golang"
	var str4 string = "World"

	//str3[0] == "G" == 71 (아스키 코드)
	fmt.Println("extract : ", str3[0:2], str3[0])
	fmt.Println("extract : ", str4[3:], str4[0])
	fmt.Println("extract : ", str3[1:3])

	//문자열 조합은 한 번 생성 후 수정 불가 이유로 새로 계속해서 생성 된다.
	//효율적 사용을 위해서 되도록 join 함수 사용 추천

	//예제1(결합 : 일반연산)
	str5 := "Go is expressive, concise, clean, and efficient. Its concurrency mechanisms make it easy to " +
		"write programs that get the most out of multicore and networked machines, while its novel " +
		"type system enables flexible and modular program construction. Go compiles quickly to machine " +
		"code yet has the convenience of garbage collection and the power of run-time reflection."
	str6 := "It's a fast, statically typed, compiled language that feels like a dynamically typed, interpreted language."

	fmt.Println("str5+str6 : ", str5+str6)

	//예제2(결합 : Join)
	strSet := []string{} //슬라이스 선언
	strSet = append(strSet, str5)
	strSet = append(strSet, str6)
	fmt.Println("strSet : ", strings.Join(strSet, ""))
}
