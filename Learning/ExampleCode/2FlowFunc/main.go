package main

import (
	"fmt"
	"example/2FlowFunc/trans"
)

var twoPi = 2 * trans.Pi

// 值传递
func Multiply(a, b int, reply *int) {
	*reply = a * b
}

// 可变长参数
func Min(a...int) int {
	if len(a) == 0 {
		return 0
	}
	min := a[0]
	for _,v := range a {
		if v < min {
			min = v
		}
	}
	return min
}

func Add(a, b int) {
	fmt.Printf("The sum of %d and %d is %d\n", a, b, a+b)
}

// 函数作为参数接收
func callback(y int, f func(int, int)){
	fmt.Println("before callback")
	f(y, 2)
	fmt.Println("end callback")
}

// 闭包

func main() {
	// 验证trase包的init
	fmt.Println("Test trans init")
	fmt.Printf("2*pi = %f\n", twoPi)

	// defer在所在函数退出前最后一个执行defer函数
	defer func(){
		println("Defer func")
	}()
	
	// 值传递，修改变量需要传递变量的指针
	n := 0
	reply := &n
	Multiply(3, 5, reply)
	fmt.Println("Multiply:", *reply)

	// 可变长参数
	data := []int{7,1,4,6,3,8,9}
	m := Min(data...)
	fmt.Printf("min of data is %d\n", m)

	// 函数可以作为参数传递
	callback(100, Add)

	// 闭包, 即函数的嵌套
	x := 10
	f := func() (func()) {
		y := 60
		sum := x + y
		return func() {
			fmt.Printf("x, y : %d, %d. Sum is: %d\n", x, y, sum)
		}
	}()

	f()

}