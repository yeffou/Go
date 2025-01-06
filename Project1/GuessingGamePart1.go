package main

import (
	"fmt"
	"math/rand"
)

func main() {
	min := 0
	max := 100
	nbr := rand.Intn(100)

	attempts := 0

	var guess int

	fmt.Println("Guess the number between 0 and 100")
	for {

		fmt.Scan(&guess)
		attempts++
		if guess > max || guess < min {
			fmt.Println("Guess Out of Range")
			continue
		}

		if guess < nbr {
			fmt.Println("Go Higher!!")
		} else if guess > nbr {
			fmt.Println("Go Lower!!")
		} else {
			fmt.Printf("Congratulations! You're right after %d guesses", attempts)

			break
		}
	}
}
