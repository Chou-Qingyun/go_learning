package try_test

import "testing"

func TestFirstTry(t *testing.T) {
	t.Log("My first try!")
}

/*
编写测试程序(单元测试)
1. 源码文件以_test 结尾：xxx_test.go
2. 测试方法名以Test开头：func TestXXX(t *testing.T) { ... }
3. 运行测试文件：go test -v xxx_test.go
*/
