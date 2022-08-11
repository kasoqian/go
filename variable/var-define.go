/* 定义变量的N种方式 */
package main

import (
	"fmt"
)

/* 包内全局变量 */
var Project = "Variable Define"

var (
	time = 2022
	why  = "learn go"
)

/* var {variable} type 定义 */
func variable() {
	var a int
	var b string
	fmt.Printf("%d %q \n", a, b)
}

/* var {variable} type = {value} 定义 */
func variableInitValue() {
	var a int = 123
	var s string = "hello"
	fmt.Println(a, s)

}

/* var {variables} = {values} */
func variableTypeDeduction() {
	var a, b, c, d = 123, "22", true, "666"

	fmt.Println(a, b, c, d)
}

/* {variable} = {values} */
func variableShorter() {
	a, b, c, d, e := "dsadsadsa", 21313, true, 666, false

	fmt.Println(a, b, c, d, e)
}

func main() {
	fmt.Println("Hello World!")
	variable()
	variableInitValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(why, time)
}
