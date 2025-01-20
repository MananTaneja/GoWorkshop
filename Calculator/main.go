package main

import "fmt"

func main() {
	var operation string
	var num1, num2 int

	fmt.Print("Choose an operation (+, -, *, /): ")
	fmt.Scanf("%s", &operation)
	fmt.Printf("Enter first number: ")
	fmt.Scanf("%d", &num1)
	fmt.Printf("Enter second number: ")
	fmt.Scanf("%d", &num2)

	switch operation {
	case "+":
		fmt.Printf("Result: %d\n", num1+num2)
	case "-":
		fmt.Printf("Result: %d\n", num1-num2)
	case "*":
		fmt.Printf("Result: %d\n", num1*num2)
	case "/":
		if num2 != 0 {
			fmt.Printf("Result: %d\n", num1/num2)
		} else {
			fmt.Println("Division by zero")
		}
	}
}
