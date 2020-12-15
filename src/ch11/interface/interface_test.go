package interface_test

import "testing"

type Programmer interface {
	WriteHelloWorld() string
}

type GoProgrammer struct {
}

func (g *GoProgrammer) WriteHelloWorld() string {
	return "fmt.Println(\"Hello World\")"
}

func TestClient(t *testing.T) {
	var p Programmer
	p = new(GoProgrammer)
	t.Log(p.WriteHelloWorld())
}

/*
Go 接口
与其他主要编程语言的差异
1. 接口为非入侵性，实现不依赖于定义
2. 所以接口的定义可以包含在接口使用者包内
*/

/*自定义类型
1. type IntConvertionFn func(n int) int
2. type MyPoint int
*/
