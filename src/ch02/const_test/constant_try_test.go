package const_test

import "testing"

/*
	常量定义
	与其他主要编程语言的差异
	快速设置连续值
*/
const (
	Monday = iota + 1
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)
const (
	Open = 1 << iota
	Close
	Pending
)
const (
	Readable   = 1 << iota // 1
	Writable               // 2
	Executable             // 4
)

func TestConstatTry(t *testing.T) {
	t.Log(Monday, Tuesday)
}

func TestConstantTry1(t *testing.T) {
	a := 7
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
	/*
		a&Readable
		  0111
		& 0001
		------
		  0001
	*/
}
