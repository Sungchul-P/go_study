package main

import "fmt"

type Dog struct {
	name   string
	weight int
}

type Cat struct {
	name   string
	weight int
}

//구조체 Dog 메소드 구현
func (d Dog) bite() {
	fmt.Println("Dog bites!")
}

func (d Dog) sounds() {
	fmt.Println("Dog barks!")
}

func (d Dog) run() {
	fmt.Println(d.name, "Dog is running!")
}

//구조체 Cat 메소드 구현
func (c Cat) bite() {
	fmt.Println("Cat bites!")
}

func (c Cat) sounds() {
	fmt.Println("Cat cries!")
}

func (c Cat) run() {
	fmt.Println(c.name, "Cat is running!")
}

//인터페이스를 파라미터로 받는다.
func act(animal Behavior) {
	animal.bite()
	animal.sounds()
	animal.run()
}

//동물의 행동 인터페이스 선언
type Behavior interface {
	bite()
	sounds()
	run()
}

func printValue(s interface{}) {
	fmt.Println(s)
}

func main() {
	//인터페이스
	//객체의 동작을 표현, 골격
	//단순히 동작에 대한 방법만 표시
	//추상화 제공
	//덕타이핑 : Go 언어의 독창적인 특성
	// => 구조체 및 변수의 값이나 타입은 상관하지 않고 오로지 구현된 메소드로만 판단하는 방식

	dog := Dog{"poll", 10}
	cat := Cat{"bob", 5}

	//개 행동 실행
	act(dog)
	//고양이 행동 실행
	act(cat)

	//빈 인터페이스 : 어떤 값이든 허용 가능(유연성 증가)
	printValue(dog)
	printValue(cat)
	printValue(15)
	printValue("Animal")
	printValue([]Dog{})
}
