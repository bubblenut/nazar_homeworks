package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func GenerateRandomRoll(dice int) int {
	return rand.Intn(dice-1) + 1
}

func ValidateInput() {
	if len(os.Args) < 2 {
		log.Fatal("Введите целое число")
	}
}

func RollAllDices() ([]int, []string) {
	var Rolls []int
	var RollsOutput []string
	for i := 1; i < len(os.Args); i++ {
		dice, err := strconv.Atoi(os.Args[i])
		if err != nil {
			log.Fatal(err.Error())
		}
		current_roll := GenerateRandomRoll(dice)
		Rolls = append(Rolls, current_roll)
		RollsOutput = append(RollsOutput, strconv.Itoa(current_roll)+"/"+strconv.Itoa(dice)+",")
	}
	return Rolls, RollsOutput
}

func SumRolls(rolls []int) int {
	sum := 0
	for i := 0; i < len(rolls); i++ {
		sum += rolls[i]
	}
	return sum
}

func PrintOutput(sum int, rolls_output []string) {
	fmt.Println("Total:", sum)
	fmt.Print("Rolls:", rolls_output)
}

func SliceDicesAsNumbers() ([]int, []int) {
	var Coefficients []int
	var DiceSize []int
	for i := 1; i < len(os.Args); i += 1 {
		current_dices := strings.Split(os.Args[i], "d")
		current_coefficient, IDK := strconv.Atoi(current_dices[0])
		current_dice, IDK := strconv.Atoi(current_dices[1])
		if IDK != nil {
			log.Fatal("Введите броски через пробел в формате [кол-во дайсов]d[значение дайса]")
		}
		Coefficients = append(Coefficients, current_coefficient)
		DiceSize = append(DiceSize, current_dice)
	}
	return Coefficients, DiceSize
}

func main() {
	ValidateInput()
	SliceDicesAsNumbers()
	nums, dices := SliceDicesAsNumbers()
	fmt.Println(nums)
	fmt.Println(dices)

	// rolls, rolls_output := RollAllDices()
	// PrintOutput(SumRolls(rolls), rolls_output)
}
