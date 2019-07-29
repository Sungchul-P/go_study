package main

import "fmt"

func main() {
	//채널(Channel)
	//동기 : 버퍼 미사용
	// ch := make(chan bool)
	// cnt := 6

	//비동기 : 버퍼 사용
	//버퍼 : 발신 -> 가득차면 대기. 비어있으면 작동
	// 수신 -> 비어있으면 대기. 가득 차있으면 작동
	// runtime.GOMAXPROCS(1)
	// ch := make(chan bool, 4)
	// cnt := 12

	// go func() {
	// 	for i := 0; i < cnt; i++ {
	// 		ch <- true
	// 		fmt.Println("Go : ", i)
	// 		//time.Sleep(1 * time.Second)
	// 	}
	// }()

	// for i := 0; i < cnt; i++ {
	// 	<-ch
	// 	fmt.Println("Main : ", i)
	// }

	//Close : 채널 닫기. 주의 -> 닫힌 채널에 값 전송 시 패닉(예외 발생)
	//Range : 채널 안에서 값을 꺼낸다.(순회)
	//채널 닫아야(Close) 반복문 종료 -> 채널이 열려 있고 값 전송하지 않으면 계속 대기!
	ch := make(chan bool)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- true
		}
		close(ch) //5회 채널에 값 전송 후 채널 닫기
	}()

	for i := range ch { //채널이 Close 될 때까지 반복
		fmt.Println("Range : ", i)
	}
}
