package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

//外部变量  注：并非全局变量而是包内变量
var (
	aa = 3
	bb = true
	ss = "abc"
)

//对应类型的默认值
func variableDefaultValue() {
	var a int
	var s string
	var c bool
	//fmt.Println(a, s, c)
	fmt.Printf("%d %q %f\n", a, s, c)
}

//给变量赋值
func variableInitValue() {
	var a, b int = 3, 4 //多个变量赋值
	var s string = "abc"
	fmt.Println(a, b, s)
}

//去除变量类型
func variableTypeRemove() {
	//方式一
	//var a, b, c, s = 3, 4, true, "abc"
	//方式二
	a, b, c, s := 3, 4, true, "abc"
	b = 5
	fmt.Println(a, b, c, s)
}

//复数定义
func euler() {
	c := 3 + 4i
	fmt.Println(cmplx.Abs(c))
	//欧拉公式
	fmt.Printf("%.3f\n", cmplx.Exp(1i*math.Pi)+1)
}

//强制类型转换
func triangle() {
	a, b := 3, 4
	var c int
	//float不准确所以需要向上取整
	c = int(math.Ceil(math.Sqrt(float64(a*a + b*b))))
	fmt.Println(c)
}

//常量
func consts() {
	const (
		filename = "app.txt"
		a, b     = 3, 4
	)
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

//枚举类型
func enums() {
	const (
		cpp = iota
		_
		java
		python
		golang
	)

	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(cpp, java, python, golang)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	fmt.Println("Hello World")
	variableDefaultValue()
	variableInitValue()
	variableTypeRemove()
	fmt.Println(aa, bb, ss)
	euler()
	triangle()
	consts()
	enums()
}
