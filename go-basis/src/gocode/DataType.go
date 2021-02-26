package main

import (
	"fmt"
	"math"
)

func main() {
	// 整型
	fmt.Println("===整型===")
	// 十进制
	var a int = 10
	fmt.Printf("%d \n", a) // 10
	fmt.Printf("%b \n", a) // 1010  占位符%b表示二进制

	// 八进制  以0开头
	var b int = 077
	fmt.Printf("%o \n", b) // 77

	// 十六进制  以0x开头
	var c int = 0xff
	fmt.Printf("%x \n", c) // ff
	fmt.Printf("%X \n", c) // FF

	// 浮点型
	fmt.Println("===浮点型===")
	fmt.Printf("%f\n", math.Pi)
	fmt.Printf("%.2f\n", math.Pi)

	// 复数	复数有实部和虚部，complex64的实部和虚部为32位，complex128的实部和虚部为64位
	fmt.Println("===复数===")
	var c1 complex64
	c1 = 1 + 2i
	var c2 complex128
	c2 = 2 + 3i
	fmt.Println(c1)
	fmt.Println(c2)

	// 布尔值 Go语言中以bool类型进行声明布尔型数据，布尔型数据只有true（真）和false（假）两个值。
	// 布尔类型变量的默认值为false。
	// Go 语言中不允许将整型强制转换为布尔型.
	// 布尔型无法参与数值运算，也无法与其他类型进行转换。
	fmt.Println("===布尔值===")
	var t1 bool
	var t2 bool = true

	fmt.Println(t1, t2)

	// 字符串
	// Go语言中的字符串以原生数据类型出现，使用字符串就像使用其他原生数据类型（int、bool、float32、float64 等）一样。
	s1 := "hello"
	s2 := "你好"
	fmt.Println("===字符串===")
	fmt.Println(s1, s2)

	// 字符串转义符
	// Go 语言的字符串常见转义符包含回车、换行、单双引号、制表符等
	fmt.Println("===字符串转义符===")
	fmt.Println("\\r: 回车符（返回行首）")
	fmt.Println("\\n: 换行符（直接跳到下一行的同列位置）")
	fmt.Println("\\t: 制表符")
	fmt.Println("\\' : 单引号")
	fmt.Println("\\\\ : 反斜杠")

	// 字符串转义符
	// Go 语言的字符串常见转义符包含回车、换行、单双引号、制表符等
	fmt.Println("===字符串转义符===")
	fmt.Println("\\r: 回车符（返回行首）")
	fmt.Println("\\n: 换行符（直接跳到下一行的同列位置）")
	fmt.Println("\\t: 制表符")
	fmt.Println("\\' : 单引号")
	fmt.Println("\\\\ : 反斜杠")

	// 多行字符串
	// Go语言中要定义一个多行字符串时，就必须使用反引号字符
	s5 := `第一行
	第二行
	第三行
`
	fmt.Println(s5)

	// 字符串的常用操作
	fmt.Println("===字符串的常用操作===")
	fmt.Println("len(str): 求长度")
	fmt.Println("+或fmt.Sprintf: 拼接字符串")
	fmt.Println("strings.Split: 分割")
	fmt.Println("strings.contains' : 判断是否包含")
	fmt.Println("strings.HasPrefix,strings.HasSuffix : 前缀/后缀判断")
	fmt.Println("strings.Index(),strings.LastIndex() : 子串出现的位置")
	fmt.Println("strings.Join(a[]string, sep string) : join操作")

	// byte和rune类型
	//uint8类型，或者叫 byte 型，代表了ASCII码的一个字符。
	//rune类型，代表一个 UTF-8字符。
	fmt.Println("===byte和rune类型===")
	a1 := '中'
	b1 := 'x'
	fmt.Println(a1, b1)
	traversalString()

	// 类型转换
	//Go语言中只有强制类型转换，没有隐式类型转换。该语法只能在两个类型之间支持相互转换的时候使用。
	//强制类型转换的基本语法如下：
	//T(表达式)	其中，T表示要转换的类型。表达式包括变量、复杂算子和函数返回值等.
	fmt.Println("===类型转换===")
	sqrtDemo()
}

// 遍历字符串
func traversalString() {
	s := "hello沙河"
	for i := 0; i < len(s); i++ { //byte
		fmt.Printf("%v(%c) ", s[i], s[i])
	}
	fmt.Println()
	for _, r := range s { //rune
		fmt.Printf("%v(%c) ", r, r)
	}
	fmt.Println()
}

func sqrtDemo() {
	var a, b = 3, 4
	var c int
	// math.Sqrt()接收的参数是float64类型，需要强制转换
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}
