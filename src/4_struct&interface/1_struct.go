package main

import (
	"fmt"
	"reflect"
)

type Account struct {
	number   string
	balance  float64
	interest float64
}

type Car struct {
	name    string "차랑명"
	color   string "색상"
	company string "제조사"
}

type totCost func(int, int) int

func describe(cnt int, price int, fn totCost) {
	fmt.Printf("cnt: %d, price: %d, orderPrice: %d", cnt, price, fn(cnt, price))
}

func (a Account) Calculate() float64 {
	return a.balance + (a.balance * a.interest)
}

func main() {
	//구조체 : Go의 필드들의 집합체 또는 컨테이너
	//필드는 갖지만 메소드는 갖지 않는다.
	//객체지향 방식을 지원 -> 리시버를 통해 메소드랑 연결
	//상속, 객체, 클래스 개념 없음
	//구조체 -> 구조체 선언 -> 구조체 인스턴스(리시버)

	kim := new(Account)
	kim.number = "245-901"
	kim.balance = 10000000
	kim.interest = 0.015

	lee := new(Account)
	lee.number = "245-902"
	lee.balance = 13000000
	lee.interest = 0.025

	park := Account{number: "245-903", interest: 0.025}

	hong := &Account{number: "245-904", balance: 15000000, interest: 0.04}

	fmt.Printf("kim : %#v\n", kim)
	fmt.Println("lee : ", lee)
	fmt.Println("park : ", park)
	fmt.Println("hong : ", hong)

	fmt.Println()

	fmt.Println("kim.Cal : ", int(kim.Calculate()))
	fmt.Println("lee.Cal : ", int(lee.Calculate()))
	fmt.Println("park.Cal : ", int(park.Calculate()))
	fmt.Println("hong.Cal : ", int(hong.Calculate()))

	fmt.Println()
	//구조체 익명 선언 및 활용
	cars := []struct{ name, color string }{{"520d", "red"}, {"220d", "white"}, {"420d", "black"}}
	for _, c := range cars {
		fmt.Printf("(%s, %s) ----- (%#v)\n", c.name, c.color, c)
	}

	fmt.Println()
	//필드 태그 사용
	tag := reflect.TypeOf(Car{})
	for i := 0; i < tag.NumField(); i++ {
		fmt.Println("tag : ", tag.Field(i).Tag, tag.Field(i).Name, tag.Field(i).Type)
	}

	fmt.Println()
	//함수 사용자 정의 타입
	var orderPrice totCost
	orderPrice = func(cnt int, price int) int {
		return (cnt * price) + 1000000
	}
	describe(3, 10000000, orderPrice)
}
