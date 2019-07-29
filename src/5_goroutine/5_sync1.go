package main

import (
	"fmt"
	"runtime"
	"sync"
)

//구조체 선언(공유 데이터)
type count struct {
	num   int
	mutex sync.Mutex
}

func (c *count) increment() {
	//공유 데이터 수정 전 뮤텍스로 보호
	c.mutex.Lock()
	c.num += 1
	//공유 데이터 수정 후 보호 해제
	c.mutex.Unlock()
}

func (c *count) result() {
	fmt.Println(c.num)
}

func main() {
	//고루틴 동기화 예제
	//실행 흐름 제어 및 변수 동기화 가능
	//공유 데이터 보호가 가장 중요
	//뮤텍스(Mutex) : 여러 고루틴에서 작업하는 공유데이터 보호
	// 상호 배제 -> Thread(고루틴)들이 서로 running time에 서로 영향을 주지 않게 하는
	// 즉, 단독으로 실행되게 하는 기술
	//sync.Mutex 선언 후 Lock, Unlock 사용

	runtime.GOMAXPROCS(runtime.NumCPU())

	c := count{num: 0}
	done := make(chan bool)

	for i := 1; i <= 10000; i++ {
		go func() {
			c.increment()
			done <- true
			runtime.Gosched() //CPU 양보
		}()
	}

	for i := 1; i <= 10000; i++ {
		<-done
	}

	c.result()
}
