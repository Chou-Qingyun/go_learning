package type_test

import "testing"

/*
基本数据类型
bool		//布尔型
string		// 字符串
int int8 int16 int32 int64
uint uint8 uint16 uint32 uint64 uintptr
byte //alias for uint8
rune //alias for int32, represents a Unicode code point
float32 float64
complex64 complex128

类型转化

与其他主要编程语言的差异
1、Go语言不允许隐式类型转换
2、别名和原有类型也不能进行隐式类型转换
*/
//测试go 对隐式类型转换的支持与约束
type MyInt int64 //别名
func TestImplicint(t *testing.T) {
	var a int = 1
	var b int64
	//b = a   // 报错 cannot use a (type int) as type int64 in assignment
	b = int64(a) //显示类型转换
	var c MyInt
	//c = b	//即便是别名也不支持隐式类型转化 报错： cannot use b (type int64) as type MyInt in assignment
	c = MyInt(b) //显示类型转换
	t.Log(a, b, c)
}

/*
类型的预定义值
1. math.MaxInt64
2. math.MaxFloat64
3. math.MaxUint32
*/

/*
指针类型

与其他主要编程语言的差异
1. 不支持指针运算
2. string是值类型，其默认的初始化值为空字符串，而不是nil
*/
func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a //&取址符
	// aPtr = aPtr + 1 //指针运算 报错  invalid operation: aPtr + 1 (mismatched types *int and int)
	t.Log(a, aPtr) // 1 	 0xc00000a1a8
	// a的值 a的地址
	t.Logf("%T %T", a, aPtr) //输出类型  int *int
}

func TestString(t *testing.T) {
	var s string         //string 初始化是一个空字符串
	t.Log("*" + s + "*") //**
	t.Log(len(s))        //0
}
