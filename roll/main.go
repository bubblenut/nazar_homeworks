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

func SumRolls(rolls, nums []int) int {
	sum := 0
	for i := 0; i < len(rolls); i++ {
		sum += rolls[i]
	}
	for k := 1; k < len(nums); k++ {
		sum += nums[k]
	}
	return sum
}

func PrintOutput(sum int, rolls_output []string, nums []int) {
	fmt.Println("Total:", sum)
	fmt.Print("Rolls:", rolls_output, "+", nums[1:])
}

func StringToInt(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
				log.Fatal("Введите бросок в формате (кол-во дайсов)d(размер дайса)")
			}
	return num
}

func CheckMassive(massive []int) []int {
	if len(massive) == 0 {
		massive = append(massive, 0)
	}
	return massive
}

func ExtractDicesFromArgs() ([]int, []int) {
	var AllDices, counts []int
	for i := 1; i < len(os.Args); i += 1 {
		if os.Args[i] == "+" || os.Args[i] == "," {
			i += 1
		}
		parts := strings.Split(os.Args[i], "d")
		if len(parts) == 1 {
			count := StringToInt(parts[0])
			counts = append(counts, count)
		} else {
			var coef, size int
			if parts[0] == "" {
				coef = 1
			} else {
				coef = StringToInt(parts[0])
			}
			size = StringToInt(parts[1])
			for i := 1; i <= coef; i++ {
				AllDices = append(AllDices, size)
			}
		}
		counts = CheckMassive(counts)
	}
	return AllDices, counts
}
func main() {
	ValidateInput()
	dices, mods := ExtractDicesFromArgs()
	rolls, text := RollAllDices(dices)
	PrintOutput(SumRolls(rolls, mods), text, mods)
	// rolls, rolls_output := RollAllDices()
	// PrintOutput(SumRolls(rolls), rolls_output)
}
