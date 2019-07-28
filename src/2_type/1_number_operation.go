package main

import (
	"fmt"
	"math"
)

func main() {
	//숫자 연산(산술, 비교)
	//타입이 같아야 가능
	//다른 타입끼리는 반드시 형 변환 후 연산
	//형 변환 없을 경우 예외 발생

	//최대, 최소값 확인
	var uN1 uint8 = math.MaxUint8
	var uN2 uint16 = math.MaxUint16
	var uN3 uint32 = math.MaxUint32
	var uN4 uint64 = math.MaxUint64

	fmt.Printf("uint Max : %d %d %d %d\n", uN1, uN2, uN3, uN4)
	fmt.Printf("int Min : %d %d %d %d\n", math.MinInt8, math.MinInt16, math.MinInt32, math.MinInt64)
	fmt.Printf("int Max : %d %d %d %d\n", math.MaxInt8, math.MaxInt16, math.MaxInt32, math.MaxInt64)
	fmt.Println("float32 Max : ", math.MaxFloat32)
	fmt.Println("float64 Max : ", math.MaxFloat64)

	//형 변환
	n5 := 100000 // int
	n6 := int16(10000)
	n7 := uint8(100)

	fmt.Println("n5 + n6 = ", n5+int(n6))
	fmt.Println("n6 + n7 = ", n6+int16(n7))
	fmt.Println("n6 > n7 = ", n6 > int16(n7))
	fmt.Println("n6 - n7 > 5000 = ", n6-int16(n7) > 5000)

	/*
		//예제1(오버플로우 에러 : 범위 초과)
		var n1 uint8 = math.MaxUint8 + 1
		var n2 uint16 = math.MaxUint16 + 1
		var n3 uint32 = math.MaxUint32 + 1
		var n4 uint64 = math.MaxUint64 + 1

		//예제2(오버플로우 에러 : 범위 미만)
		var n1 uint8 = -1
		var n2 uint16 = -1
		var n3 uint32 = -1
		var n4 uint64 = -1
	*/

}
