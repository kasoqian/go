/* 常量
* go语言中常量一般不全部大写
* 其他语言中常量一般全部大写
 */
package main

import "fmt"

/* 常量类型 */
func consts() {
	/* 数值可以做各种类型使用，动态类型 */
	const fileName = "kaso.json"
	const a, b = 1, 3

	fmt.Println(fileName, a, b)
}

/* 枚举类型 */
func enums() {
	/* 枚举的第一种方法 */
	/* 	const (
		cpp    = 1
		golang = 2
		js     = 3
	) */

	/* 自增枚举的第二种写法 */
	const (
		cpp = iota
		golang
		js
	)

	/* iota是自增值的种子 */
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
	)

	fmt.Println(cpp, golang, js)
	fmt.Println(b, kb, mb, gb)
}

func main() {
	consts()
	enums()
}
