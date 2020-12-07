package loop

import "testing"

/*
循环
与其他主要编程语言的差异
Go语言仅支持循环关键字 for
for j := 7; j<=9; j++ //此处需要注意的是go不需要使用小括号（）将循环表达式包裹

while条件循环
while (n<5)
n := 0
for n < 5 {
	n++
	fmt.Println(n)
}

无限循环
while(true)
n := 0
for {
 ...
}
*/

func TestWhileLoop(t *testing.T) {
	n := 0
	for n < 5 {
		t.Log(n)
		n++
	}
}
