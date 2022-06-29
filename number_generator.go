package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}

}

func main() {

	fmt.Printf("Enter minimum, maximum floats and amount of numbers which you want to generate, separated by space (i.e. 5000 6000 15): ")
	reader := bufio.NewScanner(os.Stdin)
	reader.Scan()
	userInputOnePiece := reader.Text()
	userInput := strings.Fields(userInputOnePiece)

	min, err := strconv.ParseFloat(userInput[0], 64)
	max, err := strconv.ParseFloat(userInput[1], 64)
	cycleAmount, err := strconv.Atoi(userInput[2])
	checkErr(err)

	for i := 0; i < cycleAmount; i++ {
		fmt.Printf("%.2f\n", (rand.Float64()*(max-min))+min)
	}

}
