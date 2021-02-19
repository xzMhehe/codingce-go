package main

import (
	"fmt"
)

// 全局变量
var n int

// var 变量名 类型 = 表达式
var name string = "Q1mi"
var age int = 18

// 一次初始化多个变量
var name1, age1 = "Q1mi", 20

// 类型推导
var name2 = "Q1mi"
var age2 = 18

// 常量
const pi = 3.1415
const e1 = 2.7182

// 多个常量
const (
	pi1 = 3.1415
	e2  = 2.7182
)

// iota是go语言的常量计数器，只能在常量的表达式中使用
// iota在const关键字出现时将被重置为0。const中每新增一行常量声明将使iota计数一次(iota可理解为const语句块中的行索引)。 使用iota能简化定义，在定义枚举时很有用。
const (
	n1 = iota //0
	n2        //1
	n3        //2
	n4        //3
)

// 使用_跳过某些值
const (
	n5 = iota //0
	n6        //1
	_
	n8 //3
)

// iota声明中间插队
const (
	n9  = iota //0
	n10 = 100  //100
	n11 = iota //2n
	n12        //3
)
const n13 = iota //0

// 位运算
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

const (
	a, b = iota + 1, iota + 2 //1,2
	c, d                      //2,3
	e, f                      //3,4
)

// 方法
func foo() (int, string) {
	return 10, "Q1mi"
}

func main() {
	// 这是单行的注释
	/*
		多行注释
	*/

	n = 100
	fmt.Println("hello 后端码匠", n)

	// 短变量声明: 在函数内部，可以使用更简略的 := 方式声明并初始化变量。
	l := 10
	m := 200 // 此处声明局部变量m: 声明后必须使用
	fmt.Println(l, m)

	// 匿名变量
	x, _ := foo()
	_, y := foo()
	fmt.Println("x = ", x)
	fmt.Println("y = ", y)

}
