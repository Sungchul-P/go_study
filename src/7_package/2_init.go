package main

import (
	"7_package/lib"
	"fmt"
)

var num int32

func init() {
	num = 30
}

func main() {
	fmt.Print("10 보다 큰 수? : ", lib.CheckNum(num))
}
