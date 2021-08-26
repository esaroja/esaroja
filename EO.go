package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Println("enter the value:")
	fmt.Scanln(&a)
	if a%2 == 0 { //checks if number is even
		fmt.Println("The number", a, "is even")
		
	}else{
	fmt.Println("The number", a, "is odd")
	}
}
