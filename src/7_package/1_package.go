package main

import (
	"7_package/lib"
	"fmt"
	"os"
)

//메인 프로그램(package main)
func main() {
	//패키지 : 코드 구조화 및 재사요
	//Go는 패키지로 구성되어 있음
	//패키지 이름 == 디렉터리 이름
	//같은 패키지 내 -> 소스 파일 같은 디렉터리 위치
	//네이밍 규칙 : 소문자, 패키지명 == 소스파일명

	var name string

	fmt.Print("이름은? : ")

	fmt.Scanf("%s", &name)

	fmt.Fprintf(os.Stdout, "Hi %s\n", name)

	fmt.Print("10 보다 큰 수? : ", lib.CheckNum(15))
}
