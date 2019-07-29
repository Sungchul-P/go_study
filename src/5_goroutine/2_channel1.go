package main

import (
	"fmt"
	"time"
)

func work1(v chan int) {
	fmt.Println("Work1 : S ---> ", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("Work1 : E ---> ", time.Now())
	v <- 1
}

func work2(v chan int) {
	fmt.Println("Work2 : S ---> ", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("Work2 : E ---> ", time.Now())
	v <- 2
}

func rangeSum(rg int, c chan int) {
	sum := 0

	for i := 1; i <= rg; i++ {
		sum += i
	}
	c <- sum
}

func main() {
	//채널(Channel)
	//고루틴간의 상호 정보(데이터) 교환 및 실행 흐름 동기화 위해 사용
	//실행 흐름 제어 가능(동기, 비동기) -> 일반 변수로 선언 후 사용 가능
	//데이터 전달 자료형 선언 후 사용(지정 된 타입만 주고받을 수 있음)
	//interface{} 전달을 통해서 자료형 상관없이 전송 및 수신가능
	//값을 전달(복사 후 : bool, int 등), 포인터(슬라이스, 맵) 등을 전달시에는 주의! -> 동기화 사용(Mutex)
	//멀티프로세싱 처리에서 교착상태(경쟁; race condition) 주의!
	// <- , -> (채널 <- 데이터 : 송신) , ( <- 채널 : 수신)

	fmt.Println("Main : S ---> ", time.Now())

	v := make(chan int) //int형 채널 선언
	go work1(v)
	go work2(v)

	<-v
	<-v
	fmt.Println("Main : E ---> ", time.Now())

	c := make(chan int)

	go rangeSum(1000, c)
	go rangeSum(7000, c)
	go rangeSum(5000, c)

	//순서대로 데이터 수신(동기) : 채널에서 값 수신 완료 될 때까지 대기
	result1 := <-c
	result2 := <-c
	result3 := <-c

	fmt.Println("result1 : ", result1)
	fmt.Println("result2 : ", result2)
	fmt.Println("result3 : ", result3)
}
