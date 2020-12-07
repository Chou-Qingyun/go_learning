package condition

import "testing"

/*
if 条件语句
if condition {
	//code to be executed if condition is true
} else {
	//code to be executed if condition is false
}

if condition-1 {
	//code to be executed if condition-1 is true
} else if condition-2 {
	//code to be executed if condition-2 is true
} else {
	//code to be excuted if both condition1 and condition2 are false
}
与其他主要编程语言的差异
1. condition 表达式结果必须为布尔值
2. 支持变量赋值
	if var declaration; condition {
		//code to be executed if condition is true
	}
*/

func TestIfMultiSec(t *testing.T) {
	if a := 1 == 1; a {
		t.Log("1==1")
	}
	/*一般的应用场景是在方法有返回值的情况
	if v,err := someFunc(); err == nil {
		t.Log("1==1")
	} else {
		t.Log("报错了")
	}*/
}

/*
switch 条件
与其他主要编程语言的差异
1. 条件表达式不限制为常量或者整数；
2. 单个case中，可以出现多个结果选项，使用逗号分隔；
3. 与C语言等规则相反，Go语言不需要用break来明确退出一个case；
4. 可以不设定switch之后的条件表达式，在此种情况下，整个switch结构与多个if...else...的逻辑作用等同

switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
		//break
	case "":
		fmt.Println("Linux.")
	default:
		//freebsd, openbsd,
		//plan9,windows...
		fmt.printf("%s.", os)
}

switch {
	case 0 <= Num && Num <= 3:
		fmt.printf("0-3")
	case 4 <= Num && Num <= 6
		fmt.printf("4-6")
	case 7 <= Num && Num <= 9
		fmt.Printf("7-9")
}
*/

func TestSwitchMultiCase(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch i {
		case 0, 2:
			t.Log("Even")
		case 1, 3:
			t.Log("Odd")
		default:
			t.Log("it is not 0-3")
		}
	}
}

func TestSwitchCaseCondition(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch {
		case i%2 == 0:
			t.Log("Even")
		case i%2 == 1:
			t.Log("Odd")
		default:
			t.Log("unknow")
		}
	}
}
