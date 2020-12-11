package my_map

import "testing"

func TestInitMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	t.Log(m1[2])
	t.Logf("len m1=%d", len(m1))
	m2 := map[int]int{}
	m2[4] = 16
	t.Logf("len m2=%d", len(m2))
	m3 := make(map[int]int, 10)
	t.Logf("len m3=%d", len(m3)) //cap函数不能用于求map的capacity
}

func TestAccessNotExistingKey(t *testing.T) {
	// 在访问的key不存在时，仍会返回零值不能通过返回nil 来判断元素是否存在
	m1 := map[int]int{}
	t.Log(m1[1])
	m1[2] = 0
	t.Log(m1[2])
	m1[3] = 0
	if v, ok := m1[3]; ok { // v为key对应的值，ok代表这个key存不存在
		t.Logf("Key 3's value is %d", v)
	} else {
		t.Log("key 3 is not existing")
	}
}

// map遍历

func TestTravelMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	for k, v := range m1 {
		t.Log(k, v)
	}
}

func TestMapFunValue(t *testing.T) {
	m := map[int]func(op int) int{}
	m[0] = func(op int) int {
		return op
	}
	m[1] = func(op int) int {
		return op * op
	}
	m[2] = func(op int) int {
		return op * op * op
	}
	t.Log(m[0](2), m[1](2), m[2](2))
}
