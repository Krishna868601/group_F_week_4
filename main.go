package main

import "fmt"

func printFullName(firstName, lastName string) {
	fmt.Println("Full Name:", firstName+" "+lastName)
}

func max(num1, num2 int) int {
	var result int
	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

func main() {
	fmt.Println("Welcome to Group F's Week 4 Project!")
	firstName := "Naga"
	lastName := "Teja"
	printFullName(firstName, lastName)

	var a int = 100
	var b int = 200
	var c int
	c = max(a, b)
	fmt.Printf("Max value is : %d\n", c)

	value := total()
	fmt.Println(value)

	var num int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)

	if IsEven(num) {
		fmt.Println(num, "is even.")
	} else {
		fmt.Println(num, "is not even.")
	}
	arr := []int{1, 2, 3, 4, 5}
	reverseArray(arr)
	fmt.Println(arr)

}
