package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func mul(a, b int) int {
	return a * b
}
func div(a, b int) int {
	return a / b
}


func main() {
	var operation string
	var number1, number2 int

	fmt.Println("Welcome to my calculator program in go")
	fmt.Println("=======================================")
	fmt.Println("What operation would you like to perform? (add, sub, mul, div)")
	fmt.Scanf("%s", &operation)
	fmt.Println("Enter the first number")
	fmt.Scanf("%d", &number1)
	fmt.Println("Enter the second number")
	fmt.Scanf("%d", &number2)


	switch operation {
		case "add":
			fmt.Println("The sum of", number1, "and", number2, "is", add(number1, number2))
		case "sub":
			fmt.Println("The difference of", number1, "and", number2, "is", sub(number1, number2))
		case "mul":
			fmt.Println("The product of", number1, "and", number2, "is", mul(number1, number2))
		case "div":
			fmt.Println("The division of", number1, "and", number2, "is", div(number1, number2))
		default:
			fmt.Println("Invalid operation")
	}
}