package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
)

var Answear int

func main() {
	if len(os.Args) != 3 {
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
		fmt.Printf("Solve: %d + %d\n", RandHundOne, RandHundTwo)
		fmt.Scan(&Answear)

		if Answear == (RandHundOne + RandHundTwo) {
			fmt.Println("Good job!")
			break
		} else {
			fmt.Println("Nope")
		}
	}
}
