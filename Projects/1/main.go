package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to my quiz game!")
	var name string
	var age int

	fmt.Printf("Enter your name: ")
	fmt.Scan(&name)
	fmt.Printf("Hello, %v, welcome to the game!\n", name)

	fmt.Printf("Enter your age: ")
	fmt.Scan(&age)
	if age < 10 {
		fmt.Println("You're not allowed to play this game.")
		return
	} else {
		fmt.Println("You're allowed to play this game.")
	}

	score := 0
	num_questions := 2

	fmt.Printf("What is better, the RTX 3080 or RTX 3090? ")
	var answer string
	var answer2 string
	fmt.Scan(&answer, &answer2)

	if answer+" "+answer2 == "RTX 3090" || answer+" "+answer2 == "rtx 3090" {
		fmt.Println("Correct!")
		score++
	} else {
		fmt.Println("Incorrect!")
	}

	fmt.Printf("How many cores does the Ryzen 9 3900x have? ")
	var cores int
	fmt.Scan(&cores)

	if cores == 12 {
		fmt.Println("Correct!")
		score++
	} else {
		fmt.Println("Incorrect!")
	}
	fmt.Printf("You scored: %v out of %v\n", score, num_questions)

	percent := (float64(score) / float64(num_questions)) * 100
	fmt.Printf("You scored: %.0f%%", percent)
}
