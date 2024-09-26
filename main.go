package main

import "fmt"

func printFullName(firstName, lastName string) {
	fmt.Println("Full Name:", firstName+" "+lastName)
}



func max(num1, num2 int) int {
	var result int
	if (num1 > num2) {
	   result = num1
	} else {
	   result = num2
	}
	return result 
 }





func main() {

	firstName := "Naga"  
	lastName := "Teja"
	printFullName(firstName, lastName)




	
	var a int = 100
    var b int = 200
    var c int
    c = max(a, b)
    fmt.Printf( "Max value is : %d\n", c )


	
}




