package main

/**
 * Go by Example 中文版：值
 */
import "fmt"

func main() {
	//字符串可以通过 + 连接。

	fmt.Println("go" + "lang")
	//整数和浮点数

	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)
	//布尔型，以及常见的布尔操作。

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
