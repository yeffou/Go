package main

import (
	"fmt"
	"sync"
)

func squareNumber(num int, wg *sync.WaitGroup) {
	defer wg.Done()
	result := num * num
	fmt.Printf("Square of %d is %d\n", num, result)
}

func main() {

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var wg sync.WaitGroup

	for _, num := range numbers {
		wg.Add(1)
		go squareNumber(num, &wg)
	}

	wg.Wait()

	fmt.Println("All goroutines have finished.")
}
