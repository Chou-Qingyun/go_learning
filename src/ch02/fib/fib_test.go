package fib

import (
	"testing"
)

// 也可以先外部声明全局变量
// var a int
func TestFibList(t *testing.T) {
	/*
		var a int = 1
		var b int = 1
	*/
	// a = 1
	var (
		a int = 1
		b int = 1
	)
	/*
		go 有类型推断能力
		以上的表达式等同于如下
		a:=1
		b:=1
	*/
	//fmt.Print(a)
	t.Log(a)
	for i := 0; i < 5; i++ {
		//fmt.Print(" ", b)
		t.Log(b)
		tmp := a
		a = b
		b = tmp + a
	}
	//fmt.Println()
}

/*
	变量赋值
	与其他的主要编程语言的差异
	· 赋值可以进行自动类型推断
	· 在一个赋值语句中可以多个变量进行同时赋值
*/
func TestExchange(t *testing.T) {
	/*最常见的交换两个变量的值*/
	a := 1
	b := 2
	/*tmp:=a
	a = b
	b = tmp*/
	a, b = b, a //go语言中更便利的方式
	t.Log(a, b)
}
