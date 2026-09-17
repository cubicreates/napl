/*
This program performs calculations using integers and floating point numbers.
*/
package main

import "fmt"

func main() {
	menu()
}

// Function for integer calculations
func resi() {

	calcint := func(a int, b int, op string) (int, bool) {
		switch op {
		case "+":
			return a + b, true
		case "-":
			return a - b, true
		case "*":
			return a * b, true
		case "/":
			if b != 0 {
				return a / b, true
			}
			return 0, false
		default:
			return 0, false
		}
	}

	fmt.Println("---- Integer Calculator ----")

	var a, b int
	var op string

	for {
		fmt.Print("Enter a number 1 > ")
		_, err := fmt.Scan(&a)
		if err == nil {
			break
		}

		fmt.Println("Invalid input!")
		var trash string
		fmt.Scan(&trash)
	}

	for {
		fmt.Print("Enter a number 2 > ")
		_, err := fmt.Scan(&b)
		if err == nil {
			break
		}

		fmt.Println("Invalid input!")
		var trash string
		fmt.Scan(&trash)
	}

	for {
		fmt.Print("Enter an operator ('+','-','*','/') > ")
		fmt.Scan(&op)

		if op == "+" || op == "-" || op == "*" || op == "/" {
			break
		}

		fmt.Println("Invalid operator!")
	}

	result, valid := calcint(a, b, op)

	if !valid {
		fmt.Println("Invalid result! Cannot divide by zero.")
	} else {
		fmt.Println("The result of", a, op, b, "is", result)
	}
}

// Function for floating point calculations
func resf() {

	calc64 := func(a float64, b float64, op string) (float64, bool) {
		switch op {
		case "+":
			return a + b, true
		case "-":
			return a - b, true
		case "*":
			return a * b, true
		case "/":
			if b != 0 {
				return a / b, true
			}
			return 0, false
		default:
			return 0, false
		}
	}

	fmt.Println("---- Floating Point Calculator ----")

	var a, b float64
	var op string

	for {
		fmt.Print("Enter a number 1 > ")
		_, err := fmt.Scan(&a)
		if err == nil {
			break
		}

		fmt.Println("Invalid input!")
		var trash string
		fmt.Scan(&trash)
	}

	for {
		fmt.Print("Enter a number 2 > ")
		_, err := fmt.Scan(&b)
		if err == nil {
			break
		}

		fmt.Println("Invalid input!")
		var trash string
		fmt.Scan(&trash)
	}

	for {
		fmt.Print("Enter an operator ('+','-','*','/') > ")
		fmt.Scan(&op)

		if op == "+" || op == "-" || op == "*" || op == "/" {
			break
		}

		fmt.Println("Invalid operator!")
	}

	result, valid := calc64(a, b, op)

	if !valid {
		fmt.Println("Invalid result! Cannot divide by zero.")
	} else {
		fmt.Println("The result of", a, op, b, "is", result)
	}
}

// Menu
func menu() {
	for {
		fmt.Println("---- Calculations Menu ----")
		fmt.Println("1. Perform Calculations using integers")
		fmt.Println("2. Perform Calculations using floating point numbers")
		fmt.Println("3. Exit")

		var choice int
		fmt.Print("Enter your choice > ")
		_, err := fmt.Scan(&choice)

		if err != nil {
			fmt.Println("Invalid choice!")
			var trash string
			fmt.Scan(&trash)
			continue
		}

		switch choice {
		case 1:
			resi()
		case 2:
			resf()
		case 3:
			fmt.Println("Thank you! Goodbye!")
			return
		default:
			fmt.Println("Invalid choice!")
		}
	}
}