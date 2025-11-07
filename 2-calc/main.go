package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const AVG = "AVG"
const SUM = "SUM"
const MED = "MED"

// Калькулятор принимает тип операции в параметре командной строки
// и список чисел через ввод в приложении
func main() {
	operation, numbers, err := parseInput(os.Args[1:])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	switch operation {
	case AVG:
		sum, count := calcSum(numbers)
		if count < 1 {
			fmt.Println("Количество чисел равно нулю, нельзя вычислить среднее")
		} else {
			fmt.Printf("AVG: %.2f\n", float64(sum)/float64(count))
		}
	case SUM:
		sum, _ := calcSum(numbers)
		fmt.Printf("SUM: %d\n", sum)
	case MED:
		med := calcMed(numbers)
		fmt.Printf("MED: %.1f\n", med)
	}
}

func parseInput(args []string) (operation string, numbers []int, err error) {
	if len(args) != 1 {
		err = fmt.Errorf("калькулятор принимает один параметр - тип операции (AVG, SUM или MED)")
		return
	}
	operation = args[0]
	if operation != AVG && operation != SUM && operation != MED {
		err = fmt.Errorf("ошибочный тип операции %s, должен быть AVG, SUM или MED", operation)
		return
	}
	reader := bufio.NewReader(os.Stdin)
	numbersLine, err := reader.ReadString('\n')
	if err != nil {
		err = fmt.Errorf("ошибка при вводе списка чисел: %v", err)
		return
	}
	numbersLineParts := strings.Split(numbersLine, ",")
	numbers = make([]int, len(numbersLineParts))
	for index, numberString := range numbersLineParts {
		numbers[index], err = strconv.Atoi(strings.TrimSpace(numberString))
		if err != nil {
			err = fmt.Errorf("ошибка в формате числа в позиции %d: %v", index, err)
			return
		}
	}
	return
}

func calcSum(numbers []int) (sum int, count int) {
	for _, number := range numbers {
		sum += number
		count++
	}
	return
}

func calcMed(numbers []int) float64 {
	sort.Ints(numbers)
	middle := len(numbers) / 2
	if len(numbers)%2 == 0 {
		return float64(numbers[middle-1]+numbers[middle]) / 2
	} else {
		return float64(numbers[middle])
	}
}
