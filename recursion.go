package main

import (
	"fmt"
)

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	n := 7
	fmt.Println(fmt.Sprintf("fact(%d) =", n), fact(n))
}
