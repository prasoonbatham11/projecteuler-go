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
	var t, n int
	fmt.Scanf("%d", &t)
	for ; t > 0; t-- {
		fmt.Scanf("%d", &n)
		fmt.Println(getLargestPF(n))
	}
}

func getLargestPF(n int) int {
	primes := func() []int {
		primes := make([]int, 0)
		seiveSize := 1000000
		seive := make([]bool, seiveSize)
		for i := range seive {
			seive[i] = true
		}
		seive[0] = false
		seive[1] = false
		var p, j int
		for p = 2; p*p < seiveSize; p++ {
			if seive[p] {
				for j = p * p; j < seiveSize; j += p {
					seive[j] = false
				}
			}
		}
		for p = 2; p < seiveSize; p++ {
			if seive[p] {
				primes = append(primes, p)
			}
		}
		return primes
	}()

	lpf := 0
	for i := 0; primes[i]*primes[i] <= n; i++ {
		if n%primes[i] == 0 {
			lpf = primes[i]
			for n%primes[i] == 0 {
				n /= primes[i]
			}
		}
	}
	if n > 1 {
		lpf = n
	}

	return lpf
}

func solvePE() {
	// projecteuler.net solution
	fmt.Println(getLargestPF(600851475143))
}
