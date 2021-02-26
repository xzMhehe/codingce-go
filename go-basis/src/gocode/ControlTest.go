package main

import "fmt"

func multiplication() {

	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(" ", j, " * ", i, " = ", i*j)
		}
		fmt.Println()
	}

}

func main() {
	multiplication()
}
