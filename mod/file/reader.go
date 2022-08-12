package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	fileName := "test.txt"

	data, err := ioutil.ReadFile(fileName)

	fmt.Printf("%s\n", data, err)

}
