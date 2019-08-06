package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

// 요청에 대한 동작이 정의된 핸들러
func defaultHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	// GET 파라미터 및 정보 출력
	fmt.Println("default : ", r.Form)
	fmt.Println("path : ", r.URL.Path)
	fmt.Println("param : ", r.Form["test_param"])

	// 파라미터 전체 출력
	for k, v := range r.Form {
		fmt.Println("key: ", k)
		fmt.Println("value: ", strings.Join(v, ""))
	}

	// 기본 출력
	fmt.Fprintf(w, "Golang WebServer Working!")
}

func main() {
	http.HandleFunc("/", defaultHandler)     // 기본 URL 에 대한 핸들러 메소드 지정
	err := http.ListenAndServe(":9090", nil) // 서버 시작
	if err != nil {                          // 예외 처리
		log.Fatal("ListenAndServe: ", err)
	} else {
		fmt.Println("ListenAndServe Started! -> Port(9090)")
	}
}
