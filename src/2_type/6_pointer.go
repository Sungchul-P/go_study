package main

import "fmt"

func rptc(n *int) {
	*n = 77
}

func vptc(n int) {
	n = 77
}

func main() {
	//포인터
	//주소의 값은 직접 변경 불가능
	//메모리 주소를 출력(값의 메모리 주소)
	//nil로 초기화

	var a *int
	var b *int = new(int)

	fmt.Println(a) //nil
	fmt.Println(b)

	i := 7
	a = &i //& 주소값 전달
	b = &i

	fmt.Printf("i : %d %p\n", i, &i)
	fmt.Printf("a : %d %p %p\n", *a, a, &a)
	fmt.Printf("b : %d %p %p\n", *b, b, &b)

	fmt.Println("==================================")

	type bag struct{ width, height, weight float32 } //구조체
	var p *int = new(int)
	var p_bag *bag = &bag{20, 50, 30}

	fmt.Println("p : ", p)
	fmt.Println("p_bag : ", p_bag)
	fmt.Println()

	j := 7
	p = &j

	fmt.Println(j, *p, &j, p)

	*p++ //1 증가
	fmt.Println(j, *p, &j, p)

	*p = 10 //포인터 변수 역 참조 값 변경
	fmt.Println(j, *p, &j, p)

	j = 77 //원본 변수 값 변경
	fmt.Println(j, *p, &j, p)

	fmt.Println("==================================")

	//포인터 값 전달
	//함수, 메서드 호출 시에 매개변수 값을 복사 전달
	//-> 함수, 메서드 내에서는 원본 값 변경 불가능
	//원본 값 변경 위해서 포인터 전달 가능
	//특히 크기가 큰 배열인 경우 값 복사시에 시스템 부담
	//-> 포인터 전달로 해결 But 슬라이스는 참조 전달

	var pa int = 10
	var pb int = 10
	fmt.Println("pa : ", pa)
	fmt.Println("pb : ", pb)
	fmt.Println()

	rptc(&pa)
	vptc(pb)

	fmt.Println("pa : ", pa)
	fmt.Println("pb : ", pb)

}
