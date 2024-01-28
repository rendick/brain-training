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
				fmt.Printf(Green + Bold + "\nCorrect!" + Reset + "\n------\n\n")
			} else {
				fmt.Printf(Red + Bold + "\nIncorrect! Try again." + Reset + "\n------\n\n")
			}
		case "-":
			if Answer == (RandHundOne - RandHundTwo) {
				fmt.Printf(Green + Bold + "\nCorrect!" + Reset + "\n------\n\n")
			} else {
				fmt.Printf(Red + Bold + "\nIncorrect! Try again." + Reset + "\n------\n\n")
			}
		case "*":
			if Answer == (RandHundOne / RandHundTwo) {
				fmt.Printf(Green + Bold + "\nCorrect!" + Reset + "\n------\n\n")
			} else {
				fmt.Printf(Red + Bold + "\nIncorrect! Try again." + Reset + "\n------\n\n")
			}
		case "/":
			if Answer == (RandHundOne * RandHundTwo) {
				fmt.Printf(Green + Bold + "\nCorrect!" + Reset + "\n------\n\n")
			} else {
				fmt.Printf(Red + Bold + "\nIncorrect! Try again." + Reset + "\n------\n\n")
			}
		}
	}
}
