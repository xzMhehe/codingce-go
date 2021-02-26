package main

import "fmt"

/*
	if 表达式1 {
		分支1
	} else if 表达式2 {
		分支2
	} else{
		分支3
	}
*/

func ifDemo1() {
	score := 65
	if score > 90 {
		fmt.Println("A")
	} else if score > 80 {
		fmt.Println("B")
	} else if score > 70 {
		fmt.Println("C")
	} else {
		fmt.Println("D")
	}
}

func ifDemo2() {
	if score := 91; score > 90 {
		fmt.Println("A")
	} else if score > 80 {
		fmt.Println("B")
	} else if score > 70 {
		fmt.Println("C")
	} else {
		fmt.Println("D")
	}
}

/*
	for 初始语句;条件表达式;结束语句{
		循环体语句
	}
*/

func forDemo() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func forDemo2() {
	i := 0
	for ; i < 10; i++ {
		fmt.Println(i)
	}
}

// 这种写法类似于其他编程语言中的while，在while后添加一个条件表达式，满足条件表达式时持续循环，否则结束循环。
func forDemo3() {
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
}

func forDemo4() {
	for {
		if false {
			break
		}
	}
}

func switchDemo1() {
	finger := 3
	switch finger {
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
	case 4:
		fmt.Println("无名指")
	case 5:
		fmt.Println("小拇指")
	default:
		fmt.Println("无效的输入！")
	}
}

// 一个分支可以有多个值，多个case值中间使用英文逗号分隔
func testSwitch3() {
	switch n := 7; n {
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")
	case 2, 4, 6, 8:
		fmt.Println("偶数")
	default:
		fmt.Println(n)
	}
}

// 分支还可以使用表达式，这时候switch语句后面不需要再跟判断变量
func switchDemo4() {
	age := 30
	switch {
	case age < 25:
		fmt.Println("好好学习吧")
	case age > 25 && age < 35:
		fmt.Println("好好工作吧")
	case age > 60:
		fmt.Println("好好享受吧")
	default:
		fmt.Println("活着真好")
	}
}

// fallthrough语法可以执行满足条件的case的下一个case，是为了兼容C语言中的case设计的
func switchDemo5() {
	s := "a"
	switch {
	case s == "a":
		fmt.Println("a")
		fallthrough
	case s == "b":
		fmt.Println("b")
	case s == "c":
		fmt.Println("c")
	default:
		fmt.Println("...")
	}
}

func gotoDemo2() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				// 设置退出标签
				goto breakTag
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
	return
	// 标签
breakTag:
	fmt.Println("结束for循环")
}

func main() {
	ifDemo1()
	ifDemo2()
	switchDemo5()
}
