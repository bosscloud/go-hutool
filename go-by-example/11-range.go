package main

import "fmt"

func main() {
	//这里我们使用 range 来对 slice 中的元素求和。 对于数组也可以采用这种方法。

	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)
	//range 在数组和 slice 中提供对每项的索引和值的访问。 上面我们不需要索引，所以我们使用 空白标识符 _ 来忽略它。有时候我们实际上是需要这个索引的。

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}
	//range 在 map 中迭代键值对。

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
	//range 也可以只遍历 map 的键。

	for k := range kvs {
		fmt.Println("key:", k)
	}
	//range 在字符串中迭代 unicode 码点(code point)。 第一个返回值是字符的起始字节位置，然后第二个是字符本身。

	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
