package string_fn_test

import (
	"strconv"
	"strings"
	"testing"
)

func TestStringFn(t *testing.T) {
	s := "A,B,C"
	parts := strings.Split(s, ",")
	t.Log(parts)
	for _, part := range parts {
		t.Log(part)
	}
	t.Log(strings.Join(parts, "_"))
}

func TestStringConvert(t *testing.T) {
	s := strconv.Itoa(10) //整型转字符串
	t.Log("str" + s)
	if i, err := strconv.Atoi("10"); err == nil {
		t.Log(10 + i)
	}
}
