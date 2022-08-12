package main

import "fmt"

func getBigger(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}

}

func main() {
	c := getBigger(3, 5)
	fmt.Println(c)
}
