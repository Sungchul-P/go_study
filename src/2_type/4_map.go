package main

import "fmt"

func main() {
	//맵(Map)
	//맵 : 해시테이블, 딕셔너리(파이썬), Key-Value로 자료 저장
	//레퍼런스 타입!(참조 값 전달)
	//비교 연산자 사용 불가능(참조 타입이므로)
	//특징 : 참조타입 키(Key)로 사용 불가능 , 값(Value)으로 모든 타입 사용 가능
	//make 함수 및 축약(리터럴)로 초기화 가능
	//순서 없음

	var map1 map[string]int = make(map[string]int) // 정석
	var map2 = make(map[string]int, 10)            // 자료형 생략
	map3 := make(map[string]int)                   // 리터럴 형
	map4 := map[string]int{
		"apple":  15,
		"banana": 30,
		"orange": 23, //콤마 주의
	}

	map1["apple"] = 25
	map1["banana"] = 40
	map1["orange"] = 33

	map2["apple"] = 5
	map2["banana"] = 10
	map2["orange"] = 15

	fmt.Println("map1 : ", map1)
	fmt.Println("map2 : ", map2)
	fmt.Println("map3 : ", map3)
	fmt.Println("map4 : ", map4)
	fmt.Println("map4[\"orange\"] : ", map4["orange"])

	fmt.Println("======================================")

	//순서가 없으므로 랜덤 출력
	for k, v := range map4 {
		fmt.Println("Iterator : ", k, v)
	}

	map5 := map[string]string{
		"daum":   "http://daum.net",
		"naver":  "http://naver.com",
		"google": "http://google.com",
		"home1":  "http://test1.com",
	}
	fmt.Println("map5 : ", map5)
	map5["home2"] = "http://test2.com" //추가
	fmt.Println("map5 : ", map5)
	map5["home2"] = "http://test2-2.com" //수정
	fmt.Println("map5 : ", map5)

	fmt.Println("======================================")

	//삭제
	delete(map5, "home2")
	fmt.Println("map5 : ", map5)
	delete(map5, "home1")
	fmt.Println("map5 : ", map5)

	fmt.Println("======================================")

	value1, ok := map1["kiwi"]

	fmt.Println("map1 : ", value1, ok) //두 번째 리턴 값으로 키 존재 유무 확인 가능

	if value, ok := map1["kiwi"]; ok {
		fmt.Println("map1 : ", value)
	} else {
		fmt.Println("map1 : kiwi is not exist!")
	}

}
