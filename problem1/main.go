package main

import "fmt"

func main() {
	const max = 1000
	total := 0
	for i := 0; i < max; i++ {
		if i%3 == 0 || i%5 == 0 {
			total += i
		}
	}
	fmt.Printf("Total: %d\n", total)
}
