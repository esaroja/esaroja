package main
import(
    "fmt"
)
func main(){
    var n int
	fmt.Println("enter the n value ")
	fmt.Scanln(&n)
   
	if n%2 != 0 { //checks if number is even
		fmt.Printf("Weird")
		
	}else if n%2 == 0  && n>20{

	fmt.Printf("Not Wierd")

	}else {
	fmt.Printf("wierd")
	}

}
