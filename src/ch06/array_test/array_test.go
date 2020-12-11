package array_test

import "testing"

/*
var a [3]int //声明并初始化为默认零值
a[0] = 1
b := [3]int{1, 2, 3}			//声明同时初始化
c := [2][2]int{{1,2}, {3,4}}	//多维数组初始化
*/

func TestArrayInit(t *testing.T) {
	var arr [3]int
	arr1 := [4]int{1, 2, 3, 4}
	arr3 := [...]int{1, 3, 4, 5}
	var arr4 = [3]int{1, 2, 3}

	t.Log(arr4)
	t.Log(arr[1], arr[2])
	t.Log(arr1, arr3)
}

/*
数组元素遍历
与其他主要编程语言的差异

*/

func TestArrayTravel(t *testing.T) {
	arr3 := [...]int{1, 3, 4, 5} //不指定元素个数
	/*for i := 0; i < len(arr3); i++ {
		t.Log(arr3[i])
	}*/
	for idx /*索引*/, e /*元素*/ := range arr3 { //此处申明了idx索引，由于go语言的编程约束性，在以下的代码中若没有使用idx变量会报错
		//在go语言中可以使用_(下划线)来替代idx代表并不关心这个值的结果，但是又知道这个值有一个返回值的占位，然后在接下来的代码中不使用_就不会报错
		t.Log(e)
		t.Log(idx)
	}
}

/*数组截取
a[开始索引（包含）,结束索引（不包括）]
a := [...]int{1, 2, 3, 4, 5}
a[1:2]	//2
a[1:3]	//2,3
a[1:len(a)]		//2,3,4,5
a[1:]			//2,3,4,5
a[:3]			//1,2,3
*/

func TestArraySection(t *testing.T) {
	arr3 := [...]int{1, 3, 4, 5}
	arr3_section := arr3[3:]
	// arr3_section_1 := arr3[:-1]		// invalid slice index -1 (index must be non-negative)此处报错，go语言不知道负数指向arr3[-1:]一样报错
	t.Log(arr3_section)
}

func TestArraySlice(t *testing.T) {
	arr := []int{} //声明一个切片
	t.Log(arr, len(arr), cap(arr))
	for i := 1; i <= 5; i++ {
		arr = append(arr, i)
		t.Log(arr, len(arr), cap(arr))
	}
}
