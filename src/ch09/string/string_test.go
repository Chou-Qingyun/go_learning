package string_fn

import "testing"

func TestString(t *testing.T) {
	var s string
	t.Log(s) //初始化默认零值
	s = "hello"
	t.Log(len(s))
	//s[1] = '3' cannot assign to s[1]  *string 是不可变的byte slice
	s = "\xE4\xB8\xA5" //可以存储任何二进制数据
	t.Log(s)
	t.Log(len(s)) //求出来的是byte数
	s = "中"
	t.Log(len(s)) //是byte数

	c := []rune(s)
	t.Log(len(c))
	t.Logf("中 Unicode %x", c[0])
	t.Logf("中 UTF8 %x", s)
}

/*
	Unicode UTF8
	1. Unicode 是一种字符集(code point)
	2. UTF8是unicode的存储实现（转换为字节序列的规则）
*/

func TestStringToRune(t *testing.T) {
	s := "中华人民共和国"
	for _, c := range s {
		t.Logf("%[1]c %[1]d", c) // %[1]c 或者%[1]d  代表的是都取得是后面跟随的第一个参数c
		t.Logf("%[1]c %[1]x", c)
	}
}
