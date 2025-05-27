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

func ResultAllDices (nums []int, dice_counts []int) []int {
	if len(nums) != len(dice_counts) {
		log.Fatal("Перед всеми дайсами должно быть указано кол-во бросков")
	}
	var rolling_dices []int
	for i := 0; i < len(nums); i += 1 {
		for k := 0; k < nums[i]; k += 1 {
			rolling_dices = append(rolling_dices, dice_counts[i])
		}
	}

	return rolling_dices
}

func RollAllDices(rolling_dices []int) ([]int, []string) {
	var Rolls []int
	var RollsOutput []string
	for i := 0; i < len(rolling_dices); i++ {
		current_roll := GenerateRandomRoll(rolling_dices[i])
		Rolls = append(Rolls, current_roll)
		RollsOutput = append(RollsOutput, strconv.Itoa(current_roll)+"/"+strconv.Itoa(rolling_dices[i])+",")
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
		current_coefficient, err := strconv.Atoi(current_dices[0])
		current_dice, err := strconv.Atoi(current_dices[1])
		if err != nil {
			log.Fatal("Введите броски через пробел в формате [кол-во дайсов]d[значение дайса]")
		}
		Coefficients = append(Coefficients, current_coefficient)
		DiceSize = append(DiceSize, current_dice)
	}
	return Coefficients, DiceSize
}

func main() {
	ValidateInput()
	rolls_counts, rolls_text := (RollAllDices(ResultAllDices(SliceDicesAsNumbers())))
	PrintOutput(SumRolls(rolls_counts), rolls_text)
	// rolls, rolls_output := RollAllDices()
	// PrintOutput(SumRolls(rolls), rolls_output)
}
