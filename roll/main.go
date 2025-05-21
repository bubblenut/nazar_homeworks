package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
)

func GenerateRandomRoll(dice int) int {
	return rand.Intn(dice-1) + 1
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Введите целое число")
	}

	var rolls []int
	var rolls_output []string
	for i := 1; i < len(os.Args); i++ {
		dice, err := strconv.Atoi(os.Args[i])
		if err != nil {
			log.Fatal(err.Error())
		}
		current_roll := GenerateRandomRoll(dice)
		rolls = append(rolls, current_roll)
		rolls_output = append(rolls_output, strconv.Itoa(current_roll) + "/" + strconv.Itoa(dice) + ",")
	}

	sum := 0
	for i := 0; i < len(rolls); i++ {
		sum += rolls[i]
	}



	fmt.Println("Total:", sum)
	fmt.Print("Rolls:", rolls_output)
}
