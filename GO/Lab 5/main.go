package main

import (
	"fmt"
	"math"
	"strings"
)

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero: cannot divide %v by zero", a)
	}
	return a / b, nil
}

func SquareRoot(x float64) (float64, error) {
	if x < 0 {
		return 0, fmt.Errorf("negative square root: value %v must be non-negative", x)
	}
	return math.Sqrt(x), nil
}

func ProcessInput(input string) (string, error) {
	trimmed := strings.TrimSpace(input)
	if trimmed == "" {
		return "", fmt.Errorf("empty input: input string cannot be empty or whitespace only")
	}
	return strings.ToUpper(trimmed), nil
}

func main() {
	divRes, err := Divide(10, 2)
	if err != nil {
		fmt.Printf("Divide error: %v\n", err)
	} else {
		fmt.Printf("10 / 2 = %v\n", divRes)
	}

	_, err = Divide(10, 0)
	if err != nil {
		fmt.Printf("Divide error: %v\n", err)
	}

	sqrtRes, err := SquareRoot(25)
	if err != nil {
		fmt.Printf("SquareRoot error: %v\n", err)
	} else {
		fmt.Printf("sqrt(25) = %v\n", sqrtRes)
	}

	_, err = SquareRoot(-9)
	if err != nil {
		fmt.Printf("SquareRoot error: %v\n", err)
	}

	procRes, err := ProcessInput("hello go")
	if err != nil {
		fmt.Printf("ProcessInput error: %v\n", err)
	} else {
		fmt.Printf("Processed: %s\n", procRes)
	}

	_, err = ProcessInput("   ")
	if err != nil {
		fmt.Printf("ProcessInput error: %v\n", err)
	}
}
