package main

import(
    "fmt"
)

func main(){
 var  mealcost float64
 var  tip float64
 var  tax float64
 var totalCost int

    fmt.Println("enter the mealcost")
	fmt.Scanln(&mealcost)
	fmt.Println("enter the tip")
	fmt.Scanln(&tip)
	tip= float64(mealcost*0.2)
	fmt.Println("percentage of meal after added the tip",tip)
	fmt.Println("enter the tax")
	fmt.Scanln(&tax)
	tax=float64(mealcost*0.08)
	fmt.Println("percentage of meal after added the tax",tax)
	totalCost=int(mealcost+tip+tax)
	fmt.Println("totalcost",totalCost)
	
}
