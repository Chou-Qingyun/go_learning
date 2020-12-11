package slice_test

import "testing"

func TestSliceInit(t *testing.T) {
	var s0 []int //初始化切片
	t.Log(len(s0), cap(s0))
	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0))

	s1 := []int{1, 2, 3, 4}
	t.Log(len(s1), cap(s1))

	s2 := make([]int, 3, 5) //第一参数是类型，第二个参数是len(长度，代表可访问的个数)，第三个参数是cap（内部数组的容量capacity/连续存储空间）
	t.Log(len(s2), cap(s2))
	//t.Log(s2[0], s2[1], s2[2], s2[3]) //此处会报错， index out of range [3] with length 3，其实内部数组的容量是5是足够的，但是初始化的时候数组的长度只有3，所以超出初始化长度范围了就报错了。
}

/**
*	切片声明
	var s0 []int
	s0 = append(s0, 1)
	s := []int{}
	s1 := []int{1,2,3}
	s2 := make([]int, 2, 4)
	[]type, len, cap其中len个元素会被初始化为默认零值，未初始化元素不可以访问
*/

func TestSliceGrowing(t *testing.T) {
	s := []int{}
	for i := 0; i < 10; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s)) //每一次cap的的增长都要基于前面的元素乘以2
	}
}

func TestSliceShareMemory(t *testing.T) {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	Q2 := year[3:6]
	t.Log(Q2, len(Q2), cap(Q2))
	summer := year[5:8]
	t.Log(summer, len(summer), cap(summer))
	// 修改summer的值
	summer[0] = "Unknow"
	t.Log(Q2)   //  [Apr May Unknow]
	t.Log(year) // [Jan Feb Mar Apr May Unknow Jul Aug Sep Oct Nov Dec] 因为切片是共享存储结构，并没有复制一份地址，而是指向同一地址的，所以修改的话，都会一起改变的。
}

func TestSliceComparing(t *testing.T) {
	a := []int{1, 2, 3, 4}
	b := []int{1, 2, 3, 4}
	t.Log(a)
	t.Log(b)
	/*if a == b {				//invalid operation: a == b (slice can only be compared to nil)
		t.Log("equal")
	}*/
}
