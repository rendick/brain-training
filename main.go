package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
)

var (
	// Colors
	Green = "\033[32m"
	Red   = "\033[31m"

	Reset = "\033[0m"

	// Text
	Bold = "\033[1m"
)

var Answer int
var Continue int

func ResultCorrect(text string) {
	fmt.Printf(Green+Bold+"\n%s"+Reset+"\n------\n\n", text)
}

func ResultIncorrect(text string) {
	fmt.Printf(Red+Bold+"\n%s"+Reset+"\n------\n\n", text)
}

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Not enough arguments!")
		os.Exit(0)
	}
	for {
		convert_first, err := strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}

		convert_second, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}

		RandHundOne := rand.Intn(convert_first)
		RandHundTwo := rand.Intn(convert_second)
		fmt.Printf("Solve: %d %s %d\n", RandHundOne, os.Args[3], RandHundTwo)
		fmt.Print("Your answer: ")
		fmt.Scan(&Answer)

		switch os.Args[3] {
		case "+":
			if Answer == (RandHundOne + RandHundTwo) {
				ResultCorrect("Correct!")
			} else {
				ResultIncorrect("Incorrect!")
			}
		case "-":
			if Answer == (RandHundOne - RandHundTwo) {
				ResultCorrect("Correct!")
			} else {
				ResultIncorrect("Incorrect!")
			}
		case "*":
			if Answer == (RandHundOne / RandHundTwo) {
				ResultCorrect("Correct!")
			} else {
				ResultIncorrect("Incorrect!")
			}
		case "/":
			if Answer == (RandHundOne * RandHundTwo) {
				ResultCorrect("Correct!")
			} else {
				ResultIncorrect("Incorrect!")
			}
		}
	}
}
