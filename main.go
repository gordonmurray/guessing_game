package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	fmt.Println("Guess the number!")

	rand.Seed(time.Now().UTC().UnixNano())

	secret_number := rand.Intn(100)

	var guess int

	for {

		fmt.Println("Please input your guess?")

		fmt.Scan(&guess)

		fmt.Println("You guessed:", guess)

		if guess == secret_number {
			fmt.Println("You win!")
			break
		}

		switch {
		case guess == secret_number:
			fmt.Println("You win!")
			break
		case guess < secret_number:
			fmt.Println("Too small!")
		case guess > secret_number:
			fmt.Println("Too big!")
		}
	}

}
