package main

import (
	"errors"
	"fmt"
	"os"
)

func total(x, y int) (int, error) {

	if y == 0 {
		return 0, errors.New("cannot divided by 0")
	}

	return x / y, nil

}

func divide(x, y int) int {

	defer func() {
		fmt.Println(recover())
	}()
	return x / y
}

func openfile() {

	_, err := os.Open("uses/file.txt")

	if err != nil {

		panic(err)
	}
}

func main() {

	result, error := total(10, 0)

	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(result)
	}

	divide(10, 1)
	divide(3, 0)
	openfile()
}
