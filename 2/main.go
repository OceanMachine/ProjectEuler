package main

import "fmt"

func main() {
	var sum int
	limit := 4000000
	sum = sumFiboEvens(limit)
	fmt.Println(sum)
}

func sumFiboEvens(limit int) int {
	var sum int
	var fibo = 1
	for lastFibo := 1; lastFibo < limit; {
		swap := fibo
		fibo += lastFibo
		if even(fibo) {
			sum += fibo
		}
		lastFibo = swap
	}
	return sum
}

func even(n int) bool {
	if n%2 == 0 {
		return true
	} else {
		return false
	}
}
