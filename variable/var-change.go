package main

import (
	"fmt"
)

func main() {
	var age uint = 20

	name := age.(string)

	fmt.Println(name)
}
