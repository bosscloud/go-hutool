package main

import "fmt"

//(int, int) 在这个函数中标志着这个函数返回 2 个 int。

func vals() (int, int) {
	return 3, 7
}
func main() {
	//这里我们通过_多赋值_操作来使用这两个不同的返回值。

	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)
	//如果你仅仅需要返回值的一部分的话，你可以使用空白标识符_。

	_, c := vals()
	fmt.Println(c)
}
