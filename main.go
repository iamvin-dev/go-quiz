package main

import (
	"fmt"
)

var (
	user string
	difficulty int
)

func login() {
	fmt.Println("Please enter your name: ")
	fmt.Scanln(&user)
	fmt.Println("Welcome to the Quiz, " + user + "!")
}

func handleQuestions() {
	switch(difficulty) {
		case 1:
			fmt.Println("You have chosen Easy!")
			fmt.Println("What is 1 + 1?")
			var answer int
			fmt.Scanln(&answer)
			if answer == 2 {
				fmt.Println("Correct!")
			} else {
				fmt.Println("Incorrect!")
			}
		case 2:
			fmt.Println("You have chosen Medium!")
			fmt.Println("What is 2 + 2?")
			var answer int
			fmt.Scanln(&answer)
			if answer == 4 {
			fmt.Println("Correct!")
			} else {
				fmt.Println("Incorrect!")
			}
		case 3:
			fmt.Println("You have chosen Hard!")
			fmt.Println("What is 3 + 3?")
			var answer int
			fmt.Scanln(&answer)
			if answer == 6 {
				fmt.Println("Correct!")
			} else {
				fmt.Println("Incorrect!")
			}
		default:
			fmt.Println("Invalid difficulty level!")
		}
}

func handleDifficulty() {
	fmt.Println("Choose a difficulty level: 1: Easy, 2: Medium, 3: Hard")
	fmt.Scanln(&difficulty)
	switch(difficulty) {
		case 1:
			fmt.Println("You have chosen Easy!")
		case 2:
			fmt.Println("You have chosen Medium!")
		case 3:
			fmt.Println("You have chosen Hard!")
		default:
			fmt.Println("Invalid difficulty level!")
	}
}

func main() {
	fmt.Println("Welcome to my quiz game!")
	login()
	handleDifficulty()
	handleQuestions()
}
