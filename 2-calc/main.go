package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Калькулятор принимает тип операции в параметре командной строки
// и список чисел через ввод в приложении
func main() {
	availableOperations := getAvailableOperations()

	operation, err := getOperation(os.Args[1:], availableOperations)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	calculator := operationFactory(operation)
	if calculator == nil {
		fmt.Fprintf(os.Stderr, "ошибочный тип операции %s, должен быть %s", operation, availableOperations)
		os.Exit(1)
	}
	numbers, err := parseInput()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	result := calculator(numbers...)
	fmt.Printf("%s: %.1f\n", operation, result)
}

func getOperation(args []string, availableOperations string) (operation string, err error) {
	if len(args) != 1 {
		err = fmt.Errorf("калькулятор принимает один параметр - тип операции (%)", availableOperations)
		return
	}
	operation = args[0]
	return
}

func parseInput() (numbers []int, err error) {
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

const AVG = "AVG"
const SUM = "SUM"
const MED = "MED"

var operations = map[string]func(...int) float64{
	AVG: calcAvg,
	SUM: calcSum,
	MED: calcMed,
}

func getAvailableOperations() string {
	keys := make([]string, 0, len(operations))
	for k := range operations {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var builder strings.Builder
	for index, op := range keys {
		if index > 0 {
			if index == len(keys)-1 {
				builder.WriteString(" и ")
			} else {
				builder.WriteString(", ")
			}
		}
		builder.WriteString(op)
	}
	return builder.String()
}

func operationFactory(operation string) func(...int) float64 {
	return operations[operation]
}

func calcSum(numbers ...int) (sum float64) {
	if len(numbers) == 0 {
		sum = math.NaN()
		return
	}
	for _, number := range numbers {
		sum += float64(number)
	}
	return
}

func calcMed(numbers ...int) float64 {
	if len(numbers) == 0 {
		return math.NaN()
	}
	sort.Ints(numbers)
	middle := len(numbers) / 2
	if len(numbers)%2 == 0 {
		return float64(numbers[middle-1]+numbers[middle]) / 2
	} else {
		return float64(numbers[middle])
	}
}

func calcAvg(numbers ...int) float64 {
	if len(numbers) < 1 {
		return math.NaN()
	}
	sum := calcSum(numbers...)
	return float64(sum) / float64(len(numbers))
}
