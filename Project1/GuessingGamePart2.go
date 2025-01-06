package main

import (
	"fmt"
	"math/rand"
)

var score int

func main() {

	fmt.Println("Set a range of difficulty :")
	var difficulty int
	fmt.Scan(&difficulty)

	fmt.Println("Set the limit of attempts: ")
	var limit int
	fmt.Scan(&limit)

	var nbr int
	var max int

	if difficulty == 1 {
		max = 100
		nbr = rand.Intn(100)
	}
	if difficulty == 2 {
		max = 200
		nbr = rand.Intn(200)
	}
	if difficulty == 3 {
		max = 300
		nbr = rand.Intn(300)
	}

	min := 0
	attempts := 0

	var guess int
	var answer int

	fmt.Println("Guess the number")
	for {

		fmt.Scan(&guess)
		attempts++

		if guess < nbr {
			fmt.Println("Go Higher!!")
		} else if guess > nbr {
			fmt.Println("Go Lower!!")
		} else {
			score = attempts
			fmt.Printf("Congratulations! You're right after %d guesses \n", score)

			fmt.Println("Do you want to play again ?")
			fmt.Scan(&answer)
			if answer == 1 {
				main()
			} else {
				break
			}

		}

		if attempts > limit-1 {
			fmt.Printf("You Lost, your highest score was %d \n", score)
			fmt.Println("Do you want to play again ?")

			fmt.Scan(&answer)
			if answer == 1 {
				main()
			} else {
				break
			}
		}

		if guess > max || guess < min {
			fmt.Println("Guess Out of Range")
			continue
		}

	}
}
