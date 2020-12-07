package operator_test

import "testing"

/*
算术运算符
Go 语言没有前置的++，-- （++a）

用 == 比较数组
· 相同维数且含有相同个数元素的数组才可以比较
· 每个元素都相同的才相等
*/

func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 3, 2, 4}
	//c := [...]int{1, 2, 3, 4, 5}
	d := [...]int{1, 2, 3, 4}
	t.Log(a == b) // false
	//t.Log(a == c)  	//invalid operation: a == c (mismatched types [4]int and [5]int)
	t.Log(a == d) // true
}

/*
位运算符
与其主要编程语言的差异
&^ 按位置零 (按位清零)
对于左右两个操作数来说，只要是右边位上的操作数为1，那么无论左边位上的操作数是0还是1，结果都是0。
也就是只要是右边的位上的操作数设为1的二进制位，都把左边对应的二进制清零
如果右边的操作位上是0，那么左边的操作位上的值原来是什么，结果就是什么
1 &^ 0 -- 1
1 &^ 1 -- 0
0 &^ 1 -- 0
0 &^ 0 -- 0
*/

const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestBitClear(t *testing.T) {
	a := 7
	a = a &^ Readable
	a = a &^ Writable
	a = a &^ Executable
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}
