# Go Language Basic

- [https://github.com/eunki7/golang_example](https://github.com/eunki7/golang_example) 에서 예제를 발췌하여 재정리하였습니다.

```
src
|---1_basic
|---2_type
|---3_function
|---4_struct&interface
|---5_goroutine
|---6_error
|---7_package
|   \---lib
|   \---lib2
```

## 1. Basic Syntax

- 상수 및 변수 선언과 정의  

```go
const a string = "const1"
const (
    x, y    int32  = 50, 90
    z, z2          = "Data", 777
)

var (
    name        string = "machine"
    height      int32
    weight      float32
    isRunning   bool
)
shortVar1 := 1908
shortVar2 := "GoStudy"
```

- 열거형(Enumeration)
  - 일정한 규칙에 따라 숫자를 증가시키는 상수의 묶음
  - `iota` : 0 부터 시작해서 1씩 증가

```go
const (
    _ = iota + 0.75 * 2 // 1.5
    DEFAULT             // 2.5
    SILVER              // 3.5
    GOLD                // 4.5
    PLATINUM            // 5.5
)
```

- 반복문(for)
  - Go 에서 유일한 반복문

```go
Loop1:
    for i := 0; i < 5; i++ {
        for j := 0; j < 5; j++ {
            if i == 2 && j == 4 {
                break Loop1
            }
            fmt.Println(i, j)
        }
    }

loc := []string{"Seoul", "Busan", "Incheon"}
for index, name := range loc {
    fmt.Println(index, name)
}

for {  // 무한반복
    fmt.Println("Infinite Loop!")
}
```

- 제어문(switch)

```go
switch a := "go"; a {
case "go":
    fmt.Println("Go!")
    fallthrough
case "java":
    fmt.Println("Java!")
    fallthrough
case "python":
    fmt.Println("Python!")
    fallthrough
case "ruby":
    fmt.Println("Scala!")
default:
    fmt.Println("일치하는 값 없음")
}
```

## 2. Data Type

- 문자열(string)

```go
var str1 string = "Hello, world!"
var str2 string = `안녕하세요`

//길이(바이트 수)
fmt.Println(len(str1)) // 13
fmt.Println(len(str2)) // 15

//길이(실제 문자 수)
fmt.Println(utf8.RuneCountInString(str1)) // 13
fmt.Println(utf8.RuneCountInString(str2)) // 5
fmt.Println([]rune(str2)) // [50504 45397 54616 49464 50836]

fmt.Println(str1[0]) // "H" == 72
```

- 배열(Array)

```go
var arr1 [5]int
arr2 := [5]int{1, 2, 3} //나머지 0 초기화
arr3 := [...]int{1, 2, 3, 4, 5}
arr4 := [5][5]int{
    {1, 2, 3, 4, 5},
    {6, 7, 8, 9, 10}, //콤마 주의
}
arr5 := [...]string{"Kim", "Lee", "Park"}

for _, v := range arr3 {
    fmt.Printf("%v \t | ", v)
}
fmt.Println()

//배열 복사
//값만 복사되고 새로 생성된다. 확인 중요 !
arrCp1 := [5]int{1, 10, 100, 1000, 10000}
arrCp2 := arrCp1

// arrCp1 : [1 10 100 1000 10000] 0xc00007c1e0
fmt.Printf("arrCp1 : %v %p\n", arrCp1, &arrCp1)
// arrCp2 : [1 10 100 1000 10000] 0xc00007c210
fmt.Printf("arrCp2 : %v %p\n", arrCp2, &arrCp2)
```

- Map
  - 레퍼런스 타입
  - 비교 연산자 사용 불가능

```go
var map1 map[string]int = make(map[string]int) // 정석
var map2 = make(map[string]int, 10)            // 자료형 생략
map3 := make(map[string]int)                   // 리터럴 형
map4 := map[string]int{
    "apple":  15,
    "banana": 30,
    "orange": 23, //콤마 주의
}

for k, v := range map4 {
    fmt.Println("Iterator : ", k, v)
}

if value, ok := map4["kiwi"]; ok {
    fmt.Println("map4 : ", value)
} else {
    fmt.Println("map4 : kiwi is not exist!")
}
```

- 슬라이스(slice)

```go
s1 := []int{1, 2, 3, 4, 5}
s2 := []int{8, 9, 10, 11, 12}
s3 := []int{13, 14, 15, 16, 17}
s1 = append(s1, 6, 7)
s2 = append(s1, s2...)      //슬라이스를 삽입할 경우 '...' 사용
s3 = append(s2, s3[0:3]...) // 추출 후 병합
```

- `len()` : 슬라이스 개수, `cap()` : 슬라이스 용량
- 용량(capacity)을 초과하는 경우 현재 용량의 2배에 해당하는 array를 생성
- 기존 배열 값들을 모두 새 배열에 복제한 후 다시 슬라이스를 할당한다.

```go
s4 := make([]int, 0, 5)
//길이 및 용량 자동 증가(용량 : 2배)
for i := 0; i < 15; i++ {
    s4 = append(s4, i)
    fmt.Printf("len : %d, cap : %d, value : %v\n", len(s4), cap(s4), s4)
}
```

- 슬라이스 복사
  - copy(복사 대상, 원본)
  - make로 공간을 할당 후 복사 해야 한다.

```go
a := []int{1, 2, 3, 4, 5}
b := make([]int, 5)
copy(b, a)
b[0] = 7
b[4] = 10

fmt.Println("a : ", a) // 원본 유지
fmt.Println("b : ", b) // 복사 된 대상 변경
// a :  [1 2 3 4 5]
// b :  [7 2 3 4 10]

c := a[0:3] // 부분 추출은 참조가 유지되므로, 원본이 변경된다.
c[1] = 7
fmt.Println("a : ", a) // 원본 변경
fmt.Println("c : ", c)
// a :  [1 7 3 4 5]
// c :  [1 7 3]
```

- 포인터(pointer)

```go
type bag struct{ width, height, weight float32 } //구조체
var p *int = new(int)
var p_bag *bag = &bag{20, 50, 30}

fmt.Println("p : ", p) // 0xc0000520c8
fmt.Println("p_bag : ", p_bag) // &{20 50 30}
```

- Call By Reference, Call By Value

```go
func rptc(n *int) {
    *n = 77
}

func vptc(n int) {
    n = 77
}

var pa int = 10
var pb int = 10

rptc(&pa)
vptc(pb)

fmt.Println("pa : ", pa)
fmt.Println("pb : ", pb)
```
