package main

import (
	"fmt"
	"os"
)

/*
To execute hackerrank projecteuler+ solution use:
	go build main.go
	./main

To execute projecteuler.net solution use:
	go build main.go
	./main 1
*/

func main() {
	switch len(os.Args) {
	case 1:
		solve()
	default:
		solvePE()
	}
}

func solve() {
	// Hackerrank PE+ solution
	var t int
	var n int
	fmt.Scanf("%d", &t)
	for ; t > 0; t-- {
		fmt.Scanf("%d", &n)
		fmt.Println(getSumMultiple(n))
	}
}

func solvePE() {
	// projecteuler.net solution
	fmt.Println(getSumMultiple(1000))
}

func getSumMultiple(n int) int {
	n--
	n3 := n / 3
	n5 := n / 5
	n15 := n / 15
	n3 = 3 * (n3 * (n3 + 1) / 2)
	n5 = 5 * (n5 * (n5 + 1) / 2)
	n15 = 15 * (n15 * (n15 + 1) / 2)
	return n3 + n5 - n15
}
