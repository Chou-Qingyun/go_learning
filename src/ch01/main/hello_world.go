package main  //包，表明代码所在的模块（包）

import (
	"fmt"
	"os"
) //引入代码依赖

//功能实现
func main() {
	//fmt.Println(os.Args)
	if len(os.Args) > 1{
		fmt.Println("Hello world",  os.Args[1])
	}
//	返回值
	os.Exit(-1) //传入异常值-1
}

/*
应用程序的入口
1. 必须是main包：package main
2. 必须是main方法：func main()
3. 文件名不一定是main.go,目录名也不一定是main

*/
/*
退出返回值

与其他的主要编程语言的差异
· Go中main函数不支持任何返回值
· 通过os.Exit来返回状态

获取命令行参数

与其他主要编程语言的差异
· main函数不支持传入参数
	 func main(arg []string)
· 在程序中直接通过os.Args获取命令行参数
*/
