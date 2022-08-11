/*
* bool                  #布尔
* string                #字符串
* (u)int8...int64...    #整数
* float32,float64       #浮点数
* complex64,complex128  #复数，用于实现特殊效果 i = √-1；3+4i
* uintptr               #指针
* byte                  #8位
* rune                  #32位，可以理解成char，字符
 */

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
