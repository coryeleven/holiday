package main

import "fmt"
import "github.com/q1mi/hello" // 导入github上的第三方包

import "holiday/summer" //导入当前项目下的包

func main() {
	fmt.Println("现在是假期时间～")
	hello.SayHi()

	summer.Diving()
}
