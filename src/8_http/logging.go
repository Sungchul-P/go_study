package main

import (
	"fmt"
	"log"
	"net/http"
)

//미들웨어 함수 선언(Return : func)
func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("MiddleWare Test(Log) : ", r.URL.Path)
		f(w, r)
	}
}

//요청1 핸들러
func reqTest1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Requested : (/reqTest1)")
}

//요청2 핸들러
func reqTest2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Requested : (/reqTest2)")
}

func main() {
	//핸들러 메소드 지정1
	http.HandleFunc("/reqTest1", logging(reqTest1))
	//핸들러 메소드 지정2
	http.HandleFunc("/reqTest2", logging(reqTest2))

	//기본 출력
	fmt.Println("Golang WebServer Working!")

	//서버 시작
	http.ListenAndServe(":9090", nil)
}
